grammar Calculator;

// Parser rules
expression
    : term                              # termOnly
    | Left=expression Op=('+'|'-') Right=term         # addSub
    ;

term
    : factor                            # factorOnly
    | Left=term Op=('*'|'/') Right=factor             # mulDiv
    ;

factor
    : NUMBER                            # number
    | '(' Expr=expression ')'                # parentheses
    ;

// Lexer rules
NUMBER: [0-9]+ ('.' [0-9]+)?;
WS: [ \t\r\n]+ -> skip;
