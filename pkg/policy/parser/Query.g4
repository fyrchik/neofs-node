grammar Query;

policy
	: repStmt repStmt*
      cbtStmt?
      selectStmt*
      filterStmt*
	;

repStmt
    : REP Count=NUMBER1 	 // number of object replicas
      (IN Selector=ident)? // optional selector name
	;

cbtStmt
	: CBF Factor=NUMBER1 // container backup factor
	;

selectStmt
	: SELECT Count=NUMBER1   	// number of nodes to select without container backup factor *)
      (IN clause? Bucket=ident)? // bucket name
      FROM Filter=identWC 	// filter reference or whole netmap
      (AS Name=ident)?         // optional selector name
	;

clause
    : 'SAME'     // nodes from the same bucket
    | 'DISTINCT' // nodes from distinct buckets
	;

filterExpr
	: F1=filterExpr Op='AND' F2=filterExpr
	| F1=filterExpr Op='OR' F2=filterExpr
	| expr
	;

filterStmt
    : FILTER Expr=filterExpr
      AS Name=ident // obligatory filter name
	;

expr
	: '@' Filter=ident   // filter reference
    | Key=ident Op=operation Value=identOrValue // attribute filter
;

operation : 'EQ' | 'NE' | 'GE' | 'GT' | 'LT' | 'LE' ;
identOrValue : ident | number | Str=STRING ;
number : ZERO | NUMBER1 ;
keyword : REP | IN | AS | SELECT | FROM | FILTER ;
ident : keyword | IDENT;
identWC : ident | WILDCARD ;

REP : 'REP' ;
IN : 'IN' ;
AS : 'AS' ;
CBF : 'CBF' ;
SELECT : 'SELECT' ;
FROM : 'FROM' ;
FILTER : 'FILTER' ;
WILDCARD : '*' ;

IDENT : Nondigit (Digit | Nondigit)* ;

fragment Digit : [0-9] ;
fragment Nondigit : [a-zA-Z_] ;

NUMBER1 : [1-9] Digit* ;
ZERO  : '0' ;

// taken from antlr4 json grammar.
STRING : '"' (ESC | SAFECODEPOINT)* '"' ;

fragment ESC : '\\' (["\\/bfnrt] | UNICODE) ;
fragment UNICODE : 'u' HEX HEX HEX HEX ;
fragment HEX : [0-9a-fA-F] ;
fragment SAFECODEPOINT : ~ ["\\\u0000-\u001F] ;

WS : [ \t\n\r] + -> skip ;
