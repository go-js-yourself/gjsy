package ast

import (
	"fmt"

	"github.com/go-js-yourself/gjsy/pkg/token"
)

type WhileExpression struct {
	Token      token.Token
	Condition  Expression
	Expression *ClosureStatement
}

func (we *WhileExpression) expressionNode() {}

func (we *WhileExpression) TokenLiteral() string {
	return we.Token.Literal
}

func (we *WhileExpression) String() string {
	return fmt.Sprintf("while (%s) %s", we.Condition.String(), we.Expression.String())
}
