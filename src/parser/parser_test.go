package parser

import (
	"fmt"
	"testing"

	"github.com/singlaanish56/Interpreter-In-Go/ast"
	"github.com/singlaanish56/Interpreter-In-Go/lexer"
)

func TestLetStatements(t *testing.T){

	input :=`
	let x=5;
	let y=10;
	let foobar=838383838;
	`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	checkForErrors(p, t)
	if prog == nil{
		t.Fatalf("parseprogram() returned nil")
	} 

	if len(prog.Statements) != 3{
		t.Fatalf("some of the statments are missing, got %d \n", len(prog.Statements))
	}

	tests := []struct{
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}


	for i, tt:= range tests{
		st := prog.Statements[i]
		if !testLetStatement(t, st, tt.expectedIdentifier){
			return 
		}

	}
}

func TestReturnStatements(t *testing.T){
	
	input :=`
		return 5;
		return y;
		return add(5,10);
	`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	checkForErrors(p, t)
	if prog == nil{
		t.Fatalf("parseprogram() returned nil")
	} 

	if len(prog.Statements) != 3{
		t.Fatalf("some of the statments are missing, got %d \n", len(prog.Statements))
	}

	for _, returnsts:= range prog.Statements{
		st, ok := returnsts.(*ast.ReturnStatement)
		if !ok{
			t.Errorf("st not of type Return , got=%T", returnsts)
		}

		if st.TokenLiteral() != "return"{
			t.Errorf("token liternal not return , got=%s", st.TokenLiteral())
		}


	}
}


func TestVariableExpression(t *testing.T){
	input := "anish"

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()

	checkForErrors(p,t)

	if len(prog.Statements) != 1{
		t.Fatalf("the program has not enough statements")
	}


	st, ok := prog.Statements[0].(*ast.ExpressionStatement)
	if !ok{
		t.Fatalf("the first statement of not type expression, got %T", prog.Statements[0])
	}

	varr, ok := st.Expression.(*ast.Variable)
	if !ok{
		t.Fatalf("the first statement of not type variable, got %T", st.Expression)
	}

	if varr.Value != "anish"{
		t.Errorf("the variable value not as expected %s, got %s","anish",varr.Value)
	}

	if varr.TokenLiteral()!="anish"{
		t.Errorf("the variable tokenliteral not as expected %s, got %s","anish",varr.TokenLiteral())
	}

}

func TestIntegerExpression(t *testing.T){
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()

	checkForErrors(p,t)

	if len(prog.Statements) != 1{
		t.Fatalf("the program has not enough statements")
	}


	st, ok := prog.Statements[0].(*ast.ExpressionStatement)
	if !ok{
		t.Fatalf("the first statement of not type expression, got %T", prog.Statements[0])
	}

	varr, ok := st.Expression.(*ast.IntegerLiteral)
	if !ok{
		t.Fatalf("the first statement of not type integer, got %T", st.Expression)
	}

	if varr.Value != 5{
		t.Errorf("the variable value not as expected %d, got %d",5,varr.Value)
	}

	if varr.TokenLiteral()!="5"{
		t.Errorf("the variable tokenliteral not as expected %s, got %s","5",varr.TokenLiteral())
	}

}

func TestPrefixExpression(t *testing.T){
	prefixTests :=[]struct{
		Input string
		Operator string
		Number interface{}
	}{
		{"!5;","!",5},
		{"-10","-",10},
		{"!true","!", true},
		{"!false","!", false},
	}

	for _, tt := range prefixTests{
		l := lexer.New(tt.Input)
		p := New(l)
		prog := p.ParseProgram()
		checkForErrors(p,t)
		if len(prog.Statements) != 1{
			t.Fatalf("the program has not enough statements")
		}
		st, ok := prog.Statements[0].(*ast.ExpressionStatement)
		if !ok{
		t.Fatalf("the first statement of not type expression, got %T", prog.Statements[0])
		}
		varr, ok := st.Expression.(*ast.PrefixExpression)
		if !ok{
			t.Fatalf("the first statement of not type prefix, got %T", st.Expression)
		}
		if varr.Operator != tt.Operator{
			t.Fatalf("the operator is not expected %s, got %s", tt.Operator, varr.Operator)
		}
		if testLiteral(t, varr.RightOperator, tt.Number){
			return
		}
	}

}

func TestInfixExpression(t *testing.T){

	infixTests := []struct{
		input string
		leftValue interface{}
		operator string
		rightValue interface{}
	}{
		{"5+5;",5,"+",5},
		{"5-5;",5,"-",5},
		{"5*5;",5,"*",5},
		{"5/5;",5,"/",5},
		{"5>5;",5,">",5},
		{"5<5;",5,"<",5},
		{"5==5;",5,"==",5},
		{"5!=5;",5,"!=",5},
		{"true == true",true,"==", true},
		{"true != false",true,"!=", false},
		{"false == false",false,"==", false},
		
	}


	for _, tt :=range infixTests{
		l := lexer.New(tt.input)
		p := New(l)

		prog := p.ParseProgram()
		checkForErrors(p,t)
		if len(prog.Statements) != 1{
			t.Fatalf("the program has not enough statements")
		}
		st, ok := prog.Statements[0].(*ast.ExpressionStatement)
		if !ok{
		t.Fatalf("the first statement of not type expression, got %T", prog.Statements[0])
		}
		varr, ok := st.Expression.(*ast.InfixExpression)
		if !ok{
			t.Fatalf("the first statement of not type infix, got %T", st.Expression)
		}

		if testLiteral(t, varr.LeftOperator, tt.leftValue){
			return
		}
		if varr.Operator != tt.operator{
			t.Fatalf("the operator is not expected %s, got %s", tt.operator, varr.Operator)
		}
		if testLiteral(t, varr.RightOperator, tt.rightValue){
			return
		}
	}
}

