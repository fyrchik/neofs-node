// Code generated from Query.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Query

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 130,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 7, 2, 33, 10, 2, 12, 2,
	14, 2, 36, 11, 2, 3, 2, 5, 2, 39, 10, 2, 3, 2, 7, 2, 42, 10, 2, 12, 2,
	14, 2, 45, 11, 2, 3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 57, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 5, 5, 66, 10, 5, 3, 5, 5, 5, 69, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5,
	5, 75, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 85,
	10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 93, 10, 7, 12, 7, 14,
	7, 96, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 109, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 116,
	10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 124, 10, 14, 3,
	15, 3, 15, 5, 15, 128, 10, 15, 3, 15, 2, 3, 12, 16, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 2, 6, 3, 2, 3, 4, 3, 2, 10, 15, 3, 2, 25,
	26, 4, 2, 16, 18, 20, 22, 2, 131, 2, 30, 3, 2, 2, 2, 4, 52, 3, 2, 2, 2,
	6, 58, 3, 2, 2, 2, 8, 61, 3, 2, 2, 2, 10, 76, 3, 2, 2, 2, 12, 84, 3, 2,
	2, 2, 14, 97, 3, 2, 2, 2, 16, 108, 3, 2, 2, 2, 18, 110, 3, 2, 2, 2, 20,
	115, 3, 2, 2, 2, 22, 117, 3, 2, 2, 2, 24, 119, 3, 2, 2, 2, 26, 123, 3,
	2, 2, 2, 28, 127, 3, 2, 2, 2, 30, 34, 5, 4, 3, 2, 31, 33, 5, 4, 3, 2, 32,
	31, 3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2,
	2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 39, 5, 6, 4, 2, 38, 37,
	3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 43, 3, 2, 2, 2, 40, 42, 5, 8, 5, 2,
	41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3,
	2, 2, 2, 44, 49, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 48, 5, 14, 8, 2, 47,
	46, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2,
	2, 50, 3, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 7, 16, 2, 2, 53, 56,
	7, 25, 2, 2, 54, 55, 7, 17, 2, 2, 55, 57, 5, 26, 14, 2, 56, 54, 3, 2, 2,
	2, 56, 57, 3, 2, 2, 2, 57, 5, 3, 2, 2, 2, 58, 59, 7, 19, 2, 2, 59, 60,
	7, 25, 2, 2, 60, 7, 3, 2, 2, 2, 61, 62, 7, 20, 2, 2, 62, 68, 7, 25, 2,
	2, 63, 65, 7, 17, 2, 2, 64, 66, 5, 10, 6, 2, 65, 64, 3, 2, 2, 2, 65, 66,
	3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 5, 26, 14, 2, 68, 63, 3, 2, 2,
	2, 68, 69, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 71, 7, 21, 2, 2, 71, 74,
	5, 28, 15, 2, 72, 73, 7, 18, 2, 2, 73, 75, 5, 26, 14, 2, 74, 72, 3, 2,
	2, 2, 74, 75, 3, 2, 2, 2, 75, 9, 3, 2, 2, 2, 76, 77, 9, 2, 2, 2, 77, 11,
	3, 2, 2, 2, 78, 79, 8, 7, 1, 2, 79, 85, 5, 16, 9, 2, 80, 81, 7, 7, 2, 2,
	81, 82, 5, 12, 7, 2, 82, 83, 7, 8, 2, 2, 83, 85, 3, 2, 2, 2, 84, 78, 3,
	2, 2, 2, 84, 80, 3, 2, 2, 2, 85, 94, 3, 2, 2, 2, 86, 87, 12, 6, 2, 2, 87,
	88, 7, 5, 2, 2, 88, 93, 5, 12, 7, 7, 89, 90, 12, 5, 2, 2, 90, 91, 7, 6,
	2, 2, 91, 93, 5, 12, 7, 6, 92, 86, 3, 2, 2, 2, 92, 89, 3, 2, 2, 2, 93,
	96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 13, 3, 2, 2,
	2, 96, 94, 3, 2, 2, 2, 97, 98, 7, 22, 2, 2, 98, 99, 5, 12, 7, 2, 99, 100,
	7, 18, 2, 2, 100, 101, 5, 26, 14, 2, 101, 15, 3, 2, 2, 2, 102, 103, 7,
	9, 2, 2, 103, 109, 5, 26, 14, 2, 104, 105, 5, 26, 14, 2, 105, 106, 5, 18,
	10, 2, 106, 107, 5, 20, 11, 2, 107, 109, 3, 2, 2, 2, 108, 102, 3, 2, 2,
	2, 108, 104, 3, 2, 2, 2, 109, 17, 3, 2, 2, 2, 110, 111, 9, 3, 2, 2, 111,
	19, 3, 2, 2, 2, 112, 116, 5, 26, 14, 2, 113, 116, 5, 22, 12, 2, 114, 116,
	7, 27, 2, 2, 115, 112, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 114, 3, 2,
	2, 2, 116, 21, 3, 2, 2, 2, 117, 118, 9, 4, 2, 2, 118, 23, 3, 2, 2, 2, 119,
	120, 9, 5, 2, 2, 120, 25, 3, 2, 2, 2, 121, 124, 5, 24, 13, 2, 122, 124,
	7, 24, 2, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 27, 3, 2,
	2, 2, 125, 128, 5, 26, 14, 2, 126, 128, 7, 23, 2, 2, 127, 125, 3, 2, 2,
	2, 127, 126, 3, 2, 2, 2, 128, 29, 3, 2, 2, 2, 17, 34, 38, 43, 49, 56, 65,
	68, 74, 84, 92, 94, 108, 115, 123, 127,
}
var literalNames = []string{
	"", "'SAME'", "'DISTINCT'", "'AND'", "'OR'", "'('", "')'", "'@'", "'EQ'",
	"'NE'", "'GE'", "'GT'", "'LT'", "'LE'", "'REP'", "'IN'", "'AS'", "'CBF'",
	"'SELECT'", "'FROM'", "'FILTER'", "'*'", "", "", "'0'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "REP", "IN", "AS",
	"CBF", "SELECT", "FROM", "FILTER", "WILDCARD", "IDENT", "NUMBER1", "ZERO",
	"STRING", "WS",
}

