// Code generated from Query.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 192,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 7, 21, 143, 10,
	21, 12, 21, 14, 21, 146, 11, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3,
	24, 7, 24, 154, 10, 24, 12, 24, 14, 24, 157, 11, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 7, 26, 164, 10, 26, 12, 26, 14, 26, 167, 11, 26, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 27, 5, 27, 174, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 6, 31, 187, 10, 31, 13,
	31, 14, 31, 188, 3, 31, 3, 31, 2, 2, 32, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 2, 45, 2, 47, 23, 49, 24,
	51, 25, 53, 2, 55, 2, 57, 2, 59, 2, 61, 26, 3, 2, 9, 3, 2, 50, 59, 5, 2,
	67, 92, 97, 97, 99, 124, 3, 2, 51, 59, 10, 2, 36, 36, 49, 49, 94, 94, 100,
	100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99,
	104, 5, 2, 2, 33, 36, 36, 94, 94, 5, 2, 11, 12, 15, 15, 34, 34, 2, 192,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41,
	3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2,
	61, 3, 2, 2, 2, 3, 63, 3, 2, 2, 2, 5, 68, 3, 2, 2, 2, 7, 77, 3, 2, 2, 2,
	9, 81, 3, 2, 2, 2, 11, 84, 3, 2, 2, 2, 13, 86, 3, 2, 2, 2, 15, 89, 3, 2,
	2, 2, 17, 92, 3, 2, 2, 2, 19, 95, 3, 2, 2, 2, 21, 98, 3, 2, 2, 2, 23, 101,
	3, 2, 2, 2, 25, 104, 3, 2, 2, 2, 27, 108, 3, 2, 2, 2, 29, 111, 3, 2, 2,
	2, 31, 114, 3, 2, 2, 2, 33, 118, 3, 2, 2, 2, 35, 125, 3, 2, 2, 2, 37, 130,
	3, 2, 2, 2, 39, 137, 3, 2, 2, 2, 41, 139, 3, 2, 2, 2, 43, 147, 3, 2, 2,
	2, 45, 149, 3, 2, 2, 2, 47, 151, 3, 2, 2, 2, 49, 158, 3, 2, 2, 2, 51, 160,
	3, 2, 2, 2, 53, 170, 3, 2, 2, 2, 55, 175, 3, 2, 2, 2, 57, 181, 3, 2, 2,
	2, 59, 183, 3, 2, 2, 2, 61, 186, 3, 2, 2, 2, 63, 64, 7, 85, 2, 2, 64, 65,
	7, 67, 2, 2, 65, 66, 7, 79, 2, 2, 66, 67, 7, 71, 2, 2, 67, 4, 3, 2, 2,
	2, 68, 69, 7, 70, 2, 2, 69, 70, 7, 75, 2, 2, 70, 71, 7, 85, 2, 2, 71, 72,
	7, 86, 2, 2, 72, 73, 7, 75, 2, 2, 73, 74, 7, 80, 2, 2, 74, 75, 7, 69, 2,
	2, 75, 76, 7, 86, 2, 2, 76, 6, 3, 2, 2, 2, 77, 78, 7, 67, 2, 2, 78, 79,
	7, 80, 2, 2, 79, 80, 7, 70, 2, 2, 80, 8, 3, 2, 2, 2, 81, 82, 7, 81, 2,
	2, 82, 83, 7, 84, 2, 2, 83, 10, 3, 2, 2, 2, 84, 85, 7, 66, 2, 2, 85, 12,
	3, 2, 2, 2, 86, 87, 7, 71, 2, 2, 87, 88, 7, 83, 2, 2, 88, 14, 3, 2, 2,
	2, 89, 90, 7, 80, 2, 2, 90, 91, 7, 71, 2, 2, 91, 16, 3, 2, 2, 2, 92, 93,
	7, 73, 2, 2, 93, 94, 7, 71, 2, 2, 94, 18, 3, 2, 2, 2, 95, 96, 7, 73, 2,
	2, 96, 97, 7, 86, 2, 2, 97, 20, 3, 2, 2, 2, 98, 99, 7, 78, 2, 2, 99, 100,
	7, 86, 2, 2, 100, 22, 3, 2, 2, 2, 101, 102, 7, 78, 2, 2, 102, 103, 7, 71,
	2, 2, 103, 24, 3, 2, 2, 2, 104, 105, 7, 84, 2, 2, 105, 106, 7, 71, 2, 2,
	106, 107, 7, 82, 2, 2, 107, 26, 3, 2, 2, 2, 108, 109, 7, 75, 2, 2, 109,
	110, 7, 80, 2, 2, 110, 28, 3, 2, 2, 2, 111, 112, 7, 67, 2, 2, 112, 113,
	7, 85, 2, 2, 113, 30, 3, 2, 2, 2, 114, 115, 7, 69, 2, 2, 115, 116, 7, 68,
	2, 2, 116, 117, 7, 72, 2, 2, 117, 32, 3, 2, 2, 2, 118, 119, 7, 85, 2, 2,
	119, 120, 7, 71, 2, 2, 120, 121, 7, 78, 2, 2, 121, 122, 7, 71, 2, 2, 122,
	123, 7, 69, 2, 2, 123, 124, 7, 86, 2, 2, 124, 34, 3, 2, 2, 2, 125, 126,
	7, 72, 2, 2, 126, 127, 7, 84, 2, 2, 127, 128, 7, 81, 2, 2, 128, 129, 7,
	79, 2, 2, 129, 36, 3, 2, 2, 2, 130, 131, 7, 72, 2, 2, 131, 132, 7, 75,
	2, 2, 132, 133, 7, 78, 2, 2, 133, 134, 7, 86, 2, 2, 134, 135, 7, 71, 2,
	2, 135, 136, 7, 84, 2, 2, 136, 38, 3, 2, 2, 2, 137, 138, 7, 44, 2, 2, 138,
	40, 3, 2, 2, 2, 139, 144, 5, 45, 23, 2, 140, 143, 5, 43, 22, 2, 141, 143,
	5, 45, 23, 2, 142, 140, 3, 2, 2, 2, 142, 141, 3, 2, 2, 2, 143, 146, 3,
	2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 42, 3, 2, 2,
	2, 146, 144, 3, 2, 2, 2, 147, 148, 9, 2, 2, 2, 148, 44, 3, 2, 2, 2, 149,
	150, 9, 3, 2, 2, 150, 46, 3, 2, 2, 2, 151, 155, 9, 4, 2, 2, 152, 154, 5,
	43, 22, 2, 153, 152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2,
	2, 2, 155, 156, 3, 2, 2, 2, 156, 48, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2,
	158, 159, 7, 50, 2, 2, 159, 50, 3, 2, 2, 2, 160, 165, 7, 36, 2, 2, 161,
	164, 5, 53, 27, 2, 162, 164, 5, 59, 30, 2, 163, 161, 3, 2, 2, 2, 163, 162,
	3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2,
	2, 2, 166, 168, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 168, 169, 7, 36, 2, 2,
	169, 52, 3, 2, 2, 2, 170, 173, 7, 94, 2, 2, 171, 174, 9, 5, 2, 2, 172,
	174, 5, 55, 28, 2, 173, 171, 3, 2, 2, 2, 173, 172, 3, 2, 2, 2, 174, 54,
	3, 2, 2, 2, 175, 176, 7, 119, 2, 2, 176, 177, 5, 57, 29, 2, 177, 178, 5,
	57, 29, 2, 178, 179, 5, 57, 29, 2, 179, 180, 5, 57, 29, 2, 180, 56, 3,
	2, 2, 2, 181, 182, 9, 6, 2, 2, 182, 58, 3, 2, 2, 2, 183, 184, 10, 7, 2,
	2, 184, 60, 3, 2, 2, 2, 185, 187, 9, 8, 2, 2, 186, 185, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190,
	3, 2, 2, 2, 190, 191, 8, 31, 2, 2, 191, 62, 3, 2, 2, 2, 10, 2, 142, 144,
	155, 163, 165, 173, 188, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'SAME'", "'DISTINCT'", "'AND'", "'OR'", "'@'", "'EQ'", "'NE'", "'GE'",
	"'GT'", "'LT'", "'LE'", "'REP'", "'IN'", "'AS'", "'CBF'", "'SELECT'", "'FROM'",
	"'FILTER'", "'*'", "", "", "'0'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "REP", "IN", "AS", "CBF",
	"SELECT", "FROM", "FILTER", "WILDCARD", "IDENT", "NUMBER1", "ZERO", "STRING",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "REP", "IN", "AS", "CBF", "SELECT", "FROM", "FILTER",
	"WILDCARD", "IDENT", "Digit", "Nondigit", "NUMBER1", "ZERO", "STRING",
	"ESC", "UNICODE", "HEX", "SAFECODEPOINT", "WS",
}

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *QueryLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewQueryLexer(input antlr.CharStream) *QueryLexer {
	l := new(QueryLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerT__0     = 1
	QueryLexerT__1     = 2
	QueryLexerT__2     = 3
	QueryLexerT__3     = 4
	QueryLexerT__4     = 5
	QueryLexerT__5     = 6
	QueryLexerT__6     = 7
	QueryLexerT__7     = 8
	QueryLexerT__8     = 9
	QueryLexerT__9     = 10
	QueryLexerT__10    = 11
	QueryLexerREP      = 12
	QueryLexerIN       = 13
	QueryLexerAS       = 14
	QueryLexerCBF      = 15
	QueryLexerSELECT   = 16
	QueryLexerFROM     = 17
	QueryLexerFILTER   = 18
	QueryLexerWILDCARD = 19
	QueryLexerIDENT    = 20
	QueryLexerNUMBER1  = 21
	QueryLexerZERO     = 22
	QueryLexerSTRING   = 23
	QueryLexerWS       = 24
)
