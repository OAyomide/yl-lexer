package token

type TokenType string

// Token is the
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	UNKNOWN = "UNKNOWN"

	// identifier and literals
	IDENTIFIER = "IDENTIFIER"
	INTEGER  = "INTEGER"

	// operators
	OPERATOR = "OPERATORS"
	// ASSIGN   = "="
	// PLUS     = "+"
	// SUBTRACT = "-"
	// BANG     = "!"
	// ASTERISK = "*"
	// SLASH    = "/"
	// LT       = "<"
	// GT       = ">"
	// EQ       = "=="
	// NOT_EQ   = "!="
	// GT_OR_EQ = ">="
	// LT_OR_EQ = "<="
	// INCR     = "++"
	// DECR     = "--"
	// MODULO   = "%"

	// delimters
	SEPARATOR  = "SEPARATORS"
	// COMMA      = ","
	// SEMICOLON  = ";"
	// LEFTPAREN  = "("
	// RIGHTPAREN = ")"
	// LEFTBRACE  = "{"
	// RIGHTBRACE = "}"

	// keywords
	KEYWORD = "KEYWORDS"
	// FUNCTION = "FUN" 
	// LET      = "LET"
	// TRUE     = "TRUE"
	// FALSE    = "FALSE"
	// IF       = "IF"
	// ELSE     = "ELSE"
	// RETURN   = "RETURN"
	BOOLEAN  = "BOOLEAN"
	SPACE    =  "SPACE"
)

var keywords = [20]string {
	"fun",
	"let",
	"true",
	"false",
	"if",
	"else",
	"return",
}

var boolean = [2]string {
	"true",
	"false",
}

// LookUpKeyword -> checks if series of character are valid keyword
func LookUpKeyword(ident string) bool {
	for _, i := range keywords {
		if i == ident {
			return true
		}
	}

	return false
}


// LookUpBoolean -> checks if series of character are valid keyword
func LookUpBoolean(ident string) bool {
	for _, i := range boolean {
		if i == ident {
			return true
		}
	}

	return false
}
