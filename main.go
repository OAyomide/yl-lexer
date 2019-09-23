package main

import "fmt"
import "./lexer"

func main() {
	fmt.Println("HELLO THERE!! THIS IS JUST A SIMPLE LEXER")
	fmt.Println(lex.Lex("jeki yorlang = ooto"))
}
