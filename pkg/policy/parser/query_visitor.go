// Code generated from Query.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QueryParser.
type QueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryParser#policy.
	VisitPolicy(ctx *PolicyContext) interface{}

	// Visit a parse tree produced by QueryParser#repStmt.
	VisitRepStmt(ctx *RepStmtContext) interface{}

	// Visit a parse tree produced by QueryParser#cbtStmt.
	VisitCbtStmt(ctx *CbtStmtContext) interface{}

	// Visit a parse tree produced by QueryParser#selectStmt.
	VisitSelectStmt(ctx *SelectStmtContext) interface{}

	// Visit a parse tree produced by QueryParser#clause.
	VisitClause(ctx *ClauseContext) interface{}

	// Visit a parse tree produced by QueryParser#filterExpr.
	VisitFilterExpr(ctx *FilterExprContext) interface{}

	// Visit a parse tree produced by QueryParser#filterStmt.
	VisitFilterStmt(ctx *FilterStmtContext) interface{}

	// Visit a parse tree produced by QueryParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by QueryParser#operation.
	VisitOperation(ctx *OperationContext) interface{}

	// Visit a parse tree produced by QueryParser#identOrValue.
	VisitIdentOrValue(ctx *IdentOrValueContext) interface{}

	// Visit a parse tree produced by QueryParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by QueryParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by QueryParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by QueryParser#identWC.
	VisitIdentWC(ctx *IdentWCContext) interface{}
}
