package evaluator

import (
	"github.com/go-js-yourself/gjsy/pkg/ast"
	"github.com/go-js-yourself/gjsy/pkg/object"
)

func evalGoFunction(node *ast.GoFunctionApplication, env *object.Environment) object.Object {
	target := Eval(node.Function, env)
	if isError(target) {
		return target
	}

	args := evalExpressions(node.Arguments, env)
	if (len(args) == 1) && isError(args[0]) {
		return args[0]
	}

	return applyGoFunction(target, args)
}

func applyGoFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		env := extendFunctionEnv(fn, args)
		go Eval(fn.Body, env)
		return UNDEFINED
	case *object.BuiltinFunc:
		go fn.Func(args...)
		return UNDEFINED
	default:
		return newError("not a function: %s", fn.Type())
	}
}

func evalFunction(node *ast.FunctionApplication, env *object.Environment) object.Object {
	target := Eval(node.Function, env)
	if isError(target) {
		return target
	}

	args := evalExpressions(node.Arguments, env)
	if (len(args) == 1) && isError(args[0]) {
		return args[0]
	}

	return applyFunction(target, args)
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		env := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, env)
		return unwrapReturnValue(evaluated)
	case *object.BuiltinFunc:
		return fn.Func(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return &object.Undefined{}
}
