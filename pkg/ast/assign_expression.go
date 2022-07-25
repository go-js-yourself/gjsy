package ast

import (
	"fmt"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type AssignExpression struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (*AssignExpression) expressionNode() {}

func (as *AssignExpression) TokenLiteral() string {
	return as.Token.Literal
}

func (as *AssignExpression) String() string {
	return fmt.Sprintf("%s = %s;", as.Name.String(), as.Value.String())
}
