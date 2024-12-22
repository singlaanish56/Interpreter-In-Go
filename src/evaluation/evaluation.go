package evaluation

import(

	"github.com/singlaanish56/Interpreter-In-Go/object"
	"github.com/singlaanish56/Interpreter-In-Go/ast"
	
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
		return evaluatePrefixExpression(node.Operator, right)
	default:
		return NULL
	}

}

func evaluateStatements(statements []ast.Statement) object.Object{

	var result object.Object

	for _, statement := range statements{
		result =Eval(statement)
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
		return NULL
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
		return NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}