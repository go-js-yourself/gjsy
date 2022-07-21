package ast

import "github.com/go-js-yourself/gjsy/pkg/token"

type Undefined struct {
	Token token.Token
}

func (*Undefined) expressionNode() {}

func (u *Undefined) TokenLiteral() string {
	return u.Token.Literal
}

func (u *Undefined) String() string {
	return u.Token.Literal
}
