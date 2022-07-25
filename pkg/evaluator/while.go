package evaluator

import (
	"github.com/go-js-yourself/gjsy/pkg/ast"
	"github.com/go-js-yourself/gjsy/pkg/object"
)

func evalWhileExpression(ie *ast.WhileExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	for isTruthy(condition) {
		Eval(ie.Expression, env)
		condition = Eval(ie.Condition, env)
	}
	return UNDEFINED
}
