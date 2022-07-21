package ast

import (
	"fmt"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type OperationExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (*OperationExpression) expressionNode() {}

func (oe *OperationExpression) TokenLiteral() string {
	return oe.Token.Literal
}

func (oe *OperationExpression) String() string {
	return fmt.Sprintf(
		"(%s %s %s);",
		oe.Left.String(),
		oe.Operator,
		oe.Right.String(),
	)
}
