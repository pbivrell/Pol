package parser

import "fmt"
import "strconv"
import "os"

//Local libs
import "../common"

var Debug = 0
var HasErrors = false
var prevError = ""

//This is a hacky solution to my array indexing problem
//I don't like it but so be it
var index = 0

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
		return &common.Token{common.ILLEGAL, "ILLEGAL", toks.Toks[toks.pos-1].Lineno}, false
	}
	pos := toks.pos
	toks.pos = toks.pos + 1
	return &toks.Toks[pos], true
}

func (toks *Tokens) peek() (*common.Token, bool) {
	if toks.pos > toks.size-1 {
		return &common.Token{common.ILLEGAL, "ILLEGAL", toks.Toks[toks.pos-1].Lineno}, false
	}
	return &toks.Toks[toks.pos], true
}

func (toks *Tokens) prev() {
	if toks.pos <= 0 {
		return
	}
	toks.pos = toks.pos - 1
}

func expect(check common.TokenType, tokens *Tokens) bool {
	next, hasNext := tokens.peek()
	if !hasNext || check != next.Type {
		printError(next, "Expected "+string(check)+" found "+string(next.Type))
		return false
	}
	return true
}

func validate(token common.TokenType, tok *common.Token) bool {
	return token == tok.Type
}

func printError(tok *common.Token, msg string) {
	HasErrors = true

	error := fmt.Sprintf("Symantic Error: %s: Line %d: Token [%s]", msg, tok.Lineno, tok.Value)

	//This is a quick fix for infinite errors
	//TODO take second pass at error system
	if error == prevError {
		os.Exit(-1)
	}
	prevError = error

	fmt.Println(error)

}

func printDebug(message string, priority int) {
	if Debug >= priority {
		fmt.Println("[DEBUG] " + message)
	}
}

func dPeek(tokens *Tokens) string {
	ret, _ := tokens.peek()
	return "[ token " + ret.String() + " ]"
}

func printTreebug(tree *common.Tree, msg string, priority int) {
	printDebug(msg, priority)
	if Debug >= priority {
		tree.PrettyPrint()
	}
}

func createSrcTree() *common.Tree {
	tree := common.MakeTree(&common.Token{common.TREE_ROOT, "POL", 0})
	node_global := common.MakeTree(&common.Token{common.TREE_GLOBAL, "GLOBAL", 0})
	node_func := common.MakeTree(&common.Token{common.TREE_FUNC, "FUNC", 0})
	node_main := common.MakeTree(&common.Token{common.TREE_MAIN, "MAIN", 0})

	tree.Append(node_global)
	tree.Append(node_func)
	tree.Append(node_main)
	return tree
}

func createFuncTree(name *common.Token, args *common.Tree, body *common.Tree) *common.Tree {
	root := common.MakeTree(&common.Token{common.TREE_FUNC, name.Value, name.Lineno})
	root.Append(args)
	root.Append(body)
	return root
}

func Parse(tokens *Tokens) *common.Tree {
	tree := createSrcTree()

	for _, hasNext := tokens.peek(); hasNext; _, hasNext = tokens.peek() {
		node, t := compilation_unit(tokens)
		if t != -1 {
			tree.GetNode(t).Append(node)
		}
	}
	return tree
}

func listify(tokens *Tokens, fn func(*Tokens, bool) *common.Tree) []*common.Tree {
	printDebug("Entering listify "+dPeek(tokens), 5)
	defer printDebug("Exiting listify", 5)
	results := make([]*common.Tree, 0)
	first := fn(tokens, false)
	if first.IsInvalidTree() {
		return results
	}
	results = append(results, first)
	next, _ := tokens.peek()
	printDebug("This should be a comma: "+next.String(), 1)

	for next, hasNext := tokens.peek(); hasNext && next.Type == common.COMMA; next, hasNext = tokens.peek() {
		_, _ = tokens.next()
		results = append(results, fn(tokens, true))
	}

	printDebug("Printing Generated List", 9)
	if Debug >= 9 {
		for index, next := range results {
			fmt.Printf("Index %d\n", index)
			next.PrettyPrint()
		}
	}
	return results
}

