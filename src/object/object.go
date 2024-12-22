package object

import "fmt"

const (
	INTEGER_VAL = "INTEGER"
		BOOLEAN_VAL = "BOOLEAN"
		NULL_VAL = "NULL"
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

