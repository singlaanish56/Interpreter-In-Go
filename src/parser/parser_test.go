package parser

import (
	"fmt"
	"testing"

	"github.com/singlaanish56/Interpreter-In-Go/ast"
	"github.com/singlaanish56/Interpreter-In-Go/lexer"
	
)

func TestLetStatements(t *testing.T){
	tests := []struct{
		input string
		expectedVariable string
		expectedValue interface{}
	}{
		{"let x=5;","x", 5},
		{"let y = true;","y", true},
		{"let foobar = y;","foobar","y"},
	}

	for _, tt := range tests{
		l := lexer.New(tt.input)
		p := New(l)
		prog := p.ParseProgram()
		checkForErrors(p,t)

		if len(prog.Statements) != 1{
			t.Fatalf("the program has not enough statements")
		}
	
	
		st := prog.Statements[0]
	
		if !testLetStatement(t, st, tt.expectedVariable){
			return
		}

		val := st.(*ast.LetStatement).Value

		if !testLiteral(t, val, tt.expectedValue){
			return
		}
	}

}

func TestReturnStatements(t *testing.T){
	tests := []struct{
		input string
		expectedValue interface{}
	}{
		{"return 5;", 5},
		{"return true;",true},
		{"return foobar;","foobar"},
	}

	for _, tt := range tests{
		l := lexer.New(tt.input)
		p := New(l)
		prog := p.ParseProgram()
		checkForErrors(p,t)

		if len(prog.Statements) != 1{
			t.Fatalf("the program has not enough statements")
		}
	
	
		st := prog.Statements[0]
	
		val := st.(*ast.ReturnStatement).ReturnValue

		if !testLiteral(t, val, tt.expectedValue){
			return
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
		{"a*(b*c)","(a*(b*c))"},
		{"(a+b)/c","((a+b)/c)"},
		{"a + add(b*c) +d", "((a+add((b*c)))+d)"},
		{"add(a,b,1,2*3,4+5,add(6,7*8))","add(a,b,1,(2*3),(4+5),add(6,(7*8)))"},
		
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

func TestIfElseExpression(t *testing.T){
	tests := []struct{
		input string
	}{
		{"if(x<y){x}"},
		{"if(x<y){x}else{y}"},
	}

	for _, tt := range tests{
		l := lexer.New(tt.input)
		p := New(l)
		prog := p.ParseProgram()
		checkForErrors(p,t)

		if len(prog.Statements) != 1{
			t.Fatalf("the program has not enough statements")
			return
		}

		st , ok := prog.Statements[0].(*ast.ExpressionStatement)
		if !ok{
			t.Errorf("the program statement is not of the expected type, got=%T",prog.Statements[0])
			return
		}

		exp, ok := st.Expression.(*ast.IfExpression)
		if !ok{
			t.Errorf("expected the if expression , got=%T", st.Expression)
			return 
		}

		if !testInfix(t, exp.Condition, "x", "<", "y"){
			return 
		}

		if len(exp.Consequence.Statements) != 1{
			t.Fatalf("the consequence has not enough statements")
			return 
		}

		con, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
		if !ok{
			t.Errorf("the consquence for the if doesnt hav the expected type got=%T", con)
			return 
		}

		if !testIdentifier(t, con.Expression, "x"){
			return
		}

		if exp.Alternative != nil{
			if len(exp.Alternative.Statements) != 1{
				t.Fatalf("the consequence has not enough statements")
				return 
			}
	
			alt, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
			if !ok{
				t.Errorf("the consquence for the if doesnt hav the expected type got=%T", exp.Alternative.Statements[0])
				return 
			}
	
			if !testIdentifier(t, alt.Expression, "y"){
				return
			}
		}


	}
}

func TestFuncExpression(t *testing.T){

	input := `fn(x, y){x+y;}`

	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()
	checkForErrors(p,t)

	if len(prog.Statements) !=1{
		t.Errorf("the number of statements not as expected")
		return 
	}

	st, ok := prog.Statements[0].(*ast.ExpressionStatement)
	if !ok{
		t.Errorf("the expression  type not as expected got=%T", prog.Statements[0])
		return 
	}

	fst, ok := st.Expression.(*ast.FunctionExpression)
	if !ok{
		t.Errorf("the type expected is function literal got=%T", st.Expression)
		return
	}

	testLiteral(t, fst.Parameters[0],"x")
	testLiteral(t, fst.Parameters[1],"y")

	if !testIdentifier(t, fst.Parameters[1], "y"){
		return
	}

	if len(fst.Body.Statements) !=1{
		t.Errorf("the number of statement in the function body not as expected , got=%d",len(fst.Body.Statements))
		return
	}

	body, ok := fst.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok{
		t.Errorf("expected as expression got something else, got=%T", fst.Body.Statements[0])
		return
	}

	testInfix(t, body.Expression, "x","+","y")
}

func TestFunctionParameters(t *testing.T){
	tests := []struct{
		input string
		expectedParams []string
	}{
		{"fn(){}", []string{}},
		{"fn(x){}",[]string{"x"}},
		{"fn(x,y,z){}",[]string{"x","y","z"}},
	}

	for _, tt := range tests{
		l := lexer.New(tt.input)
		p := New(l)
		prog := p.ParseProgram()
		checkForErrors(p,t)

		if len(prog.Statements)!=1{
			t.Errorf("the number of statements not as expected")
			return 
		}

		st := prog.Statements[0].(*ast.ExpressionStatement)
		fn := st.Expression.(*ast.FunctionExpression)

		if len(fn.Parameters) != len(tt.expectedParams){
			t.Errorf("the lemght pf the paranmeters not as expected=%d, got=%d",len(fn.Parameters),len(tt.expectedParams))
		}

		for i, arg := range fn.Parameters{
			testLiteral(t, arg, tt.expectedParams[i])
		}
	}
}

func TestCallExpressions(t *testing.T){
	input := `add(1, 2*3, 4+5)`

	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()
	checkForErrors(p,t)
	
	if len(prog.Statements)!=1{
		t.Errorf("the number of statements not as expected")
		return 
	}

	st , ok := prog.Statements[0].(*ast.ExpressionStatement)
	if !ok{
		t.Errorf("the expression  type not as expected got=%T", prog.Statements[0])
		return 
	}

	exp , ok := st.Expression.(*ast.CallExpression)
	if !ok{
		t.Errorf("the expression  type not as expected got=%T", prog.Statements[0])
		return 
	}

	if !testIdentifier(t, exp.Function, "add"){
		return
	} 

	if len(exp.Arguments) != 3{
		t.Errorf("the number of arg not as exprected 3, got=%d", len(exp.Arguments))
		return
	}

	testLiteral(t, exp.Arguments[0], 1)
	testInfix(t, exp.Arguments[1],2,"*",3)
	testInfix(t, exp.Arguments[2],4,"+",5)
}	


//helpers
func testIdentifier(t *testing.T, exp ast.Expression, value string)bool{
	ident, ok := exp.(*ast.Variable)
	if!ok{
		t.Errorf("exp not Variable, got=%T", exp)
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
		t.Errorf("the integer value not as expected %d , got %d", expected, num.Value)
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