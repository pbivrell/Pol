package main

import "io/ioutil"
import "fmt"
import "os"
import "strconv"

//import "strings"

//Local Libs
import "../lexer"
import "../parser"

//import "../common"

/* Pol main */
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [.pol src file]\n", os.Args[0])
		os.Exit(-1)
	}

	if len(os.Args) == 4 {
		lexer.Debug, _ = strconv.Atoi(os.Args[3])
		parser.Debug, _ = strconv.Atoi(os.Args[2])
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

	//tokens = tokens
	//os.Exit(3)

	/*
		//TODO Make output writing a runtime option
		//TODO replace this horrible mess of string spliting
		file := "." + strings.Split(strings.Split(os.Args[1], "/")[len(strings.Split(os.Args[1], "/"))-1], ".pol")[0] + ".lex_pol"
		common.WriteJSON(tokens, file)

	*/
	//for _,next := range tokens{
	//	fmt.Printf("%s\n",next)
	//}

	//fmt.Printf("THIS IS WHERE WE ARE NOW\n")

	absTree := parser.Parse(parser.NewTokens(tokens))
	/*if parser.HasErrors {
		fmt.Println("Pol: Fix errors above before continuing")
		return
	}*/
	absTree.PrettyPrint()

}
