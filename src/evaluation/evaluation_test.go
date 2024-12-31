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
		{"5-5",0},
		{"5+5",10},
		{"-5-5",-10},
		{"5*5",25},
		{"5/5",1},
		{"((6/3)*2)-5",-1},
		{"((6/3)*2)-(4*(2+3))",-16},
	}

	for _, tt:= range tests{
		eval := testEval(tt.input)
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
		{"5==5", true},
		{"5==6", false},
		{"5!=10", true},
		{"5!=5", false},
		{"(2*(10/5)+3)==((6/2)*6-9)", false},
		{"1<2", true},
		{"2<1", false},
		{"true==false", false},
		{"true==true", true},
		{"false==false", true},
		{"(2<1)==true", false},
		{"(2<1)==false", true},
		{`"are we equal"=="are we equal"`, true},
		{`"are we equal"=="no we are not"`, false},
		{`"are we equal"!="no we are not"`, true},
		{`"are we equal">"no we are not"`, false},
		{`"are we equal"<"no we are not"`, true},
	}

	for _, tt:= range tests{
		eval := testEval(tt.input)
		testBooleanObject(t, eval, tt.expected)
	}
}

func TestEvalStringEvaluation(t *testing.T){
	tests := []struct{
		input string
		expected string
	}{
		{`"((6/3)*2)-5"`,"((6/3)*2)-5"},
		{`"((6/3)*2)-(4*(2+3))"`,"((6/3)*2)-(4*(2+3))"},
		{`"this is a b ugdjdsljlsfdlkfdkjlhdkjfdjkfdfjkdhafjks string"`,"this is a b ugdjdsljlsfdlkfdkjlhdkjfdjkfdfjkdhafjks string"},
		{`"this is it"+", is it?"`,"this is it, is it?"},
	}

	for _, tt:= range tests{
		eval := testEval(tt.input)
		testStringObject(t, eval, tt.expected)
	}
}

func TestEvalIfExpression(t *testing.T){
	tests :=[]struct{
		input string
		expected interface{}
	}{
		{"if(true){10}",10},
		{"if(false){10}",nil},
		{"if(1){10}",10},
		{"if(1 < 2){10}",10},
		{"if(2 > 1){10}",10},
		{"if(1==1){10}",10},
		{"if(1!=2){10}",10},
		{"if(1==2){10}else{5}",5},
	}

	for _,tt := range tests{
		eval := testEval(tt.input)
		integ , ok := tt.expected.(int)
		if ok{
			testIntegerObject(t, eval, int64(integ))
		}else{
			testNullObject(t, eval)
		}
	}
}

func TestEvalReturnExpression(t *testing.T){
	tests:=[] struct{
		input string
		expected interface{}
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2*5; 9;", 10},
		{"9; return 2*5; 9;", 10},
		{"if(10){if(9){return 9;}} return 10;", 9},
	}

	for _,tt := range tests{
		eval := testEval(tt.input)
		integ , ok := tt.expected.(int)
		if ok{
			testIntegerObject(t, eval, int64(integ))
		}else{
			testNullObject(t, eval)
		}
	}
}

func TestErrorHandling(t *testing.T){
	tests := []struct{
		input string
		expected string
	}{
		{
			"5 + true",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"false + true",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if(10>1){ true+false;}",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{	
			"if(10){if(9){return true+false;}} return 10;", 
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"variable not found: foobar",
		},
		{
			`"hello" - "worls"`,
			"unknown operator: STRING - STRING",
		},
		{"[1,2,3][3]","array out of bound index, min index=0, max index=2, got=3"},
		{"[1,2,3][-1]","array out of bound index, min index=0, max index=2, got=-1"},
	}
	for _,tt := range tests{
		eval := testEval(tt.input)
		errObj, ok := eval.(*object.Error)
		if !ok{
			t.Errorf("no error object returned , got=%T", eval)
			continue
		}


		if errObj.Message != tt.expected{
			t.Errorf("wrong error message, expected=%q, got=%q", tt.expected, errObj.Message)
		}
	}
}


