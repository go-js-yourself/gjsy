package object

const (
	BUILTIN_FUNC_OBJ = "BUILTIN_FUNC"
	BUILTIN_OBJ      = "BUILTIN"
)

type BuiltinFunction func(args ...Object) Object

type BuiltinFunc struct {
	Func BuiltinFunction
}

func (*BuiltinFunc) Type() ObjectType {
	return BUILTIN_FUNC_OBJ
}

func (*BuiltinFunc) Inspect() string {
	return "builtin function"
}

type BuiltinObj struct {
	Funcs map[string]*BuiltinFunc
}

func (*BuiltinObj) Type() ObjectType {
	return BUILTIN_OBJ
}

func (*BuiltinObj) Inspect() string {
	return "builtin object"
}
