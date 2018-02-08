package interpreter

import "fmt"
import "../common"

type funcCall struct {
    Args []Arg
    Body *common.Tree
}

type Arg struct {
    Ref bool
    Name string
}

type FuncTable map[string]map[int]*funcCall

func (f funcCall) String() string{
    result := "["
    for _, next := range f.Args{
        result = result + next.Name
        if next.Ref {
            result = result + "*"
        }
        result = result  + ","
    }
    result += f.Body.GetChildren()[0].Tok.String()
    result = result + "]"
    return result
}

func MakeFuncTable() FuncTable{
    return new(FuncTable).Init()
}

func (f FuncTable) Call(name string, argCount int) *funcCall {
    item, hasItem := f[name][argCount]
    if hasItem {
        return item
    }
    return nil
}

func (f FuncTable) Init() FuncTable {
    var pointer_f FuncTable = make(map[string]map[int]*funcCall)
    return pointer_f
}

func (f FuncTable) String() string{
    result := ""
    for key, value := range f{
        result = result + key
        for key2, value2 := range value {
            result = result + "[" + fmt.Sprintf("%d",key2) + ":" + value2.String() + "], "
        }
        result = result + "\n"
    }
    return result
}

func (f FuncTable) Insert(name string, argsCount int, call* funcCall){
    //fmt.Println("A")
    //(*call).Body.PrettyPrint()

    if f == nil {
        f = make(map[string]map[int]*funcCall)
    }
    _, hasItem := f[name]
	if hasItem {
		_, hasItem := f[name][argsCount]
		if hasItem {
			fmt.Printf("[Error] Func %s with %d arguments has already been declared\n",name,argsCount)
		}else{
			f[name][argsCount] = call
		}
	}else{
		f[name] = make(map[int]*funcCall)
		f[name][argsCount] = call
	}
}

func (f FuncTable) Add(funcTree *common.Tree){
    current := funcTree.Tok
    if current.Type != common.TREE_FUNC {
        fmt.Printf("Called addFunction with a node that was not FUNC: %s\n", current)
    }
    children := funcTree.GetChildren()

    args := &children[0]
    body := &children[1]
    argsList, count := evalArgs(args)

    f.Insert(current.Value, count, &funcCall{argsList, body})
}

func evalArgs(argsTree *common.Tree) ([]Arg, int){
    current := argsTree.Tok
    if current.Type != common.TREE_ARGS {
        fmt.Printf("Called evalArgs with a node that was not ARGS: %s\n", current)
    }

    argTree := argsTree.GetChildren()
    count := len(argTree)
    args := make([]Arg,count)

    for index, next := range argTree{
        if next.Tok.Type == common.REFRENCE_OP {
            args[index] = Arg{true,next.GetChildren()[0].Tok.Value}
        }else{
            args[index] = Arg{false,next.Tok.Value}
        }
    }
    return args,count
}
