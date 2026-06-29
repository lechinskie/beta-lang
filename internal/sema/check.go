package sema

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
	exprTypes   map[antlr.ParserRuleContext]string
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
		exprTypes: make(map[antlr.ParserRuleContext]string),
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
		c.symbols.ExitScope()
		for _, g := range c.gotoStmts {
			if !c.labels[g.name] {
				c.addError(g.line, g.col, fmt.Sprintf("label '%s' undeclared", g.name))
			}
		}
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

func (c *SemanticChecker) getType(ctx parser.IExprContext) string {
	if ctx == nil {
		return ""
	}
	exprCtx, ok := ctx.(*parser.ExprContext)
	if !ok {
		return ""
	}
	t, ok := c.exprTypes[exprCtx]
	if ok {
		return t
	}
	return ""
}

func (c *SemanticChecker) ExitExpr(ctx *parser.ExprContext) {
	if ctx.STRING() != nil && len(ctx.AllExpr()) == 0 {
		c.exprTypes[ctx] = "string"
		return
	}

	if ctx.DECIMAL() != nil {
		c.exprTypes[ctx] = "int"
		return
	}
	if ctx.OCTAL() != nil {
		c.exprTypes[ctx] = "int"
		return
	}
	if ctx.CHAR() != nil {
		c.exprTypes[ctx] = "int"
		return
	}

	if ctx.ID() != nil && len(ctx.AllExpr()) == 0 {
		name := ctx.ID().GetText()
		sym := c.symbols.Lookup(name)
		if sym != nil {
			if sym.Type == "string" {
				c.exprTypes[ctx] = "string"
			} else {
				c.exprTypes[ctx] = "int"
			}
		}
		return
	}

	if len(ctx.AllExpr()) == 2 {
		left := ctx.AllExpr()[0]
		right := ctx.AllExpr()[1]
		leftType := c.getType(left)
		rightType := c.getType(right)

		if ctx.Assign_op() != nil {
			c.exprTypes[ctx] = rightType
			e, ok := left.(*parser.ExprContext)
			if ok && e.ID() != nil {
				sym := c.symbols.Lookup(e.ID().GetText())
				if sym != nil {
					sym.Type = rightType
				}
			}
			return
		}

		text := ctx.GetText()
		isArith := strings.Contains(text, "+") || strings.Contains(text, "-") || strings.Contains(text, "*") || strings.Contains(text, "/") || strings.Contains(text, "%") || strings.Contains(text, "<<") || strings.Contains(text, ">>") || strings.Contains(text, "&") || strings.Contains(text, "|") || strings.Contains(text, "^")

		if isArith {
			if leftType == "string" && rightType == "int" {
				line := ctx.GetStart().GetLine()
				col := ctx.GetStart().GetColumn() + 1
				c.addError(line, col, fmt.Sprintf("incompatible types: cannot mix '%s' with '%s' in arithmetic", leftType, rightType))
			}
			if leftType == "int" && rightType == "string" {
				line := ctx.GetStart().GetLine()
				col := ctx.GetStart().GetColumn() + 1
				c.addError(line, col, fmt.Sprintf("incompatible types: cannot mix '%s' with '%s' in arithmetic", leftType, rightType))
			}
			if leftType == "string" && rightType == "string" {
				line := ctx.GetStart().GetLine()
				col := ctx.GetStart().GetColumn() + 1
				c.addError(line, col, "incompatible types: cannot do arithmetic with strings")
			}
			c.exprTypes[ctx] = "int"
			return
		}

		c.exprTypes[ctx] = "int"
		return
	}
}

func CheckSemantic(source string, errors *ErrorList, tree antlr.ParseTree) *SymbolTable {
	symbols := NewSymbolTable()
	checker := NewSemanticChecker(symbols, errors, source)
	antlr.ParseTreeWalkerDefault.Walk(checker, tree)
	return symbols
}
