package object

import "fmt"

const (
	INTEGER_VAL = "INTEGER"
	BOOLEAN_VAL = "BOOLEAN"
	NULL_VAL = "NULL"
	RETURN_VAL ="RETURN"
	ERROR_OBJ ="ERROR"
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



