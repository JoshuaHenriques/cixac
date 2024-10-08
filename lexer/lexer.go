package lexer

import (
	"strings"

	"github.com/joshuahenriques/cixac/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = l.newTwoCharToken(token.EQ)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		switch l.peekChar() {
		case '+':
			tok = l.newTwoCharToken(token.INCR)
		case '=':
			tok = l.newTwoCharToken(token.ADD_ASSIGN)
		default:
			tok = newToken(token.PLUS, l.ch)
		}
	case '-':
		switch l.peekChar() {
		case '-':
			tok = l.newTwoCharToken(token.DECR)
		case '=':
			tok = l.newTwoCharToken(token.SUB_ASSIGN)
		default:
			tok = newToken(token.MINUS, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = l.newTwoCharToken(token.NOT_EQ)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		switch l.peekChar() {
		case '/':
			tok = l.newTwoCharToken(token.COMMENT_END)
		case '=':
			tok = l.newTwoCharToken(token.MUL_ASSIGN)
		default:
			tok = newToken(token.ASTERISK, l.ch)
		}
	case '/':
		switch l.peekChar() {
		case '/':
			tok = l.newTwoCharToken(token.COMMENT)
			l.skipComment()
		case '*':
			tok = l.newTwoCharToken(token.COMMENT_START)
			l.skipMultiComment()
			return tok
		case '=':
			tok = l.newTwoCharToken(token.DIV_ASSIGN)
		default:
			tok = newToken(token.SLASH, l.ch)
		}
	case '%':
		tok = newToken(token.MOD, l.ch)
	case '&':
		if l.peekChar() == '&' {
			tok = l.newTwoCharToken(token.AND)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			tok = l.newTwoCharToken(token.OR)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			tok = l.newTwoCharToken(token.LT_EQ)
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = l.newTwoCharToken(token.GT_EQ)
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '.':
		if !isDigit(l.peekChar()) {
			tok = newToken(token.PERIOD, l.ch)
		} else {
			return l.readIdentOrNumber()
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isAlphaNum(l.ch) {
			return l.readIdentOrNumber()
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) newTwoCharToken(tokenType token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	literal := string(ch) + string(l.ch)
	return token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentOrNumber() token.Token {
	position := l.position

	for isAlphaNum(l.ch) {
		if isLetter(l.ch) && l.peekChar() == '.' {
			l.readChar()
			return token.Token{Literal: l.input[position:l.position], Type: token.LookupIdent(l.input[position:l.position])}
		}
		l.readChar()
	}

	input := l.input[position:l.position]

	if isNumber(input) {
		return l.readNumber(input)
	} else {
		return token.Token{Literal: input, Type: token.LookupIdent(input)}
	}
}

func isNumber(input string) bool {
	for i := range input {
		if !isDigit(input[i]) && input[i] != '.' {
			return false
		}
	}
	return true
}

func (l *Lexer) readNumber(num string) token.Token {
	var literal strings.Builder
	var isFloat bool

	if num[0] == '.' {
		literal.WriteString("0.")
		isFloat = true
		num = num[1:]
	}

	for i := range num {
		if num[i] == '.' && !isFloat {
			literal.WriteByte(num[i])
			isFloat = true
			if i >= len(num)-1 {
				literal.WriteString("0")
				break
			}
		} else if num[i] == '.' && isFloat {
			break
		} else {
			literal.WriteByte(num[i])
		}
	}

	tok := token.Token{Type: token.INT, Literal: literal.String()}

	if isFloat {
		tok.Type = token.FLOAT
	}

	return tok
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != '\r' && l.peekChar() != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipMultiComment() {
	l.readChar()
	l.readChar()
	for l.ch != '*' || l.peekChar() != '/' {
		l.readChar()
	}
}

func isAlphaNum(ch byte) bool {
	return isDigit(ch) || isLetter(ch) || ch == '.'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
