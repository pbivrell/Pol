package lexer

import "fmt"
import "unicode"
import "strings"

//Local libs
import "../common"

/* Set by command line arguments. Enables non-error print statements */
var Debug = 0

/* Lexer is a struct containing Lexer information
   - file is an array of runes read from a .pol file
   - pos is the index of the current rune
   - size is the size of file[]
   - lineno is the number of \n chars we have encountered
   - tokens is an array of tokens that is built as we tokenize
*/
type Lexer struct {
	File      []rune
	Pos       int
	size      int
	lineno    int
	HasErrors bool
}

/* Lexer constructor
   creates and initalizes as Lexer
*/

func NewLexer(file string) *Lexer {
	return &Lexer{[]rune(file), 0, len(file), 1, false}
}

/* Next is a method of the Lexer struct. It does not
   take any parameters and returns two types: a bool
   that represents whether or not we have reached
   the end of the file; and a rune that is the next
   utf-8 char in the file. After next has been called
   the position of the current rune is incremented
*/
func (lex *Lexer) next() (rune, bool) {
	if lex.Pos > lex.size-1 {
		return 0, false
	}
	pos := lex.Pos
	lex.Pos = lex.Pos + 1
	return lex.File[pos], true
}

func (lex *Lexer) prev() {
	if lex.Pos <= 0 {
		return
	}
	lex.Pos = lex.Pos - 1
}

/* Peek is a method of the Lexer struct. It does not
   take any parameters and returns two types; a bool
   that represents whether or not we have reached
   the end of the file; and a rune that is the next
   utf-8 char in the file. After peek has been called
   the position of the current rune remains the same.
   (IE: subsequent calls to peek return the same results)
*/
func (lex *Lexer) peek() (rune, bool) {
	if lex.Pos > lex.size-1 {
		return 0, false
	}
	return lex.File[lex.Pos], true
}

/* Error is a function of the takes
   three parameters a pointer to a Lexer, an error message
   and a token. Then it prints a syntax error.
*/
func printError(lex *Lexer, msg string, token string) {
	token = strings.Replace(token, "\n", "\\n", -1)
	if len(token) > 10 {
		token = token[0:9]
	}

	fmt.Printf("Sytnax error: %s line %d: Token [%s]\n", msg, lex.lineno, token)
	//Eat tokens while until whitespace or seperator
	//Should decrease likelyhood of cascading errors
	lex.HasErrors = true
	//This causes problems with the new line count. TODO Revisit
	/*for char, hasNext := lex.next(); hasNext && !((isSeparator(char)) || unicode.IsSpace(char)); char, hasNext = lex.next() {
		if char == '\n'{
			lex.lineno = lex.lineno + 1
		}
	}*/

}

func printDebug(message string, priority int) {
	if Debug >= priority {
		fmt.Println("[DEBUG] " + message)
	}
}

/* Tokenize is a method of Lexer that iterates over
   the Lexer taking identifying runes and calling appropriate
   function to tokenize the identified token.
*/
func (lex *Lexer) Tokenize() []common.Token {
	resTokens := []common.Token{}
	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		token := common.Token{common.ILLEGAL, "", 0}

		//printDebug("BEFORE COMMENT [" + string(test) + "]",1)
		//Eat any comments
		if lex.GetBlockComment() || lex.GetLineComment() {
			continue
		}
		//printDebug("AFTER COMMENT [" + string(test) + "]",1)

		//New line
		if char == '\n' {
			lex.lineno = lex.lineno + 1
			_, _ = lex.next()
			printDebug("New line", 3)
			continue
		} else if unicode.IsSpace(char) {
			_, _ = lex.next()
			printDebug("Whitespace", 2)
			continue
		} else if token = lex.GetNum(); token.Type != common.ILLEGAL {
			//String: returns string token.type
		} else if token = lex.GetString(); token.Type != common.ILLEGAL {
			//Indentifier: returns identifier token.type
		} else if token = lex.GetID(); token.Type != common.ILLEGAL {
			//Operators
		} else if token = lex.GetOp(); token.Type != common.ILLEGAL {
			//Special Chars '{', '}', '(', ')', '[',']', ';'
		} else if token = lex.GetSpecial(); token.Type != common.ILLEGAL {

		}

		if token.Type == common.ILLEGAL {
			printError(lex, "invalid character", string(char))
		} else {
			printDebug(token.String(), 1)
			resTokens = append(resTokens, token)
		}
	}
	return resTokens
}

