package codegen

import (
	"github.com/eramoss/b/internal/parser"
)

func (g *Generator) ExitStatement(ctx *parser.StatementContext) {
	if ctx.Expr() != nil {
		g.evalExpr(ctx.Expr())
		return
	}
}