var ruleNames = []string{
	"policy", "repStmt", "cbtStmt", "selectStmt", "clause", "filterExpr", "filterStmt",
	"expr", "operation", "identOrValue", "number", "keyword", "ident", "identWC",
}

type QueryParser struct {
	*antlr.BaseParser
}

// NewQueryParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *QueryParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewQueryParser(input antlr.TokenStream) *QueryParser {
	this := new(QueryParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Query.g4"

	return this
}

// QueryParser tokens.
const (
	QueryParserEOF      = antlr.TokenEOF
	QueryParserT__0     = 1
	QueryParserT__1     = 2
	QueryParserT__2     = 3
	QueryParserT__3     = 4
	QueryParserT__4     = 5
	QueryParserT__5     = 6
	QueryParserT__6     = 7
	QueryParserT__7     = 8
	QueryParserT__8     = 9
	QueryParserT__9     = 10
	QueryParserT__10    = 11
	QueryParserT__11    = 12
	QueryParserT__12    = 13
	QueryParserREP      = 14
	QueryParserIN       = 15
	QueryParserAS       = 16
	QueryParserCBF      = 17
	QueryParserSELECT   = 18
	QueryParserFROM     = 19
	QueryParserFILTER   = 20
	QueryParserWILDCARD = 21
	QueryParserIDENT    = 22
	QueryParserNUMBER1  = 23
	QueryParserZERO     = 24
	QueryParserSTRING   = 25
	QueryParserWS       = 26
)