func TestOPeratorPrecendenceParsing(t *testing.T){
	tests := []struct{
		input string
		expected string
	}{
		{"-a*b","((-a)*b)"},
		{"!-a","(!(-a))"},
		{"a+b+c","((a+b)+c)"},
		{"a+b-c","((a+b)-c)"},
		{"a*b*c","((a*b)*c)"},
		{"a+b/c","(a+(b/c))"},
		{"a+b/c","(a+(b/c))"},
		{"a+b*c+d/e-f","(((a+(b*c))+(d/e))-f)"},
		{"3+4; -5*5","(3+4)((-5)*5)"},
		{"5>4 == 3<4","((5>4)==(3<4))"},
		{"5<4 == 3>4","((5<4)==(3>4))"},
		{"5>4 != 3<4","((5>4)!=(3<4))"},
		{"3+4*5==3*1+4*5","((3+(4*5))==((3*1)+(4*5)))"},
		{"true","true"},
		{"false","false"},
		{"3>5 == false","((3>5)==false)"},
		{"3<5 != true","((3<5)!=true)"},
	}

	for _, tt:= range tests{
		l := lexer.New(tt.input)
		p := New(l)
		prog := p.ParseProgram()
		checkForErrors(p  ,t)

		actual := prog.String()

		if actual != tt.expected{
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

//helpers
func testIdentifier(t *testing.T, exp ast.Expression, value string)bool{
	ident, ok := exp.(*ast.Variable)
	if!ok{
		t.Errorf("exp npt Variable, got=%T", exp)
		return false
	}

	if ident.Value != value{
		t.Errorf("ident. value wrong expected=%q, got=%q", value, ident.Value)
		return false
	}

	if ident.TokenLiteral()!= value{
		t.Errorf("ident. value wrong expected=%q, got=%q", value, ident.TokenLiteral())
		return false
	}

	return true
}

func testBoolean(t *testing.T, exp ast.Expression, value bool) bool{
	bo , ok := exp.(*ast.BooleanLiteral)
	if !ok{
		t.Errorf("the expression not as expected , got=%T", exp)
		return false
	}

	if bo.Value != value{
		t.Errorf("the value not as expected=%t, got=%t", value, bo.Value)
		return false
	}

	if bo.TokenLiteral() != fmt.Sprintf("%t", value){
		t.Errorf("the token literal not as expected=%q, got=%q", bo.TokenLiteral() ,fmt.Sprintf("%t", value))
		return false
	}

	return true
}

func testLiteral(t *testing.T, exp ast.Expression, expected interface{}) bool{
	switch v := expected.(type){
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBoolean(t, exp, v)
	}

	t.Errorf("type of exp not recognized, got=%T", exp)
	return false
}


func testInfix(t *testing.T, exp ast.Expression, left interface{}, operator string , right interface{}) bool{

	inexp, ok:= exp.(*ast.InfixExpression)
	if !ok{
		t.Errorf("exp npt infix, got=%T", inexp)
		return false
	}

	if !testLiteral(t, inexp.LeftOperator, left){
		return false
	}
	if inexp.Operator != operator{
		t.Errorf("operator not matching for the infix expression , expected=%q, got=%q", operator, inexp.Operator)
		return false
	}

	if !testLiteral(t, inexp.RightOperator, right){
		return false
	}

	return true
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool{
	if s.TokenLiteral() != "let"{
		t.Errorf("the token type is not let")
		return false
	}
	
	letst, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("the statement type is not let")
		return false
	} 
	
	if letst.Variable.Value != name{
		t.Errorf("the let statement variable name doesnt not match")
		return false
	}
	
	if letst.Variable.TokenLiteral() != name{
		t.Errorf("the let statement variabe token no store properly got %s, expected %s", letst.Variable.TokenLiteral(), name)
		return false
	}
	
	return true
	}


func testIntegerLiteral(t *testing.T, ro  ast.Expression, expected int64) bool{
	num , ok := ro.(*ast.IntegerLiteral)
	if !ok{
		t.Errorf("the integer type literal not found , got %T", ro)
		return false
	}

	if num.Value != expected{
		t.Errorf("the integer value not expected %d , got %d", expected, num.Value)
		return false	
	}


	if num.TokenLiteral()!=fmt.Sprintf("%d", expected){
		t.Errorf("the integer tokenLiteral not expected %d , got %s",expected, num.TokenLiteral(), )
		return false
	}

	return true
}

func checkForErrors(parser *Parser, t *testing.T){
	errors := parser.Errors()

	if len(errors)==0{
		return
	}

	for _, error := range errors{
		t.Errorf(error.Error())
	}

	t.FailNow()
}