package ast

import (
	"bytes"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (*ReturnStatement) statementNode()          {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
