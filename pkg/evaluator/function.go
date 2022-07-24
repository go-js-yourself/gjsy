package evaluator

import (
	"github.com/go-js-yourself/gjsy/pkg/ast"
	"github.com/go-js-yourself/gjsy/pkg/object"
)

func evalFunction(node *ast.FunctionApplication, env *object.Environment) object.Object {
	target := Eval(node.Function, env)
	if isError(target) {
		return target
	}
	function, ok := target.(*object.Function)

	if !ok {
		return newError("not a function: %s", target.Type())
	}

	args := evalExpressions(node.Arguments, env)
	if (len(args) == 1) && isError(args[0]) {
		return args[0]
	}

	return applyFunction(function, args)
}

func applyFunction(fn *object.Function, args []object.Object) object.Object {
	env := extendFunctionEnv(fn, args)
	evaluated := Eval(fn.Body, env)
	return unwrapReturnValue(evaluated)
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

	return obj
}
