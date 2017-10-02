package parser

//import "fmt"
import "common"

type Parser struct {
    tokens []common.Token
    Root common.Tree
}

func NewParser(tokens []common.Token) *Parser{
    return &Parser{tokens,common.Tree{nil,common.Token{},nil}}
}
