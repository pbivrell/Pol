package main

import "io/ioutil"
import "fmt"
import "os"
import "strings"

//Local Libs
import "../lexer"
//import "../parser"
import "../common"

/* Pol main */
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [.pol src file]\n", os.Args[0])
		os.Exit(-1)
	}

	if len(os.Args) == 3 {
		lexer.Debug = true
	}

	//Read input file
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	//Lex file data into slice of Tokens
	lex := lexer.NewLexer(string(data))
	tokens := lex.Tokenize()
	if lex.HasErrors {
		fmt.Println("Pol: Fix errors above before continuing")
		return
	}

	//TODO Make output writing a runtime option
	//TODO replace this horrible mess of string spliting
	file := "." + strings.Split(strings.Split(os.Args[1], "/")[len(strings.Split(os.Args[1], "/"))-1], ".pol")[0] + ".lex_pol"
	common.WriteJSON(tokens, file)

//	absTree := parser.Parse(parser.NewTokens(tokens))
	//fmt.Println(absTree)
}
