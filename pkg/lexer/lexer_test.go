package lexer

import (
	"testing"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let a = 10;
	var b = 20;
	let c = a + b;

	function test(x,y) {
		if (1 == 1) {
			return "foo bar";
		} else {
			1 != 2;
		}
		while (!(true && false)) {
			return 1 * a - b / c;
		}
	}

	console.log("test");
	go test(null,undefined,'foo bar');
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "c"},
		{token.ASSIGN, "="},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "function"},
		{token.IDENT, "test"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.EQ, "=="},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.STRING, "foo bar"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.INT, "1"},
		{token.NEQ, "!="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.WHILE, "while"},
		{token.LPAREN, "("},
		{token.NOT, "!"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.AND, "&&"},
		{token.FALSE, "false"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "1"},
		{token.TIMES, "*"},
		{token.IDENT, "a"},
		{token.MINUS, "-"},
		{token.IDENT, "b"},
		{token.DIV, "/"},
		{token.IDENT, "c"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.IDENT, "console"},
		{token.DOT, "."},
		{token.IDENT, "log"},
		{token.LPAREN, "("},
		{token.STRING, "test"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.GO, "go"},
		{token.IDENT, "test"},
		{token.LPAREN, "("},
		{token.NULL, "null"},
		{token.COMMA, ","},
		{token.UNDEF, "undefined"},
		{token.COMMA, ","},
		{token.STRING, "foo bar"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
