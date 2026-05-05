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
}

func NewSemanticChecker(symbols *SymbolTable, errors *ErrorList, source string) *SemanticChecker {
	return &SemanticChecker{
		symbols: symbols,
		errors:  errors,
		source:  source,
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
	// Check if this is an extrn definition
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

	// Check if this is a variadic definition
	if ctx.VARIADIC() != nil {
		return
	}

	// Check if this is a naked function definition with asm
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

	// Check if this is a function definition: ID '(' arg_list? ')' statement
	if strings.Contains(ctx.GetText(), "(") && strings.Contains(ctx.GetText(), ")") {
		// Function definition
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

		// Enter function scope
		c.symbols.EnterScope()
		c.inFuncScope = true

		// Register parameters
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

	// Check if this is an array declaration: ID '[' expr? ']' ival_list? ';'
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

	// Simple global declaration: ID ';'
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

		// Skip if this is a function call being declared (handled elsewhere)
		// Check if variable exists in symbol table
		if c.symbols.Lookup(name) == nil {
			c.addError(line, col, fmt.Sprintf("'%s' undeclared", name))
		}
	}
}

func CheckSemantic(source string, errors *ErrorList, tree antlr.ParseTree) {
	symbols := NewSymbolTable()
	checker := NewSemanticChecker(symbols, errors, source)
	antlr.ParseTreeWalkerDefault.Walk(checker, tree)
}
