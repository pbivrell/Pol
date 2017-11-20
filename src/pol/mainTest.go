package main

import "../lexer"
import "fmt"

func main() {
	lex := lexer.NewLexer("+")
	fmt.Println(lex.GetOp())

}
