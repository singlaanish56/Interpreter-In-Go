package evaluation

import (
	"testing"

	"github.com/singlaanish56/Interpreter-In-Go/lexer"
	"github.com/singlaanish56/Interpreter-In-Go/object"
	"github.com/singlaanish56/Interpreter-In-Go/parser"
)

func TestEvalIntegerEvaluation(t *testing.T){
	tests := []struct{
		input string
		expected int64
	}{
		{"5",5},
		{"10",10},
		{"-10",-10},
		{"-5",-5},
	}

	for _, tt:= range tests{
		lexer := lexer.New(tt.input)
		p := parser.New(lexer)
		prog := p.ParseProgram()
		eval := Eval(prog)
		testIntegerObject(t, eval, tt.expected)
	}
}

func TestEvalBooleanEvaluation(t *testing.T){
	tests := []struct{
		input string
		expected bool
	}{
		{"true",true},
		{"false",false},
		{"!true",false},
		{"!false",true},
		{"!5",false},
		{"!!true",true},
		{"!!false",false},
		{"!!5",true},
	}

	for _, tt:= range tests{
		lexer := lexer.New(tt.input)
		p := parser.New(lexer)
		prog := p.ParseProgram()
		eval := Eval(prog)
		testBooleanObject(t, eval, tt.expected)
	}
}

//helpers

func testIntegerObject(t *testing.T, eval object.Object, expected int64) bool{
	result, ok := eval.(*object.Integer)
	if !ok{
		t.Errorf("object is not integer , got=%T", eval)
		return false
	}

	if result.Value != expected{
		t.Errorf("integer values dont match expected=%d , got=%d",expected, result.Value)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, eval object.Object, expected bool) bool{
	result, ok := eval.(*object.Boolean)
	if !ok{
		t.Errorf("object is not boolean , got=%T", eval)
		return false
	}

	if result.Value != expected{
		t.Errorf("boolean values dont match expected=%t , got=%t",expected, result.Value)
		return false
	}

	return true
}