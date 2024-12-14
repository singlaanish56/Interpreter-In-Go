package ast

import (
	"github.com/singlaanish56/Interpreter-In-Go/token"
)

//basic node struct for the tree
type ASTNode interface{
	TokenLiteral() string
}


type Statement interface{
	ASTNode
	statementNode()
}


type Expression interface{
	ASTNode
	expressionNode()
}

//parseProgram returns this which contains all the Statemen trees in the list
type ASTRootNode struct{
	Statements []Statement
}

//basic Node in itself
func (program *ASTRootNode) TokenLiteral() string{
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}


type LetStatement struct{
	Token       token.Token
	Variable    *Variable
	Value        Expression
}

func (letStatment *LetStatement) statementNode() {}
func (letStatement *LetStatement) TokenLiteral() string {return letStatement.Token.Identifier}



type Variable struct{
	Token token.Token
	Value string
}

func (variable *Variable) expressionNode() {}
func (variable *Variable) TokenLiteral() string {return variable.Token.Identifier}



type ReturnStatement struct{
	Token token.Token
	ReturnValue  Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Identifier}