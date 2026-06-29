package codegen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eramoss/b/internal/parser"
	"github.com/eramoss/b/internal/sema"
)

func (g *Generator) evalExpr(ctx parser.IExprContext) {
	ec := ctx.(*parser.ExprContext)
	n := len(ec.AllExpr())

	switch {
	case ec.DECIMAL() != nil:
		val, _ := strconv.Atoi(ec.GetText())
		g.emitLi("a0", val)

	case ec.OCTAL() != nil:
		val, _ := strconv.ParseInt(ec.GetText(), 8, 64)
		g.emitLi("a0", int(val))

	case ec.STRING() != nil:

	case ec.ID() != nil && n == 0:
		g.emitLoadVar(ec.ID().GetText())

	case ec.Expr_list() != nil && ec.Expr(0).ID() != nil:
		name := ec.Expr(0).ID().GetText()
		sym := g.st.Lookup(name)
		if sym != nil && sym.Kind == sema.KindExtrn {
			g.handleExtrnCall(ec)
			return
		}
		if sym != nil && sym.Kind == sema.KindFunc {
			g.handleFuncCall(name, ec)
			return
		}
		g.evalExpr(ec.Expr(0))

	case ec.Assign_op() != nil:
		g.handleAssign(ec)

	case n == 2 && ec.GetToken(5, 0) != nil:
		g.handleArraySubscript(ec)

	case n == 1 && strings.HasPrefix(ec.GetText(), "("):
		g.evalExpr(ec.Expr(0))

	case n == 1 && strings.HasPrefix(ec.GetText(), "*"):
		g.evalExpr(ec.Expr(0))
		g.emit("lw a0, 0(a0)")

	case n == 1 && strings.HasPrefix(ec.GetText(), "&"):
		g.handleAddrOf(ec)

	case n == 1 && strings.HasPrefix(ec.GetText(), "-"):
		g.evalExpr(ec.Expr(0))
		g.emit("sub a0, zero, a0")

	case n == 1 && strings.HasPrefix(ec.GetText(), "!"):
		g.evalExpr(ec.Expr(0))
		g.emit("sltiu a0, a0, 1")

	case n == 1 && strings.HasPrefix(ec.GetText(), "~"):
		g.evalExpr(ec.Expr(0))
		g.emit("xori a0, a0, -1")

	case n == 2:
		g.handleBinaryOp(ec)
	}
}

func (g *Generator) handleExtrnCall(ec *parser.ExprContext) {
	name := ec.Expr(0).ID().GetText()
	args := ec.Expr_list().AllExpr()

	switch name {
	case "print":
		if len(args) >= 1 {
			g.evalExpr(args[0])
			g.emit("addi a7, zero, 1")
			g.emit("ecall")
			g.emit("la a0, __nl__")
			g.emit("addi a7, zero, 4")
			g.emit("ecall")
		}
	case "read":
		if len(args) >= 1 {
			g.emit("addi a7, zero, 5")
			g.emit("ecall")
			arg := args[0]
			if strings.Contains(arg.GetText(), "[") {
				g.pushA0()
				g.evalArrayAddr(arg)
				g.popA0()
				g.emit("sw a0, 0(t0)")
			} else if arg.ID() != nil {
				g.emitStoreVar(arg.ID().GetText())
			}
		}
	}
}

func (g *Generator) handleFuncCall(name string, ec *parser.ExprContext) {
	args := []parser.IExprContext{}
	if ec.Expr_list() != nil {
		args = ec.Expr_list().AllExpr()
	}

	for _, arg := range args {
		g.evalExpr(arg)
		g.pushA0()
	}

	n := len(args)
	for i := n - 1; i >= 0; i-- {
		g.popT0()
		reg := fmt.Sprintf("a%d", i)
		g.emit("addi %s, t0, 0", reg)
	}

	g.emit("jal ra, %s", mangle(name))
}

