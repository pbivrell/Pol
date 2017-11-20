package lexer

import "testing"
import "../common"
import "../lexer"

type testpair struct {
	lex *lexer.Lexer
	tok common.Token
}

//Negative number constants don't exist
var numericTests = []testpair{
	{lexer.NewLexer("10"), common.Token{common.INTEGER_CONST, "10", 0}},
	{lexer.NewLexer("088"), common.Token{common.INTEGER_CONST, "088", 0}},
	{lexer.NewLexer("0"), common.Token{common.INTEGER_CONST, "0", 0}},
	{lexer.NewLexer("0.23"), common.Token{common.DECIMAL_CONST, "0.23", 0}},
	{lexer.NewLexer("100.3"), common.Token{common.DECIMAL_CONST, "100.3", 0}},
	{lexer.NewLexer("100.3.2"), common.Token{common.DECIMAL_CONST, "100.3", 0}},
	{lexer.NewLexer("A"), common.Token{common.ILLEGAL, "", 0}},
}

func TestGetNum(t *testing.T) {
	for _, pair := range numericTests {
		v := pair.lex.GetNum()
		if v.Type != pair.tok.Type && v.Value != pair.tok.Value {
			t.Error(
				"For", string(pair.lex.File),
				"expected", pair.tok.String(),
				"got", v.String(),
			)
		}
	}
}

var opTests = []testpair{
	{lexer.NewLexer("+"), common.Token{common.ADDITIVE_OP, "+", 0}},
	{lexer.NewLexer("-"), common.Token{common.ADDITIVE_OP, "-", 0}},
	{lexer.NewLexer("*"), common.Token{common.MULTIPLICATIVE_OP, "*", 0}},
	{lexer.NewLexer("/"), common.Token{common.MULTIPLICATIVE_OP, "/", 0}},
	{lexer.NewLexer("%"), common.Token{common.MULTIPLICATIVE_OP, "%", 0}},
	{lexer.NewLexer("^"), common.Token{common.MULTIPLICATIVE_OP, "^", 0}},
	{lexer.NewLexer("!"), common.Token{common.UNARY_OP, "!", 0}},
	{lexer.NewLexer("="), common.Token{common.ASSIGNMENT_OP, "=", 0}},
	{lexer.NewLexer("+="), common.Token{common.ASSIGNMENT_OP, "+=", 0}},
	{lexer.NewLexer("-="), common.Token{common.ASSIGNMENT_OP, "-=", 0}},
	{lexer.NewLexer("*="), common.Token{common.ASSIGNMENT_OP, "*=", 0}},
	{lexer.NewLexer("/="), common.Token{common.ASSIGNMENT_OP, "/=", 0}},
	{lexer.NewLexer("%="), common.Token{common.ASSIGNMENT_OP, "%=", 0}},
	{lexer.NewLexer("=="), common.Token{common.CONDITIONAL_OP, "==", 0}},
	{lexer.NewLexer("!="), common.Token{common.CONDITIONAL_OP, "!=", 0}},
	{lexer.NewLexer("<"), common.Token{common.CONDITIONAL_OP, "<", 0}},
	{lexer.NewLexer("<="), common.Token{common.CONDITIONAL_OP, "<=", 0}},
	{lexer.NewLexer(">"), common.Token{common.CONDITIONAL_OP, ">", 0}},
	{lexer.NewLexer(">="), common.Token{common.CONDITIONAL_OP, ">=", 0}},
	{lexer.NewLexer("&&"), common.Token{common.CONDITIONAL_OP, "&&", 0}},
	{lexer.NewLexer("||"), common.Token{common.CONDITIONAL_OP, "||", 0}},
	{lexer.NewLexer("&"), common.Token{common.REFRENCE_OP, "&", 0}},
	{lexer.NewLexer("{}"), common.Token{common.REFRENCE_OP, "{}", 0}},
	{lexer.NewLexer("[]"), common.Token{common.REFRENCE_OP, "[]", 0}},
	//Invalid Ops
	{lexer.NewLexer("@"), common.Token{common.ILLEGAL, "", 0}},
	{lexer.NewLexer("["), common.Token{common.ILLEGAL, "", 0}},
	{lexer.NewLexer("|"), common.Token{common.ILLEGAL, "", 0}},
	{lexer.NewLexer("("), common.Token{common.ILLEGAL, "", 0}},
	{lexer.NewLexer("APPLE"), common.Token{common.ILLEGAL, "", 0}},
}

func TestGetOp(t *testing.T) {
	for _, pair := range opTests {
		v := pair.lex.GetOp()
		if v.Type != pair.tok.Type && v.Value != pair.tok.Value {
			t.Error(
				"For", string(pair.lex.File),
				"expected", pair.tok.String(),
				"got", v.String(),
			)
		}
	}
}

