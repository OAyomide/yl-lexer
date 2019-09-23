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
	OPERATOR = "OPERATOR"
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
	SEPARATOR  = "SEPARATOR"
	// COMMA      = ","
	// SEMICOLON  = ";"
	// LEFTPAREN  = "("
	// RIGHTPAREN = ")"
	// LEFTBRACE  = "{"
	// RIGHTBRACE = "}"

	// keywords
	KEYWORD = "KEYWORD"
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

// Note make sure u give the exact number of token in array,
// else you gonna spend ages debugging invain
// Total_time_wasted = 2hours
var keywords = [14]string {
	"jeki", // LET
    "nigbati",
    "sope",
    "tabi", // IF
    "ise",
    "fun",
    "pada", // RETURN
    "kuro",
    "se",
    "yi",
    "iru",
    "padasi",
    "gbewole", // IMPORT
    "woke",
}

var boolean = [2]string {
	"ooto", // TRUE
	"iro", // FALSE
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
