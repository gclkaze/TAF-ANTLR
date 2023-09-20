// Tafexpr.g4
grammar Tafexpr;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
INTEGER: [0-9]+;
DOUBLE: [0]'.'[0]|[1-9][0-9]*'.'[0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
LBR : '[';
RBR : ']';
CON : '.';

fragment LOWERCASE  :   [a-z];
fragment UPPERCASE  :   [A-Z];
fragment UNDERSCORE :   '_';
DOLLAR :   '$';  
fragment SINGLELETTER   :   ( LOWERCASE | UPPERCASE);
fragment VARIABLECON : '.';


NUMBER : INTEGER;

VARIABLE_NAME : 
    DOLLAR SINGLELETTER (SINGLELETTER|INTEGER|UNDERSCORE)*
;

/*VARIABLE_PATH:
    VARIABLE_NAME ('.'  PROP  )+
    ;
*/
taf_expression: expression EOF;

expression
   : 
    expression op=('*'|'/') expression #MulDiv
   | expression op=('+'|'-') expression #AddSub
   | INTEGER                             #Number
   | DOUBLE                              #DoubleValue
   | parenthesisExpression              #OrderedEvaluation
   | var_expression #HandleVarExpression
   ;

/**

var_path: 
CON  PROP ( LBR expression RBR  )? (CON  PROP ( LBR expression RBR  )? )* #VarPath
;

var_expression :
|  VARIABLE_NAME var_path|;
 */

 /**old
var_expression :
|  VARIABLE_NAME (CON  PROP ( LBR expression RBR  )? )*;
 
  */
index_expression : expression #IndexExpression
;

var_path: 
  PROP ( LBR index_expression RBR  )? (CON  PROP ( LBR index_expression RBR  )? )* #VarPath
;

var_expression :
| VARIABLE_NAME CON var_path
|  VARIABLE_NAME;

PROP: SINGLELETTER (SINGLELETTER|INTEGER|UNDERSCORE)*;


parenthesisExpression: 
                      '(' (parenthesisExpression | expression) ')' 
                      ;
 