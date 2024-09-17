package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // token/character we don't know
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"

	// Operators
	ASSIGN     = "="
	ADD_ASSIGN = "+="
	SUB_ASSIGN = "-="
	MUL_ASSIGN = "*="
	DIV_ASSIGN = "/="

	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MOD      = "%"

	INCR = "++"
	DECR = "--"

	AND = "&&"
	OR  = "||"

	LT     = "<"
	LT_EQ  = "<="
	GT     = ">"
	GT_EQ  = ">="
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	PERIOD        = "."
	COMMA         = ","
	SEMICOLON     = ";"
	COLON         = ":"
	COMMENT       = "//"
	COMMENT_START = "/*"
	COMMENT_END   = "*/"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NULL     = "NULL"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	CONST    = "CONST"
	FOR      = "FOR"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	WHILE    = "WHILE"
	IN       = "IN"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"null":     NULL,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"const":    CONST,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"while":    WHILE,
	"in":       IN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
