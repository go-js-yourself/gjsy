package evaluator

import "github.com/go-js-yourself/gjsy/pkg/object"

func evalOperationExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntPair(operator, left.(*object.Integer), right.(*object.Integer))
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBoolPair(operator, left.(*object.Boolean), right.(*object.Boolean))
	default:
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntPair(operator string, left *object.Integer, right *object.Integer) object.Object {
	lval := left.Value
	rval := right.Value

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
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalBoolPair(operator string, left *object.Boolean, right *object.Boolean) object.Object {
	switch operator {
	case "==":
		if left.Value == right.Value {
			return TRUE
		} else {
			return FALSE
		}
	case "!=":
		if left.Value != right.Value {
			return TRUE
		} else {
			return FALSE
		}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
