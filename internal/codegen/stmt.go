package codegen

import (
	"github.com/eramoss/b/internal/parser"
)

func (g *Generator) ExitStatement(ctx *parser.StatementContext) {
	switch {
	case ctx.IF() != nil:
		if ctx.ELSE() != nil {
			endLabel := g.exitLabels[len(g.exitLabels)-1]
			g.exitLabels = g.exitLabels[:len(g.exitLabels)-1]
			g.emitLabel(endLabel)
		} else {
			endLabel := g.exitLabels[len(g.exitLabels)-1]
			g.exitLabels = g.exitLabels[:len(g.exitLabels)-1]
			g.emitLabel(endLabel)
		}

	case ctx.WHILE() != nil:
		startLabel := g.exitLabels[len(g.exitLabels)-2]
		endLabel := g.exitLabels[len(g.exitLabels)-1]
		g.exitLabels = g.exitLabels[:len(g.exitLabels)-2]
		if len(g.breakLabels) > 0 {
			g.breakLabels = g.breakLabels[:len(g.breakLabels)-1]
		}
		g.emit("j %s", startLabel)
		g.emitLabel(endLabel)

	case ctx.BREAK() != nil:

	case ctx.RETURN() != nil:

	default:
		if ctx.Expr() != nil {
			g.evalExpr(ctx.Expr())
		}
	}

	if parent, ok := ctx.GetParent().(*parser.StatementContext); ok {
		if parent.IF() != nil && parent.ELSE() != nil && parent.Statement(0) == ctx {
			if len(g.ifElseStack) > 0 {
				pair := g.ifElseStack[len(g.ifElseStack)-1]
				g.ifElseStack = g.ifElseStack[:len(g.ifElseStack)-1]
				elseLabel, endLabel := pair[0], pair[1]
				g.emit("j %s", endLabel)
				g.emitLabel(elseLabel)
			}
		}
	}
}

func (g *Generator) EnterStatement(ctx *parser.StatementContext) {
	switch {
	case ctx.IF() != nil:
		g.evalExpr(ctx.Expr())
		if ctx.ELSE() != nil {
			elseLabel := g.newLabel()
			endLabel := g.newLabel()
			g.ifElseStack = append(g.ifElseStack, [2]string{elseLabel, endLabel})
			g.exitLabels = append(g.exitLabels, endLabel)
			g.emit("beq a0, zero, %s", elseLabel)
		} else {
			endLabel := g.newLabel()
			g.exitLabels = append(g.exitLabels, endLabel)
			g.emit("beq a0, zero, %s", endLabel)
		}

	case ctx.WHILE() != nil:
		startLabel := g.newLabel()
		endLabel := g.newLabel()
		g.exitLabels = append(g.exitLabels, startLabel, endLabel)
		g.breakLabels = append(g.breakLabels, endLabel)
		g.emitLabel(startLabel)
		g.evalExpr(ctx.Expr())
		g.emit("beq a0, zero, %s", endLabel)

	case ctx.BREAK() != nil:
		if len(g.breakLabels) > 0 {
			endLabel := g.breakLabels[len(g.breakLabels)-1]
			g.emit("j %s", endLabel)
		}

	case ctx.RETURN() != nil:
		if ctx.Expr() != nil {
			g.evalExpr(ctx.Expr())
		} else {
			g.emitLi("a0", 0)
		}
		if g.isFirst {
			g.emitExit()
		} else {
			g.emit("lw ra, %d(sp)", g.frameSize-4)
			g.emit("addi sp, sp, %d", g.frameSize)
			g.emit("jalr zero, ra, 0")
		}
	}
}
