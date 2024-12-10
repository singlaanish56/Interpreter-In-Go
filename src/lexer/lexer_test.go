package lexer

import (
	"fmt"
	"testing"

	"github.com/singlaanish56/Interpreter-In-Go/token"
)

func TestNextToken(t *testing.T){

	input := `let five=5;let ten=10;let add = fn(x,y){x+y;}let str = "this is a string";`

	tests:=[]struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.VARIABLE, "five"},
		{token.EQUALTO, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.VARIABLE, "ten"},
		{token.EQUALTO, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.VARIABLE, "add"},
		{token.EQUALTO, "="},
		{token.FUNCTION, "fn"},
		{token.OROUNDBR, "("},
		{token.VARIABLE, "x"},
		{token.COMMA, ","},
		{token.VARIABLE, "y"},
		{token.CROUNDBR, ")"},
		{token.OCURLYBR, "{"},
		{token.VARIABLE, "x"},
		{token.PLUS, "+"},
		{token.VARIABLE, "y"},
		{token.SEMICOLON, ";"},
		{token.CCURLYBR, "}"},
		{token.LET, "let"},
		{token.VARIABLE, "str"},
		{token.EQUALTO, "="},
		{token.STRING, "this is a string"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests{
		tok := lexer.GetToken()

		if tok.Type !=tt.expectedType{
			t.Fatalf("tests[%d] - token type wrong, expected=%q, got=%q, expecedliteral=%q",i,tt.expectedType, tok.Type, tt.expectedLiteral)
		}
		if tok.Identifier !=tt.expectedLiteral{
			t.Fatalf("tests[%d] - literal type wrong, expected=%q, got=%q",i,tt.expectedType, tok.Identifier)
		}

		fmt.Printf("tokenizeliteral : %q\n", tt.expectedLiteral)
	}
}