package interpreter

import "strings"
import "../common"
import "../pol_type"
import "fmt"
import "reflect"
import "../go_pol"

var funcTable FuncTable
var scope *Scope

var Debug = 0

func Interpret(syntaxTree *common.Tree){

    funcTable = MakeFuncTable()
    scope = MakeScope()

    AddLibararyFunction("print","print",[]Arg{Arg{false,"a_pol_variable"}})
    AddLibararyFunction("println","println",[]Arg{Arg{false,"a_pol_variable"}})
    AddLibararyFunction("flush","flush",[]Arg{})

    for _, subTree := range syntaxTree.GetChildren(){
        if subTree.Tok.Type == common.TREE_FUNC {
            funcTable.Add(&subTree)
        }
    }

    for _, subTree := range syntaxTree.GetChildren(){
        if subTree.Tok.Type != common.TREE_FUNC {
            execAssignment(&subTree)
        }
    }


    //fmt.Println(funcTable.String())
    main := funcTable.Call("main",0)
    if main == nil {
        printError("Could not find main function")
    }
    execFunc(main, nil)

    go_pol.WriteBuffer()
    //fmt.Println((*funcTable).String())
}

func AddLibararyFunction(name string, where string, params []Arg) {
	body := common.MakeTree(&common.Token{common.SHELL_CODE, where, -1})
    funcTable.Insert(name,len(params),&funcCall{params,body})
}

func printError(msg string){
    common.RuntimeError(fmt.Sprintf("[Runtime Error] %s\n",msg))
}

func execFunc(call *funcCall, callArgs []common.Tree) pol_type.Pol_Type{
    new, old := scope.Call()

    types := make([]pol_type.Pol_Type,len(callArgs))

    for i, arg := range callArgs {
        types[i] = execExpression(&arg)
    }

    scope = new

    for i, item := range types {
        scope.Assign(call.Args[i].Name, item)
    }

    defer exit(old, call, callArgs)
    //call.Body.PrettyPrint()
    //TODO
    //Evaulate args of func_call

    return execBody(call.Body)
}

func exit(old *Scope, call *funcCall, callArgs []common.Tree) {

    types := make([]pol_type.Pol_Type,len(call.Args))

    for i, arg := range call.Args {
        if arg.Ref {
            types[i] = scope.ValueOf(call.Args[i].Name)
        }
    }
    scope = old

    for i, arg := range call.Args {
        if arg.Ref {
            scope.Assign(callArgs[i].Tok.Value, types[i])
        }
    }
}

func execBody(subTree *common.Tree) pol_type.Pol_Type {
    if subTree.Tok.Type == common.TREE_BODY {
        for _, next := range subTree.GetChildren(){
            if ret := execBody(&next); ret.GetType() != pol_type.NOTHING{
                return ret
            }
        }
    }
    current := subTree.Tok
    //fmt.Println(current)
    //subTree.PrettyPrint()
    if current.Type == common.RETURN {
        return execExpression(&subTree.GetChildren()[0])
    }else if current.Type == common.IF{
        return execIf(subTree)
    }else if current.Type == common.FOR {
        return execForLoop(subTree)
    }else if current.Type == common.WHILE {
        return execWhileLoop(subTree)
    }else if current.Type == common.TREE_INIT_LIST || current.Value == "=" {
        execAssignment(subTree)
        //Change this to anything else
        return pol_type.NewNothing()
    }else if current.Type == common.FUNC_CALL {
        call := funcTable.Call(current.Value,len(subTree.GetChildren()))
        if call == nil {
            printError(fmt.Sprintf("Unable to call function %s with %d arguments", current.Value, len(subTree.GetChildren())))
        }
        return execFunc(call, subTree.GetChildren())
    }else if current.Type == common.TREE_BODY {

    }else if current.Type == common.SHELL_CODE {
        types := scope.Dump()
        //Change this so that it is called variadiaclly
        args := reflect.ValueOf(types)
        var t go_pol.Go_Pol
        reflect.ValueOf(&t).MethodByName(strings.Title(current.Value)).Call([]reflect.Value{args})

    }else{
        printError(fmt.Sprintf("Unsupported Thing maybe: %s", current))
        subTree.PrettyPrint()
    }
    return pol_type.NewNothing()

}