var IDTests = []testpair{
	//keywords
	{lexer.NewLexer("if"), common.Token{common.IF, "if", 0}},
	{lexer.NewLexer("else"), common.Token{common.ELSE, "else", 0}},
	{lexer.NewLexer("for"), common.Token{common.FOR, "for", 0}},
	{lexer.NewLexer("while"), common.Token{common.WHILE, "while", 0}},
	{lexer.NewLexer("continue"), common.Token{common.CONTINUE, "continue", 0}},
	{lexer.NewLexer("break"), common.Token{common.BREAK, "break", 0}},
	{lexer.NewLexer("return"), common.Token{common.RETURN, "return", 0}},
	//random identifiers
	{lexer.NewLexer("all"), common.Token{common.IDENTIFIER, "all", 0}},
	{lexer.NewLexer("a123"), common.Token{common.IDENTIFIER, "a123", 0}},
	{lexer.NewLexer("a_"), common.Token{common.IDENTIFIER, "a_", 0}},
	{lexer.NewLexer("a__2"), common.Token{common.IDENTIFIER, "a__2", 0}},
	{lexer.NewLexer("a_if_2"), common.Token{common.IDENTIFIER, "a_if_2", 0}},
	//Legal followed by illegal
	{lexer.NewLexer("test\"Apple\""), common.Token{common.IDENTIFIER, "test", 0}},

	//Illegal
	{lexer.NewLexer("_a"), common.Token{common.ILLEGAL, "", 0}},
	{lexer.NewLexer("1abc"), common.Token{common.ILLEGAL, "", 0}},
}

func TestGetId(t *testing.T) {
	for _, pair := range IDTests {
		v := pair.lex.GetID()
		if v.Type != pair.tok.Type && v.Value != pair.tok.Value {
			t.Error(
				"For", string(pair.lex.File),
				"expected", pair.tok.String(),
				"got", v.String(),
			)
		}
	}
}

var StringTests = []testpair{
	{lexer.NewLexer("\"1abc\""), common.Token{common.STRING_CONST, "1abc", 0}},
	{lexer.NewLexer("\"@  & \n 9\""), common.Token{common.STRING_CONST, "@ & \n 9", 0}},
	{lexer.NewLexer("\"@  & \\n 9\""), common.Token{common.STRING_CONST, "@ & \\n 9", 0}},
	{lexer.NewLexer("\"\\\"A2\""), common.Token{common.STRING_CONST, "\"2", 0}},
	{lexer.NewLexer("\"This is a test string\""), common.Token{common.STRING_CONST, "This is a test string", 0}},

	//Illegal
	{lexer.NewLexer("This is not a string"), common.Token{common.ILLEGAL, "", 0}},
	{lexer.NewLexer("\"ENDLESS"), common.Token{common.ILLEGAL, "", 0}},
}

func TestGetString(t *testing.T) {
	for _, pair := range StringTests {
		v := pair.lex.GetString()
		if v.Type != pair.tok.Type && v.Value != pair.tok.Value {
			t.Error(
				"For", string(pair.lex.File),
				"expected", pair.tok.String(),
				"got", v.String(),
			)
		}
	}

}

var SpecialTests = []testpair{
	{lexer.NewLexer(","), common.Token{common.COMMA, ",", 0}},
	{lexer.NewLexer("("), common.Token{common.LEFT_PAREN, "(", 0}},
	{lexer.NewLexer(")"), common.Token{common.RIGHT_PAREN, ")", 0}},
	{lexer.NewLexer("["), common.Token{common.LEFT_BRACE, "[", 0}},
	{lexer.NewLexer("]"), common.Token{common.RIGHT_BRACE, "]", 0}},
	{lexer.NewLexer("{"), common.Token{common.LEFT_BRACKET, "{", 0}},
	{lexer.NewLexer("}"), common.Token{common.RIGHT_BRACKET, "}", 0}},
	{lexer.NewLexer(";"), common.Token{common.SEMI_COLON, ";", 0}},
	{lexer.NewLexer(";a"), common.Token{common.SEMI_COLON, "", 0}},
	//This is an operator because of the order of evaluation it will never be
	//identified as a special character
	{lexer.NewLexer("{}"), common.Token{common.LEFT_BRACKET, "{", 0}},
	//Illegal
	{lexer.NewLexer("a123"), common.Token{common.ILLEGAL, "", 0}},
}

func TestGetSpecial(t *testing.T) {
	for _, pair := range SpecialTests {
		v := pair.lex.GetSpecial()
		if v.Type != pair.tok.Type && v.Value != pair.tok.Value {
			t.Error(
				"For", string(pair.lex.File),
				"expected", pair.tok.String(),
				"got", v.String(),
			)
		}
	}
}

type testpair2 struct {
	lex *lexer.Lexer
	pos int
}

var BlockCommentTests = []testpair2{
	{lexer.NewLexer("/*T*/"), 5},
	{lexer.NewLexer("/*Test*/A"), 8},
	{lexer.NewLexer("/*Test comment that doesn't end \n Apple \n 1234"), 46},
}

func TestGetBlockComment(t *testing.T) {
	for _, pair := range BlockCommentTests {
		pair.lex.GetBlockComment()
		if pair.lex.Pos != pair.pos {
			t.Error(
				"For", string(pair.lex.File),
				"expected", pair.pos,
				"got", pair.lex.Pos,
			)
		}
	}
}

var LineCommentTests = []testpair2{
	{lexer.NewLexer("//Apple\n"), 8},
	{lexer.NewLexer("//Apple Never ENDING second line!@ \"\" 123\n"), 42},
}

func TestGetLineComment(t *testing.T) {
	for _, pair := range LineCommentTests {
		pair.lex.GetLineComment()
		if pair.lex.Pos != pair.pos {
			t.Error(
				"For", string(pair.lex.File),
				"expected", pair.pos,
				"got", pair.lex.Pos,
			)
		}
	}

}
