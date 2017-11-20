package common

/* token is a struct containing token information
   - typ is the token type
   - value is the token value
*/

type TokenType string

const (
	ILLEGAL    TokenType = "ILLEGAL"
	IDENTIFIER           = "IDENTIFIER"

	//Litterals
	STRING_CONST  = "STRING_CONST"
	INTEGER_CONST = "INTEGER_CONST"
	DECIMAL_CONST = "DECIMAL_CONST"
	ZERO_VALUE    = "ZERO_VALUE"

	//Operators
	HASH_OP           = "HASH_OP"
	ARRAY_OP          = "ARRAY_OP"
	ASSIGNMENT_OP     = "ASSIGNMENT_OP"
	ADDITIVE_OP       = "ADDITIVE_OP"
	MULTIPLICATIVE_OP = "MULTIPLICATIVE_OP"
	CONDITIONAL_OP    = "CONDITIONAL_OP"
	UNARY_OP          = "UNARY_OP"
	REFRENCE_OP       = "REFRENCE_OP"

	//Keywords
	IF       = "IF"
	FOR      = "FOR"
	WHILE    = "WHILE"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	//SPECIAL characters
	COMMA         = "COMMA"
	SEMI_COLON    = "SEMI_COLON"
	LEFT_PAREN    = "LEFT_PAREN"
	RIGHT_PAREN   = "RIGHT_PAREN"
	LEFT_BRACKET  = "LEFT_BRACKET"
	RIGHT_BRACKET = "RIGHT_BRACKET"
	LEFT_BRACE    = "LEFT_BRACE"
	RIGHT_BRACE   = "RIGHT_BRACE"

	//TREE TOKENS
	TREE_ROOT      = "POL"
	TREE_MAIN      = "MAIN"
	TREE_FUNC      = "FUNC"
	TREE_GLOBAL    = "GLOBAL"
	TREE_BODY      = "BODY"
	TREE_ARGS      = "ARGS"
	TREE_TEMP      = "TEMP"
	TREE_INIT_LIST = "INIT_LIST"
	TREE_LOOK_UP   = "LOOK_UP"
	HASH           = "HASH"
	ARRAY          = "ARRAY"
	PRIMATIVE      = "PRIMATIVE"
	INDEX          = "INDEX"
)

type Token struct {
	Type   TokenType
	Value  string
	Lineno int
}

func InvalidToken() *Token {
	return &Token{ILLEGAL, "ILLEGAL", -1}
}

func (t Token) String() string {
	return "[T: " + string(t.Type) + " V: " + t.Value + "]"
}
