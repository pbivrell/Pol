package parser

//import "fmt"
import "common"

type Tokens struct {
    Toks []common.Token
    pos int
    size int
}

func NewTokens(tokens []common.Token) *Tokens{
    return &Tokens{tokens,0,len(tokens)}
}

func (toks *Tokens) next() (common.Token, bool) {
    if toks.pos > toks.size - 1 {
		return common.Token{}, false
	}
	pos := toks.pos
	toks.pos = toks.pos + 1
	return toks.Toks[pos], true
}

func (toks *Tokens) peek() (common.Token, bool) {
    if toks.pos > toks.size - 1 {
		return common.Token{}, false
	}
	return toks.Toks[toks.pos], true
}

func Parse(tokens *Tokens) common.Tree {
    return common.Tree{}
}
