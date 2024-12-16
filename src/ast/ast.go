package ast

import (
	"bytes"

	"github.com/singlaanish56/Interpreter-In-Go/token"
)

//basic node struct for the tree
type ASTNode interface{
	TokenLiteral() string
	String() string
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
func (astRootNode *ASTRootNode) String() string{
	var out bytes.Buffer
	for _,s := range astRootNode.Statements{
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct{
	Token       token.Token
	Variable    *Variable
	Value        Expression
}

func (letStatment *LetStatement) statementNode() {}
func (letStatement *LetStatement) TokenLiteral() string {return letStatement.Token.Identifier}
func (letStatment *LetStatement) String() string{
	var out bytes.Buffer

	out.WriteString(letStatment.TokenLiteral()+" ")
	out.WriteString(letStatment.Variable.String())
	out.WriteString(" = ")

	if letStatment.Value != nil{
		out.WriteString(letStatment.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type Variable struct{
	Token token.Token
	Value string
}

func (variable *Variable) expressionNode() {}
func (variable *Variable) TokenLiteral() string {return variable.Token.Identifier}
func (variable *Variable) String() string{ return variable.Value}


type ReturnStatement struct{
	Token token.Token
	ReturnValue  Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Identifier}
func (rs *ReturnStatement) String() string{
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil{
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct{
	Token token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {return es.Token.Identifier}
func (es *ExpressionStatement) String() string{
	
	if es.Expression != nil{
		return es.Expression.String()
	}

	return ""
}

type IntegerLiteral struct{
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode(){}
func (il *IntegerLiteral) TokenLiteral() string {return il.Token.Identifier}
func (il *IntegerLiteral) String() string {return il.Token.Identifier}


type PrefixExpression struct{
	Token token.Token
	Operator string
	RightOperator Expression
}

func (pe *PrefixExpression) expressionNode(){}
func (pe *PrefixExpression) TokenLiteral() string{return pe.Token.Identifier}
func (pe *PrefixExpression) String() string{
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.RightOperator.String())
	out.WriteString(")")

	return out.String()
}