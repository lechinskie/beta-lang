grammar b;

program
    : topLevel* EOF
    ;

topLevel
    : definition
    | extrndecl
    | variadicdecl
    | ';'
    ;

definition
    : name constant? (ival (',' ival)*)* ';'
    | name '__asm__' '(' stringlist ')' ';'
    | name '(' (name (',' name)*)? ')' statement
    ;

extrndecl
    : 'extrn' name (',' name)* ';'
    ;

variadicdecl
    : '__variadic__' '(' name ',' INT ')' ';'
    ;

ival
    : constant
    | name
    ;

statement
    : externsmt
    | autosmt
    | name ':' statement
    | casestmt
    | blockstmt
    | ifstmt
    | whilestmt
    | switchstmt
    | gotostmt
    | returnstmt
    | asmstmt
    | expressionstmt
    | nullstmt
    ;

nullstmt
    : ';'
    ;

expressionstmt
    : rvalue ';'
    ;

blockstmt
    : '{' statement* '}'
    ;

returnstmt
    : 'return' ('(' rvalue ')')? ';'
    ;

gotostmt
    : 'goto' rvalue ';'
    ;

switchstmt
    : 'switch' rvalue statement
    ;

whilestmt
    : 'while' '(' rvalue ')' statement
    ;

ifstmt
    : 'if' '(' rvalue ')' statement ('else' statement)?
    ;

casestmt
    : 'case' constant ':' statement
    ;

externsmt
    : 'extrn' name (',' name)* ';'
    ;

autosmt
    : 'auto' name constant? (',' name constant?)* ';'
    ;

asmstmt
    : '__asm__' '(' stringlist ')' ';'
    ;

stringlist
    : STRING1 (',' STRING1)*
    ;

rvalue
    : expression
    | comparison
    | ternary
    | assignment
    ;

ternary
    : expression '?' rvalue ':' rvalue
    ;

comparison
    : expression binary rvalue
    ;

assignment
    : name assign rvalue
    ;

expression
    : '(' rvalue ')'
    | name
    | constant
    | incdec name
    | name incdec
    | unary rvalue
    | '&' name
    | functioninvocation
    ;

functioninvocation
    : name '(' functionparameters? ')'
    ;

functionparameters
    : rvalue (',' rvalue)*
    ;

assign
    : '=' binary?
    ;

incdec
    : '++'
    | '--'
    ;

unary
    : '-'
    | '!'
    ;

binary
    : '|'
    | '&'
    | '=='
    | '!='
    | '<'
    | '<='
    | '>'
    | '>='
    | '<<'
    | '>>'
    | '-'
    | '+'
    | '%'
    | '*'
    | '/'
    ;

lvalue
    : name
    | '*' rvalue
    | rvalue '[' rvalue ']'
    ;

constant
    : INT
    | STRING1
    | STRING2
    ;

name
    : NAME
    ;

NAME
    : [a-zA-Z_] [a-zA-Z0-9_]*
    ;

INT
    : [0-9]+
    ;

STRING1
    : '"' ~ ["\r\n]* '"'
    ;

STRING2
    : '\'' ~ ['\r\n]* '\''
    ;

LINECOMMENT
    : '//' ~[\r\n]* -> skip
    ;

BLOCKCOMMENT
    : '/*' .*? '*/' -> skip
    ;

WS
    : [ \t\r\n]+ -> skip
    ;