// QueryParser rules.
const (
	QueryParserRULE_policy       = 0
	QueryParserRULE_repStmt      = 1
	QueryParserRULE_cbtStmt      = 2
	QueryParserRULE_selectStmt   = 3
	QueryParserRULE_clause       = 4
	QueryParserRULE_filterExpr   = 5
	QueryParserRULE_filterStmt   = 6
	QueryParserRULE_expr         = 7
	QueryParserRULE_operation    = 8
	QueryParserRULE_identOrValue = 9
	QueryParserRULE_number       = 10
	QueryParserRULE_keyword      = 11
	QueryParserRULE_ident        = 12
	QueryParserRULE_identWC      = 13
)

// IPolicyContext is an interface to support dynamic dispatch.
type IPolicyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolicyContext differentiates from other interfaces.
	IsPolicyContext()
}

type PolicyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolicyContext() *PolicyContext {
	var p = new(PolicyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_policy
	return p
}

func (*PolicyContext) IsPolicyContext() {}

func NewPolicyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolicyContext {
	var p = new(PolicyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_policy

	return p
}

func (s *PolicyContext) GetParser() antlr.Parser { return s.parser }

func (s *PolicyContext) AllRepStmt() []IRepStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRepStmtContext)(nil)).Elem())
	var tst = make([]IRepStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRepStmtContext)
		}
	}

	return tst
}

func (s *PolicyContext) RepStmt(i int) IRepStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRepStmtContext)
}

func (s *PolicyContext) CbtStmt() ICbtStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICbtStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICbtStmtContext)
}

func (s *PolicyContext) AllSelectStmt() []ISelectStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectStmtContext)(nil)).Elem())
	var tst = make([]ISelectStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectStmtContext)
		}
	}

	return tst
}

func (s *PolicyContext) SelectStmt(i int) ISelectStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectStmtContext)
}

func (s *PolicyContext) AllFilterStmt() []IFilterStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterStmtContext)(nil)).Elem())
	var tst = make([]IFilterStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterStmtContext)
		}
	}

	return tst
}

func (s *PolicyContext) FilterStmt(i int) IFilterStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterStmtContext)
}

func (s *PolicyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolicyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolicyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterPolicy(s)
	}
}

func (s *PolicyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitPolicy(s)
	}
}

