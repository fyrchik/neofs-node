// Code generated from Query.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Query

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQueryVisitor) VisitPolicy(ctx *PolicyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitRepStmt(ctx *RepStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitCbtStmt(ctx *CbtStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitSelectStmt(ctx *SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitClause(ctx *ClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitFilterExpr(ctx *FilterExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitFilterStmt(ctx *FilterStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitOperation(ctx *OperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitIdentOrValue(ctx *IdentOrValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryVisitor) VisitIdentWC(ctx *IdentWCContext) interface{} {
	return v.VisitChildren(ctx)
}
