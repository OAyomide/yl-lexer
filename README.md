### AIM
The aim of this project is to eventually write the interpreter for [yorlang](https://github.com/anoniscoding/yorlang/). From start to finish, this is purely for academic/learning purposes. Currently, this is just the lexer for the interpreter. The aim is to change the tokens to use Yoruba words like yorlang.

I am writing this without changing the tokens and all that because i first want to get the basic token working then extend/change as deemed fit. If you want to help with making this faster, sure pull request are welcome! 

Also, this lexer wouldn't be possible without Thorsten Ball's (correct?) book on building interpreters. I simply read the book very well, learned from it, read more, polished the code and added an idea of two of mine. Again he's inspired another person to build something (Mark Bates being another person!). Its a must read


#### Whats inside
Inside the lexer package, we have the code for the lexer itself. The lexer takes the input code, tokenizes it and returns a key/value pair of the token type and the token literal (A.K.A string representation). Nothing special (atleast yet).

The token package contains the tokens we want to support. In this package, we'll eventually add the yoruba tokens we want to support.


#### Whats the long game?
Well, I am hugely a fan of [Anoniscoding](https://github.com/anoniscoding/)! like, a huge fan! So the aim of the project is to write a proper interpreter (and even probably full fledge compiler) for Yorlang. Just as expected, we're looking at a recursive-descent parser here.

First, after making sure our tokens support Yoruba tokens, the next thing is the repl to test things out (well, just print out the tokens and the token literal). Then the parser and the parser will come later.


### Maintenance?
Yeah well, this is dear to me! So unlike most of my projects (tsk tsk), I'll actually pay mind to this and pursue it as best as I can.