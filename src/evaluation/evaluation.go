package evaluation

import (
	"fmt"

	"github.com/singlaanish56/Interpreter-In-Go/ast"
	"github.com/singlaanish56/Interpreter-In-Go/object"
)

var (
	TRUE = &object.Boolean{Value:true}
	FALSE = &object.Boolean{Value:false}
	NULL = &object.Null{}
)

func Eval(node ast.ASTNode) object.Object{

	switch node := node.(type){
	case *ast.ASTRootNode:
		return evaluateStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return evaluateBoolean(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.RightOperator)
		if isError(right){
			return right
		}
		return evaluatePrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left:= Eval(node.LeftOperator)
		if isError(left){
			return left
		}
		right:=Eval(node.RightOperator)

		if isError(right){
			return right
		}
		return evaluateInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evaluateBlockStatements(node.Statements)
	case *ast.IfExpression:
		return evaluateIfExpression(node)
	case *ast.ReturnStatement:
		returnVal := Eval(node.ReturnValue)
		if isError(returnVal){
			return returnVal
		}
		return &object.ReturnValue{Value : returnVal}
	default:
		return NULL
	}

}

func evaluateStatements(statements []ast.Statement) object.Object{

	var result object.Object

	for _, statement := range statements{
		result =Eval(statement)
		
		switch result := result.(type){
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result	
}

func evaluateBlockStatements(statements []ast.Statement) object.Object{

	var result object.Object

	for _, statement := range statements{
		result =Eval(statement)
		
		if result != nil && (result.Type()==object.RETURN_VAL || result.Type()==object.ERROR_OBJ){
			return result
		}

	}

	return result	
}

func evaluateBoolean(val bool) object.Object{
	if val{
		return TRUE
	}

	return FALSE
}

func evaluatePrefixExpression(operator string, right object.Object) object.Object{
	switch operator{
	case "!":
		return evaluateExclamationExpression(right)
	case "-":
		return evaluateMinusExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evaluateExclamationExpression(right object.Object) object.Object{
	switch right{
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evaluateMinusExpression(right object.Object) object.Object{
	if right.Type() != object.INTEGER_VAL{
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evaluateInfixExpression(operator string, left, right object.Object) object.Object{

	switch {
	case left.Type() == object.INTEGER_VAL && right.Type()==object.INTEGER_VAL:
		return evaluateIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return evaluateBoolean(left == right)
	case operator == "!=":
		return evaluateBoolean(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evaluateIntegerInfixExpression(operator string , left, right object.Object) object.Object{
	lval := left.(*object.Integer).Value
	rval := right.(*object.Integer).Value

	switch (operator){
	case "+":
		return &object.Integer{Value: lval + rval}
	case "-":
		return &object.Integer{Value: lval-rval}
	case "*":
		return &object.Integer{Value: lval*rval}
	case "/":
		return &object.Integer{Value: lval/rval}
	case ">":
		return evaluateBoolean(lval > rval)
	case "<":
		return evaluateBoolean(lval < rval)
	case "==":
		return evaluateBoolean(lval == rval)
	case "!=":
		return evaluateBoolean(lval != rval)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}


func evaluateIfExpression(node *ast.IfExpression) object.Object{
	condition := Eval(node.Condition)
	if isError(condition){
		return condition
	}
	if isTruthful(condition){
		return Eval(node.Consequence)
	}else if node.Alternative != nil{
		return Eval(node.Alternative)
	}else{
		return NULL
	}

}

func isTruthful(obj object.Object) bool{
	switch obj {
	case NULL:
		return false
	case FALSE:
		return false
	case TRUE:
		return true
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error{
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool{
	if obj != nil{
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}