package codegen

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/eramoss/b/internal/parser"
	"github.com/eramoss/b/internal/sema"
)

const maxTemps = 8

type scanner struct {
	*parser.BasebListener
	st         *sema.SymbolTable
	localOffs  map[string]int
	arraySizes map[string]int
	nextOff    int
}

func (s *scanner) EnterAuto_decl(ctx *parser.Auto_declContext) {
	for _, def := range ctx.AllAuto_def() {
		name := def.ID().GetText()
		if _, exists := s.localOffs[name]; !exists {
			s.localOffs[name] = s.nextOff
			s.nextOff += 4
		}
	}
}

func (s *scanner) EnterExt_def(ctx *parser.Ext_defContext) {
	if ctx.EXTRN() != nil || ctx.VARIADIC() != nil || ctx.ASM() != nil {
		return
	}
	if ctx.Statement() != nil {
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
	sym := s.st.Lookup(ctx.ID().GetText())
	if sym != nil && sym.Kind == sema.KindVar && sym.Scope <= 1 {
		s.arraySizes[ctx.ID().GetText()] = 1
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
	tempDepth  int
}

func Generate(source string, errors *sema.ErrorList, tree antlr.ParseTree, st *sema.SymbolTable) *strings.Builder {
	scan := &scanner{
		st:         st,
		localOffs:  make(map[string]int),
		arraySizes: make(map[string]int),
	}
	antlr.ParseTreeWalkerDefault.Walk(scan, tree)

	locals := len(scan.localOffs)
	needed := max((locals+maxTemps)*4, 16)
	needed = (needed + 15) & ^15

	for name := range scan.localOffs {
		scan.localOffs[name] += maxTemps * 4
	}

	g := &Generator{
		st:         st,
		source:     source,
		errors:     errors,
		localOffs:  scan.localOffs,
		frameSize:  needed,
		arraySizes: scan.arraySizes,
	}

	g.emitData()

	g.emitDirective(".text")

	antlr.ParseTreeWalkerDefault.Walk(g, tree)

	if g.inFunc {
		g.emitExit()
	}

	return &g.out
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
		g.emit("la a0, %s", name)
		return
	}
	g.emit("la t0, %s", name)
	g.emit("lw a0, 0(t0)")
}

func (g *Generator) emitStoreVar(name string) {
	if off, ok := g.localOffs[name]; ok {
		g.emit("sw a0, %d(sp)", off)
		return
	}
	g.emit("la t0, %s", name)
	g.emit("sw a0, 0(t0)")
}

func (g *Generator) emitGetAddr(name string) {
	if off, ok := g.localOffs[name]; ok {
		g.emit("addi t0, sp, %d", off)
	} else {
		g.emit("la t0, %s", name)
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
