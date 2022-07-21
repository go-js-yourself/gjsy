package ast

import (
	"strings"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type FunctionExpression struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Expression *ClosureStatement
}

func (*FunctionExpression) expressionNode() {}

func (fe *FunctionExpression) TokenLiteral() string {
	return fe.Token.Literal
}

func (fe *FunctionExpression) String() string {
	out := fe.TokenLiteral()

	if fe.Name != nil {
		out += " " + fe.Name.String()
	}

	params := make([]string, len(fe.Parameters))
	for i, p := range fe.Parameters {
		params[i] = p.String()
	}

	return out + "(" + strings.Join(params, ", ") + ")" + fe.Expression.String()
}
