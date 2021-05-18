package policy

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	"github.com/nspcc-dev/neofs-node/pkg/policy/parser"
)

var (
	ErrInvalidNumber   = errors.New("policy: expected positive integer")
	ErrUnknownClause   = errors.New("policy: unknown clause")
	ErrUnknownOp       = errors.New("policy: unknown operation")
	ErrUnknownFilter   = errors.New("policy: filter not found")
	ErrUnknownSelector = errors.New("policy: selector not found")
)

type policyVisitor struct {
	errors []error
	parser.BaseQueryVisitor
}

func newPolicyVisitor() *policyVisitor {
	return &policyVisitor{}
}

func (p *policyVisitor) VisitPolicy(ctx *parser.PolicyContext) interface{} {
	pl := new(netmap.PlacementPolicy)

	repStmts := ctx.AllRepStmt()
	rs := make([]*netmap.Replica, 0, len(repStmts))
	for _, r := range repStmts {
		rep := r.Accept(p)
		if rep == nil {
			break
		}

		rs = append(rs, rep.(*netmap.Replica))
	}
	pl.SetReplicas(rs...)

	if cbfStmt := ctx.CbtStmt(); cbfStmt != nil {
		cbf := cbfStmt.Accept(p)
		if cbf != nil {
			pl.SetContainerBackupFactor(cbf.(uint32))
		}
	}

	selStmts := ctx.AllSelectStmt()
	ss := make([]*netmap.Selector, 0, len(selStmts))
	for _, s := range selStmts {
		sel := s.Accept(p)
		if sel == nil {
			break
		}

		ss = append(ss, sel.(*netmap.Selector))
	}
	pl.SetSelectors(ss...)

	filtStmts := ctx.AllFilterStmt()
	fs := make([]*netmap.Filter, 0, len(filtStmts))
	for _, f := range filtStmts {
		rep := f.Accept(p)
		if rep == nil {
			break
		}

		fs = append(fs, rep.(*netmap.Filter))
	}
	pl.SetFilters(fs...)

	return pl
}

func (p *policyVisitor) VisitRepStmt(ctx *parser.RepStmtContext) interface{} {
	cnt := ctx.GetCount().GetText()
	num, err := strconv.ParseUint(cnt, 10, 32)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("REP count must be uint32"))
		return nil
	}

	rs := new(netmap.Replica)
	rs.SetCount(uint32(num))
	if sel := ctx.GetSelector(); sel != nil {
		rs.SetSelector(sel.GetText())
	}
	return rs
}

func (p *policyVisitor) VisitSelectStmt(ctx *parser.SelectStmtContext) interface{} {
	s := new(netmap.Selector)
	s.SetCount(p.number(ctx.GetCount()))

	if clStmt := ctx.Clause(); clStmt != nil {
		s.SetClause(clauseFromString(clStmt.GetText()))
	}

	if bStmt := ctx.GetBucket(); bStmt != nil {
		s.SetAttribute(ctx.GetBucket().GetText())
	}

	s.SetFilter(ctx.GetFilter().GetText())

	if ctx.AS() != nil {
		s.SetName(ctx.GetName().GetText())
	}
	return s
}

func (p *policyVisitor) VisitFilterStmt(ctx *parser.FilterStmtContext) interface{} {
	f := p.fillFilter(ctx.GetExpr().(*parser.FilterExprContext))
	f.SetName(ctx.GetName().GetText())
	return f
}

func (p *policyVisitor) fillFilter(e *parser.FilterExprContext) *netmap.Filter {
	if e == nil {
		return nil
	}

	if inner := e.GetInner(); inner != nil {
		return p.fillFilter(inner.(*parser.FilterExprContext))
	}

	f := new(netmap.Filter)
	if eCtx := e.Expr(); eCtx != nil {
		if flt := eCtx.GetFilter(); flt != nil {
			f.SetName(flt.GetText())
		} else {
			f.SetKey(eCtx.GetKey().GetText())
			f.SetOperation(operationFromString(eCtx.GetOp().GetText()))

			valCtx := eCtx.GetValue()
			value := valCtx.GetText()
			if valCtx.GetStr() != nil {
				value = strings.Trim(value, `"`)
			}
			f.SetValue(value)
		}
		return f
	}

	op := operationFromString(e.GetOp().GetText())
	f.SetOperation(op)
	f1 := p.fillFilter(e.GetF1().(*parser.FilterExprContext))
	f2 := p.fillFilter(e.GetF2().(*parser.FilterExprContext))
	if f1.Operation() == op {
		f.SetInnerFilters(append(f1.InnerFilters(), f2)...)
		return f
	}
	f.SetInnerFilters(f1, f2)

	return f
}

func operationFromString(op string) netmap.Operation {
	switch strings.ToUpper(op) {
	case "AND":
		return netmap.OpAND
	case "OR":
		return netmap.OpOR
	case "EQ":
		return netmap.OpEQ
	case "NE":
		return netmap.OpNE
	case "GE":
		return netmap.OpGE
	case "GT":
		return netmap.OpGT
	case "LE":
		return netmap.OpLE
	case "LT":
		return netmap.OpLT
	default:
		return 0
	}
}

func (p *policyVisitor) number(tok antlr.Token) uint32 {
	if tok != nil {
		res, err := strconv.ParseUint(tok.GetText(), 10, 32)
		if err == nil {
			return uint32(res)
		}
		p.errors = append(p.errors, err)
	}
	return 0
}

