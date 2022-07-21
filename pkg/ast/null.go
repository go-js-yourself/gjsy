package ast

import "github.com/go-js-yourself/gjsy/pkg/token"

type Null struct {
	Token token.Token
}

func (*Null) expressionNode() {}

func (n *Null) TokenLiteral() string {
	return n.Token.Literal
}

func (n *Null) String() string {
	return n.Token.Literal
}