func (s *PolicyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitPolicy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Policy() (localctx IPolicyContext) {
	localctx = NewPolicyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryParserRULE_policy)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.RepStmt()
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QueryParserREP {
		{
			p.SetState(29)
			p.RepStmt()
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QueryParserCBF {
		{
			p.SetState(35)
			p.CbtStmt()
		}

	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QueryParserSELECT {
		{
			p.SetState(38)
			p.SelectStmt()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QueryParserFILTER {
		{
			p.SetState(44)
			p.FilterStmt()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRepStmtContext is an interface to support dynamic dispatch.
type IRepStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCount returns the Count token.
	GetCount() antlr.Token

	// SetCount sets the Count token.
	SetCount(antlr.Token)

	// GetSelector returns the Selector rule contexts.
	GetSelector() IIdentContext

	// SetSelector sets the Selector rule contexts.
	SetSelector(IIdentContext)

	// IsRepStmtContext differentiates from other interfaces.
	IsRepStmtContext()
}

type RepStmtContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Count    antlr.Token
	Selector IIdentContext
}

func NewEmptyRepStmtContext() *RepStmtContext {
	var p = new(RepStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_repStmt
	return p
}

func (*RepStmtContext) IsRepStmtContext() {}

func NewRepStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepStmtContext {
	var p = new(RepStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_repStmt

	return p
}

func (s *RepStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RepStmtContext) GetCount() antlr.Token { return s.Count }

func (s *RepStmtContext) SetCount(v antlr.Token) { s.Count = v }

func (s *RepStmtContext) GetSelector() IIdentContext { return s.Selector }

func (s *RepStmtContext) SetSelector(v IIdentContext) { s.Selector = v }

func (s *RepStmtContext) REP() antlr.TerminalNode {
	return s.GetToken(QueryParserREP, 0)
}

func (s *RepStmtContext) NUMBER1() antlr.TerminalNode {
	return s.GetToken(QueryParserNUMBER1, 0)
}

func (s *RepStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(QueryParserIN, 0)
}

func (s *RepStmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *RepStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterRepStmt(s)
	}
}

func (s *RepStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitRepStmt(s)
	}
}

func (s *RepStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitRepStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) RepStmt() (localctx IRepStmtContext) {
	localctx = NewRepStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QueryParserRULE_repStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(QueryParserREP)
	}
	{
		p.SetState(51)

		var _m = p.Match(QueryParserNUMBER1)

		localctx.(*RepStmtContext).Count = _m
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QueryParserIN {
		{
			p.SetState(52)
			p.Match(QueryParserIN)
		}
		{
			p.SetState(53)

			var _x = p.Ident()

			localctx.(*RepStmtContext).Selector = _x
		}

	}

	return localctx
}

// ICbtStmtContext is an interface to support dynamic dispatch.
type ICbtStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFactor returns the Factor token.
	GetFactor() antlr.Token

	// SetFactor sets the Factor token.
	SetFactor(antlr.Token)

	// IsCbtStmtContext differentiates from other interfaces.
	IsCbtStmtContext()
}

type CbtStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Factor antlr.Token
}

func NewEmptyCbtStmtContext() *CbtStmtContext {
	var p = new(CbtStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_cbtStmt
	return p
}

func (*CbtStmtContext) IsCbtStmtContext() {}

func NewCbtStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CbtStmtContext {
	var p = new(CbtStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_cbtStmt

	return p
}

func (s *CbtStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CbtStmtContext) GetFactor() antlr.Token { return s.Factor }

func (s *CbtStmtContext) SetFactor(v antlr.Token) { s.Factor = v }

func (s *CbtStmtContext) CBF() antlr.TerminalNode {
	return s.GetToken(QueryParserCBF, 0)
}

func (s *CbtStmtContext) NUMBER1() antlr.TerminalNode {
	return s.GetToken(QueryParserNUMBER1, 0)
}

func (s *CbtStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CbtStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CbtStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterCbtStmt(s)
	}
}

func (s *CbtStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitCbtStmt(s)
	}
}

func (s *CbtStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitCbtStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) CbtStmt() (localctx ICbtStmtContext) {
	localctx = NewCbtStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryParserRULE_cbtStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(QueryParserCBF)
	}
	{
		p.SetState(57)

		var _m = p.Match(QueryParserNUMBER1)

		localctx.(*CbtStmtContext).Factor = _m
	}

	return localctx
}

// ISelectStmtContext is an interface to support dynamic dispatch.
type ISelectStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCount returns the Count token.
	GetCount() antlr.Token

	// SetCount sets the Count token.
	SetCount(antlr.Token)

	// GetBucket returns the Bucket rule contexts.
	GetBucket() IIdentContext

	// GetFilter returns the Filter rule contexts.
	GetFilter() IIdentWCContext

	// GetName returns the Name rule contexts.
	GetName() IIdentContext

	// SetBucket sets the Bucket rule contexts.
	SetBucket(IIdentContext)

	// SetFilter sets the Filter rule contexts.
	SetFilter(IIdentWCContext)

	// SetName sets the Name rule contexts.
	SetName(IIdentContext)

	// IsSelectStmtContext differentiates from other interfaces.
	IsSelectStmtContext()
}

type SelectStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Count  antlr.Token
	Bucket IIdentContext
	Filter IIdentWCContext
	Name   IIdentContext
}

func NewEmptySelectStmtContext() *SelectStmtContext {
	var p = new(SelectStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_selectStmt
	return p
}

func (*SelectStmtContext) IsSelectStmtContext() {}

func NewSelectStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStmtContext {
	var p = new(SelectStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_selectStmt

	return p
}

func (s *SelectStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStmtContext) GetCount() antlr.Token { return s.Count }

func (s *SelectStmtContext) SetCount(v antlr.Token) { s.Count = v }

func (s *SelectStmtContext) GetBucket() IIdentContext { return s.Bucket }

func (s *SelectStmtContext) GetFilter() IIdentWCContext { return s.Filter }

func (s *SelectStmtContext) GetName() IIdentContext { return s.Name }

func (s *SelectStmtContext) SetBucket(v IIdentContext) { s.Bucket = v }

func (s *SelectStmtContext) SetFilter(v IIdentWCContext) { s.Filter = v }

func (s *SelectStmtContext) SetName(v IIdentContext) { s.Name = v }

func (s *SelectStmtContext) SELECT() antlr.TerminalNode {
	return s.GetToken(QueryParserSELECT, 0)
}

func (s *SelectStmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(QueryParserFROM, 0)
}

func (s *SelectStmtContext) NUMBER1() antlr.TerminalNode {
	return s.GetToken(QueryParserNUMBER1, 0)
}

func (s *SelectStmtContext) IdentWC() IIdentWCContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentWCContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentWCContext)
}

func (s *SelectStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(QueryParserIN, 0)
}

func (s *SelectStmtContext) AS() antlr.TerminalNode {
	return s.GetToken(QueryParserAS, 0)
}

func (s *SelectStmtContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *SelectStmtContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *SelectStmtContext) Clause() IClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *SelectStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterSelectStmt(s)
	}
}