type LRpair struct{
    lhs string
    rhs pol_type.Pol_Type
    isLookUp bool
}

func evalAssignment(assignmentTree *common.Tree) LRpair{
    if assignmentTree.Tok.Value != "=" {
        fmt.Printf("Called evaluateAssignment with a node that was not = : %s\n", assignmentTree.Tok)
    }

    assignExpr := assignmentTree.GetChildren()

    return LRpair{assignExpr[0].Tok.Value, execExpression(&assignExpr[1]),false}
}

type CollectionPair struct{
    structure string
    pol_type pol_type.Pol_Type
}

func execAssignment(assignmentTree *common.Tree){
    items := len(assignmentTree.GetChildren())

    pairs := make([]LRpair,items)
    collections := make([]CollectionPair,items)

    count := 0
    for i, next := range assignmentTree.GetChildren(){
        if next.GetChildren()[0].Tok.Type == common.TREE_LOOK_UP {
            pairs[i] = LRpair{"", execExpression(&next.GetChildren()[1]), true }
            children := next.GetChildren()[0].GetChildren()
            collections[count] = CollectionPair{children[0].Tok.Value, execExpression(&children[1]) }
            count++
        }else{
            pairs[i] = evalAssignment(&next)
        }
    }

    count = 0
    for _, next := range pairs {
        if next.isLookUp {
            collection := collections[count]
            count++
            scope.AssignCollection(collection.structure, collection.pol_type, next.rhs)
        }else{
            scope.Assign(next.lhs,next.rhs)
        }
    }

    /*if lhs := common.LOOK_UP {
        where := lhs.GetChildren()[0].Tok.Value
        key := execExpression(lhs.GetChildren()[1])
        scope.AssignCollection(where,key,rhs)
    }else{
        scope.Assign(lhs.Tok.Value,rhs)
    }*/
}

func execIf(ifTree *common.Tree) pol_type.Pol_Type {
    //scope.Enter()
    //defer scope.Exit()
    ifExpr := execExpression(&ifTree.GetChildren()[0])
    if ifExpr.GetType() != pol_type.BOOL {
        printError(fmt.Sprintf("if condition at %d can not be evaluated in boolean context",ifTree.Tok.Lineno))
    }
    condition := ifExpr.(pol_type.Pol_Bool)
    if condition.GetValue(){
        return execBody(&ifTree.GetChildren()[1])
    } else if !(&ifTree.GetChildren()[2]).IsInvalidTree(){
        return execBody(&ifTree.GetChildren()[2])
    }
    return pol_type.NewNothing()
}

func execForLoop(loopTree *common.Tree) pol_type.Pol_Type{
    forExpr := loopTree.GetChildren()
    execAssignment(&forExpr[0])

    FOR_LOOP:
    cond := execExpression(&forExpr[1])
    if cond.GetType() != pol_type.BOOL {
        printError(fmt.Sprintf("for condition at %d can not be evaluated in boolean context",loopTree.Tok.Lineno))
    }

    condition := cond.(pol_type.Pol_Bool)
    if condition.GetValue() {
        if retVal := execBody(&forExpr[3]); retVal.GetType() != pol_type.NOTHING {
            return retVal
        }
        execAssignment(&forExpr[2])
        goto FOR_LOOP
    }

    return pol_type.NewNothing()
}

