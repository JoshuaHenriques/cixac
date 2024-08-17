package lexer

import (
	"testing"

	"github.com/joshuahenriques/cixac/token"
)

func TestNextToken(t *testing.T) {
	input := `
    let five = 5;
    let ten = 10;

    let add = fn(x, y) {
      x + y
    }

    let result = add(five, ten)
    !*-/5
    5 < 10 > 5

    if (5 < 10) {
      return true
    } else if (5 == 5) {
      return true
    } else {
      return false
    }

    10 == 10
    10 != 9
    "foobar"
    "foo bar"
    [1, 2]
    {"foo": "bar"}
    let nil = null
    10 >= 10
    11 <= 10
    /*
    this is a multi-line comment
    */
    true && false
    true || false
    // this is a comment
    5 % 4

    const w = 5
    const x = 5.5
    const y = 5.
    const z = .5
    
    i++
    i--

    for (let i = 0; i < 5; i++) {
      break
      continue
    }

    while (i < 10) {}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
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
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.BANG, "!"},
		{token.ASTERISK, "*"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.INT, "5"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.EQ, "=="},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "nil"},
		{token.ASSIGN, "="},
		{token.NULL, "null"},
		{token.INT, "10"},
		{token.GT_EQ, ">="},
		{token.INT, "10"},
		{token.INT, "11"},
		{token.LT_EQ, "<="},
		{token.INT, "10"},
		{token.COMMENT_START, "/*"},
		{token.COMMENT_END, "*/"},
		{token.TRUE, "true"},
		{token.AND, "&&"},
		{token.FALSE, "false"},
		{token.TRUE, "true"},
		{token.OR, "||"},
		{token.FALSE, "false"},
		{token.COMMENT, "//"},
		{token.INT, "5"},
		{token.MOD, "%"},
		{token.INT, "4"},
		{token.CONST, "const"},
		{token.IDENT, "w"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.CONST, "const"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.FLOAT, "5.5"},
		{token.CONST, "const"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.FLOAT, "5.0"},
		{token.CONST, "const"},
		{token.IDENT, "z"},
		{token.ASSIGN, "="},
		{token.FLOAT, "0.5"},
		{token.IDENT, "i"},
		{token.INCR, "++"},
		{token.IDENT, "i"},
		{token.DECR, "--"},
		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.LET, "let"},
		{token.IDENT, "i"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.LT, "<"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.INCR, "++"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.BREAK, "break"},
		{token.CONTINUE, "continue"},
		{token.RBRACE, "}"},
		{token.WHILE, "while"},
		{token.LPAREN, "("},
		{token.IDENT, "i"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
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