func (lex *Lexer) GetLineComment() bool {
	//fmt.Println("WE ARE LOOKING AT: " + lex.peek())
	if next, hasNext := lex.peek(); !hasNext || next != '/' {
		return false
	}
	_, _ = lex.next()
	//printDebug("Found first /",1)
	if next, hasNext := lex.peek(); !hasNext || next != '/' {
		lex.prev()
		return false
	}
	_, _ = lex.next()
	//printDebug("Found second /",1)
	printDebug("Found Line comment", 2)
	for next, hasNext := lex.next(); hasNext && next != '\n'; next, hasNext = lex.next() {
	}
	return true
}

func (lex *Lexer) GetBlockComment() bool {
	if next, hasNext := lex.peek(); !hasNext || next != '/' {
		return false
	}
	_, _ = lex.next()
	if next, hasNext := lex.peek(); !hasNext || next != '*' {
		lex.prev()
		return false
	}
	_, _ = lex.next()
	printDebug("Found block comment", 2)
	found := false
	for next, hasNext := lex.next(); hasNext; next, hasNext = lex.next() {

		if next == '*' {
			if next2, hasNext2 := lex.next(); hasNext2 && next2 == '/' {
				found = true
				break
			} else {
				lex.prev()
			}
		}
	}
	if !found {
		printError(lex, "Unclosed block comment", "/*")
	}
	return true
}

func (lex *Lexer) GetSpecial() common.Token {
	var t common.TokenType = common.ILLEGAL

	next, _ := lex.peek()
	switch next {
	case '{':
		t = common.LEFT_BRACE
	case '}':
		t = common.RIGHT_BRACE
	case '[':
		t = common.LEFT_BRACKET
	case ']':
		t = common.RIGHT_BRACKET
	case '(':
		t = common.LEFT_PAREN
	case ')':
		t = common.RIGHT_PAREN
	case ';':
		t = common.SEMI_COLON
	case ',':
		t = common.COMMA
	}
	_, _ = lex.next()

	return common.Token{t, string(next), lex.lineno}
}

/* getOp is a method of the Lexer that
   returns a new token. getOp assumes that you have
   already verified that the first rune is an operator
*/

func (lex *Lexer) GetOp() common.Token {
	op, _ := lex.next()
	op2, hasNext := lex.peek()
	new_op := string(op) + string(op2)

	if token := getOpType(new_op); hasNext && token != common.ILLEGAL {
		_, _ = lex.next()
		return common.Token{token, new_op, lex.lineno}
	} else if token := getOpType(string(op)); token != common.ILLEGAL {
		return common.Token{token, string(op), lex.lineno}
	} else {
		lex.prev()
		return common.Token{common.ILLEGAL, "", 0}
	}
}

func (lex *Lexer) GetNum() common.Token {
	raw_res, hasNext := lex.next()
	if !isDigit(raw_res) || !hasNext {
		lex.prev()
		return common.Token{common.ILLEGAL, "", 0}
	}
	res := string(raw_res)

	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		if isDigit(char) {
			res = res + string(char)
		} else if char == '.' {
			res = res + lex.getDecimal()
			return common.Token{common.DECIMAL_CONST, res, lex.lineno}
		} else {
			break
		}
		_, _ = lex.next()
	}
	return common.Token{common.INTEGER_CONST, res, lex.lineno}
}

/* getDecimal is a method of the Lexer that will eat all of the runes
   that legally follow a decimal point of a number. Returns results as
   a string
*/
func (lex *Lexer) getDecimal() string {
	next, _ := lex.next()
	res := string(next)
	if num, hasNext := lex.peek(); !isDigit(num) || !hasNext {
		printError(lex, "decimal point must be followed by number", string(num))
		return ""
	}
	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		if isDigit(char) {
			res = res + string(char)
		} else {
			return res
		}
		_, _ = lex.next()
	}
	return res
}

/* isDigit() takes rune returns bool if it is between 0 and 9 */
func isDigit(char rune) bool {
	return (char >= '0' && char <= '9')
}

