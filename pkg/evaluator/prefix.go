package evaluator

import "github.com/go-js-yourself/gjsy/pkg/object"

func unknownOperatorError(operator string, obj object.Object) *object.Error {
	return newError("unknown operator: %s%s", operator, obj.Type())
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return unknownOperatorError(operator, right)
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return unknownOperatorError("-", right)
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}
