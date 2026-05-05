package analysis

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/eramoss/b/internal/parser"
)

type SemanticChecker struct {
	*parser.BasebListener
	symbols     *SymbolTable
	errors      *ErrorList
	source      string
	inFuncScope bool
	labels      map[string]bool
	gotoStmts   []gotoInfo
	inGoto      bool
}

type gotoInfo struct {
	name string
	line int
	col  int
}

func NewSemanticChecker(symbols *SymbolTable, errors *ErrorList, source string) *SemanticChecker {
	return &SemanticChecker{
		symbols:   symbols,
		errors:    errors,
		source:    source,
		labels:    make(map[string]bool),
		gotoStmts: make([]gotoInfo, 0),
	}
}

func (c *SemanticChecker) addError(line, col int, msg string) {
	c.errors.Add(&Error{
		Kind:   KindParser,
		Line:   line,
		Column: col,
		Text:   msg,
	})
}

func (c *SemanticChecker) EnterProgram(ctx *parser.ProgramContext) {
	c.symbols.EnterScope()
}

func (c *SemanticChecker) ExitProgram(ctx *parser.ProgramContext) {
	c.symbols.ExitScope()
}

func (c *SemanticChecker) EnterExt_def(ctx *parser.Ext_defContext) {
	if ctx.EXTRN() != nil {
		nameList := ctx.Name_list()
		if nameList != nil {
			for _, idTok := range nameList.AllID() {
				n := idTok.GetText()
				l := idTok.GetSymbol().GetLine()
				co := idTok.GetSymbol().GetColumn() + 1

				sym := &Symbol{
					Name:       n,
					Type:       "word",
					Scope:      c.symbols.CurrentDepth(),
					IsDeclared: true,
					Kind:       KindExtrn,
					ParamPos:   -1,
				}
				if err := c.symbols.Insert(n, sym); err != nil {
					c.addError(l, co, fmt.Sprintf("'%s' already declared", n))
				}
			}
		}
		return
	}

	if ctx.VARIADIC() != nil {
		return
	}

	if ctx.ASM() != nil && ctx.String_list() != nil {
		idCtx := ctx.ID()
		if idCtx == nil {
			return
		}
		name := idCtx.GetText()
		line := idCtx.GetSymbol().GetLine()
		col := idCtx.GetSymbol().GetColumn() + 1

		sym := &Symbol{
			Name:       name,
			Type:       "func",
			Scope:      c.symbols.CurrentDepth(),
			IsDeclared: true,
			Kind:       KindFunc,
			ParamPos:   -1,
		}
		if err := c.symbols.Insert(name, sym); err != nil {
			c.addError(line, col, fmt.Sprintf("'%s' already declared", name))
		}
		return
	}

	idCtx := ctx.ID()
	if idCtx == nil {
		return
	}
	name := idCtx.GetText()
	line := idCtx.GetSymbol().GetLine()
	col := idCtx.GetSymbol().GetColumn() + 1

	if strings.Contains(ctx.GetText(), "(") && strings.Contains(ctx.GetText(), ")") {
		sym := &Symbol{
			Name:       name,
			Type:       "func",
			Scope:      c.symbols.CurrentDepth(),
			IsDeclared: true,
			Kind:       KindFunc,
			ParamPos:   -1,
		}
		if err := c.symbols.Insert(name, sym); err != nil {
			c.addError(line, col, fmt.Sprintf("'%s' already declared", name))
		}

		c.symbols.EnterScope()
		c.inFuncScope = true
		c.labels = make(map[string]bool)
		c.gotoStmts = make([]gotoInfo, 0)

		argList := ctx.Arg_list()
		if argList != nil {
			for i, paramID := range argList.AllID() {
				paramName := paramID.GetText()
				paramLine := paramID.GetSymbol().GetLine()
				paramCol := paramID.GetSymbol().GetColumn() + 1

				paramSym := &Symbol{
					Name:       paramName,
					Type:       "word",
					Scope:      c.symbols.CurrentDepth(),
					IsDeclared: true,
					Kind:       KindParam,
					ParamPos:   i,
				}
				if err := c.symbols.Insert(paramName, paramSym); err != nil {
					c.addError(paramLine, paramCol, fmt.Sprintf("parameter '%s' already declared", paramName))
				}
			}
		}
		return
	}

	if strings.Contains(ctx.GetText(), "[") {
		sym := &Symbol{
			Name:       name,
			Type:       "array",
			Scope:      c.symbols.CurrentDepth(),
			IsDeclared: true,
			Kind:       KindArray,
			ParamPos:   -1,
		}
		if err := c.symbols.Insert(name, sym); err != nil {
			c.addError(line, col, fmt.Sprintf("'%s' already declared", name))
		}
		return
	}

	sym := &Symbol{
		Name:       name,
		Type:       "word",
		Scope:      c.symbols.CurrentDepth(),
		IsDeclared: true,
		Kind:       KindVar,
		ParamPos:   -1,
	}
	if err := c.symbols.Insert(name, sym); err != nil {
		c.addError(line, col, fmt.Sprintf("'%s' already declared", name))
	}
}