/* getID is a method of the Lexer that returns
   a token. getID eat runes calling lex.next() until it reaches an
   invalid identifier character ![0-9a-zA-Z]. getID assumes that
   the first rune is already a valid ID
*/

func (lex *Lexer) GetID() common.Token {

	raw_res, hasNext := lex.next()
	if !((raw_res >= 'a' && raw_res <= 'z') || (raw_res >= 'A' && raw_res <= 'Z')) || !hasNext {
		lex.prev()
		return common.Token{common.ILLEGAL, "", 0}
	}

	res := string(raw_res)

	for char, hasNext := lex.peek(); hasNext && isID(char); char, hasNext = lex.peek() {
		_, _ = lex.next()
		res = res + string(char)
	}
	if tokType := isKeyword(res); tokType != common.ILLEGAL {
		return common.Token{tokType, res, lex.lineno}
	}

	return common.Token{common.IDENTIFIER, res, lex.lineno}
}

func isKeyword(keyword string) common.TokenType {
	switch keyword {
	case "if":
		return common.IF
	case "for":
		return common.FOR
	case "else":
		return common.ELSE
	case "return":
		return common.RETURN
	case "break":
		return common.BREAK
	case "continue":
		return common.CONTINUE
	case "while":
		return common.WHILE
	case "true":
		return common.TRUE
	case "false":
		return common.FALSE
	default:
		return common.ILLEGAL
	}
}

/* isID takes a rune returns bool if it is [0-9a-zA-Z_] */
func isID(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '_'
}

/* getString is a function that takes a Lexer and returns a token.
   getString will eat runes until it reaches a quote rune '"'. Note
   quotes can be escaped with a backslash allowing the rune '"' to
   be in a string. getString assumes that the first rune is a quote.
*/

func (lex *Lexer) GetString() common.Token {
	raw_res, hasNext := lex.next()
	if hasNext && raw_res != '"' {
		lex.prev()
		return common.Token{common.ILLEGAL, "", 0}
	}

	res := ""
	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		if char == '"' {
			_, _ = lex.next()
			return common.Token{common.STRING_CONST, res, lex.lineno}
		} else if char == '\\' {
			tmp, _ := lex.next()
			if char, hasNext = lex.peek(); hasNext && char == '"' {
				res = res + "\""
				_, _ = lex.next()
			} else {
				res = res + string(tmp)
			}
		} else {
			res = res + string(char)
			_, _ = lex.next()
		}
	}
	printError(lex, "string missing closing quote", res)
	return common.Token{common.ILLEGAL, "", 0}
}

func getOpType(op string) common.TokenType {
	switch op {
	//Operators on Expressions
	case "+":
		return common.ADDITIVE_OP
	case "-":
		return common.ADDITIVE_OP
	case "*":
		return common.MULTIPLICATIVE_OP
	case "/":
		return common.MULTIPLICATIVE_OP
	case "^":
		return common.MULTIPLICATIVE_OP
	case "%":
		return common.MULTIPLICATIVE_OP
	case "!":
		return common.UNARY_OP
	//Operators on Assignment Expressions
	case "=":
		return common.ASSIGNMENT_OP
	case "+=":
		return common.ASSIGNMENT_OP
	case "-=":
		return common.ASSIGNMENT_OP
	case "*=":
		return common.ASSIGNMENT_OP
	case "/=":
		return common.ASSIGNMENT_OP
	case "%=":
		return common.ASSIGNMENT_OP
	//Operators on Conditional Expressions
	case "==":
		return common.CONDITIONAL_OP
	case "!=":
		return common.CONDITIONAL_OP
	case "<":
		return common.CONDITIONAL_OP
	case "<=":
		return common.CONDITIONAL_OP
	case ">":
		return common.CONDITIONAL_OP
	case ">=":
		return common.CONDITIONAL_OP
	case "&&":
		return common.CONDITIONAL_OP
	case "||":
		return common.CONDITIONAL_OP
	//Special Ops
	case "&":
		return common.REFRENCE_OP
	case "{}":
		return common.HASH_OP
	case "[]":
		return common.ARRAY_OP
	default:
		return common.ILLEGAL
	}
}
