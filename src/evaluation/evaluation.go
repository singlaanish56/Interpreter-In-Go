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

func Eval(node ast.ASTNode, env *object.Environment) object.Object{

	switch node := node.(type){
	case *ast.ASTRootNode:
		return evaluateStatements(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return evaluateBoolean(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.RightOperator, env)
		if isError(right){
			return right
		}
		return evaluatePrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left:= Eval(node.LeftOperator, env)
		if isError(left){
			return left
		}
		right:=Eval(node.RightOperator, env)

		if isError(right){
			return right
		}
		return evaluateInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evaluateBlockStatements(node.Statements, env)
	case *ast.IfExpression:
		return evaluateIfExpression(node, env)
	case *ast.ReturnStatement:
		returnVal := Eval(node.ReturnValue, env)
		if isError(returnVal){
			return returnVal
		}
		return &object.ReturnValue{Value : returnVal}
	case *ast.LetStatement:
		letVal := Eval(node.Value, env)
		if isError(letVal){
			return letVal
		}
		return env.Set(node.Variable.Value, letVal)
	case *ast.Variable:
		return evalVariable(node, env)
	case *ast.FunctionExpression:
		params := node.Parameters
		body := node.Body
		return &object.Function{Params: params, Body: body, Env: env}
	case *ast.CallExpression:
		fnc := Eval(node.Function, env)
		if isError(fnc){
			return fnc
		}
		args := evalArguments(node.Arguments, env)
		if len(args)==1 && isError(args[0]){
			return args[0]
		}
		return applyFunction(fnc, args)
	default:
		return NULL
	}

}

func evaluateStatements(statements []ast.Statement, env *object.Environment) object.Object{

	var result object.Object

	for _, statement := range statements{
		result =Eval(statement, env)
		
		switch result := result.(type){
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result	
}

func evaluateBlockStatements(statements []ast.Statement,env *object.Environment) object.Object{

	var result object.Object

	for _, statement := range statements{
		result =Eval(statement, env)
		
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


func evaluateIfExpression(node *ast.IfExpression, env *object.Environment) object.Object{
	condition := Eval(node.Condition, env)
	if isError(condition){
		return condition
	}
	if isTruthful(condition){
		return Eval(node.Consequence, env)
	}else if node.Alternative != nil{
		return Eval(node.Alternative, env)
	}else{
		return NULL
	}

}

func evalVariable(node *ast.Variable, env *object.Environment) object.Object{
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("variable not found: %s", node.Value)
	}

	return val
}

func evalArguments(args []ast.Expression, env *object.Environment) []object.Object{
	var result []object.Object

	for _, e := range args{
		eval := Eval(e, env)
		if isError(eval){
			return []object.Object{eval}
		}

		result=append(result, eval)
	}

	return result
}

func applyFunction(fnc object.Object, args []object.Object) object.Object{
	fn, ok := fnc.(*object.Function)
	if !ok{
		return newError("not a function: %s", fnc.Type())
	}

	fnEnv := newFunctionEnvironment(fn , args)
	eval := Eval(fn.Body, fnEnv)
	return unwrap(eval)
}

func newFunctionEnvironment(fn *object.Function, args []object.Object) *object.Environment{

	extendedEnv := object.NewEnclosedEnvironment(fn.Env)

	for argIdx, arg := range fn.Params{
		extendedEnv.Set(arg.Value, args[argIdx])
	}

	return extendedEnv
}

func unwrap(obj object.Object) object.Object{
	if returnVal ,ok := obj.(*object.ReturnValue); ok{
		return returnVal.Value
	}

	return obj
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