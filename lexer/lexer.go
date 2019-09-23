package lex

import "strings"
// import "fmt"
import "../token"

// there should be a way to import fikes locally
// import "github.com/oayomide/chi/token"

var input string =  ""               
var lexerPosition int = 0
var column int = 0
var line int = 1
var digitDecimal string = "0123456789"
var identifierBeginChar string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var identifierEndChar string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"
var operatorChar = "+-/=*%"
var separatorChar = "{}[]<>();,"
var spaceChar = " "

// check if lexer position is yet to exceed input length
func isBound() bool {
	if lexerPosition < len(input) {
		return true
	}

	return false
}

// peek character in current lexer position
func peekChar() string {

	// check if we are yet to exceed input length
	if isBound() {
		return string(input[lexerPosition])
	} 

	return ""
}

// consume character in current lexer position ans increment lexer position
func eatChar() string {
	position := lexerPosition

	// check if we are yrt to exceed input length
	if isBound() {

		// increment lexer position
		lexerPosition++
		column++
		return string(input[position])
	}

	return ""
}

// revert lexer position if lexer method does not identify character
func revert(position int) {
	lexerPosition = position
}

// check if series of characters are valid identifiers
func checkValidIdentifier() token.Token {
	identifier := ""
	var tokens token.Token

	// check if first character is a valid letter
	if isBound() && strings.Contains(identifierBeginChar, peekChar()) {
		identifier += eatChar()
		
		// characeters after the first one can either be a letter or a digit
		for isBound() && strings.Contains(identifierEndChar, peekChar()) {
			identifier += eatChar()
		}
	}

	if len(identifier) > 0 {
		tokens = token.Token{ token.IDENTIFIER, identifier }
		return tokens
	}

	return token.Token{ token.UNKNOWN, identifier }
}

// check if series of characters are valid keyword
func checkKeywordValid() token.Token {
	keyword := ""
	var tokens token.Token
	position := lexerPosition

	// check if first character is a valid letter
	for isBound() && strings.Contains(identifierBeginChar, peekChar()) {
		keyword += eatChar()
	}

	if token.LookUpKeyword(keyword) {
		tokens = token.Token{ token.KEYWORD, keyword }
		return tokens
	}

	revert(position)
	return token.Token{ token.UNKNOWN, keyword }
}


// check if series of characters are valid boolean
func checkBooleanValid() token.Token {
	boolean := ""
	position := lexerPosition
	var tokens token.Token

	// check if first character is a valid letter
	for isBound() && strings.Contains(identifierBeginChar, peekChar()) {
		boolean += eatChar()
	}

	if token.LookUpBoolean(boolean) {
		tokens = token.Token{ token.BOOLEAN, boolean }
		return tokens
	}

	revert(position)
	return token.Token{ token.UNKNOWN, boolean }
}

// check if character is a valid operator
func checkOperatorValid() token.Token {
	operator := ""
	var tokens token.Token

	if isBound() && strings.Contains(operatorChar, peekChar()) {
		operator += eatChar()
	}

	if len(operator) > 0 {
		tokens = token.Token{ token.OPERATOR, operator }
		return tokens
	}

	return token.Token{ token.UNKNOWN, operator }
}

// check if character is a valid separator
func checkSeparatorValid() token.Token {
	separator := ""
	var tokens token.Token

	if isBound() && strings.Contains(separatorChar, peekChar()) {
		separator += eatChar()
	}

	if len(separator) > 0 {
		tokens = token.Token{ token.SEPARATOR, separator }
		return tokens
	}

	return token.Token{ token.UNKNOWN, separator }
}

// check if character is a valid space
func checkSpaceValid() token.Token {
	space := ""
	var tokens token.Token

	for isBound() && peekChar() == " " || peekChar() == "\t" {
		space += eatChar()
	}

	if len(space) > 0 {
		tokens = token.Token{ token.SPACE, space }
		return tokens
	}

	return token.Token{ token.UNKNOWN, space }
}




// run all lexer functions
func lexNext() token.Token {

	var boolean = checkBooleanValid()
    if boolean.Type != token.UNKNOWN {
        return boolean
	}

	var keyword = checkKeywordValid()
    if keyword.Type != token.UNKNOWN {
        return keyword
	}
	
	var identifier = checkValidIdentifier()
    if identifier.Type != token.UNKNOWN {
        return identifier
	}

	var space = checkSpaceValid()
    if space.Type != token.UNKNOWN {
        return space
	}

	var operator = checkOperatorValid()
    if operator.Type != token.UNKNOWN {
        return operator
	}

	var separator = checkSeparatorValid()
    if separator.Type != token.UNKNOWN {
        return separator
	}

	return token.Token{ token.UNKNOWN, eatChar() }
} 

// Lex -> run lexNext function
func Lex(code string) []token.Token {
	input = code

	var tokens []token.Token
	for isBound() {
		temp := lexNext()

		if temp.Type != token.UNKNOWN {
			tokens = append(tokens, temp)
		}
	}
	return tokens
}

