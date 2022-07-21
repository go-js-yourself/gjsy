package ast

import (
	"github.com/go-js-yourself/gjsy/pkg/token"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	out := ls.TokenLiteral() + " " + ls.Name.String()

	if ls.Value != nil {
		out += " = " + ls.Value.String()
	}

	return out + ";"
}
