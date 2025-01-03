package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/singlaanish56/Interpreter-In-Go/ast"
)

const (
	INTEGER_VAL = "INTEGER"
	BOOLEAN_VAL = "BOOLEAN"
	STRING_VAL="STRING"
	NULL_VAL = "NULL"
	RETURN_VAL ="RETURN"
	ERROR_OBJ ="ERROR"
	FUNCTION ="FUNCTION"
	BUILTIN_OBJ="BUILTIN"
	ARRAY_OBJ="ARRAY"
	HASHPAIR_OBJ="HASHPAIR"
)

type HashKey struct{
	Type ObjectType
	Value uint64
}

type Hashable interface{
	HashKey() HashKey
}

type HashPair struct{
	Key Object
	Value Object
}


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
func (i *Integer) HashKey() HashKey{
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

type Boolean struct{
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_VAL }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t",b.Value)}
func (b *Boolean) HashKey() HashKey{
	if b.Value{
		return HashKey{Type: b.Type(), Value: uint64(1)}
	}

	return HashKey{Type: b.Type(), Value: uint64(0)}
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_VAL }
func (s *String) Inspect() string { return fmt.Sprintf("%s",s.Value)}
func (s *String) HashKey() HashKey{
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
type Array struct{
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements :=[]string{}
	for _, e := range a.Elements{
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements,","))
	out.WriteString("]")

	return out.String()
}

type Hash struct{
	Pairs map[HashKey]HashPair
}


func (hm *Hash) Type() ObjectType { return HASHPAIR_OBJ }
func (hm *Hash) Inspect() string {
	var out bytes.Buffer

	elements :=[]string{}
	for _, p := range hm.Pairs{
		elements = append(elements, p.Key.Inspect()+":"+p.Value.Inspect())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(elements,","))
	out.WriteString("}")

	return out.String()
}

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


type BuiltinFunction func(args ...Object) Object

type Builtin struct{
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ}
func (b *Builtin) Inspect() string { return "builtin function"}