func (g *Generator) handleAssign(ec *parser.ExprContext) {
	lhs := ec.Expr(0)
	rhs := ec.Expr(1)

	isCompound := false
	op := ""
	if ec.Assign_op() != nil {
		op = ec.Assign_op().GetText()
		isCompound = op != "="
	}

	if isCompound {
		if strings.Contains(lhs.GetText(), "[") {
			g.evalArrayAddr(lhs)
			g.emit("lw a0, 0(t0)")
			g.pushA0()
			g.evalExpr(rhs)
			g.popT0()
			g.applyCompoundOp(op)
			g.popT0()
			g.emit("sw a0, 0(t0)")
			return
		}

		if strings.HasPrefix(lhs.GetText(), "*") {
			g.evalExpr(lhs.Expr(0))
			g.emit("lw a0, 0(a0)")
			g.pushA0()
			g.evalExpr(rhs)
			g.popT0()
			g.applyCompoundOp(op)
			g.popT0()
			g.emit("sw a0, 0(t0)")
			return
		}

		if lhs.ID() != nil {
			g.emitLoadVar(lhs.ID().GetText())
			g.pushA0()
			g.evalExpr(rhs)
			g.popT0()
			g.applyCompoundOp(op)
			g.emitStoreVar(lhs.ID().GetText())
			return
		}
	}

	if strings.Contains(lhs.GetText(), "[") {
		g.evalArrayAddr(lhs)
		g.pushT0()
		g.evalExpr(rhs)
		g.popT0()
		g.emit("sw a0, 0(t0)")
		return
	}

	if strings.HasPrefix(lhs.GetText(), "*") {
		g.evalExpr(lhs.Expr(0))
		g.pushA0()
		g.evalExpr(rhs)
		g.popT0()
		g.emit("sw a0, 0(t0)")
		return
	}

	g.evalExpr(rhs)
	if lhs.ID() != nil {
		g.emitStoreVar(lhs.ID().GetText())
	}
}

func (g *Generator) applyCompoundOp(op string) {
	switch op {
	case "=+":
		g.emit("add a0, t0, a0")
	case "=-":
		g.emit("sub a0, t0, a0")
	case "=*":
		g.emit("mul a0, t0, a0")
	case "=/":
		g.emit("div a0, t0, a0")
	case "=%":
		g.emit("rem a0, t0, a0")
	case "=&":
		g.emit("and a0, t0, a0")
	case "=|":
		g.emit("or a0, t0, a0")
	case "=^":
		g.emit("xor a0, t0, a0")
	case "=<<":
		g.emit("sll a0, t0, a0")
	case "=>>":
		g.emit("srl a0, t0, a0")
	}
}

func (g *Generator) evalArrayAddr(ctx parser.IExprContext) {
	ec := ctx.(*parser.ExprContext)
	base := ec.Expr(0)
	index := ec.Expr(1)

	if base.ID() != nil {
		g.emitGetAddr(base.ID().GetText())
	} else {
		g.evalExpr(base)
		g.emit("addi t0, a0, 0")
	}
	g.pushT0()
	g.evalExpr(index)
	g.popT0()
	g.emit("slli a0, a0, 2")
	g.emit("add t0, t0, a0")
}

func (g *Generator) handleArraySubscript(ec *parser.ExprContext) {
	g.evalArrayAddr(ec)
	g.emit("lw a0, 0(t0)")
}

func (g *Generator) handleAddrOf(ec *parser.ExprContext) {
	operand := ec.Expr(0)
	if operand.ID() != nil {
		g.emitGetAddr(operand.ID().GetText())
		g.emit("addi a0, t0, 0")
	}
}

func (g *Generator) handleBinaryOp(ec *parser.ExprContext) {
	g.evalExpr(ec.Expr(0))
	g.pushA0()
	g.evalExpr(ec.Expr(1))
	g.popT0()

	switch {
	case ec.GetToken(17, 0) != nil:
		g.emit("add a0, t0, a0")
	case ec.GetToken(12, 0) != nil:
		g.emit("sub a0, t0, a0")
	case ec.GetToken(10, 0) != nil:
		g.emit("mul a0, t0, a0")
	case ec.GetToken(15, 0) != nil:
		g.emit("div a0, t0, a0")
	case ec.GetToken(16, 0) != nil:
		g.emit("rem a0, t0, a0")
	case ec.SHL() != nil:
		g.emit("sll a0, t0, a0")
	case ec.SHR() != nil:
		g.emit("srl a0, t0, a0")
	case ec.GetToken(11, 0) != nil:
		g.emit("and a0, t0, a0")
	case ec.GetToken(20, 0) != nil:
		g.emit("xor a0, t0, a0")
	case ec.GetToken(21, 0) != nil:
		g.emit("or a0, t0, a0")
	case ec.GetToken(18, 0) != nil:
		g.emit("slt a0, t0, a0")
	case ec.GetToken(19, 0) != nil:
		g.emit("slt a0, a0, t0")
	case ec.LE() != nil:
		g.emit("slt a0, a0, t0")
		g.emit("xori a0, a0, 1")
	case ec.GE() != nil:
		g.emit("slt a0, t0, a0")
		g.emit("xori a0, a0, 1")
	case ec.EQ() != nil:
		g.emit("sub a0, t0, a0")
		g.emit("sltiu a0, a0, 1")
	case ec.NE() != nil:
		g.emit("sub a0, t0, a0")
		g.emit("sltu a0, zero, a0")
	}
}
