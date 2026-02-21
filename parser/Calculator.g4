grammar Calculator;

// Parser rules
expression
    : term                                   # termOnly
    | Left=expression Op=FN_NAME Right=term  # dyaCall
    | Left=expression Op=FN_NAME             # monCall
    | Left=expression '{' Body=expression '}' Right=term  # dyaCallBody
    | Left=expression '{' Body=expression '}'             # monCallBody
    ;

term
    : NUMBER                   # number
    | VAL_NAME                 # valueName
    | '(' Body=expression ')'  # parentheses
    | '{' Body=expression '}'  # funcBody
    ;

// Lexer rules
NUMBER: [0-9]+ ('.' [0-9]+)?;
WS: [ \t\r\n]+ -> skip;
FN_NAME: [A-Z][A-Za-z0-9]*|[*/+-];
VAL_NAME: [a-z][A-Za-z0-9]*;
