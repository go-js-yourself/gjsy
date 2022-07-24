package evaluator

import "github.com/go-js-yourself/gjsy/pkg/object"

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntPair(operator, left, right)
	default:
		return NULL
	}
}

func evalIntPair(operator string, left object.Object, right object.Object) object.Object {
	lval := left.(*object.Integer).Value
	rval := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: lval + rval}
	case "-":
		return &object.Integer{Value: lval - rval}
	case "*":
		return &object.Integer{Value: lval * rval}
	case "/":
		return &object.Integer{Value: lval / rval}
	case "<":
		if lval < rval {
			return TRUE
		} else {
			return FALSE
		}
	case ">":
		if lval > rval {
			return TRUE
		} else {
			return FALSE
		}
	case "==":
		if lval == rval {
			return TRUE
		} else {
			return FALSE
		}
	case "!=":
		if lval != rval {
			return TRUE
		} else {
			return FALSE
		}
	default:
		return NULL
	}
}