func (s *SelectStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitSelectStmt(s)
	}
}

func (s *SelectStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitSelectStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) SelectStmt() (localctx ISelectStmtContext) {
	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryParserRULE_selectStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(QueryParserSELECT)
	}
	{
		p.SetState(60)

		var _m = p.Match(QueryParserNUMBER1)

		localctx.(*SelectStmtContext).Count = _m
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QueryParserIN {
		{
			p.SetState(61)
			p.Match(QueryParserIN)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == QueryParserT__0 || _la == QueryParserT__1 {
			{
				p.SetState(62)
				p.Clause()
			}

		}
		{
			p.SetState(65)

			var _x = p.Ident()

			localctx.(*SelectStmtContext).Bucket = _x
		}

	}
	{
		p.SetState(68)
		p.Match(QueryParserFROM)
	}
	{
		p.SetState(69)

		var _x = p.IdentWC()

		localctx.(*SelectStmtContext).Filter = _x
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QueryParserAS {
		{
			p.SetState(70)
			p.Match(QueryParserAS)
		}
		{
			p.SetState(71)

			var _x = p.Ident()

			localctx.(*SelectStmtContext).Name = _x
		}

	}

	return localctx
}

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }
func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitClause(s)
	}
}

func (s *ClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Clause() (localctx IClauseContext) {
	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryParserRULE_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryParserT__0 || _la == QueryParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFilterExprContext is an interface to support dynamic dispatch.
type IFilterExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the Op token.
	GetOp() antlr.Token

	// SetOp sets the Op token.
	SetOp(antlr.Token)

	// GetF1 returns the F1 rule contexts.
	GetF1() IFilterExprContext

	// GetInner returns the Inner rule contexts.
	GetInner() IFilterExprContext

	// GetF2 returns the F2 rule contexts.
	GetF2() IFilterExprContext

	// SetF1 sets the F1 rule contexts.
	SetF1(IFilterExprContext)

	// SetInner sets the Inner rule contexts.
	SetInner(IFilterExprContext)

	// SetF2 sets the F2 rule contexts.
	SetF2(IFilterExprContext)

	// IsFilterExprContext differentiates from other interfaces.
	IsFilterExprContext()
}

type FilterExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	F1     IFilterExprContext
	Inner  IFilterExprContext
	Op     antlr.Token
	F2     IFilterExprContext
}

