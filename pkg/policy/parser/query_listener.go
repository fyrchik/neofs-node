// Code generated from Query.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QueryListener is a complete listener for a parse tree produced by QueryParser.
type QueryListener interface {
	antlr.ParseTreeListener

	// EnterPolicy is called when entering the policy production.
	EnterPolicy(c *PolicyContext)

	// EnterRepStmt is called when entering the repStmt production.
	EnterRepStmt(c *RepStmtContext)

	// EnterCbtStmt is called when entering the cbtStmt production.
	EnterCbtStmt(c *CbtStmtContext)

	// EnterSelectStmt is called when entering the selectStmt production.
	EnterSelectStmt(c *SelectStmtContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterFilterExpr is called when entering the filterExpr production.
	EnterFilterExpr(c *FilterExprContext)

	// EnterFilterStmt is called when entering the filterStmt production.
	EnterFilterStmt(c *FilterStmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterIdentOrValue is called when entering the identOrValue production.
	EnterIdentOrValue(c *IdentOrValueContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterIdentWC is called when entering the identWC production.
	EnterIdentWC(c *IdentWCContext)

	// ExitPolicy is called when exiting the policy production.
	ExitPolicy(c *PolicyContext)

	// ExitRepStmt is called when exiting the repStmt production.
	ExitRepStmt(c *RepStmtContext)

	// ExitCbtStmt is called when exiting the cbtStmt production.
	ExitCbtStmt(c *CbtStmtContext)

	// ExitSelectStmt is called when exiting the selectStmt production.
	ExitSelectStmt(c *SelectStmtContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitFilterExpr is called when exiting the filterExpr production.
	ExitFilterExpr(c *FilterExprContext)

	// ExitFilterStmt is called when exiting the filterStmt production.
	ExitFilterStmt(c *FilterStmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitIdentOrValue is called when exiting the identOrValue production.
	ExitIdentOrValue(c *IdentOrValueContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitIdentWC is called when exiting the identWC production.
	ExitIdentWC(c *IdentWCContext)
}
