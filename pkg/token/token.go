package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="

	// Binary Operators
	PLUS  = "+"
	MINUS = "-"
	TIMES = "*"
	DIV   = "/"
	MOD   = "%"

	// Logical Operators
	NOT = "!"
	AND = "&&"
	OR  = "||"

	// Comparison Operators
	LT  = "<"
	LTE = "<="
	GT  = ">"
	GTE = ">="
	EQ  = "=="
	NEQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	TRUE     = "true"
	FALSE    = "false"
	LET      = "let"
	FUNCTION = "function"
	RETURN   = "return"
)

var keywords = map[string]TokenType{
	"true":     TRUE,
	"false":    FALSE,
	"function": FUNCTION,
	"let":      LET,
	"return":   RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
