package ast

import "github.com/go-js-yourself/gjsy/pkg/token"

type GoFunctionApplication struct {
	Token token.Token
	*FunctionApplication
}
