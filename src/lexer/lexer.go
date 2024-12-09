package lexer

import "github.com/singlaanish56/Interpreter-In-Go/token"

type Lexer struct {
	input            []rune
	char             rune
	currentPostion   int
	nextReadPosition int
}

func New(input string) *Lexer {
	lexer := &Lexer{input: []rune(input), nextReadPosition: 0}
	lexer.nextChar()
	return lexer
}

func (lexer *Lexer) nextChar() {
	if lexer.nextReadPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.nextReadPosition]
	}

	lexer.currentPostion = lexer.nextReadPosition
	lexer.nextReadPosition++
}

func (lexer *Lexer) GetToken() token.Token{

	//white space

	//switch statement to handle the basics
	
}