func compilation_unit(tokens *Tokens) (*common.Tree, int) {
	printDebug("Entering compilation_unit"+dPeek(tokens), 5)
	defer printDebug("Exiting compilation_unit", 5)
	node := common.MakeInvalidTree()

	id, hasNext := tokens.peek()
	if !hasNext || id.Type != common.IDENTIFIER {
		printError(id, "Expected global variable declaration or function defination")
		return node, -1
	} else if node = funcDef(tokens); !node.IsInvalidTree() {
		if node.Tok.Value == "main" {
			return node, 2
		} else {
			return node, 1
		}
	} else if node = declaration(tokens); !node.IsInvalidTree() {
		return node, 0
	} else {
		return node, -1
	}
}

func funcDef(tokens *Tokens) *common.Tree {
	printDebug("Entering funcDef"+dPeek(tokens), 5)
	defer printDebug("Exiting funcDef", 5)
	if !expect(common.IDENTIFIER, tokens) {
		return common.MakeInvalidTree()
	}
	id, _ := tokens.next()

	next, _ := tokens.peek()
	if next.Type != common.LEFT_PAREN {
		tokens.prev()
		return common.MakeInvalidTree()
	}

	args := arguments(tokens)

	body := body(tokens)
	return createFuncTree(id, args, body)
}

func arguments(tokens *Tokens) *common.Tree {
	printDebug("Entering arguments"+dPeek(tokens), 5)
	defer printDebug("Exiting argument", 5)
	expect(common.LEFT_PAREN, tokens)
	_, _ = tokens.next()
	node := common.MakeTree(&common.Token{common.TREE_ARGS, "ARGS", -1})
	args := listify(tokens, listify_idOrRef)
	for _, next := range args {
		node = node.Append(next)
	}
	expect(common.RIGHT_PAREN, tokens)

	a, _ := tokens.next()
	fmt.Printf("TEST %s\n", a)
	printTreebug(node, "Arguments to function", 9)
	return node
}

