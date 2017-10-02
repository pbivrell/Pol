package lexer

import "fmt"
import "unicode"
import "strings"
import "common"

/* Set by command line arguments. Enables non-error print statements */
var Debug = false

/* Lexer is a struct containing Lexer information
   - file is an array of runes read from a .pol file
   - pos is the index of the current rune
   - size is the size of file[]
   - lineno is the number of \n chars we have encountered
   - tokens is an array of tokens that is built as we tokenize
*/
type Lexer struct {
	file      []rune
	pos       int
	size      int
	lineno    int
	HasErrors bool
}

/* Lexer constructor
   creates and initalizes as Lexer
*/

func NewLexer(file string) *Lexer{
	return &Lexer{[]rune(file),0,len(file),1,false}
}


/* Next is a method of the Lexer struct. It does not
   take any parameters and returns two types: a bool
   that represents whether or not we have reached
   the end of the file; and a rune that is the next
   utf-8 char in the file. After next has been called
   the position of the current rune is incremented
*/
func (lex *Lexer) next() (rune, bool) {
	if lex.pos > lex.size-1 {
		return 0, false
	}
	pos := lex.pos
	lex.pos = lex.pos + 1
	return lex.file[pos], true
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
	if lex.pos > lex.size-1 {
		return 0, false
	}
	return lex.file[lex.pos], true
}

/* Error is a function of the takes
   three parameters a pointer to a Lexer, an error message
   and a token. Then it prints a syntax error.
*/
func printError(lex *Lexer, msg string, token string) {
	token = strings.Replace(token,"\n","\\n",-1)
	if len(token) > 10{
		token = token[0:9]
	}

	fmt.Printf("Sytnax error: %s: Line %d: Token [%s]\n", msg,lex.lineno, token)
	//Eat tokens while until whitespace or seperator
	//Should decrease likelyhood of cascading errors
	lex.HasErrors = true
	//This causes problems with the new line count
	/*for char, hasNext := lex.next(); hasNext && !((isSeparator(char)) || unicode.IsSpace(char)); char, hasNext = lex.next() {
		if char == '\n'{
			lex.lineno = lex.lineno + 1
		}
	}*/

}

/* Tokenize is a function that takes
   a Lexer pointer and iterates over the Lexer taking identifying
   runes and calling appropriate function to tokenize the identified
   token.
*/
func Tokenize(lex *Lexer) []common.Token {
	resTokens := []common.Token{}
	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		//New line
		if char == '\n' {
			// Printing when theres a new line feels like to much
			/*if Debug {
				fmt.Println("New line")
			}*/
			lex.lineno = lex.lineno + 1
			_, _ = lex.next()

			//Digit
		} else if isDigit(char) {
			tNum := lex.GetNum()
			if Debug {
				fmt.Printf("Number: %+v\n", tNum)
			}
			resTokens = append(resTokens, tNum)

			//String
		} else if char == '"' {
			tString := lex.GetString()
			if Debug {
				fmt.Printf("String: %+v\n", tString)
			}
			resTokens = append(resTokens, tString)

			//Indentifier
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			tId := lex.GetID()
			if Debug {
				fmt.Printf("ID: %+v\n", tId)
			}
			resTokens = append(resTokens, tId)

			//Operators
		} else if isOp := isOperator(char); isOp > 0 { //+ - * / ^ % || && = < > <= >= == != !
			tOp := lex.GetOp()
			if Debug {
				fmt.Printf("OP: %+v\n", tOp)
			}
			resTokens = append(resTokens, tOp)

			//Seperators
		} else if isSeparator(char) {
			tSeparator := common.Token{"Separator", string(char)}
			if Debug {
				fmt.Printf("Separator: %+v\n", tSeparator)
			}
			_, _ = lex.next()
			resTokens = append(resTokens, tSeparator)

			//Should be the catch all for only whitespace
		} else if unicode.IsSpace(char) {
			if Debug {
				fmt.Println("Whitespace")
			}
			_, _ = lex.next()

		} else {
			printError(lex, "invalid character ["+string(char)+"]", "FIX THIS")
			_, _ = lex.next()
		}
	}
	return resTokens
}

