package ast

import "github.com/go-js-yourself/gjsy/pkg/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

func (*StringLiteral) expressionNode() {}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}
