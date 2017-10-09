package parser

import "fmt"

//Local libs
import "../common"

var Debug = false
var HasErrors = false

//TODO create so sort of generic iterator structure
//for code from line 8-32 this code is practically
//identical to code in lexer
type Tokens struct {
	Toks []common.Token
	pos  int
	size int
}

func NewTokens(tokens []common.Token) *Tokens {
	return &Tokens{tokens, 0, len(tokens)}
}

func (toks *Tokens) next() (*common.Token, bool) {
	if toks.pos > toks.size-1 {
		return &common.Token{}, false
	}
	pos := toks.pos
	toks.pos = toks.pos + 1
	return &toks.Toks[pos], true
}

func (toks *Tokens) peek() (*common.Token, bool) {
	if toks.pos > toks.size-1 {
		return &common.Token{}, false
	}
	return &toks.Toks[toks.pos], true
}

func printError(tok *common.Token, msg string) {
	HasErrors = true
	fmt.Printf("Symantic Error: %s: Line %d: Token [%s]\n", msg, tok.Lineno, tok.Value)
}

func Parse(tokens *Tokens) *common.Tree {
	tree := common.MakeTree(&common.Token{"POL", "POL", -1})
	for _, hasNext := tokens.peek(); hasNext; _, hasNext = tokens.peek() {
		subTree := asgn_expr(tokens)
		tree = tree.Append(subTree)
	}
	return tree
}

func asgn_expr(tokens *Tokens) *common.Tree {
	tree := generic_expr(tokens, cond_expr, "a_op")
	if token, hasNext := tokens.peek(); hasNext && token.Value == ";" {
		if Debug {
			fmt.Printf("Found %s\n", token)
		}
	} else {
		printError(token, "Semi colon expected")
	}
	_, _ = tokens.next()
	return tree
}

func cond_expr(tokens *Tokens) *common.Tree {
	return generic_expr(tokens, expr, "c_op")
}

func expr(tokens *Tokens) *common.Tree {
	return generic_expr(tokens, term, "t_op")
}

func term(tokens *Tokens) *common.Tree {
	return generic_expr(tokens, factor, "f_op")
}

func generic_expr(tokens *Tokens, fn func(*Tokens) *common.Tree, tokenType string) *common.Tree {
	token, hasNext := tokens.peek()

	if !hasNext {
		printError(token, tokenType+" expected")
		return &common.Tree{}
	}

	node := fn(tokens)

	for token, hasNext := tokens.peek(); hasNext && token.Type == tokenType; token, hasNext = tokens.peek() {
		if Debug {
			fmt.Printf("Found %s: %s\n", tokenType, token)
		}
		temp := node
		node = common.MakeTree(token)
		node = node.Append(temp)
		_, _ = tokens.next()
		subTree := fn(tokens)
		node = node.Append(subTree)
	}
	return node
}

func factor(tokens *Tokens) *common.Tree {
	token, hasNext := tokens.peek()

	if Debug {
		fmt.Printf("Found Factor: %+v\n", token)
	}

	if !hasNext {
		printError(token, "Factor expected")
		return &common.Tree{}

	} else if token.Type == "num" || token.Type == "id" || token.Type == "string" {
		node := common.MakeTree(token)
		_, _ = tokens.next()
		return node
	} else if token.Value == "-" {
		node := common.MakeTree(token)
		_, _ = tokens.next()
		subTree := factor(tokens)
		node = node.Append(subTree)
		return node
	} else if token.Value == "!" {
		node := common.MakeTree(token)
		_, _ = tokens.next()
		subTree := cond_expr(tokens)
		node = node.Append(subTree)
		return node
	} else if token.Value == "(" {
		_, _ = tokens.next()
		node := expr(tokens)
		if next, hasNext := tokens.peek(); !hasNext || next.Value != ")" {
			printError(token, " Expected closing paren )")
		} else {
			_, _ = tokens.next()
		}
		return node
	} else {
		return &common.Tree{} //For type checker
	}
}
