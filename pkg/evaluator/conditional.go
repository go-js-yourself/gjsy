package evaluator

import (
	"github.com/go-js-yourself/gjsy/pkg/ast"
	"github.com/go-js-yourself/gjsy/pkg/object"
)

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Expression, env)
	} else if ie.ElseExpression != nil {
		return Eval(ie.ElseExpression, env)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
