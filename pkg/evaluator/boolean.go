package evaluator

import "github.com/go-js-yourself/gjsy/pkg/object"

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
