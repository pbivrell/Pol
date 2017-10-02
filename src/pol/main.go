package main

import "lexer"
import "io/ioutil"
import "fmt"
import "os"

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
	tokens := lexer.Tokenize(lex)
    lexer.WriteTokens(tokens, lex.HasErrors, os.Args[1])
}
