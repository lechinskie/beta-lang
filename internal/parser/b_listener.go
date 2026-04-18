// Code generated from b.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // b

import "github.com/antlr4-go/antlr/v4"

// bListener is a complete listener for a parse tree produced by bParser.
type bListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterExt_def is called when entering the ext_def production.
	EnterExt_def(c *Ext_defContext)

	// EnterIval is called when entering the ival production.
	EnterIval(c *IvalContext)

	// EnterIval_list is called when entering the ival_list production.
	EnterIval_list(c *Ival_listContext)

	// EnterArg_list is called when entering the arg_list production.
	EnterArg_list(c *Arg_listContext)

	// EnterName_list is called when entering the name_list production.
	EnterName_list(c *Name_listContext)

	// EnterString_list is called when entering the string_list production.
	EnterString_list(c *String_listContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterCompound_stmt is called when entering the compound_stmt production.
	EnterCompound_stmt(c *Compound_stmtContext)

	// EnterAuto_decl is called when entering the auto_decl production.
	EnterAuto_decl(c *Auto_declContext)

	// EnterAuto_def is called when entering the auto_def production.
	EnterAuto_def(c *Auto_defContext)

	// EnterExtrn_decl is called when entering the extrn_decl production.
	EnterExtrn_decl(c *Extrn_declContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr_list is called when entering the expr_list production.
	EnterExpr_list(c *Expr_listContext)

	// EnterAssign_op is called when entering the assign_op production.
	EnterAssign_op(c *Assign_opContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitExt_def is called when exiting the ext_def production.
	ExitExt_def(c *Ext_defContext)

	// ExitIval is called when exiting the ival production.
	ExitIval(c *IvalContext)

	// ExitIval_list is called when exiting the ival_list production.
	ExitIval_list(c *Ival_listContext)

	// ExitArg_list is called when exiting the arg_list production.
	ExitArg_list(c *Arg_listContext)

	// ExitName_list is called when exiting the name_list production.
	ExitName_list(c *Name_listContext)

	// ExitString_list is called when exiting the string_list production.
	ExitString_list(c *String_listContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitCompound_stmt is called when exiting the compound_stmt production.
	ExitCompound_stmt(c *Compound_stmtContext)

	// ExitAuto_decl is called when exiting the auto_decl production.
	ExitAuto_decl(c *Auto_declContext)

	// ExitAuto_def is called when exiting the auto_def production.
	ExitAuto_def(c *Auto_defContext)

	// ExitExtrn_decl is called when exiting the extrn_decl production.
	ExitExtrn_decl(c *Extrn_declContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr_list is called when exiting the expr_list production.
	ExitExpr_list(c *Expr_listContext)

	// ExitAssign_op is called when exiting the assign_op production.
	ExitAssign_op(c *Assign_opContext)
}
