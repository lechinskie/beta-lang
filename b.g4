grammar b;

program : ext_def+ EOF ;

ext_def
    : EXTRN name_list ';'
    | VARIADIC '(' ID ',' expr ')' ';'
    | ID ASM '(' string_list ')' ';'
    | ID '(' arg_list? ')' statement
    | ID '[' expr? ']' ival_list? ';'
    | ID ival ';'
    | ID ';'
    ;

ival : expr ;
ival_list : ival (',' ival)* ;

arg_list : ID (',' ID)* ;
name_list : ID (',' ID)* ;
string_list : STRING (',' STRING)* ;

statement
    : compound_stmt
    | IF '(' expr ')' statement (ELSE statement)?
    | WHILE '(' expr ')' statement
    | SWITCH '(' expr ')' statement
    | CASE expr ':' statement
    | DEFAULT ':' statement
    | ID ':' statement
    | GOTO expr ';'
    | BREAK ';'
    | RETURN ';'
    | RETURN '(' expr ')' ';'
    | ASM '(' string_list ')' ';'
    | auto_decl
    | extrn_decl
    | expr ';'
    | ';'
    ;

compound_stmt : '{' statement* '}' ;

auto_decl : AUTO auto_def (',' auto_def)* ';' ;
auto_def : ID | ID '[' expr ']' ;

extrn_decl : EXTRN name_list ';' ;

expr
    : expr '[' expr ']'
    | expr '(' expr_list? ')'
    | expr (INC | DEC)
    | ('*' | '&' | '-' | '!' | INC | DEC | '~') expr
    | expr ('*' | '/' | '%') expr
    | expr ('+' | '-') expr
    | expr (SHL | SHR) expr
    | expr ('<' | LE | '>' | GE) expr
    | expr (EQ | NE) expr
    | expr '&' expr
    | expr '^' expr
    | expr '|' expr
    | <assoc=right> expr '?' expr ':' expr
    | <assoc=right> expr assign_op expr
    | '(' expr ')'
    | ID
    | DECIMAL
    | OCTAL
    | CHAR
    | STRING
    ;

expr_list : expr (',' expr)* ;

assign_op
    : ASSIGN | ASS_MUL | ASS_DIV | ASS_MOD | ASS_ADD | ASS_SUB 
    | ASS_SHL | ASS_SHR | ASS_LT | ASS_LE | ASS_GT | ASS_GE 
    | ASS_EQ | ASS_NE | ASS_AND | ASS_XOR | ASS_OR
    ;

// Keywords
AUTO : 'auto' ;
BREAK : 'break' ;
CASE : 'case' ;
DEFAULT : 'default' ;
ELSE : 'else' ;
EXTRN : 'extrn' ;
GOTO : 'goto' ;
IF : 'if' ;
RETURN : 'return' ;
SWITCH : 'switch' ;
WHILE : 'while' ;

ASM : '__asm__' ;
VARIADIC : '__variadic__' ;

// Operators
INC : '++' ;
DEC : '--' ;
SHL : '<<' ;
SHR : '>>' ;
LE  : '<=' ;
GE  : '>=' ;
EQ  : '==' ;
NE  : '!=' ;

ASS_MUL : '=*' ;
ASS_DIV : '=/' ;
ASS_MOD : '=%' ;
ASS_ADD : '=+' ;
ASS_SUB : '=-' ;
ASS_SHL : '=<<' ;
ASS_SHR : '=>>' ;
ASS_LE  : '=<=' ;
ASS_LT  : '=<' ;
ASS_GE  : '=>=' ;
ASS_GT  : '=>' ;
ASS_EQ  : '===' ;
ASS_NE  : '=!=' ;
ASS_AND : '=&' ;
ASS_XOR : '=^' ;
ASS_OR  : '=|' ;
ASSIGN  : '=' ;

ID : [a-zA-Z_.] [a-zA-Z_.0-9]* ;

DECIMAL : [1-9][0-9]* ;
OCTAL : '0' [0-7]* ;
CHAR : '\'' ( ESC | ~['\\] )* '\'' ;
STRING : '"' ( ESC | ~["\\] )* '"' ;

fragment ESC : '\\' [0e()t*'"n] ;
LINECOMMENT
    : '//' ~[\r\n]* -> skip
    ;
COMMENT : '/*' .*? '*/' -> skip ;
WS : [ \t\r\n]+ -> skip ;
