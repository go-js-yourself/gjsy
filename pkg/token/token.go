package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN = "="
	DOT    = "."

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
	VAR      = "var"
	FUNCTION = "function"
	RETURN   = "return"
	IF       = "if"
	ELSE     = "else"
	NULL     = "null"
	UNDEF    = "undefined"
	GO       = "go"
	WHILE    = "while"
)

var keywords = map[string]TokenType{
	"true":      TRUE,
	"false":     FALSE,
	"function":  FUNCTION,
	"var":       VAR,
	"let":       LET,
	"return":    RETURN,
	"if":        IF,
	"else":      ELSE,
	"null":      NULL,
	"undefined": UNDEF,
	"go":        GO,
	"while":     WHILE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
