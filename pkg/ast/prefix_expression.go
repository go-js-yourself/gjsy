package ast

import (
	"fmt"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (*PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s %s);", pe.Operator, pe.Right.String())
}
