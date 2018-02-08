package main

import "../lexer"
import "../parser"
import "../interpreter"
import "../common"

import "io/ioutil"
import "flag"
import "fmt"
import "os"

type argsContainer struct {
	errPath string
	treePath string
	tree bool
	src string
	tail []string
}

var args argsContainer

/* ERROR CODES
	0 = OKAY
	1 = No source code provided
	2 = Empty File
	3 = Unable to read input file
	4 = lexer errors
	5 = lexer couldn't find any valid tokens
	6 = printing abs tree no interpreting
	7 = parser errors
	8 = run time errors
*/

func main(){
	arguments()
	src := getSource()
	tokens := lex(src)
	absTree := parse(&tokens)
	interpreter.Interpret(absTree)
}

func parse(tokens *[]common.Token) *common.Tree{
	absTree := parser.Parse(parser.NewTokens(*tokens))

	if errorCheckAndPrint(){
		os.Exit(7)
	}

	if args.tree {
		absTree.PrettyPrint()
		os.Exit(6)
	}

	return absTree
}

func lex(src string) []common.Token{
	lex := lexer.NewLexer(src)
	tokens := lex.Tokenize()

	if len(tokens) <= 0 {
		os.Exit(5)
	}

	if errorCheckAndPrint() {
		os.Exit(4)
	}

	return tokens
}

func getSource() string {
	if args.src != "" {
		//Source code was passed as command line string
		return args.src
	}else if len(args.tail) > 0 {
		//source code comes from a file
		data, err := ioutil.ReadFile(args.tail[0])
		if err != nil {
			fmt.Println("[File Error] Failed to read src file!")
			os.Exit(3)
		}else if len(data) < 1{
			fmt.Println("[File Error] Empty input file!")
			os.Exit(2)
		}
		return string(data)
	}else {
		fmt.Println("[Args Error] No source code provided!")
		os.Exit(0)
		return ""
	}
}


func errorCheckAndPrint() bool{
	if common.HasErrors(){
		fmt.Print(common.GetErrors())
		return true
	}
	return false
}


func arguments(){
	//err := flag.String("errorPath","stderr","prints errmessages to specified path.")
 	tree := flag.Bool("tree", false, "prints POL AST to stdout")
	//treePath := flag.String("treePath","stdout","prints POL AST to specified path.")
	debugParser := flag.Int("dParser",0,"enables debug output for parser at specified level. [0 = no output]")
	debugLexer := flag.Int("dLexer",0,"enables debug output for lexer at specified level. [0 = no output]")
	debugInterpreter := flag.Int("dInterpreter",0,"enables debug output for interpreter at specified level. [0 = no output]")
	stringSrc := flag.String("src","","interprets provided string as source code inplace of a file")

	flag.Parse()

	lexer.Debug = *debugLexer
	parser.Debug = *debugParser
	interpreter.Debug = *debugInterpreter
	args = argsContainer{"","",*tree,*stringSrc,flag.Args()}
}
