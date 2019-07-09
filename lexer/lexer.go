package lexer

import "github.com/oayomide/chi/token"

type Lexer struct {
	input             string
	tokenPosition     int
	character         byte
	readTokenPosition int
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	// then we want to read the character
	lex.readInputCharacter()
	return lex
}

func (lex *Lexer) readInputCharacter() {
	// first, check if we have read everythin in the input string
	if lex.readTokenPosition >= len(lex.input) {
		lex.character = 0
	} else {
		// the read character/token is at the position of the index`lex.readTokenPosition` on the input string.
		// E.g input string is Ade. position of d in Ade is: [1] where an array/slice starts from 0 and 1 is the `lex.readTokenPosition`
		// to simply put, the postion of the current token in the input string
		lex.character = lex.input[lex.readTokenPosition]
	}

	// the position of the current token is the position of the token
	// we've read
	lex.tokenPosition = lex.readTokenPosition
	lex.readTokenPosition++ // increase the position by 1. i.e, continue to the next token in input string
}

// NextToken returns the next token in the input string
func (lex *Lexer) NextToken() token.Token {
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