func body(tokens *Tokens) *common.Tree {
	printDebug("Entering body"+dPeek(tokens), 5)
	defer printDebug("Exiting body", 5)

	if !expect(common.LEFT_BRACE, tokens) {
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	body := common.MakeTree(&common.Token{common.TREE_BODY, "BODY", -1})

	for next, hasNext := tokens.peek(); hasNext && next.Type != common.RIGHT_BRACE; next, hasNext = tokens.peek() {
		body = body.Append(single_body(tokens))
	}

	_, _ = tokens.next()
	printTreebug(body, "Body of function", 9)

	return body
}

func single_body(tokens *Tokens) *common.Tree {
	printDebug("Entering single_body"+dPeek(tokens), 5)
	defer printDebug("Exiting single_body", 5)

	first, _ := tokens.next()
	second, _ := tokens.peek()
	tokens.prev()

	if first.Type == common.IDENTIFIER && second.Type == common.LEFT_PAREN {
		return getFuncCall(tokens)
	} else if first.Type == common.IDENTIFIER {
		return declaration(tokens)
	} else if first.Type == common.IF {
		return getIfStatement(tokens)
	} else if first.Type == common.FOR {
		return getForStatement(tokens)
	} else if first.Type == common.WHILE {
		return getWhileStatment(tokens)
	} else if node := getControlFlow(tokens); !node.IsInvalidTree() {
		return node
	}

	return common.MakeInvalidTree()
}

func getIfStatement(tokens *Tokens) *common.Tree {
	printDebug("Entering getIfStatement"+dPeek(tokens), 5)
	defer printDebug("Exiting getIfStatement", 5)

	if !expect(common.IF, tokens) {
		return common.MakeInvalidTree()
	}

	ifTok, _ := tokens.next()

	if !expect(common.LEFT_PAREN, tokens) {
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	expr := expression(tokens)

	if !expect(common.RIGHT_PAREN, tokens) {
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	body := body(tokens)

	next, _ := tokens.peek()

	node := common.MakeTree(ifTok)
	node = node.Append(expr)
	node = node.Append(body)

	if next.Type != common.ELSE {
		node = node.Append(common.MakeInvalidTree())
		return node
	}

	node = node.Append(getElseStatment(tokens))
	return node
}

func getElseStatment(tokens *Tokens) *common.Tree {
	if !expect(common.ELSE, tokens) {
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	isIf, _ := tokens.peek()

	if isIf.Type == common.IF {
		return getIfStatement(tokens)
	}

	return body(tokens)
}

func getForStatement(tokens *Tokens) *common.Tree {
	printDebug("Entering forStatement"+dPeek(tokens), 5)
	defer printDebug("Exiting forStatement", 5)
	if !expect(common.FOR, tokens) {
		return common.MakeInvalidTree()
	}

	forTok, _ := tokens.next()

	node := common.MakeTree(forTok)

	if !expect(common.LEFT_PAREN, tokens) {
		return common.MakeInvalidTree()
	}
	_, _ = tokens.next()

	init := forInit(tokens)

	if !expect(common.SEMI_COLON, tokens) {
		return common.MakeInvalidTree()
	}
	_, _ = tokens.next()

	cond := forExpress(tokens)
	if !expect(common.SEMI_COLON, tokens) {
		return common.MakeInvalidTree()
	}
	_, _ = tokens.next()

	incr := forInit(tokens)

	if !expect(common.RIGHT_PAREN, tokens) {
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	body := body(tokens)

	node = node.Append(init)
	node = node.Append(cond)
	node = node.Append(incr)
	node = node.Append(body)

	return node
}

func forInit(tokens *Tokens) *common.Tree {
	printDebug("Entering forInit"+dPeek(tokens), 5)
	defer printDebug("Exiting forInit", 5)

	next, _ := tokens.peek()
	if next.Type == common.SEMI_COLON {
		return common.MakeTree(&common.Token{common.TREE_INIT_LIST, "INIT_LIST", -1})
	}
	return var_declaration(tokens)
}

func forExpress(tokens *Tokens) *common.Tree {
	printDebug("Entering forExpress"+dPeek(tokens), 5)
	defer printDebug("Exiting forExpress", 5)
	next, _ := tokens.peek()
	if next.Type == common.SEMI_COLON {
		return common.MakeTree(&common.Token{common.ZERO_VALUE, "EMPTY", -1})
	}
	return expression(tokens)
}

func getWhileStatment(tokens *Tokens) *common.Tree {
	printDebug("Entering whileStatment"+dPeek(tokens), 5)
	defer printDebug("Exiting whileStatment", 5)

	if !expect(common.WHILE, tokens) {
		return common.MakeInvalidTree()
	}

	whileTok, _ := tokens.next()

	if !expect(common.LEFT_PAREN, tokens) {
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	cond := expression(tokens)

	if !expect(common.RIGHT_PAREN, tokens) {
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	body := body(tokens)

	node := common.MakeTree(whileTok)

	node = node.Append(cond)
	node = node.Append(body)

	return node
}

func getControlFlow(tokens *Tokens) *common.Tree {
	printDebug("Entering getControlFlow"+dPeek(tokens), 5)
	defer printDebug("Exiting getControlFlow", 5)

	nodeTok, _ := tokens.peek()
	if nodeTok.Type == common.CONTINUE {
		_, _ = tokens.next()
		return common.MakeTree(&common.Token{common.CONTINUE, "CONTINUE", -1})
	} else if nodeTok.Type == common.BREAK {
		_, _ = tokens.next()
		return common.MakeTree(&common.Token{common.BREAK, "BREAK", -1})
	} else if nodeTok.Type == common.RETURN {
		retVal := common.MakeTree(&common.Token{common.RETURN, "RETURN", -1})
		_, _ = tokens.next()
		expr := expression(tokens)
		retVal = retVal.Append(expr)
		return retVal
	}

	return common.MakeInvalidTree()
}

func listify_idOrRef(tokens *Tokens, verbose bool) *common.Tree {
	printDebug("Entering listify_idOrRef"+dPeek(tokens), 5)
	defer printDebug("Exiting listify_idOrRef", 5)
	err := "Expecting identifier or reference"
	token, _ := tokens.peek()
	node := common.MakeTree(token)
	if token.Type == common.IDENTIFIER {
		_, _ = tokens.next()
		return node
	} else if token.Type == common.REFRENCE_OP {
		_, _ = tokens.next()
		next, _ := tokens.next()
		if next.Type != common.IDENTIFIER {
			err = "identifier expected after reference op"
		} else {
			return node.Append(common.MakeTree(next))
		}
	}
	if verbose {
		printError(token, err)
	}
	return common.MakeInvalidTree()
}

func declaration(tokens *Tokens) *common.Tree {
	printDebug("Entering declaration"+dPeek(tokens), 5)
	defer printDebug("Exiting declaration", 5)
	node := common.MakeInvalidTree()
	if !expect(common.IDENTIFIER, tokens) {
		return node

	} else if node = collection_assignment(tokens); !node.IsInvalidTree() {

	} else if node = collection_declaration(tokens); !node.IsInvalidTree() {

	} else if node = var_declaration(tokens); !node.IsInvalidTree() {

	} else {

	}
	return node
}

func collection_assignment(tokens *Tokens) *common.Tree {
	printDebug("Entering collection_assignment"+dPeek(tokens), 5)
	defer printDebug("Exiting collection_assignment", 5)
	if !expect(common.IDENTIFIER, tokens) {
		return common.MakeInvalidTree()
	}
	_, _ = tokens.next()
	op, _ := tokens.peek()

	if !(op.Type == common.LEFT_BRACE || op.Type == common.LEFT_BRACKET) {
		tokens.prev()
		return common.MakeInvalidTree()
	}

	tokens.prev()

	lookUpNode := getLookUp(tokens)

	if !expect(common.ASSIGNMENT_OP, tokens) {
		return common.MakeInvalidTree()
	}
	_, _ = tokens.next()

	rhs := expression(tokens)

	result := common.MakeTree(&common.Token{common.ASSIGNMENT_OP, "=", -1})
	result = result.Append(lookUpNode)
	result = result.Append(rhs)

	return result

}

func collection_declaration(tokens *Tokens) *common.Tree {
	printDebug("Entering collection_declaration"+dPeek(tokens), 5)
	defer printDebug("Exiting collection_declaration", 5)
	if !expect(common.IDENTIFIER, tokens) {
		return common.MakeInvalidTree()
	}
	id, _ := tokens.next()
	op, _ := tokens.peek()

	var TREE_TYPE common.TokenType = common.HASH
	LISTIFY := listify_hash
	var LIST_START common.TokenType = common.LEFT_BRACE
	var LIST_END common.TokenType = common.RIGHT_BRACE
	if op.Type == common.ARRAY_OP {
		TREE_TYPE = common.ARRAY
		LISTIFY = listify_array
		LIST_START = common.LEFT_BRACKET
		LIST_END = common.RIGHT_BRACKET
	} else if op.Type != common.HASH_OP {
		tokens.prev()
		return common.MakeInvalidTree()
	}

	idNode := common.MakeTree(&common.Token{TREE_TYPE, id.Value, id.Lineno})
	_, _ = tokens.next()

	next, _ := tokens.peek()
	if next.Type != common.ASSIGNMENT_OP {
		lhs := make([]*common.Tree, 0)
		lhs = append(lhs, idNode)
		return createInitalizationTree(lhs, make([]*common.Tree, 0), &common.Token{common.ASSIGNMENT_OP, "=", -1})
	}
	_, _ = tokens.next()

	if !expect(LIST_START, tokens) {
		//Invalid but we successfully attempted to eat an init list it was just malformed
		return common.MakeTree(&common.Token{common.TREE_TEMP, "BAD", -1})
	}

	_, _ = tokens.next()

	pair := listify(tokens, LISTIFY)
	index = 0

	if !expect(LIST_END, tokens) {
		//Invalid but we successfully attempted to eat an init list it was just malformed
		return common.MakeTree(&common.Token{common.TREE_TEMP, "BAD", -1})
	}

	_, _ = tokens.next()

	lhs := make([]*common.Tree, 0)
	rhs := make([]*common.Tree, 0)

	for _, next := range pair {
		single_lhs, single_rhs := createCollectionPair(next, idNode)
		lhs = append(lhs, single_lhs)
		rhs = append(rhs, single_rhs)
	}

	return createInitalizationTree(lhs, rhs, &common.Token{common.ASSIGNMENT_OP, "=", -1})

}

func createCollectionPair(pair *common.Tree, id *common.Tree) (*common.Tree, *common.Tree) {
	printDebug("Entering createCollectionPair", 5)
	defer printDebug("Exiting createCollectionPair", 5)

	children := pair.GetChildren()
	printDebug("Key: "+children[0].String(), 5)
	lhs := createLookUpTree(id, &children[0])
	printDebug("Value: "+children[1].String(), 5)
	rhs := &children[1]
	return lhs, rhs
}

func listify_hash(tokens *Tokens, verbose bool) *common.Tree {
	printDebug("Entering listify_hash"+dPeek(tokens), 5)
	defer printDebug("Exiting listify_hash", 5)

	if !expect(common.LEFT_BRACKET, tokens) {
		return common.MakeInvalidTree()
	}
	_, _ = tokens.next()

	key := expression(tokens)

	if !expect(common.COMMA, tokens) {
		return common.MakeInvalidTree()
	}
	_, _ = tokens.next()

	value := expression(tokens)

	if !expect(common.RIGHT_BRACKET, tokens) {
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	node := common.MakeTree(&common.Token{common.TREE_TEMP, "TEMP", -1})
	node = node.Append(key)
	node = node.Append(value)

	return node
}

func listify_array(tokens *Tokens, verbose bool) *common.Tree {
	printDebug("Entering listify_array"+dPeek(tokens), 5)
	defer printDebug("Exiting listify_array", 5)

	value := expression(tokens)

	node := common.MakeTree(&common.Token{common.TREE_TEMP, "TEMP", -1})
	node = node.Append(common.MakeTree(&common.Token{common.INDEX, strconv.Itoa(index), -1}))
	node = node.Append(value)
	index = index + 1
	return node
}

func var_declaration(tokens *Tokens) *common.Tree {
	printDebug("Entering var_declaration"+dPeek(tokens), 5)
	defer printDebug("Exiting var_declaration", 5)
	if !expect(common.IDENTIFIER, tokens) {
		return common.MakeInvalidTree()
	}
	lhs := listify(tokens, listify_identifier)

	op, _ := tokens.peek()
	printDebug("This should not be an identifier: "+op.String(), 1)
	if op.Type == common.ASSIGNMENT_OP {
		_, _ = tokens.next()
		rhs := listify(tokens, listify_expression)
		return createInitalizationTree(lhs, rhs, op)
	}
	return createInitalizationTree(lhs, make([]*common.Tree, 0), &common.Token{common.ASSIGNMENT_OP, "=", -1})
}

func createInitalizationTree(lhs []*common.Tree, rhs []*common.Tree, op *common.Token) *common.Tree {
	printDebug("Entering createInitalizationTree", 5)
	defer printDebug("Exiting createInitalizationTree", 5)
	node := common.MakeTree(&common.Token{common.TREE_INIT_LIST, "INIT_LIST", -1})

	if len(lhs) < len(rhs) {
		printError(op, "More values on right hand side of '"+op.Value+"' then variables on the left")
		return common.MakeInvalidTree()
	} else {
		for i, next := range lhs {
			single_asgn := common.MakeTree(op)
			single_asgn = single_asgn.Append(next)
			if i < len(rhs) {
				single_asgn = single_asgn.Append(rhs[i])
			} else {
				single_asgn = single_asgn.Append(common.MakeTree(&common.Token{common.ZERO_VALUE, "EMPTY", -1}))
			}
			node.Append(single_asgn)
		}
	}
	printTreebug(node, "CREATED INIT TREE", 9)
	return node
}

func listify_identifier(tokens *Tokens, verbose bool) *common.Tree {
	printDebug("Entering listify_identifier"+dPeek(tokens), 5)
	defer printDebug("Exiting listify_identifier", 5)
	next, hasNext := tokens.peek()
	if !hasNext || next.Type != common.IDENTIFIER {
		if verbose {
			expect(common.IDENTIFIER, tokens)
		}
		return common.MakeInvalidTree()
	} else {
		_, _ = tokens.next()
		return common.MakeTree(next)
	}
}

func listify_expression(tokens *Tokens, verbose bool) *common.Tree {
	printDebug("Entering listify_expression"+dPeek(tokens), 5)
	defer printDebug("Exiting listify_expression", 5)
	return expression(tokens)
}

func expression(tokens *Tokens) *common.Tree {
	printDebug("Entering expression"+dPeek(tokens), 5)
	defer printDebug("Exiting expression", 5)
	return cond_expression(tokens)
}

func cond_expression(tokens *Tokens) *common.Tree {
	printDebug("Entering cond_expression"+dPeek(tokens), 5)
	defer printDebug("Exiting cond_expression", 5)
	return eval_expression(tokens, add_expression, common.CONDITIONAL_OP)
}

func add_expression(tokens *Tokens) *common.Tree {
	printDebug("Entering add_expression"+dPeek(tokens), 5)
	defer printDebug("Exiting add_expression", 5)
	return eval_expression(tokens, mult_expression, common.ADDITIVE_OP)
}

func mult_expression(tokens *Tokens) *common.Tree {
	printDebug("Entering mult_expression"+dPeek(tokens), 5)
	defer printDebug("Exiting mult_expression", 5)
	return eval_expression(tokens, atom, common.MULTIPLICATIVE_OP)
}

func atom(tokens *Tokens) *common.Tree {
	printDebug("Entering atom"+dPeek(tokens), 5)
	defer printDebug("Exiting atom", 5)
	first, _ := tokens.peek()
	node := common.MakeInvalidTree()

	if node = getConstant(tokens); !node.IsInvalidTree() {

	} else if node = getParens(tokens); !node.IsInvalidTree() {

	} else if node = getUnary(tokens); !node.IsInvalidTree() {

	} else if node = getIdentifier(tokens); !node.IsInvalidTree() {

	} else {
		printError(first, "Expected identifier, primative, unary expression, or parens")
	}
	return node
}

func getLookUp(tokens *Tokens) *common.Tree {
	printDebug("Entering getLookUp"+dPeek(tokens), 5)
	defer printDebug("Exiting getLookUp", 5)
	id, _ := tokens.next()
	next, _ := tokens.next()

	var closing common.TokenType = common.RIGHT_BRACKET
	var TYPE common.TokenType = common.ARRAY
	if next.Type == common.LEFT_BRACE {
		closing = common.RIGHT_BRACE
		TYPE = common.HASH
	}
	where := common.MakeTree(&common.Token{TYPE, id.Value, -1})

	value := expression(tokens)
	if !expect(closing, tokens) {
		return common.MakeInvalidTree()
	}
	_, _ = tokens.next()
	return createLookUpTree(where, value)

}

func createLookUpTree(where *common.Tree, value *common.Tree) *common.Tree {
	lookUpNode := common.MakeTree(&common.Token{common.TREE_LOOK_UP, "LOOK UP", -1})
	lookUpNode = lookUpNode.Append(where)
	lookUpNode = lookUpNode.Append(value)
	return lookUpNode
}

func getIdentifier(tokens *Tokens) *common.Tree {
	printDebug("Entering getIdentifier"+dPeek(tokens), 5)
	defer printDebug("Exiting getIdentifier", 5)
	token, _ := tokens.peek()
	if token.Type == common.IDENTIFIER {
		id, _ := tokens.next()
		node := common.MakeTree(id)
		next, _ := tokens.peek()

		if next.Type == common.LEFT_BRACE || next.Type == common.LEFT_BRACKET {
			tokens.prev()
			return getLookUp(tokens)
		} else if next.Type == common.LEFT_PAREN {
			tokens.prev()
			return getFuncCall(tokens)
		}
		return node
	}
	return common.MakeInvalidTree()
}

func getFuncCall(tokens *Tokens) *common.Tree {
	printDebug("Entering getFuncCall"+dPeek(tokens), 5)
	defer printDebug("Exiting getFuncCall", 5)

	if !expect(common.IDENTIFIER, tokens) {
		return common.MakeInvalidTree()
	}

	id, _ := tokens.next()

	if !expect(common.LEFT_PAREN, tokens) {
		tokens.prev()
		return common.MakeInvalidTree()
	}

	_, _ = tokens.next()

	arguments := listify(tokens, listify_expression)

	argsNode := common.MakeTree(&common.Token{common.TREE_ARGS, "ARGS", -1})
	for _, next := range arguments {
		argsNode = argsNode.Append(next)
	}

	if !expect(common.RIGHT_PAREN, tokens) {
		return common.MakeTree(&common.Token{common.TREE_TEMP, "BAD", -1})
	}

	_, _ = tokens.next()

	node := common.MakeTree(id)
	node = node.Append(argsNode)

	return node

}

func getUnary(tokens *Tokens) *common.Tree {
	printDebug("Entering getUnary"+dPeek(tokens), 5)
	defer printDebug("Exiting getUnary", 5)
	token, _ := tokens.peek()
	if (token.Value == "-" || token.Value == "!") && token.Type != common.ILLEGAL {
		next, _ := tokens.next()
		node := common.MakeTree(next)
		subNode := expression(tokens)
		node = node.Append(subNode)
		return node
	}
	return common.MakeInvalidTree()
}

func getParens(tokens *Tokens) *common.Tree {
	printDebug("Entering getParens"+dPeek(tokens), 5)
	defer printDebug("Exiting getParens", 5)
	token, _ := tokens.peek()
	if token.Type == common.LEFT_PAREN {
		_, _ = tokens.next()
		node := expression(tokens)
		if !expect(common.RIGHT_PAREN, tokens) {
			return common.MakeInvalidTree()
		}
		_, _ = tokens.next()
		return node
	}
	return common.MakeInvalidTree()
}

func getConstant(tokens *Tokens) *common.Tree {
	printDebug("Entering getConstant"+dPeek(tokens), 5)
	defer printDebug("Exiting getConstant", 5)
	token, _ := tokens.peek()
	if token.Type == common.STRING_CONST ||
		token.Type == common.INTEGER_CONST ||
		token.Type == common.DECIMAL_CONST ||
		token.Type == common.TRUE ||
		token.Type == common.FALSE {
		_, _ = tokens.next()
		return common.MakeTree(token)
	}

	return common.MakeInvalidTree()
}

func eval_expression(tokens *Tokens, fn func(*Tokens) *common.Tree, tokenType common.TokenType) *common.Tree {
	printDebug("Entering eval_expression"+dPeek(tokens), 5)
	defer printDebug("Exiting eval_expression", 5)
	token, hasNext := tokens.peek()

	if !hasNext {
		printError(token, string(tokenType)+" expected")
		return common.MakeInvalidTree()
	}

	node := fn(tokens)

	for token, hasNext := tokens.peek(); hasNext && token.Type == tokenType; token, hasNext = tokens.peek() {
		temp := node
		node = common.MakeTree(token)
		node = node.Append(temp)
		_, _ = tokens.next()
		subTree := fn(tokens)
		node = node.Append(subTree)
	}
	return node
}
