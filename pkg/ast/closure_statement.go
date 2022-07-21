package ast

import (
	"bytes"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type ClosureStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *ClosureStatement) statementNode() {}

func (bs *ClosureStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *ClosureStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
