// Code generated from Query.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQueryListener is a complete listener for a parse tree produced by QueryParser.
type BaseQueryListener struct{}

var _ QueryListener = &BaseQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPolicy is called when production policy is entered.
func (s *BaseQueryListener) EnterPolicy(ctx *PolicyContext) {}

// ExitPolicy is called when production policy is exited.
func (s *BaseQueryListener) ExitPolicy(ctx *PolicyContext) {}

// EnterRepStmt is called when production repStmt is entered.
func (s *BaseQueryListener) EnterRepStmt(ctx *RepStmtContext) {}

// ExitRepStmt is called when production repStmt is exited.
func (s *BaseQueryListener) ExitRepStmt(ctx *RepStmtContext) {}

// EnterCbtStmt is called when production cbtStmt is entered.
func (s *BaseQueryListener) EnterCbtStmt(ctx *CbtStmtContext) {}

// ExitCbtStmt is called when production cbtStmt is exited.
func (s *BaseQueryListener) ExitCbtStmt(ctx *CbtStmtContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseQueryListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseQueryListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterClause is called when production clause is entered.
func (s *BaseQueryListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BaseQueryListener) ExitClause(ctx *ClauseContext) {}

// EnterFilterExpr is called when production filterExpr is entered.
func (s *BaseQueryListener) EnterFilterExpr(ctx *FilterExprContext) {}

// ExitFilterExpr is called when production filterExpr is exited.
func (s *BaseQueryListener) ExitFilterExpr(ctx *FilterExprContext) {}

// EnterFilterStmt is called when production filterStmt is entered.
func (s *BaseQueryListener) EnterFilterStmt(ctx *FilterStmtContext) {}

// ExitFilterStmt is called when production filterStmt is exited.
func (s *BaseQueryListener) ExitFilterStmt(ctx *FilterStmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseQueryListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseQueryListener) ExitExpr(ctx *ExprContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseQueryListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseQueryListener) ExitOperation(ctx *OperationContext) {}

// EnterIdentOrValue is called when production identOrValue is entered.
func (s *BaseQueryListener) EnterIdentOrValue(ctx *IdentOrValueContext) {}

// ExitIdentOrValue is called when production identOrValue is exited.
func (s *BaseQueryListener) ExitIdentOrValue(ctx *IdentOrValueContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseQueryListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseQueryListener) ExitNumber(ctx *NumberContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseQueryListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseQueryListener) ExitKeyword(ctx *KeywordContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseQueryListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseQueryListener) ExitIdent(ctx *IdentContext) {}

// EnterIdentWC is called when production identWC is entered.
func (s *BaseQueryListener) EnterIdentWC(ctx *IdentWCContext) {}

// ExitIdentWC is called when production identWC is exited.
func (s *BaseQueryListener) ExitIdentWC(ctx *IdentWCContext) {}