func NewEmptyFilterExprContext() *FilterExprContext {
	var p = new(FilterExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_filterExpr
	return p
}

func (*FilterExprContext) IsFilterExprContext() {}

func NewFilterExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterExprContext {
	var p = new(FilterExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_filterExpr

	return p
}

func (s *FilterExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterExprContext) GetOp() antlr.Token { return s.Op }

func (s *FilterExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *FilterExprContext) GetF1() IFilterExprContext { return s.F1 }

func (s *FilterExprContext) GetInner() IFilterExprContext { return s.Inner }

func (s *FilterExprContext) GetF2() IFilterExprContext { return s.F2 }

func (s *FilterExprContext) SetF1(v IFilterExprContext) { s.F1 = v }

func (s *FilterExprContext) SetInner(v IFilterExprContext) { s.Inner = v }

func (s *FilterExprContext) SetF2(v IFilterExprContext) { s.F2 = v }

func (s *FilterExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FilterExprContext) AllFilterExpr() []IFilterExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterExprContext)(nil)).Elem())
	var tst = make([]IFilterExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterExprContext)
		}
	}

	return tst
}

func (s *FilterExprContext) FilterExpr(i int) IFilterExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterExprContext)
}

func (s *FilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterFilterExpr(s)
	}
}

func (s *FilterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitFilterExpr(s)
	}
}

func (s *FilterExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitFilterExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) FilterExpr() (localctx IFilterExprContext) {
	return p.filterExpr(0)
}

func (p *QueryParser) filterExpr(_p int) (localctx IFilterExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFilterExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFilterExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, QueryParserRULE_filterExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserT__6, QueryParserREP, QueryParserIN, QueryParserAS, QueryParserSELECT, QueryParserFROM, QueryParserFILTER, QueryParserIDENT:
		{
			p.SetState(77)
			p.Expr()
		}

	case QueryParserT__4:
		{
			p.SetState(78)
			p.Match(QueryParserT__4)
		}
		{
			p.SetState(79)

			var _x = p.filterExpr(0)

			localctx.(*FilterExprContext).Inner = _x
		}
		{
			p.SetState(80)
			p.Match(QueryParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFilterExprContext(p, _parentctx, _parentState)
				localctx.(*FilterExprContext).F1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_filterExpr)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(85)

					var _m = p.Match(QueryParserT__2)

					localctx.(*FilterExprContext).Op = _m
				}
				{
					p.SetState(86)

					var _x = p.filterExpr(5)

					localctx.(*FilterExprContext).F2 = _x
				}

			case 2:
				localctx = NewFilterExprContext(p, _parentctx, _parentState)
				localctx.(*FilterExprContext).F1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, QueryParserRULE_filterExpr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(88)

					var _m = p.Match(QueryParserT__3)

					localctx.(*FilterExprContext).Op = _m
				}
				{
					p.SetState(89)

					var _x = p.filterExpr(4)

					localctx.(*FilterExprContext).F2 = _x
				}

			}

		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IFilterStmtContext is an interface to support dynamic dispatch.
type IFilterStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the Expr rule contexts.
	GetExpr() IFilterExprContext

	// GetName returns the Name rule contexts.
	GetName() IIdentContext

	// SetExpr sets the Expr rule contexts.
	SetExpr(IFilterExprContext)

	// SetName sets the Name rule contexts.
	SetName(IIdentContext)

	// IsFilterStmtContext differentiates from other interfaces.
	IsFilterStmtContext()
}

type FilterStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Expr   IFilterExprContext
	Name   IIdentContext
}

func NewEmptyFilterStmtContext() *FilterStmtContext {
	var p = new(FilterStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_filterStmt
	return p
}

func (*FilterStmtContext) IsFilterStmtContext() {}

func NewFilterStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterStmtContext {
	var p = new(FilterStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_filterStmt

	return p
}

func (s *FilterStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterStmtContext) GetExpr() IFilterExprContext { return s.Expr }

func (s *FilterStmtContext) GetName() IIdentContext { return s.Name }

func (s *FilterStmtContext) SetExpr(v IFilterExprContext) { s.Expr = v }

func (s *FilterStmtContext) SetName(v IIdentContext) { s.Name = v }

func (s *FilterStmtContext) FILTER() antlr.TerminalNode {
	return s.GetToken(QueryParserFILTER, 0)
}

func (s *FilterStmtContext) AS() antlr.TerminalNode {
	return s.GetToken(QueryParserAS, 0)
}

func (s *FilterStmtContext) FilterExpr() IFilterExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterExprContext)
}

