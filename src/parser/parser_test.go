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
		Number int64
	}{
		{"!5;","!",5},
		{"-10","-",10},
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
		if testIntegerLiteral(t, varr.RightOperator, tt.Number){
			return
		}
	}

}


//helpers
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