func (p *policyVisitor) VisitExpr(c *parser.ExprContext) interface{} {
	f := new(netmap.Filter)
	if c.Filter != nil {
		f.SetName(c.Filter.GetText())
	} else if c.Key != nil {
		p.errors = append(p.errors, fmt.Errorf("missing key"))
		return nil
	}
	return f
}

func (p *policyVisitor) ExitEveryRule() {

}

func (p *policyVisitor) VisitTerminal(_ antlr.TerminalNode) interface{} {
	return nil
}

func (p *policyVisitor) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	return nil
}

func parseAntlr(s string) (*netmap.PlacementPolicy, error) {
	input := antlr.NewInputStream(s)
	lexer := parser.NewQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewQueryParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	v := newPolicyVisitor()
	pl := p.Policy().Accept(v)
	return pl.(*netmap.PlacementPolicy), nil
}

func parse(s string) (*query, error) {
	q := new(query)
	err := partparser.Parse(strings.NewReader(s), q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

// Parse parses s into a placement policy.
func Parse(s string) (*netmap.PlacementPolicy, error) {
	return parseAntlr(s)
	q, err := parse(s)
	if err != nil {
		return nil, err
	}

	seenFilters := map[string]bool{}
	fs := make([]*netmap.Filter, 0, len(q.Filters))
	for _, qf := range q.Filters {
		f, err := filterFromOrChain(qf.Value, seenFilters)
		if err != nil {
			return nil, err
		}
		f.SetName(qf.Name)
		fs = append(fs, f)
		seenFilters[qf.Name] = true
	}

	seenSelectors := map[string]bool{}
	ss := make([]*netmap.Selector, 0, len(q.Selectors))
	for _, qs := range q.Selectors {
		if qs.Filter != netmap.MainFilterName && !seenFilters[qs.Filter] {
			return nil, fmt.Errorf("%w: '%s'", ErrUnknownFilter, qs.Filter)
		}
		s := netmap.NewSelector()
		switch len(qs.Bucket) {
		case 1: // only bucket
			s.SetAttribute(qs.Bucket[0])
		case 2: // clause + bucket
			s.SetClause(clauseFromString(qs.Bucket[0]))
			s.SetAttribute(qs.Bucket[1])
		}
		s.SetName(qs.Name)
		seenSelectors[qs.Name] = true
		s.SetFilter(qs.Filter)
		if qs.Count == 0 {
			return nil, fmt.Errorf("%w: SELECT", ErrInvalidNumber)
		}
		s.SetCount(qs.Count)
		ss = append(ss, s)
	}

	rs := make([]*netmap.Replica, 0, len(q.Replicas))
	for _, qr := range q.Replicas {
		r := netmap.NewReplica()
		if qr.Selector != "" {
			if !seenSelectors[qr.Selector] {
				return nil, fmt.Errorf("%w: '%s'", ErrUnknownSelector, qr.Selector)
			}
			r.SetSelector(qr.Selector)
		}
		if qr.Count == 0 {
			return nil, fmt.Errorf("%w: REP", ErrInvalidNumber)
		}
		r.SetCount(uint32(qr.Count))
		rs = append(rs, r)
	}

	p := new(netmap.PlacementPolicy)
	p.SetFilters(fs...)
	p.SetSelectors(ss...)
	p.SetReplicas(rs...)
	p.SetContainerBackupFactor(q.CBF)

	return p, nil
}

func clauseFromString(s string) netmap.Clause {
	switch strings.ToUpper(s) {
	case "SAME":
		return netmap.ClauseSame
	case "DISTINCT":
		return netmap.ClauseDistinct
	default:
		return 0
	}
}

func filterFromOrChain(expr *orChain, seen map[string]bool) (*netmap.Filter, error) {
	var fs []*netmap.Filter
	for _, ac := range expr.Clauses {
		f, err := filterFromAndChain(ac, seen)
		if err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	if len(fs) == 1 {
		return fs[0], nil
	}

	f := netmap.NewFilter()
	f.SetOperation(netmap.OpOR)
	f.SetInnerFilters(fs...)
	return f, nil
}

func filterFromAndChain(expr *andChain, seen map[string]bool) (*netmap.Filter, error) {
	var fs []*netmap.Filter
	for _, fe := range expr.Clauses {
		var f *netmap.Filter
		var err error
		if fe.Expr != nil {
			f, err = filterFromSimpleExpr(fe.Expr, seen)
		} else {
			f = netmap.NewFilter()
			f.SetName(fe.Reference)
		}
		if err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	if len(fs) == 1 {
		return fs[0], nil
	}

	f := netmap.NewFilter()
	f.SetOperation(netmap.OpAND)
	f.SetInnerFilters(fs...)
	return f, nil
}

func filterFromSimpleExpr(se *simpleExpr, seen map[string]bool) (*netmap.Filter, error) {
	f := netmap.NewFilter()
	f.SetKey(se.Key)
	switch se.Op {
	case "EQ":
		f.SetOperation(netmap.OpEQ)
	case "NE":
		f.SetOperation(netmap.OpNE)
	case "GE":
		f.SetOperation(netmap.OpGE)
	case "GT":
		f.SetOperation(netmap.OpGT)
	case "LE":
		f.SetOperation(netmap.OpLE)
	case "LT":
		f.SetOperation(netmap.OpLT)
	default:
		return nil, fmt.Errorf("%w: '%s'", ErrUnknownOp, se.Op)
	}
	f.SetValue(se.Value)
	return f, nil
}
