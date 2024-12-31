package ast

import (
	"bytes"
	"strings"

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

type BlockStatement struct{
	Token token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string {return bs.Token.Identifier}
func (bs *BlockStatement) String() string{
	var out bytes.Buffer

	for _,s := range bs.Statements{
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

type BooleanLiteral struct{
	Token token.Token
	Value bool
}


func (bl *BooleanLiteral) expressionNode(){}
func (bl *BooleanLiteral) TokenLiteral() string {return bl.Token.Identifier}
func (bl *BooleanLiteral) String() string {return bl.Token.Identifier}

type StringLiteral struct{
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode(){}
func (sl *StringLiteral) TokenLiteral() string {return sl.Token.Identifier}
func (sl *StringLiteral) String() string {return sl.Token.Identifier}

type ArrayLiteral struct{
	Token token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode(){}
func (al *ArrayLiteral) TokenLiteral() string {return al.Token.Identifier}
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements:= []string{}
	for _,el:= range al.Elements{
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements,","))
	out.WriteString("]")

	return out.String()
}

type HashLiteral struct{
	Token token.Token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode(){}
func (hl *HashLiteral) TokenLiteral() string {return hl.Token.Identifier}
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	elements:= []string{}
	for k,v:= range hl.Pairs{
		elements = append(elements, k.String()+":"+v.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(elements,","))
	out.WriteString("}")

	return out.String()
}

type IfExpression struct{
	Token token.Token
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode(){}
func (ie *IfExpression) TokenLiteral() string{return ie.Token.Identifier}
func (ie *IfExpression) String() string{
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(ie.Consequence.String())
	
	if ie.Alternative != nil{
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type FunctionExpression struct{
	Token token.Token
	Parameters []*Variable
	Body *BlockStatement
}

func (fe *FunctionExpression) expressionNode(){}
func (fe *FunctionExpression) TokenLiteral() string{return fe.Token.Identifier}
func (fe *FunctionExpression) String() string{
	var out bytes.Buffer

	params := []string{}
	for _, p := range fe.Parameters{
		params = append(params, p.String())
	}


	out.WriteString("if")
	out.WriteString("(")
	out.WriteString(strings.Join(params,","))
	out.WriteString(")")
	out.WriteString(fe.Body.String())

	return out.String()
}

type CallExpression struct{
	Token token.Token
	Function Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode(){}
func (ce *CallExpression) TokenLiteral() string{return ce.Token.Identifier}
func (ce *CallExpression) String() string{
	var out bytes.Buffer

	args := []string{}
	for _, arg := range ce.Arguments{
		args = append(args, arg.String())
	}
	
	
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args,","))
	out.WriteString(")")

	return out.String()
}

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


type InfixExpression struct{
	Token token.Token
	LeftOperator Expression
	Operator string
	RightOperator Expression
}

func (pe *InfixExpression) expressionNode(){}
func (pe *InfixExpression) TokenLiteral() string{return pe.Token.Identifier}
func (pe *InfixExpression) String() string{
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.LeftOperator.String())
	out.WriteString(pe.Operator)
	out.WriteString(pe.RightOperator.String())
	out.WriteString(")")

	return out.String()
}

type IndexExpression struct{
	Token token.Token
	Left Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode(){}
func (ie *IndexExpression) TokenLiteral() string{return ie.Token.Identifier}
func (ie *IndexExpression) String() string{
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("]")
	out.WriteString(")")

	return out.String()
}