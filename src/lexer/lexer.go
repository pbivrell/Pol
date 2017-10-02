package main

import "fmt"
import "io/ioutil"
import "os"
import "unicode"
import "encoding/json"
import "strings"

/* Set by command line arguments. Enables non-error print statements */
var debug = false

/* lexer is a struct containing lexer information
   - file is an array of runes read from a .pol file
   - pos is the index of the current rune
   - size is the size of file[]
   - lineno is the number of \n chars we have encountered
   - tokens is an array of tokens that is built as we tokenize
*/
type lexer struct {
	file      []rune
	pos       int
	size      int
	lineno    int
	hasErrors bool
}

/* token is a struct containing token information
   - typ is the token type
   - value is the token value
*/

type Token struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

/* Next is a method of the lexer struct. It does not
   take any parameters and returns two types: a bool
   that represents whether or not we have reached
   the end of the file; and a rune that is the next
   utf-8 char in the file. After next has been called
   the position of the current rune is incremented
*/
func (lex *lexer) next() (rune, bool) {
	if lex.pos > lex.size-1 {
		return 0, false
	}
	pos := lex.pos
	lex.pos = lex.pos + 1
	return lex.file[pos], true
}

/* Peek is a method of the lexer struct. It does not
   take any parameters and returns two types; a bool
   that represents whether or not we have reached
   the end of the file; and a rune that is the next
   utf-8 char in the file. After peek has been called
   the position of the current rune remains the same.
   (IE: subsequent calls to peek return the same results)
*/
func (lex *lexer) peek() (rune, bool) {
	if lex.pos > lex.size-1 {
		return 0, false
	}
	return lex.file[lex.pos], true
}

/* Error is a function of the takes
   three parameters a pointer to a lexer, an error message
   and a token. Then it prints a syntax error.
*/
func printError(lex *lexer, msg string, token string) {
	fmt.Printf("Sytnax error [line %d]: %s at token: [%s]\n", lex.lineno, msg, token)
	//Eat tokens while until whitespace or seperator
	//Should decrease likelyhood of cascading errors
	lex.hasErrors = true
	for char, hasNext := lex.next(); hasNext && !((isSeperator(char)) || unicode.IsSpace(char)); char, hasNext = lex.next() {

	}

}

/* Tokenize is a function that takes
a lexer pointer and iterates over the lexer taking idenifying
   runes and calling appropriate function to tokenize the idenified
   token.
*/
func tokenize(lex *lexer) []Token {
	resTokens := []Token{}
	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		//New line
		if char == '\n' {
			lex.lineno = lex.lineno + 1
			_, _ = lex.next()

			//Digit
		} else if isDigit(char) {
			tNum := GetNum(lex)
			if debug {
				fmt.Printf("Number: %+v\n", tNum)
			}
			resTokens = append(resTokens, tNum)

			//String
		} else if char == '"' {
			tString := GetString(lex)
			if debug {
				fmt.Printf("String: %+v\n", tString)
			}
			resTokens = append(resTokens, tString)

			//Indentifier
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			tId := GetID(lex)
			if debug {
				fmt.Printf("ID: %+v\n", tId)
			}
			resTokens = append(resTokens, tId)

			//Operators
		} else if isOp := isOperator(char); isOp > 0 { //+ - * / ^ % || && = < > <= >= == != !
			tOp := GetOp(lex)
			if debug {
				fmt.Printf("OP: %+v\n", tOp)
			}
			resTokens = append(resTokens, tOp)

			//Seperators
		} else if isSeparator(char) {
			tSeperator := Token{"Seperator", string(char)}
			if debug {
				fmt.Printf("Seperator: %+v\n", tSeperator)
			}
			_, _ = lex.next()
			resTokens = append(resTokens, tSeperator)

			//Should be the catch all for only whitespace
		} else if unicode.IsSpace(char) {
			if debug {
				fmt.Printf("Whitespace\n")
			}
			_, _ = lex.next()

		} else {
			printError(lex, "invalid character ["+string(char)+"]", "FIX THIS PAUL")
			_, _ = lex.next()
		}
	}
	return resTokens
}

/* getOp is a function that takes a lexer and
   returns a new token. getOp assumes that you have
   already verified that the first rune is an operator
*/

func GetOp(lex *lexer) Token {
	num, _ := lex.next()

	//TODO rewrite this if statement it is ugly. Use some sort of Operator Map?
	if next, hasNext := lex.peek(); isOperator(num) == 2 && hasNext && isDoubleOperator(string(num)+string(next)) {
		_, _ = lex.next()
		return Token{"op", string(num) + string(next)}
	}
	return Token{"op", string(num)}
}

/* getNum is a function that takes a lexer pointer
   and returns a token. getNum will eat runes by calling
   lexer.next() until the result no longer matches [0-9]+(.[0-9]+)?
   getNum assumes that the first rune is a valid number
*/
func GetNum(lex *lexer) Token {
	first, _ := lex.next()
	res := string(first)

	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		if isDigit(char) {
			res = res + string(char)
		} else if char == '.' {
			res = res + getDecimal(lex)
			break
		} else {
			break
		}
		_, _ = lex.next()
	}
	return Token{"num", res}
}

/* getDecimal is a function that will eat all of the runes
   that can be in a decimal point of a number
*/
func getDecimal(lex *lexer) string {
	next, _ := lex.next()
	res := string(next)
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

/* getID is a function that takes a lexer pointer and returns
   a token. getID eat runes calling lex.next() until it reaches an
   invalid identifier character ![0-9a-zA-Z]. getID assumes that
   the first rune is already a valid ID
*/

func GetID(lex *lexer) Token {
	first, _ := lex.next()
	res := string(first)
  
	for char, hasNext := lex.peek(); hasNext && isID(char); char, hasNext = lex.peek() {
		_, _ = lex.next()
		res = res + string(char)
	}
	return Token{"id", res}
}

/* isID takes a rune returns bool if it is [0-9a-zA-Z] */
func isID(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

/* getString is a function that takes a lexer and returns a token.
   getString will eat runes until it reaches a quote rune '"'. Note
   quotes can be escaped with a backslash allowing the rune '"' to
   be in a string. getString assumes that the first rune is a quote.
*/

func GetString(lex *lexer) Token {
	_, _ = lex.next()
	res := ""
	for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
		if char == '"' {
			_, _ = lex.next()
			return Token{"string", res}
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
	return Token{}
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

/* isSeperator takes a rune and returns a bool if it is a valid seperator */
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

/* Writes tokens to a json file. Will not write if some errors
   were generated during tokenization
*/
func writeTokens(toks []Token, hasErrors bool, filename string) {
	if hasErrors {
		fmt.Printf("Lexer: Failed to lex file %s see errors above!\n", filename)
		return
	}
	//Convert Exported struct Token to Json
	jsonData, err := json.Marshal(toks)
	if err != nil {
		panic(err)
	}

	//TODO make run time argument so that output path is not static
	//but can be defined by the user
	file := strings.Split(filename, "/")
	filePath := "." + strings.Split(file[len(file)-1], ".pol")[0] + ".lex_pol"

	fh, errw := os.Create(filePath)
	defer fh.Close()
	if err != nil {
		panic(errw)
	}
	_, write_err := fh.Write(jsonData)
	if write_err != nil {
		panic(write_err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [.pol src file]\n", os.Args[0])
		os.Exit(-1)
	}

	if len(os.Args) == 3 {
		debug = true
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	lex := lexer{[]rune(string(data)), 0, len(data), 1, false}
	pLex := &lex

	tokens := tokenize(pLex)
	writeTokens(tokens, pLex.hasErrors, os.Args[1])
}
