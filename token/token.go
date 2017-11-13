package token

//TokenType is used to save the type of literal
type TokenType string

//Token is the struct used to capture token type
type Token struct {
	Type    TokenType
	Literal string
}

//constants are exported
const (
	ILLegal = "ILLEGAL"
	EOF     = "EOF"

	//identifiers + literals
	IDENT = "IDENTIFIER"
	INT   = "INT"

	//OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//functions
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
