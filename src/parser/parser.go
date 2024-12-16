package parser

import (
	"fmt"

	"github.com/singlaanish56/Interpreter-In-Go/ast"
	"github.com/singlaanish56/Interpreter-In-Go/lexer"
	"github.com/singlaanish56/Interpreter-In-Go/token"
)

//pratt parser says that every token would have two functions to choose from prefix/postfix and infix
type (
	prefixParseFn  func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression // the argument for the left operator
)


type Parser struct{
	currToken  token.Token
	peekToken  token.Token
	errorList  []error
	lexer *lexer.Lexer

	prefixParseFnMap  map[token.TokenType]prefixParseFn
	infixParseFnMap map[token.TokenType]infixParseFn
}

func New(lx *lexer.Lexer) *Parser{

	parser := &Parser{lexer: lx, errorList: []error{}}
	parser.nextToken()
	parser.nextToken()

	parser.prefixParseFnMap = make(map[token.TokenType]prefixParseFn)
	parser.addPrefix(token.VARIABLE, parser.parseVariable)
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

func (parser *Parser) Errors() []error{
	return parser.errorList
}


func (parser *Parser) nextToken() {
	parser.currToken = parser.peekToken
	parser.peekToken = parser.lexer.GetToken()
}

func (parser *Parser) parseStatement() ast.Statement{

	switch parser.currToken.Type{
	case token.LET:
		return parser.parseLetStatement()
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatment()
	}
}

func (parser *Parser) parseExpressionStatment() *ast.ExpressionStatement{
	st := &ast.ExpressionStatement{Token: parser.currToken}

	st.Expression = parser.parseExpression(LOWEST)

	if parser.peekTokenIs(token.SEMICOLON){
		parser.nextToken()
	}

	return st
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

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement{
	st := &ast.ReturnStatement{Token: parser.currToken}

	for !parser.currTokenIs(token.SEMICOLON){
		parser.nextToken()
	}

	return st
}


func (parser *Parser) parseExpression(precendence int) ast.Expression{
	prefix := parser.prefixParseFnMap[parser.currToken.Type]

	if prefix ==nil{
		return nil
	}

	leftExp := prefix()
	return leftExp
}

func (parser *Parser) parseVariable() ast.Expression{
	return &ast.Variable{Token: parser.currToken, Value: parser.currToken.Identifier}
}

//helpers
func (parser *Parser) checkPeekToken(tokenType token.TokenType) bool{
	if parser.peekTokenIs(tokenType){
		parser.nextToken()
		return true
	}

	parser.peekError(tokenType)
	return false
}

func (parser *Parser) peekError(tokenType token.TokenType){
	err := fmt.Errorf("expected next token to be %s, got %s", tokenType, parser.peekToken.Type)
	parser.errorList = append(parser.errorList, err)
}

func (parser *Parser) peekTokenIs(tokenType token.TokenType) bool{
	return parser.peekToken.Type == tokenType
}

func (parser *Parser) currTokenIs(tokenType token.TokenType) bool{
	return parser.currToken.Type == tokenType
}

func (parser *Parser) addPrefix(tokenType token.TokenType, fn prefixParseFn){
	parser.prefixParseFnMap[tokenType]=fn
}

func (parser *Parser) addInfix(tokenType token.TokenType, fn infixParseFn){
	parser.infixParseFnMap[tokenType]=fn
}

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)		