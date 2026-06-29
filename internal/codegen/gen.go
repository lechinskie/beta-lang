package codegen

import (
	"fmt"
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/eramoss/b/internal/parser"
	"github.com/eramoss/b/internal/sema"
)

const maxTemps = 8

type scanner struct {
	*parser.BasebListener
	st          *sema.SymbolTable
	arraySizes  map[string]int
	curFunc     string
	funcLocals  map[string]map[string]int
	funcParams  map[string]int
	funcOrder   []string
	nextOff     int
}

func (s *scanner) EnterExt_def(ctx *parser.Ext_defContext) {
	if ctx.EXTRN() != nil || ctx.VARIADIC() != nil {
		return
	}

	if ctx.Statement() != nil {
		name := ctx.ID().GetText()
		s.curFunc = name
		s.funcLocals[name] = make(map[string]int)
		s.funcOrder = append(s.funcOrder, name)
		s.nextOff = 0

		if ctx.Arg_list() != nil {
			for _, id := range ctx.Arg_list().AllID() {
				pname := id.GetText()
				s.funcLocals[name][pname] = s.nextOff
				s.nextOff += 4
			}
			s.funcParams[name] = len(ctx.Arg_list().AllID())
		}
		return
	}

	if ctx.ASM() != nil {
		return
	}

	if strings.Contains(ctx.GetText(), "[") {
		name := ctx.ID().GetText()
		sizeExpr := ctx.Expr()
		size := 0
		if sizeExpr != nil {
			fmt.Sscanf(sizeExpr.GetText(), "%d", &size)
		}
		if size <= 0 {
			size = 1
		}
		s.arraySizes[name] = size
		return
	}

	if s.curFunc == "" {
		s.arraySizes[ctx.ID().GetText()] = 1
	}
}

func (s *scanner) ExitExt_def(ctx *parser.Ext_defContext) {
	if ctx.Statement() != nil {
		s.curFunc = ""
	}
}

func (s *scanner) EnterAuto_decl(ctx *parser.Auto_declContext) {
	if s.curFunc == "" {
		return
	}
	locals := s.funcLocals[s.curFunc]
	for _, def := range ctx.AllAuto_def() {
		name := def.ID().GetText()
		if _, exists := locals[name]; !exists {
			locals[name] = s.nextOff
			s.nextOff += 4
		}
	}
}

type Generator struct {
	*parser.BasebListener
	out    strings.Builder
	st     *sema.SymbolTable
	source string
	errors *sema.ErrorList

	localOffs  map[string]int
	frameSize  int
	arraySizes map[string]int
	funcName   string
	inFunc     bool
	isFirst    bool
	tempDepth  int

	funcLocals map[string]map[string]int
	funcParams map[string]int
	funcOrder  []string

	labelCounter int
	exitLabels   []string
	breakLabels  []string
	ifElseStack  [][2]string
}

func Generate(source string, errors *sema.ErrorList, tree antlr.ParseTree, st *sema.SymbolTable) *strings.Builder {
	scan := &scanner{
		st:         st,
		arraySizes: make(map[string]int),
		funcLocals: make(map[string]map[string]int),
		funcParams: make(map[string]int),
		funcOrder:  make([]string, 0),
	}
	antlr.ParseTreeWalkerDefault.Walk(scan, tree)

	g := &Generator{
		st:         st,
		source:     source,
		errors:     errors,
		arraySizes: scan.arraySizes,
		funcLocals: scan.funcLocals,
		funcParams: scan.funcParams,
		funcOrder:  scan.funcOrder,
		isFirst:    false,
	}

	g.emitData()
	g.emitDirective(".text")

	prog := tree.(*parser.ProgramContext)
	type funcEntry struct {
		ctx   parser.IExt_defContext
		order int
	}
	funcs := make([]funcEntry, 0)
	for i, ed := range prog.AllExt_def() {
		if ed.Statement() != nil {
			funcs = append(funcs, funcEntry{ed, i})
		}
	}
	sort.Slice(funcs, func(i, j int) bool {
		ni := funcs[i].ctx.ID().GetText()
		nj := funcs[j].ctx.ID().GetText()
		if ni == "main" {
			return true
		}
		if nj == "main" {
			return false
		}
		return funcs[i].order < funcs[j].order
	})
	for _, f := range funcs {
		antlr.ParseTreeWalkerDefault.Walk(g, f.ctx)
	}

	return &g.out
}

func mangle(name string) string {
	return "_" + name
}

func (g *Generator) emit(format string, args ...any) {
	fmt.Fprintf(&g.out, "  "+format+"\n", args...)
}

func (g *Generator) emitDirective(format string, args ...any) {
	fmt.Fprintf(&g.out, format+"\n", args...)
}

func (g *Generator) emitLabel(label string) {
	fmt.Fprintf(&g.out, "%s:\n", label)
}

func (g *Generator) emitLi(rd string, val int) {
	if val >= -2048 && val <= 2047 {
		g.emit("addi %s, zero, %d", rd, val)
	} else {
		upper := (val + 0x800) >> 12
		lower := val - (upper << 12)
		g.emit("lui %s, %d", rd, upper)
		if lower != 0 {
			if lower >= 2048 {
				lower -= 4096
			}
			g.emit("addi %s, %s, %d", rd, rd, lower)
		}
	}
}

func (g *Generator) pushA0() {
	g.emit("sw a0, %d(sp)", g.tempDepth*4)
	g.tempDepth++
}

func (g *Generator) popT0() {
	g.tempDepth--
	g.emit("lw t0, %d(sp)", g.tempDepth*4)
}

func (g *Generator) pushT0() {
	g.emit("sw t0, %d(sp)", g.tempDepth*4)
	g.tempDepth++
}

func (g *Generator) popA0() {
	g.tempDepth--
	g.emit("lw a0, %d(sp)", g.tempDepth*4)
}

func (g *Generator) emitLoadVar(name string) {
	if off, ok := g.localOffs[name]; ok {
		g.emit("lw a0, %d(sp)", off)
		return
	}
	sym := g.st.Lookup(name)
	if sym == nil {
		g.emitLi("a0", 0)
		return
	}
	if sym.Kind == sema.KindArray {
		g.emit("la a0, %s", mangle(name))
		return
	}
	g.emit("la t0, %s", mangle(name))
	g.emit("lw a0, 0(t0)")
}

func (g *Generator) emitStoreVar(name string) {
	if off, ok := g.localOffs[name]; ok {
		g.emit("sw a0, %d(sp)", off)
		return
	}
	g.emit("la t0, %s", mangle(name))
	g.emit("sw a0, 0(t0)")
}

func (g *Generator) emitGetAddr(name string) {
	if off, ok := g.localOffs[name]; ok {
		g.emit("addi t0, sp, %d", off)
	} else {
		g.emit("la t0, %s", mangle(name))
	}
}

func (g *Generator) emitExit() {
	if g.frameSize > 0 {
		g.emit("addi sp, sp, %d", g.frameSize)
	}
	g.emitLi("a0", 0)
	g.emit("addi a7, zero, 93")
	g.emit("ecall")
}

func (g *Generator) newLabel() string {
	label := fmt.Sprintf("_L%d", g.labelCounter)
	g.labelCounter++
	return label
}