func (s *FilterStmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FilterStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterFilterStmt(s)
	}
}

func (s *FilterStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitFilterStmt(s)
	}
}

func (s *FilterStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitFilterStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) FilterStmt() (localctx IFilterStmtContext) {
	localctx = NewFilterStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QueryParserRULE_filterStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(QueryParserFILTER)
	}
	{
		p.SetState(96)

		var _x = p.filterExpr(0)

		localctx.(*FilterStmtContext).Expr = _x
	}
	{
		p.SetState(97)
		p.Match(QueryParserAS)
	}
	{
		p.SetState(98)

		var _x = p.Ident()

		localctx.(*FilterStmtContext).Name = _x
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFilter returns the Filter rule contexts.
	GetFilter() IIdentContext

	// GetKey returns the Key rule contexts.
	GetKey() IIdentContext

	// GetOp returns the Op rule contexts.
	GetOp() IOperationContext

	// GetValue returns the Value rule contexts.
	GetValue() IIdentOrValueContext

	// SetFilter sets the Filter rule contexts.
	SetFilter(IIdentContext)

	// SetKey sets the Key rule contexts.
	SetKey(IIdentContext)

	// SetOp sets the Op rule contexts.
	SetOp(IOperationContext)

	// SetValue sets the Value rule contexts.
	SetValue(IIdentOrValueContext)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Filter IIdentContext
	Key    IIdentContext
	Op     IOperationContext
	Value  IIdentOrValueContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetFilter() IIdentContext { return s.Filter }

func (s *ExprContext) GetKey() IIdentContext { return s.Key }

func (s *ExprContext) GetOp() IOperationContext { return s.Op }

func (s *ExprContext) GetValue() IIdentOrValueContext { return s.Value }

func (s *ExprContext) SetFilter(v IIdentContext) { s.Filter = v }

func (s *ExprContext) SetKey(v IIdentContext) { s.Key = v }

func (s *ExprContext) SetOp(v IOperationContext) { s.Op = v }

func (s *ExprContext) SetValue(v IIdentOrValueContext) { s.Value = v }

func (s *ExprContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ExprContext) Operation() IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *ExprContext) IdentOrValue() IIdentOrValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentOrValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentOrValueContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QueryParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(QueryParserT__6)
		}
		{
			p.SetState(101)

			var _x = p.Ident()

			localctx.(*ExprContext).Filter = _x
		}

	case QueryParserREP, QueryParserIN, QueryParserAS, QueryParserSELECT, QueryParserFROM, QueryParserFILTER, QueryParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)

			var _x = p.Ident()

			localctx.(*ExprContext).Key = _x
		}
		{
			p.SetState(103)

			var _x = p.Operation()

			localctx.(*ExprContext).Op = _x
		}
		{
			p.SetState(104)

			var _x = p.IdentOrValue()

			localctx.(*ExprContext).Value = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_operation
	return p
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }
func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (s *OperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Operation() (localctx IOperationContext) {
	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QueryParserRULE_operation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QueryParserT__7)|(1<<QueryParserT__8)|(1<<QueryParserT__9)|(1<<QueryParserT__10)|(1<<QueryParserT__11)|(1<<QueryParserT__12))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentOrValueContext is an interface to support dynamic dispatch.
type IIdentOrValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStr returns the Str token.
	GetStr() antlr.Token

	// SetStr sets the Str token.
	SetStr(antlr.Token)

	// IsIdentOrValueContext differentiates from other interfaces.
	IsIdentOrValueContext()
}

type IdentOrValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Str    antlr.Token
}

