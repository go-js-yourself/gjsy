package ast

import (
	"fmt"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type DotExpression struct {
	Token token.Token
	Left  Expression
	Right Expression
}

func (*DotExpression) expressionNode() {}

func (de *DotExpression) TokenLiteral() string {
	return de.Token.Literal
}

func (de *DotExpression) String() string {
	return fmt.Sprintf("%s.%s", de.Left.String(), de.Right.String())
}
