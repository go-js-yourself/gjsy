package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	UNDEFINED_OBJ    = "UNDEFINED"
	NULL_OBJ         = "NULL"
	FUNCTION_OBJ     = "FUNCTION"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	STRING_OBJ       = "STRING"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (*Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (*Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// I hate that we're about to do this, but if you're going to JS yourself,
// you might as well really JS yourself

type Null struct{}

func (*Null) Type() ObjectType {
	return NULL_OBJ
}

func (*Null) Inspect() string {
	return "null"
}

type Undefined struct{}

func (*Undefined) Type() ObjectType {
	return UNDEFINED_OBJ
}
func (*Undefined) Inspect() string {
	return "undefined"
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (*ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

type Error struct {
	Message string
}

func (*Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "Error: " + e.Message
}

type String struct {
	Value string
}

func (*String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}