func NewEmptyIdentOrValueContext() *IdentOrValueContext {
	var p = new(IdentOrValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_identOrValue
	return p
}

func (*IdentOrValueContext) IsIdentOrValueContext() {}

func NewIdentOrValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentOrValueContext {
	var p = new(IdentOrValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_identOrValue

	return p
}

func (s *IdentOrValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentOrValueContext) GetStr() antlr.Token { return s.Str }

func (s *IdentOrValueContext) SetStr(v antlr.Token) { s.Str = v }

func (s *IdentOrValueContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *IdentOrValueContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *IdentOrValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(QueryParserSTRING, 0)
}

func (s *IdentOrValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentOrValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentOrValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterIdentOrValue(s)
	}
}

func (s *IdentOrValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitIdentOrValue(s)
	}
}

func (s *IdentOrValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitIdentOrValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) IdentOrValue() (localctx IIdentOrValueContext) {
	localctx = NewIdentOrValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QueryParserRULE_identOrValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserREP, QueryParserIN, QueryParserAS, QueryParserSELECT, QueryParserFROM, QueryParserFILTER, QueryParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Ident()
		}

	case QueryParserNUMBER1, QueryParserZERO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Number()
		}

	case QueryParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)

			var _m = p.Match(QueryParserSTRING)

			localctx.(*IdentOrValueContext).Str = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) ZERO() antlr.TerminalNode {
	return s.GetToken(QueryParserZERO, 0)
}

func (s *NumberContext) NUMBER1() antlr.TerminalNode {
	return s.GetToken(QueryParserNUMBER1, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QueryParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QueryParserNUMBER1 || _la == QueryParserZERO) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) REP() antlr.TerminalNode {
	return s.GetToken(QueryParserREP, 0)
}

func (s *KeywordContext) IN() antlr.TerminalNode {
	return s.GetToken(QueryParserIN, 0)
}

func (s *KeywordContext) AS() antlr.TerminalNode {
	return s.GetToken(QueryParserAS, 0)
}

func (s *KeywordContext) SELECT() antlr.TerminalNode {
	return s.GetToken(QueryParserSELECT, 0)
}

func (s *KeywordContext) FROM() antlr.TerminalNode {
	return s.GetToken(QueryParserFROM, 0)
}

func (s *KeywordContext) FILTER() antlr.TerminalNode {
	return s.GetToken(QueryParserFILTER, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (s *KeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, QueryParserRULE_keyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QueryParserREP)|(1<<QueryParserIN)|(1<<QueryParserAS)|(1<<QueryParserSELECT)|(1<<QueryParserFROM)|(1<<QueryParserFILTER))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) Keyword() IKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *IdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(QueryParserIDENT, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, QueryParserRULE_ident)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserREP, QueryParserIN, QueryParserAS, QueryParserSELECT, QueryParserFROM, QueryParserFILTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Keyword()
		}

	case QueryParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(QueryParserIDENT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentWCContext is an interface to support dynamic dispatch.
type IIdentWCContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentWCContext differentiates from other interfaces.
	IsIdentWCContext()
}

type IdentWCContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentWCContext() *IdentWCContext {
	var p = new(IdentWCContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryParserRULE_identWC
	return p
}

func (*IdentWCContext) IsIdentWCContext() {}

func NewIdentWCContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentWCContext {
	var p = new(IdentWCContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryParserRULE_identWC

	return p
}

func (s *IdentWCContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentWCContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *IdentWCContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(QueryParserWILDCARD, 0)
}

func (s *IdentWCContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentWCContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentWCContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.EnterIdentWC(s)
	}
}

func (s *IdentWCContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryListener); ok {
		listenerT.ExitIdentWC(s)
	}
}

func (s *IdentWCContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryVisitor:
		return t.VisitIdentWC(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *QueryParser) IdentWC() (localctx IIdentWCContext) {
	localctx = NewIdentWCContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, QueryParserRULE_identWC)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryParserREP, QueryParserIN, QueryParserAS, QueryParserSELECT, QueryParserFROM, QueryParserFILTER, QueryParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Ident()
		}

	case QueryParserWILDCARD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Match(QueryParserWILDCARD)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *QueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *FilterExprContext = nil
		if localctx != nil {
			t = localctx.(*FilterExprContext)
		}
		return p.FilterExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QueryParser) FilterExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
