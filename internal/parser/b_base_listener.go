// Code generated from b.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // b

import "github.com/antlr4-go/antlr/v4"

// BasebListener is a complete listener for a parse tree produced by bParser.
type BasebListener struct{}

var _ bListener = &BasebListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasebListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasebListener) ExitProgram(ctx *ProgramContext) {}

// EnterExt_def is called when production ext_def is entered.
func (s *BasebListener) EnterExt_def(ctx *Ext_defContext) {}

// ExitExt_def is called when production ext_def is exited.
func (s *BasebListener) ExitExt_def(ctx *Ext_defContext) {}

// EnterIval is called when production ival is entered.
func (s *BasebListener) EnterIval(ctx *IvalContext) {}

// ExitIval is called when production ival is exited.
func (s *BasebListener) ExitIval(ctx *IvalContext) {}

// EnterIval_list is called when production ival_list is entered.
func (s *BasebListener) EnterIval_list(ctx *Ival_listContext) {}

// ExitIval_list is called when production ival_list is exited.
func (s *BasebListener) ExitIval_list(ctx *Ival_listContext) {}

// EnterArg_list is called when production arg_list is entered.
func (s *BasebListener) EnterArg_list(ctx *Arg_listContext) {}

// ExitArg_list is called when production arg_list is exited.
func (s *BasebListener) ExitArg_list(ctx *Arg_listContext) {}

// EnterName_list is called when production name_list is entered.
func (s *BasebListener) EnterName_list(ctx *Name_listContext) {}

// ExitName_list is called when production name_list is exited.
func (s *BasebListener) ExitName_list(ctx *Name_listContext) {}

// EnterString_list is called when production string_list is entered.
func (s *BasebListener) EnterString_list(ctx *String_listContext) {}

// ExitString_list is called when production string_list is exited.
func (s *BasebListener) ExitString_list(ctx *String_listContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasebListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasebListener) ExitStatement(ctx *StatementContext) {}

// EnterCompound_stmt is called when production compound_stmt is entered.
func (s *BasebListener) EnterCompound_stmt(ctx *Compound_stmtContext) {}

// ExitCompound_stmt is called when production compound_stmt is exited.
func (s *BasebListener) ExitCompound_stmt(ctx *Compound_stmtContext) {}

// EnterAuto_decl is called when production auto_decl is entered.
func (s *BasebListener) EnterAuto_decl(ctx *Auto_declContext) {}

// ExitAuto_decl is called when production auto_decl is exited.
func (s *BasebListener) ExitAuto_decl(ctx *Auto_declContext) {}

// EnterAuto_def is called when production auto_def is entered.
func (s *BasebListener) EnterAuto_def(ctx *Auto_defContext) {}

// ExitAuto_def is called when production auto_def is exited.
func (s *BasebListener) ExitAuto_def(ctx *Auto_defContext) {}

// EnterExtrn_decl is called when production extrn_decl is entered.
func (s *BasebListener) EnterExtrn_decl(ctx *Extrn_declContext) {}

// ExitExtrn_decl is called when production extrn_decl is exited.
func (s *BasebListener) ExitExtrn_decl(ctx *Extrn_declContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasebListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasebListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BasebListener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BasebListener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *BasebListener) EnterAssign_op(ctx *Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *BasebListener) ExitAssign_op(ctx *Assign_opContext) {}
