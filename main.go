package main

import "github.com/kris/token"

type Lexer struct {
	input             string
	tokenPosition     int
	character         byte
	readTokenPosition int
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	// then we want to read the character
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
	var token token.Token

	lex.eatWhiteSpace()
}
