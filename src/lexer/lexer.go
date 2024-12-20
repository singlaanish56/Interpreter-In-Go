 package lexer

import (

	"github.com/singlaanish56/Interpreter-In-Go/token"
)

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

func (lexer *Lexer) peekChar() rune{
	if lexer.currentPostion >= len(lexer.input){
		return 0
	}

	return lexer.input[lexer.nextReadPosition]
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
	for isEsapceSequence(lexer.char){
		lexer.nextChar()
	}

	//is it a number //TODO: move to a function
	if(lexer.char >='0' && lexer.char<='9'){
	
		start := lexer.currentPostion

		for(lexer.char >='0' && lexer.char <= '9'){
			lexer.nextChar()
		}

		return token.Token{Type: token.NUMBER,Identifier: string(lexer.input[start:lexer.currentPostion]), StartPosition: start, EndPosition: lexer.currentPostion}
	}

	// is it  a string 
	if(lexer.char=='"'){
		
		start := lexer.currentPostion+1
		var strBuilder []rune
		for{
			lexer.nextChar()

			if(lexer.char=='"' || lexer.char==0){
				break;
			}
			if lexer.char == '\n' || lexer.char == '\r' || lexer.char == '\t'{
				continue
			}

			strBuilder=append(strBuilder, lexer.char)
		}

		str := string(strBuilder)
		endIndex := lexer.currentPostion
		lexer.nextChar()
		return token.Token{Type:token.STRING, Identifier: str, StartPosition: start, EndPosition: endIndex}
	}


	//any variable names or keywords
	if((lexer.char>='a' && lexer.char<='z') || (lexer.char>='A' && lexer.char<='Z')){
		
		start := lexer.currentPostion
	
		for (lexer.char>='a' && lexer.char<='z') || (lexer.char>='A' && lexer.char<='Z'){
			lexer.nextChar()
		}

		str := string(lexer.input[start:lexer.currentPostion])
		
		tt := token.KeywordMap["var"]
		if tokenType, exists := token.KeywordMap[str]; exists{
			tt = tokenType
		}

		return token.Token{Type: tt, Identifier: str, StartPosition: start, EndPosition: lexer.currentPostion}
	}


	var tk token.Token
	switch lexer.char{
	case '=':
		str:= string(lexer.char)
		
		if lexer.peekChar() == '='{
			lexer.nextChar()
			str+=string(lexer.char)	
			tk = token.Token{Type: token.DOUBLEEQUALTO, Identifier: str, StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}

		}else{
			tk = token.Token{Type: token.EQUALTO, Identifier: str, StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
		}
	case '_':
		tk = token.Token{Type: token.UNDERSCORE, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case ';':
		tk = token.Token{Type: token.SEMICOLON, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case ':':
		tk = token.Token{Type: token.COLON, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '{':
		tk = token.Token{Type: token.OCURLYBR, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '}':
		tk = token.Token{Type: token.CCURLYBR, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '(':
		tk = token.Token{Type: token.OROUNDBR, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case ')':
		tk = token.Token{Type: token.CROUNDBR, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '[':
		tk = token.Token{Type: token.OSQAUREBR, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case ']':
		tk = token.Token{Type: token.CSQUAREBR, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '+':
		tk = token.Token{Type: token.PLUS, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case ',':
		tk = token.Token{Type: token.COMMA, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '!':
		str:= string(lexer.char)
		
		if lexer.peekChar() == '='{
			lexer.nextChar()
			str+=string(lexer.char)	
			tk = token.Token{Type: token.EXCLAMATIONEQUALTO, Identifier: str, StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}

		}else{
			tk = token.Token{Type: token.EXCLAMATION, Identifier: str, StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
		}
	case '-':
		tk = token.Token{Type: token.MINUS, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '/':
		tk = token.Token{Type: token.DIVIDE, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '*':
		tk = token.Token{Type: token.MULTIPLY, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '<':
		tk = token.Token{Type: token.OANGLEDBR, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	case '>':
		tk = token.Token{Type: token.CANGLEDBR, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
	default:
	
		if lexer.char==0{
			tk = token.Token{Type: token.EOF, Identifier: "", StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
		}else{
		tk = token.Token{Type: token.INV, Identifier: string(lexer.char), StartPosition: lexer.currentPostion, EndPosition: lexer.currentPostion+1}
		}
	}
	
	lexer.nextChar()

	return tk

}


func isEsapceSequence(currentChar rune) bool{
	return currentChar==' ' || currentChar=='\n' || currentChar=='\t' || currentChar=='\r'
}
