package main

import "lexer"
import "io/ioutil"
import "fmt"
import "os"
import "common"
import "strings"

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

	if lex.HasErrors {
		fmt.Printf("Pol: Fix errors above before continuing")
		return
	}

	//TODO Make output writing a runtime option
	//TODO replace this horrible mess of string spliting
	file := "." + strings.Split(strings.Split(os.Args[1],"/")[len(strings.Split(os.Args[1],"/"))-1],".pol")[0] + ".lex_pol"
	common.WriteJSON(tokens, file)
}