/* getOp is a method of the Lexer that
   returns a new token. getOp assumes that you have
   already verified that the first rune is an operator
*/

func (lex *Lexer) GetOp() common.Token {
	num, _ := lex.next()

	//TODO rewrite this if statement it is ugly. Use some sort of Operator Map?
	if next, hasNext := lex.peek(); isOperator(num) == 2 && hasNext && isDoubleOperator(string(num)+string(next)) {
		_, _ = lex.next()
		return common.Token{"op", string(num) + string(next)}
	}
	return common.Token{"op", string(num)}
}

/* getNum is a method of the Lexer that
   returns a token. getNum will eat runes by calling
   Lexer.next() until the result no longer matches [0-9]+(.[0-9]+)?
   getNum assumes that the first rune is a valid number
*/
func (lex *Lexer) GetNum() common.Token {
	raw_res, _ := lex.next()
	res := string(raw_res)

	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		if isDigit(char) {
			res = res + string(char)
		} else if char == '.' {
			res = res + lex.getDecimal()
			break
		} else {
			break
		}
		_, _ = lex.next()
	}
	return common.Token{"num", res}
}

/* getDecimal is a method of the Lexer that will eat all of the runes
   that legally follow a decimal point of a number. Returns results as
   a string
*/
func (lex *Lexer) getDecimal() string {
	next, _ := lex.next()
	res := string(next)
	if num, hasNext := lex.peek(); !isDigit(num) || !hasNext{
		printError(lex,"decimal point must be followed by number",string(num))
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
	raw_res, _ := lex.next()
	res := string(raw_res)

	for char, hasNext := lex.peek(); hasNext && isID(char); char, hasNext = lex.peek() {
		_, _ = lex.next()
		res = res + string(char)
	}
	return common.Token{"id", res}
}

/* isID takes a rune returns bool if it is [0-9a-zA-Z] */
func isID(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

/* getString is a function that takes a Lexer and returns a token.
   getString will eat runes until it reaches a quote rune '"'. Note
   quotes can be escaped with a backslash allowing the rune '"' to
   be in a string. getString assumes that the first rune is a quote.
*/

func (lex *Lexer) GetString() common.Token {
	_, _ = lex.next()
	res := ""
	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		if char == '"' {
			_, _ = lex.next()
			return common.Token{"string", res}
		} else if char == '\\' {
			if char, hasNext = lex.peek(); hasNext && char == '"' {
				res = res + "\""
				_, _ = lex.next()
			}
		} else {
			res = res + string(char)
			_, _ = lex.next()
		}
	}
	printError(lex, "string missing closing quote", res)
	return common.Token{}
}

/* isDoubleOperator take a string operator and returns
   a bool if it is valid doubleOperator
*/
func isDoubleOperator(op string) bool {
	switch op {
	case "||":
		return true
	case "&&":
		return true
	case "<=":
		return true
	case ">=":
		return true
	case "==":
		return true
	case "!=":
		return true
	default:
		return false
	}
}

/* isOperator takes a rune and returns and int
   - returns 0 if its an invalid operator
   - returns 1 if its a single operator
   - returns 2 if it could be part of a compound operator
   (compound operator: ==, !=, &&)
*/
func isOperator(char rune) int {
	switch char {
	case '+':
		return 1
	case '-':
		return 1
	case '*':
		return 1
	case '/':
		return 1
	case '^':
		return 1
	case '%':
		return 1
	case '|':
		return 2
	case '&':
		return 2
	case '=':
		return 2
	case '<':
		return 2
	case '>':
		return 2
	case '!':
		return 2
	default:
		return 0
	}
}

/* isSeparator takes a rune and returns a bool if it is a valid separator */
func isSeparator(char rune) bool {
	switch char {
	case ';':
		return true
	case '(':
		return true
	case ')':
		return true
	default:
		return false
	}
}
