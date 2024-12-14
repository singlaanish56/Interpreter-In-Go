package parser

import (
	"testing"

	"github.com/singlaanish56/Interpreter-In-Go/lexer"
	"github.com/singlaanish56/Interpreter-In-Go/ast"
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