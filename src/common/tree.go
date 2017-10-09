package common

import "fmt"

type Tree struct {
	Tok   Token
	Nodes []Tree
	NodeCount int
}

func (t* Tree) String() string{
	return t.Tok.Value
}

// Mimic behavior of Append method for other
// Go data structures.
func (t* Tree) Append(new* Tree) *Tree{
	if t.NodeCount + 1 > cap(t.Nodes){
		newSlice := make([]Tree, (cap(t.Nodes)+1)*2)
		copy(newSlice, t.Nodes)
		t.Nodes = newSlice
	}
	t.Nodes[t.NodeCount] = *new
	t.NodeCount++
	return t
}

func (t* Tree) Init() *Tree {
	t.Tok = Token{}
	t.Nodes = make([]Tree, 1)
	t.NodeCount = 0
	return t
}

func MakeTree(tok* Token) *Tree{
	retVal := new(Tree).Init()
	retVal.Tok = *tok
	return retVal
}

//func Tree() *Tree{
//	return new(Tree).Init()
//}

func (t* Tree) PrettyPrint(){
	prettyPrint(*t,"",0)
}

func prettyPrint(t Tree,tab string, depth int){
	if t.Tok.Value != "" {
		fmt.Printf("%s%d: %s\n",tab,depth,t.Tok.Value)
	}
	if t.NodeCount <= 0{
		return
	}

	for _,node := range t.Nodes{
		prettyPrint(node,tab + "\t",depth + 1)
	}
}
