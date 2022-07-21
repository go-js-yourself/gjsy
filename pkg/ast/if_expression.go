package ast

import (
	"github.com/go-js-yourself/gjsy/pkg/token"
)

type IfExpression struct {
	Token          token.Token
	Condition      Expression
	Expression     *ClosureStatement
	ElseExpression *ClosureStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	out := "if " + ie.Condition.String() + " " + ie.Expression.String()

	if ie.ElseExpression != nil {
		out += " else " + ie.ElseExpression.String()
	}

	return out
}