func (c *SemanticChecker) ExitExt_def(ctx *parser.Ext_defContext) {
	if c.inFuncScope && !strings.Contains(ctx.GetText(), "extrn") && !strings.Contains(ctx.GetText(), "__variadic__") {
		for _, g := range c.gotoStmts {
			if !c.labels[g.name] {
				c.addError(g.line, g.col, fmt.Sprintf("label '%s' undeclared", g.name))
			}
		}
		c.symbols.ExitScope()
		c.inFuncScope = false
	}
}

func (c *SemanticChecker) EnterAuto_decl(ctx *parser.Auto_declContext) {
	for _, autoDef := range ctx.AllAuto_def() {
		idCtx := autoDef.ID()
		if idCtx == nil {
			continue
		}
		name := idCtx.GetText()
		line := idCtx.GetSymbol().GetLine()
		col := idCtx.GetSymbol().GetColumn() + 1

		sym := &Symbol{
			Name:       name,
			Type:       "word",
			Scope:      c.symbols.CurrentDepth(),
			IsDeclared: true,
			Kind:       KindVar,
			ParamPos:   -1,
		}
		if err := c.symbols.Insert(name, sym); err != nil {
			c.addError(line, col, fmt.Sprintf("'%s' already declared", name))
		}
	}
}

func (c *SemanticChecker) EnterStatement(ctx *parser.StatementContext) {
	childCount := ctx.GetChildCount()
	if childCount >= 2 {
		first := ctx.GetChild(0)
		second := ctx.GetChild(1)

		if t, ok := first.(antlr.TerminalNode); ok && t.GetSymbol().GetTokenType() == 61 {
			if s, ok := second.(antlr.TerminalNode); ok && s.GetSymbol().GetText() == ":" {
				labelName := t.GetSymbol().GetText()
				line := t.GetSymbol().GetLine()
				col := t.GetSymbol().GetColumn() + 1

				if c.labels[labelName] {
					c.addError(line, col, fmt.Sprintf("label '%s' already declared", labelName))
					return
				}
				c.labels[labelName] = true
				return
			}
		}
	}

	if ctx.GOTO() != nil {
		c.inGoto = true
		expr := ctx.Expr()
		if expr != nil {
			idTok := expr.ID()
			if idTok != nil {
				c.gotoStmts = append(c.gotoStmts, gotoInfo{
					name: idTok.GetText(),
					line: idTok.GetSymbol().GetLine(),
					col:  idTok.GetSymbol().GetColumn() + 1,
				})
			}
		}
	}
}

func (c *SemanticChecker) ExitStatement(ctx *parser.StatementContext) {
	if ctx.GOTO() != nil {
		c.inGoto = false
	}
}

func (c *SemanticChecker) EnterCompound_stmt(ctx *parser.Compound_stmtContext) {
	c.symbols.EnterScope()
}

func (c *SemanticChecker) ExitCompound_stmt(ctx *parser.Compound_stmtContext) {
	c.symbols.ExitScope()
}

func (c *SemanticChecker) EnterExpr(ctx *parser.ExprContext) {
	idCtx := ctx.ID()
	if idCtx != nil {
		name := idCtx.GetText()
		line := idCtx.GetSymbol().GetLine()
		col := idCtx.GetSymbol().GetColumn() + 1

		if c.inGoto {
			return
		}

		if c.labels[name] {
			return
		}

		sym := c.symbols.Lookup(name)
		if sym == nil {
			c.addError(line, col, fmt.Sprintf("'%s' undeclared", name))
		}
	}
}

func CheckSemantic(source string, errors *ErrorList, tree antlr.ParseTree) {
	symbols := NewSymbolTable()
	checker := NewSemanticChecker(symbols, errors, source)
	antlr.ParseTreeWalkerDefault.Walk(checker, tree)
}