func execWhileLoop(whileTree *common.Tree) pol_type.Pol_Type{
    whileExpr := whileTree.GetChildren()
    WHILE_LOOP:
    cond := execExpression(&whileExpr[0])
    if cond.GetType() != pol_type.BOOL {
        printError(fmt.Sprintf("while condition at %d can not be evaluated in boolean context",whileTree.Tok.Lineno))
    }

    condition := cond.(pol_type.Pol_Bool)

    if condition.GetValue() {
        if retVal := execBody(&whileExpr[1]); retVal.GetType() != pol_type.NOTHING {
            return retVal
        }
        goto WHILE_LOOP
    }
    return pol_type.NewNothing()
}

func execExpression(exprTree *common.Tree) pol_type.Pol_Type{
    current := exprTree.Tok

    //fmt.Println("EXPR: ", current)
    if current.Type == common.VARIABLE {
        vari := scope.ValueOf(current.Value)
        if vari == nil {
            printError(fmt.Sprintf("Use of uninitallized value in variable %s",current.Value))
        }
        return vari
    }else if current.Type == common.STRING_CONST {
        return pol_type.NewString(current.Value)
    } else if current.Type == common.INTEGER_CONST {
        return pol_type.NewInt(current.Value)
    } else if current.Type == common.DECIMAL_CONST {
        return pol_type.NewRational(current.Value)
    } else if current.Type == common.TREE_LOOK_UP {
        return scope.LookUp(exprTree.GetChildren()[0].Tok.Value, execExpression(&exprTree.GetChildren()[1]))
    } else if current.Type == common.TRUE || current.Type == common.FALSE{
        return pol_type.NewBool(current.Type == common.TRUE)
    } else if isOperator(exprTree){
        return execOperator(exprTree)
    } else if current.Type == common.FUNC_CALL {
        call := funcTable.Call(current.Value,len(exprTree.GetChildren()))
        if call == nil {
            printError(fmt.Sprintf("Unable to call function %s with %d arguments", current.Value, len(exprTree.GetChildren())))
        }
        return execFunc(call, exprTree.GetChildren())

    } else {
        //fmt.Println("SOMETHING WENT WRONG")
    }
    return pol_type.NewNothing()
}

func isOperator(tree *common.Tree) bool {
    switch tree.Tok.Type {
    case common.ASSIGNMENT_OP: return true
    case common.ADDITIVE_OP: return true
    case common.MULTIPLICATIVE_OP: return true
    case common.CONDITIONAL_OP: return true
    case common.UNARY_OP: return true
    default: return false
    }
}

func execOperator(tree *common.Tree) pol_type.Pol_Type {
    args := make([]pol_type.Pol_Type,0)
    for _, next := range tree.GetChildren(){
        args = append(args,execExpression(&next))
    }
    args = promoteType(args)

    switch tree.Tok.Value {
    case "+": return args[0].Add(args[1])
    case "-": return args[0].Subtract(args[1])
    case "/": return args[0].Divide(args[1])
    case "*": return args[0].Multiply(args[1])
    case "^": return args[0].Exponent(args[1])
    case "%": return args[0].Mod(args[1])
    case "==": return args[0].Equals(args[1])
    case "!=": return args[0].NotEquals(args[1])
    case "<": return args[0].Less(args[1])
    case ">": return args[0].Greater(args[1])
    case "<=": return args[0].LessEquals(args[1])
    case ">=": return args[0].GreaterEquals(args[1])
    case "!": return args[0].Not()
    case "||": return args[0].Or(args[1])
    case "&&": return args[0].And(args[1])
    default:
        printError("Undefined operator " + tree.Tok.Value)
        return pol_type.NewNothing()
    }
}

func promoteType(args []pol_type.Pol_Type) []pol_type.Pol_Type {
    if len(args) == 2{
        if args[0].GetType() == args[1].GetType(){

        }else if converted := args[1].Convert(args[0]); converted.GetType() != pol_type.NOTHING {
            args[1] = converted
        }else if converted := args[0].Convert(args[1]); converted.GetType() != pol_type.NOTHING {
            args[0] = converted
        }else{
            printError("Neither argument could be converted to the other")
        }
    }
    return args
}
