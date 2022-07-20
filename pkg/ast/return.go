package ast

import "github.com/go-js-yourself/gjsy/pkg/token"

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (*ReturnStatement) statementNode()          {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
