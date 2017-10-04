package parser

import "fmt"
//Local libs
import "../common"

//TODO create so sort of generic iterator structure
//for code from line 8-32 this code is practically
//identical to code in lexer
type Tokens struct {
	Toks      []common.Token
	pos       int
	size      int
}

func NewTokens(tokens []common.Token) *Tokens {
	return &Tokens{tokens, 0, len(tokens)}
}

func (toks *Tokens) next() (common.Token, bool) {
	if toks.pos > toks.size-1 {
		return common.Token{}, false
	}
	pos := toks.pos
	toks.pos = toks.pos + 1
	return toks.Toks[pos], true
}

func (toks *Tokens) peek() (common.Token, bool) {
	if toks.pos > toks.size-1 {
		return common.Token{}, false
	}
	return toks.Toks[toks.pos], true
}

func printError(tok common.Token, msg string) {
    fmt.Printf("Symantic Error: %s: Line %d: Token [%s]\n",msg,tok.Lineno,tok.Value)
}

func Parse(tokens *Tokens) *common.Tree {
    var resTree common.Tree
    token, hasNext := tokens.peek()
	if !hasNext {
        printError(common.Token{},"Expected assignment statement or expression")
        return &common.Tree{}
    }else if token.Type == "id"{
        return asgn_expr(tokens)
    }else {
        return cond_expr(tokens)
    }
}

func asgn_expr(tokens *Tokens) *common.Tree {
    token, _ := tokens.next()
    lhs := &common.Tree{nil,token,nil}
    rhs := expr(tokens)
    return &common.Tree{lhs,token,rhs}
}

func cond_expr(tokens *Tokens) *common.Tree {
    token, _ := tokens.next()
    lhs := expr(tokens)

    for token, hasNext := tokens.peek(); hasNext && token.Type == "c_op"; {
		cop ,_ = tokens.next()
		rhs := expr(tokens)
		lhs := common.Tree{lhs,cop,rhs}
    }
	return lhs
}

func expr(tokens *Tokens) *common.Tree {
	token, _ := tokens.next()
	lhs := term(tokens)

	for token, hasNext := tokens.peek(); hasNext && token.Type = "t_op"; {
		op, _ = tokens.next()
		rhs := expr(tokens)
		lhs := common.Tree{lhs,op, rhs}
	}

	return lhs
}

func term(tokens *Tokens) *common.Tree {
	token, _ := tokens.next()
	lhs := factor(tokens)

	for token, hasNext := tokens.peek(); hasNext && token.Type = "f_op" {
		op, _ = tokens.next()
		rhs := factor(tokens *Tokens)
		lhs := common.Tree{lhs,op,rhs}
	}

	return lhs
}

func factor(tokens *Tokens) *common.Tree {

}
