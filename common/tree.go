package common

import "fmt"

type Tree struct {
	Tok       Token
	Nodes     []Tree
	NodeCount int
}

func (t *Tree) String() string {
	return t.Tok.Value
}

// Mimic behavior of Append method for other
// Go data structures.
func (t *Tree) Append(new *Tree) *Tree {
	t.Nodes = append(t.Nodes,*new)
	/*if t.NodeCount+1 > cap(t.Nodes) {
		//Do not double size because it zero initalizes unused elements
		//which makes looping over elements harder because we have to check
		//if values are non-zero first.
		newSlice := make([]Tree,(cap(t.Nodes)+1))
		copy(newSlice, t.Nodes)
		t.Nodes = newSlice
	}
	t.Nodes[t.NodeCount] = *new*/
	t.NodeCount++
	return t
}

func (t *Tree) Init() *Tree {
	t.Tok = Token{}
	t.Nodes = make([]Tree, 0)
	t.NodeCount = 0
	return t
}

func MakeTree(tok *Token) *Tree {
	retVal := new(Tree).Init()
	retVal.Tok = *tok
	return retVal
}

func (t *Tree) GetNode(i int) *Tree {
	if t.NodeCount <= i {
		fmt.Printf("TREE DATA STRUCTURE ERROR: Getting node %d with a count of %d", i, (*t).NodeCount)
		return &Tree{}
	}
	return &t.Nodes[i]
}

func (t *Tree) GetChildren() []Tree {
	return t.Nodes
}

func MakeInvalidTree() *Tree {
	return MakeTree(InvalidToken())
}

func (t *Tree) IsInvalidTree() bool {
	return t.Tok.Type == ILLEGAL
}

//func Tree() *Tree{
//	return new(Tree).Init()
//}

func (t *Tree) PrettyPrint() {
	prettyPrint(*t, "", 0)
}

func prettyPrint(t Tree, tab string, depth int) {
	fmt.Printf("%s%d: %s %s\n", tab, depth, t.Tok.Value, t.Tok.Type)
	if t.NodeCount <= 0 {
		return
	}

	for _, node := range t.Nodes {
		prettyPrint(node, tab+"\t", depth+1)
	}
}