/**
// NextToken returns the next token in the input string
func (lex *Lexer) NextToken() Token {
	var tokn token.Token

	lex.eatWhiteSpace()
	switch lex.character {
	case '=':
		// we are checking if the nextcharacter contains more than one "="
		// since our interpreter will understand equality sign, we check if
		// we encounter another "=" immediatetely after encountering "="
		// in that case, its a token type of equality (==) and not assign (=)
		// if we were to support strict equality as seen in javascript
		// we could simple check again.
		if lex.peekNextCharacter() == '=' {
			// we are assigning the character encountered by our lexer
			// to a variable character before we continue so that it'd form the
			// complete token type
			// remember increment is ++. so we get the first + then we assign that to a variable
			// our lex.readInputCharacter() method then get the next character which is also +
			// then we concat this to form the increment, "++"!
			// same applies for "==" and "--"
			character := lex.character
			// continue reading the input
			lex.readInputCharacter()
			// this makes our token become "==" jus as we want
			tokn = token.Token{Type: token.EQ, Literal: string(character) + string(lex.character)}
		} else {
			tokn = createNewToken(token.ASSIGN, lex.character)
		}
	case ';':
		tokn = createNewToken(token.SEMICOLON, lex.character)
	case '+':
		// same as equality and assignment above
		if lex.peekNextCharacter() == '+' {
			character := lex.character
			lex.readInputCharacter()
			tokn = token.Token{Type: token.INCR, Literal: string(character) + string(lex.character)}
		} else {
			tokn = createNewToken(token.PLUS, lex.character)
		}
	case '-':
		if lex.peekNextCharacter() == '-' {
			character := lex.character
			lex.readInputCharacter()
			tokn = token.Token{Type: token.DECR, Literal: string(character) + string(lex.character)}
		} else {
			tokn = createNewToken(token.SUBTRACT, lex.character)
		}
	case ')':
		tokn = createNewToken(token.RIGHTPAREN, lex.character)
	case '(':
		tokn = createNewToken(token.LEFTPAREN, lex.character)
	case '}':
		tokn = createNewToken(token.RIGHTBRACE, lex.character)
	case '{':
		tokn = createNewToken(token.LEFTBRACE, lex.character)
	case ',':
		tokn = createNewToken(token.COMMA, lex.character)
	case '!':
		if lex.peekNextCharacter() == '!' {
			character := lex.character
			lex.readInputCharacter()
			tokn = token.Token{Type: token.NOT_EQ, Literal: string(character) + string(lex.character)}
		} else {
			tokn = createNewToken(token.BANG, lex.character)
		}

	case '*':
		tokn = createNewToken(token.ASTERISK, lex.character)
	case '/':
		tokn = createNewToken(token.SLASH, lex.character)
	case '<':
		if lex.peekNextCharacter() == '=' {
			character := lex.character
			lex.readInputCharacter()
			tokn = token.Token{Type: token.LT_OR_EQ, Literal: string(character) + string(lex.character)}
		} else {
			tokn = createNewToken(token.LT, lex.character)
		}
	case '>':
		if lex.peekNextCharacter() == '=' {
			character := lex.character
			lex.readInputCharacter()
			tokn = token.Token{Type: token.GT_OR_EQ, Literal: string(character) + string(lex.character)}
		} else {
			tokn = createNewToken(token.GT, lex.character)
		}
	case '%':
		tokn = createNewToken(token.PERC, lex.character)
	case '0':
		tokn.Type = token.EOF // end of the file/input
		tokn.Literal = ""     // empty string

	default:
		if isLetter(lex.character) {
			tokn.Literal = lex.readIdentifier()
			tokn.Type = token.LookupIndententifier(tokn.Literal)
			return tokn
		} else if isDigit(lex.character) {
			tokn.Type = token.INT
			tokn.Literal = lex.readNumber()
			return tokn
		} else {
			tokn = createNewToken(token.ILLEGAL, lex.character)
		}
	}
	lex.readInputCharacter()
	return tokn
}

// eatWhiteSpace eats the whitespace in the input string.
// this is because <<kris>> is not case sensitive. so \n, \t, \r are all ignored
// so while the lexer encounters these chars, it simply ignores/skips and continue to other tokens in the input
func (lex *Lexer) eatWhiteSpace() {
	for lex.character == ' ' || lex.character == '\t' || lex.character == '\n' || lex.character == '\r' {
		lex.readInputCharacter()
	}
}

func (lex *Lexer) peekNextCharacter() byte {
	// if the position of the currently read token is greater than or equal
	// to the length of our input string. meaning that the input string
	// has been fully read and there's nothing left to read. in this case
	// return 0 else return the character in the index of the current token position
	if lex.readTokenPosition >= len(lex.input) {
		return 0
	}

	return lex.input[lex.readTokenPosition]
}

// newToken creates a new token using the token struct
func createNewToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(character),
	}
}

// the ascii code of the alphabet range from
// 097 to 122 where 097 is for a and 122 is for z
// it ranges from 065 to 090 for capital letters.
// 065 for A and 090 for Z
// so we are simply sayin if the ascii value of the value
// is less than or equal to these or equal to the ascii value of _ which is 95
// then its a letter
func isLetter(character byte) bool {
	return character <= 'a' && character <= 'z' || character <= 'A' && character <= 'Z' || character == '_'
}

// checks if the character is a number
func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lex *Lexer) readNumber() string {
	position := lex.tokenPosition
	for isDigit(lex.character) {
		lex.readInputCharacter()
	}
	return lex.input[position:lex.tokenPosition]
}

// readIdentifier returns the identifier of the input string
// since we have read the input string and we've not encountered
// an operator or delimiter, that means the rest of our input
// is a reserved keyword (identifier) like ```let```.
// we want to look for the identifier here by looking for strings
// that are between our remaining strings that are not delimiters and operators
func (lex *Lexer) readIdentifier() string {
	position := lex.tokenPosition
	// while our token is a letter, continue reading
	for isLetter(lex.character) {
		lex.readInputCharacter()
	}
	// returns the  the tokens from the
	// position of the token in the inpout to the next tokenposition ??
	// NB: will word this better later.. caffeined up right now
	return lex.input[position:lex.tokenPosition]
}
**/