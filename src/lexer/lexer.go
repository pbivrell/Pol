package main

import "fmt"
import "io/ioutil"
import "os"
import "unicode"


/* lexer is a struct containing lexer information
   - file is an array of runes read from a .pol file
   - pos is the index of the current rune
   - size is the size of file[]
   - lineno is the number of \n chars we have encountered
   - tokens is an array of tokens that is built as we tokenize
*/
type lexer struct{
    file []rune
    pos int
    size int
    lineno int
    tokens []token
}

/* token is a struct containing token information
   - typ is the token type
   - value is the token value
*/
type token struct{
    typ string
    value string
}

/* Next is a method of the lexer struct. It does not
   take any parameters and returns two types: a bool
   that represents whether or not we have reached
   the end of the file; and a rune that is the next
   utf-8 char in the file. After next has been called
   the position of the current rune is incremented
*/
func (lex *lexer) next() (rune,bool) {
    if lex.pos > lex.size - 1{
        return 0,false
    }
    pos := lex.pos
    lex.pos = lex.pos + 1
    return lex.file[pos],true
}

/* Peek is a method of the lexer struct. It does not
   take any parameters and returns two types; a bool
   that represents whether or not we have reached
   the end of the file; and a rune that is the next
   utf-8 char in the file. After peek has been called
   the position of the current rune remains the same.
   (IE: subsequent calls to peek return the same results)
*/
func (lex *lexer) peek() (rune,bool) {
    if lex.pos > lex.size - 1{
        return 0, false
    }
    return lex.file[lex.pos],true
}


/* Error is a function of the takes
   three parameters a pointer to a lexer, an error message
   and a token. Then it prints a syntax error.

*/
func  error(lex *lexer, msg string, token string) {
    fmt.Printf("Sytnax error [line %d]: %s at token: [%s]\n",lex.lineno, msg, token)
}

/* Tokenize is a function that takes
   a lexer pointer and iterates over the lexer taking idenifying
   runes and calling appropriate function to tokenize the idenified
   token.
*/
func tokenize(lex *lexer)  {
    for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
        //New line
        if char == '\n'{
            lex.lineno = lex.lineno + 1;
            _,_ = lex.next()

        //Digit
        }else if isDigit(char){
            fmt.Printf("Number: %+v\n",getNum(lex))

        //String
        }else if char == '"'{
            fmt.Printf("String: %+v\n",getString(lex))

        //Indentifier
        }else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z'){
            fmt.Printf("ID: %+v\n",getID(lex))

        //Operators
        }else if isOp := isOperator(char); isOp > 0 { //+ - * / ^ % || && = < > <= >= == != !
            fmt.Printf("OP: %+v\n",getOp(lex))

        //Seperators
        }else if isSeperator(char){
            fmt.Printf("Seperator: %+v\n",token{"Seperator",string(char)})
            _,_ = lex.next()

        //Should be the catch all for only whitespace
        }else if unicode.IsSpace(char){
            fmt.Printf("whitespace\n")
            _,_ = lex.next()

        }else{
            error(lex,"invalid character [" + string(char) + "]","FIX THIS PAUL")
        }
    }
}

/* getOp is a function that takes a lexer and
   returns a new token. getOp assumes that you have
   already verfied that the first rune is an operator
*/
func getOp(lex *lexer) token{
    num,_ = lex.next()

    //TODO rewrite this if statement it is ugly. Use some sort of Operator Map?
    if next, hasNext := lex.peek(); isOperator(num) == 2 && hasNext && isDoubleOperator(string(num) + string(next)){
        _, _ = lex.next()
        return token{"op",string(first) + string(next)}
    }
    return token{"op",string(first)}
}

/* getNum is a function that takes a lexer pointer
   and returns a token. getNum will eat runes by calling
   lexer.next() until the result no longer matches [0-9]+(.[0-9]+)?
   getNum assumes that the first rune is a valid number
*/
func getNum(lex *lexer) token {
    res,_ := string(lex.next())
    hasDec := false

    for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek(){
        //This if statement cheks logic for decimals.
        //TODO This is terrible please rewrite
        if char == '.' {
            if hasDec{
                error(lex,"number can not have multiple decimals",res + ".")
                break
            }else{
                res = res + string(char)
                _, _ = lex.next()
                char, _ = lex.peek()
                if char, hasNext = lex.peek(); !hasNext || !isDigit(char){
                    error(lex,"dot must be followed by a number", res + string(char))
                }
                hasDec = true
            }
        }
        if isDigit(char){
            res = res + string(char)
        }else{
            break
        }
    }
    return token{"num",res}
}

/* isDigit() takes rune returns bool if it is between 0 and 9 */
func isDigit(char rune) bool{
    return (char >= '0' && char <= '9')
}

/* getID is a function that takes a lexer pointer and returns
   a token. getID eat runes calling lex.next() until it reaches an
   invalid identifier character ![0-9a-zA-Z]. getID assumes that
   the first rune is already a valid ID
*/
func getID(lex *lexer) token{
    res := string(lex.next())
    char, hasNext := lex.peek()
    for char, hasNext := lex.peek(); hasNext && isID(char); char, hasNext = lex.peek() {
        _, _ = lex.next()
        res = res + string(char)
    }
    return token{"id",res}
}

/* isID takes a rune returns bool if it is [0-9a-zA-Z] */
func isID(char rune) bool{
    return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

/* getString is a function that takes a lexer and returns a token.
   getString will eat runes until it reaches a quote rune '"'. Note
   quotes can be escaped with a backslash allowing the rune '"' to
   be in a string. getString assumes that the first rune is a quote.
*/
func getString(lex *lexer) token{
    _,_ = lex.next()
    res := ""
    for char, hasNext := lex.peek(); hasNext; char, hasNext = lex.peek() {
        if char == '"'{
            return token{"string",res}
        }else if char == '\\' {
            if char, hasNext = lex.peek(); hasNext && char == '"' {
                res = res + "\""
                _,_ = lex.next()
            }
        }else{
            res = res + string(char)
            _,_ = lex.next()
        }
    }
    error(lex,"string missing closing quote",res)
    return token{}
}

/* isDoubleOperator take a string operator and returns
   a bool if it is valid doubleOperator
*/
func isDoubleOperator(op string) bool{
    switch op{
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
func isOperator(char rune) int{
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
func isSeperator(char rune) bool {
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

func main()  {
    if len(os.Args) < 2{
        fmt.Printf("Usage: %s [.pol src file]\n", os.Args[0])
        os.Exit(-1)
    }
    data, err := ioutil.ReadFile(os.Args[1])
    if err != nil{
        panic(err)
    }

    lex := lexer{[]rune(string(data)),0,len(data),1, make([]token,2)}
    pLex := &lex

    //fmt.Printf("%s\n",pLex.file)
    tokenize(pLex)
}
