package ast

import (
	"bytes"

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
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}
