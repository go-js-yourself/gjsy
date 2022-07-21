package ast

import (
	"fmt"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type GoFunctionApplication struct {
	Token token.Token
	*FunctionApplication
}

func (gfa *GoFunctionApplication) String() string {
	return fmt.Sprintf("go %s", gfa.FunctionApplication.String())
}
