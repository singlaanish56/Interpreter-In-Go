package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/singlaanish56/Interpreter-In-Go/ast"
)

const (
	INTEGER_VAL = "INTEGER"
	BOOLEAN_VAL = "BOOLEAN"
	NULL_VAL = "NULL"
	RETURN_VAL ="RETURN"
	ERROR_OBJ ="ERROR"
	FUNCTION ="FUNCTION"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_VAL }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d",i.Value)}

type Boolean struct{
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_VAL }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t",b.Value)}

type Null struct{}

func (b *Null) Type() ObjectType { return NULL_VAL }
func (b *Null) Inspect() string { return "null"}


type ReturnValue struct{
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VAL }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect()}

type Error struct{
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { return e.Message}


type Environment struct{
	env map[string]Object
	outer *Environment
}

func NewEnv() *Environment{
	s:= make(map[string]Object)
	return &Environment{env: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment{
	env := NewEnv()
	env.outer = outer

	return env
}

func (e *Environment) Get(name string) (Object, bool){
	obj, ok := e.env[name]
	if !ok && e.outer!=nil{
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string , obj Object) Object{
	e.env[name] = obj;
	return obj
}


type Function struct{
	Params []*ast.Variable
	Body *ast.BlockStatement
	Env *Environment	
}

func (f *Function) Type() ObjectType { return FUNCTION}
func (f *Function) Inspect() string{
	var out bytes.Buffer

	params := []string{}
	for _, p:= range f.Params{
		params = append(params, p.String())
	}

	out.WriteString("fn(")
	out.WriteString(strings.Join(params,","))
	out.WriteString("){\n")
	out.WriteString(f.Body.String())
	out.WriteString("}\n")

	return out.String()
}