func TestEvalLetStatements(t * testing.T){
	tests:= []struct{
		input string
		expected int64
	}{
		{"let a=5;a;", 5},
		{"let a=5*5;a;", 25},
		{"let a=5; let b= a;b;", 5},
		{"let a=5;let b = a; let c = b+a;c;", 10},
	}

	for _, tt:= range tests{
		eval := testEval(tt.input)
		testIntegerObject(t, eval, tt.expected)
	}
}


func TestEvalFunctionObject(t *testing.T){
	input := "fn(x, y){x+y;};"
	eval := testEval(input)
	fn , ok := eval.(*object.Function)
	if !ok{
		t.Fatalf("object is not of type function got=%T", eval)
		return
	}

	if len(fn.Params) !=2{
		t.Fatalf("the number of parameters not as expected=2, got=%d", len(fn.Params))
		return
	}

	if fn.Params[0].String()!="x" || fn.Params[1].String()!="y"{
		t.Fatalf("one of the parameters not as expected=x,y got=%s,%s", fn.Params[0].String(), fn.Params[1].String())
		return
	}

	if fn.Body.String() != "(x+y)"{
		t.Fatalf("body not as expected=%s, got=%s","(x+y)", fn.Body.String())
		return
	}
}

func TestEvalFunctionApplication(t *testing.T){
	tests := []struct{
		input string
		expected int64
	}{
		{"let identity = fn(x){x;} identity(5);",5},
		{"let identity = fn(x){return x;} identity(5);",5},
		{"let double = fn(x){x*2;} double(5);",10},
		{"let add = fn(x, y){x+y;} add(5, 5);",10},
		{"let add = fn(x, y){x+y;} add(5+5, add(5,5));",20},
		{"fn(x){x;}(5)",5},
	}

	for _, tt := range tests{
		eval := testEval(tt.input)
		testIntegerObject(t, eval, tt.expected)
	}
}

func TestEvalBuiltInFunction(t *testing.T){
	tests := []struct{
		input string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument for the len builtin not supported, got INTEGER"},
	}

	for _, tt := range tests{
		eval := testEval(tt.input)

		switch expected:= tt.expected.(type){
		case int:
			testIntegerObject(t, eval, int64(expected))
		case string:
			ee, ok := eval.(*object.Error)
			if !ok{
				t.Errorf("object is not of error type, got=%T", eval)
				continue
			}

			if ee.Message != expected{
				t.Errorf("wrong error message , expected=%s, got=%s", expected, ee.Message)
				continue
			}
		}
	}
}

func TestEvalArrayLiteral(t * testing.T){
	input :="[1+2, 2, 3*3, 4-4]"

	eval := testEval(input)

	result , ok := eval.(*object.Array)
	if !ok{
		t.Fatalf("object is not array, got=%T", eval)
	}

	if len(result.Elements) != 4{
		t.Fatalf("the number of elements not as expected=3, got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 3)
	testIntegerObject(t, result.Elements[1], 2)
	testIntegerObject(t, result.Elements[2], 9)
	testIntegerObject(t, result.Elements[3], 0)
}

func TestEvalArrayIndex(t *testing.T){
	tests := []struct{
		input string
		expected interface{}
	}{
		{"[1,2,3][0]", 1},
		{"[1,2,3][1]", 2},
		{"let arr = [1,2,3]; arr[2];", 3},
		{"let arr = [1,2,3]; arr[0]+arr[1]+arr[2];", 6},
	}

	for _, tt := range tests{
		eval := testEval(tt.input)

		result, ok := tt.expected.(int)
		if ok{
			testIntegerObject(t, eval, int64(result))
		}
	}
}








//helpers
func testNullObject(t *testing.T, eval object.Object) bool{
	if eval != NULL{
		t.Errorf("the object is not null , got=%T", eval)
		return false
	}

	return true
}


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

func testStringObject(t *testing.T, eval object.Object, expected string) bool{
	result, ok := eval.(*object.String)
	if !ok{
		t.Errorf("object is not string , got=%T", eval)
		return false
	}

	if result.Value != expected{
		t.Errorf("integer values dont match expected=%s , got=%s",expected, result.Value)
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

func testEval(input string) object.Object{
	lexer := lexer.New(input)
	p := parser.New(lexer)
	prog := p.ParseProgram()
	e := object.NewEnv()
	eval := Eval(prog, e)
	return eval
}