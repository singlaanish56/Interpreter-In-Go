package parser

import (
	"github.com/singlaanish56/Interpreter-In-Go/lexer"
	"github.com/singlaanish56/Interpreter-In-Go/token"
	"github.com/singlaanish56/Interpreter-In-Go/ast"
)




type Parser struct{
	currToken  token.Token
	peekToken  token.Token
	errorList  []error
	lexer *lexer.Lexer
}

func New(lx *lexer.Lexer) *Parser{

	parser := &Parser{lexer: lx, errorList: []error{}}
	parser.nextToken()
	parser.nextToken()

	return parser
}


func (parser *Parser) ParseProgram() *ast.ASTRootNode{
	topNode := &ast.ASTRootNode{Statements: []ast.Statement{}}

	for parser.currToken.Type != token.EOF{
		st := parser.parseStatement()
		if st != nil{
			topNode.Statements =append(topNode.Statements, st)
		}

		parser.nextToken()
	}

	return topNode
}	


func (parser *Parser) nextToken() {
	parser.currToken = parser.peekToken
	parser.peekToken = parser.lexer.GetToken()
}

func (parser *Parser) parseStatement() ast.Statement{

	switch parser.currToken.Type{
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}


func (parser *Parser) parseLetStatement() *ast.LetStatement{

	st := &ast.LetStatement{Token: parser.currToken}

	if ! parser.checkPeekToken( token.VARIABLE){
		return nil
	}

	st.Variable = &ast.Variable{Token: parser.currToken, Value: parser.currToken.Identifier}

	if !parser.checkPeekToken(token.EQUALTO){
		return nil
	}

	for !parser.currTokenIs(token.SEMICOLON){
		parser.nextToken()
	}

	return st
}

func (parser *Parser) checkPeekToken(tokenType token.TokenType) bool{
	if parser.peekTokenIs(tokenType){
		parser.nextToken()
		return true
	}

	return false
}

func (parser *Parser) peekTokenIs(tokenType token.TokenType) bool{
	return parser.peekToken.Type == tokenType
}

func (parser *Parser) currTokenIs(tokenType token.TokenType) bool{
	return parser.currToken.Type == tokenType
}