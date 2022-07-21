package ast

import (
	"fmt"
	"strings"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type FunctionApplication struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (*FunctionApplication) expressionNode() {}

func (fa *FunctionApplication) TokenLiteral() string {
	return fa.Token.Literal
}

func (fa *FunctionApplication) String() string {
	args := make([]string, len(fa.Arguments))
	for i, a := range fa.Arguments {
		args[i] = a.String()
	}
	return fmt.Sprintf("%s(%s)", fa.Function.String(), strings.Join(args, ", "))
}
