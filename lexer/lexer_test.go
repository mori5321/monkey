package lexer

import (
	"testing"

	"example.com/monkey/token"
)

type Expected struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func runTokenTestCase(t *testing.T, input string, cs []Expected) {
	l := New(input)

	for i, expected := range cs {
		tok := l.NextToken()

		verifyToken(t, i, tok, expected)
	}
}

func verifyToken(t *testing.T, index int, tok token.Token, expected Expected) {
	if tok.Type != expected.expectedType {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", index, expected.expectedType, tok.Type)
	}

	if tok.Literal != expected.expectedLiteral {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", index, expected.expectedLiteral, tok.Literal)
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []Expected{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runTokenTestCase(t, input, tests)
}

func TestNextToken2(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;
		let add = fn(x, y) {
			x + y;
		}

		let result = add(five, ten);

		if (5 < 10) {
			return true
		} else {
			return false
		}
	`

	tests := []Expected{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
	}

	runTokenTestCase(t, input, tests)
}

// If it's valid as sentence, lexer just parse it as token.
func TestNextToken3(t *testing.T) {
	input := `
		!-/*5;
		5 < 10 > 5;

		10 == 10;
		10 != 9;
	`

	tests := []Expected{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}

	runTokenTestCase(t, input, tests)
}

func TestNextToken4(t *testing.T) {
	input := `
		"foobar";
		"foo bar";
		[1, 2];
	`

	tests := []Expected{
		{token.STRING, "foobar"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foo bar"},
		{token.SEMICOLON, ";"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
	}

	runTokenTestCase(t, input, tests)
}
