package main

import "fmt"
import "io/ioutil"
import "os"
import "unicode"


type lexer struct{
    file []rune
    pos int
    size int
    lineno int
    tokens []token
}

type token struct{
    typ string
    value string
}

func  error(lex *lexer, msg string, token string) {
    fmt.Printf("Sytnax error [line %d]: %s at token: [%s]\n",lex.lineno, msg, token)
}

func (lex *lexer) next() (rune,bool) {
    if lex.pos > lex.size - 1{
        return 0,false
    }
    pos := lex.pos
    lex.pos = lex.pos + 1
    return lex.file[pos],true
}


func (lex *lexer) peek() (rune,bool) {
    if lex.pos > lex.size - 1{
        return 0, false
    }
    return lex.file[lex.pos],true
}

/* Tokens in POL
  identifier: [A-Za-z][A-Za-z0-9]
  num: [0-9]+(.[0-9]+)?
  string: "*" You can not have a quote within a string
  operator: + - * / ^ % && ||
*/
func tokenize(lex *lexer)  {
    char, hasNext := lex.next()
    for hasNext {
        //New line
        if char == '\n'{
            lex.lineno = lex.lineno + 1;

        //Digit
        }else if isDigit(char){
            //lex.tokens.append(getNum(lex))
            fmt.Printf("Number: %+v\n",getNum(lex,char))

        //String
        }else if char == '"'{
            //lex.tokens.append(getString(lex))
            fmt.Printf("String: %+v\n",getString(lex))

        //Indentifier
        }else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z'){
            fmt.Printf("ID: %+v\n",getID(lex, char))

        //Operators
        }else if isOp := isOperator(char); isOp > 0 { //+ - * / ^ % || && = < > <= >= == != !
            fmt.Printf("OP: %+v\n",getOp(lex,char,isOp))
            //lex.tokens.append(getOp(lex))

        //Seperators
        }else if isSeperator(char){
            fmt.Printf("Seperator: %+v\n",token{"Seperator",string(char)})

        //Should be the catch all for only whitespace
        }else if unicode.IsSpace(char){
            fmt.Printf("whitespace\n")

        }else{
            error(lex,"invalid character [" + string(char) + "]","FIX THIS PAUL")
        }
        char, hasNext = lex.next()
    }
}

func getOp(lex *lexer, first rune, opType int) token{
    if opType == 2{
        if next, hasNext := lex.peek(); hasNext && isDoubleOperator(string(first) + string(next)){
            _, _ = lex.next()
            return token{"op",string(first) + string(next)}
        }
    }
    return token{"op",string(first)}
}

func getNum(lex *lexer, first rune) token {
    res := string(first)
    hasDec := false
    char, hasNext := lex.next();

    for hasNext {
        if char == '.'{
            if hasDec{
                error(lex,"number can not have multiple decimals",res + ".")
                break
            }else{
                res = res + string(char)
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
        char, hasNext = lex.next()
    }
    return token{"num",res}
}

func isDigit(char rune) bool{
    return (char >= '0' && char <= '9')
}

func getID(lex *lexer, first rune) token{
    res := string(first)
    char, hasNext := lex.peek()
    for hasNext && isID(char) {
        _, _ = lex.next()
        res = res + string(char)
        char, hasNext = lex.next()
    }
    return token{"id",res}
}

func isID(char rune) bool{
    return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func getString(lex *lexer) token{
    res := ""
    char, hasNext := lex.next()
    for hasNext {
        if char == '"'{
            return token{"string",res}
        }else if char == '\\' {
            if char, hasNext = lex.peek(); hasNext && char == '"' {
                res = res + "\""
                _,_ = lex.next()
            }
        }else{
            res = res + string(char)
        }
        char, hasNext = lex.next()
    }
    error(lex,"string missing closing quote",res)
    return token{}
}

/* Valid ops
   + - * / ^ % || && = < > <= >= == != !
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
