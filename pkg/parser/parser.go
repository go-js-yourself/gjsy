package parser

import (
	"fmt"
	"strconv"

	"github.com/go-js-yourself/gjsy/pkg/ast"
	"github.com/go-js-yourself/gjsy/pkg/lexer"
	"github.com/go-js-yourself/gjsy/pkg/token"
)

const (
	_ int = iota
	LOWEST
	EQ      // ==
	LTGT    // > or <
	OR      // ||
	AND     // &&
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	APP     // func(X)
)

var undefined = &ast.Undefined{Token: token.Token{
	Type:    token.UNDEF,
	Literal: string("undefined"),
}}

var precedences = map[token.TokenType]int{
	token.EQ:     EQ,
	token.NEQ:    EQ,
	token.LT:     LTGT,
	token.LTE:    LTGT,
	token.GT:     LTGT,
	token.GTE:    LTGT,
	token.OR:     AND,
	token.AND:    OR,
	token.PLUS:   SUM,
	token.MINUS:  SUM,
	token.MOD:    PRODUCT,
	token.DIV:    PRODUCT,
	token.TIMES:  PRODUCT,
	token.LPAREN: APP,
}

type (
	expr   func() ast.Expression
	opExpr func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	exprs   map[token.TokenType]expr
	opExprs map[token.TokenType]opExpr
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: make([]string, 0),
	}

	p.exprs = make(map[token.TokenType]expr)
	p.registerExpr(token.IDENT, p.parseIdentifier)
	p.registerExpr(token.INT, p.parseIntegerLiteral)
	p.registerExpr(token.TRUE, p.parseBoolean)
	p.registerExpr(token.FALSE, p.parseBoolean)
	p.registerExpr(token.NOT, p.parsePrefixExpr)
	p.registerExpr(token.MINUS, p.parsePrefixExpr)
	p.registerExpr(token.IF, p.parseIfExpression)
	p.registerExpr(token.WHILE, p.parseWhileExpression)
	p.registerExpr(token.FUNCTION, p.parseFunctionExpression)
	p.registerExpr(token.UNDEF, p.parseUndefined)
	p.registerExpr(token.NULL, p.parseNull)
	p.registerExpr(token.GO, p.parseGoExpression)
	p.registerExpr(token.STRING, p.parseStringLiteral)

	p.opExprs = make(map[token.TokenType]opExpr)
	p.registerOpExpr(token.PLUS, p.parseOpExpr)
	p.registerOpExpr(token.MINUS, p.parseOpExpr)
	p.registerOpExpr(token.DIV, p.parseOpExpr)
	p.registerOpExpr(token.TIMES, p.parseOpExpr)
	p.registerOpExpr(token.MOD, p.parseOpExpr)
	p.registerOpExpr(token.EQ, p.parseOpExpr)
	p.registerOpExpr(token.NEQ, p.parseOpExpr)
	p.registerOpExpr(token.LT, p.parseOpExpr)
	p.registerOpExpr(token.GT, p.parseOpExpr)
	p.registerOpExpr(token.LTE, p.parseOpExpr)
	p.registerOpExpr(token.GTE, p.parseOpExpr)
	p.registerOpExpr(token.AND, p.parseOpExpr)
	p.registerOpExpr(token.OR, p.parseOpExpr)
	p.registerOpExpr(token.LPAREN, p.parseApplicationExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	p.errors = append(
		p.errors,
		fmt.Sprintf("expected next token to be %s, got %s", t, p.peekToken.Type),
	)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.VAR:
		return p.parseLetStatement()
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.SEMICOLON:
		return nil
	}
	return p.parseExpressionStatement()
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.peekTokenIs(token.ASSIGN) {
		if !p.expectPeek(token.SEMICOLON) {
			return nil
		}
		stmt.Value = undefined
		return stmt
	}
	p.nextToken()
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for p.curTokenIs(token.SEMICOLON) {
		stmt.ReturnValue = undefined
		p.nextToken()
		return stmt
	}

	stmt.ReturnValue = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) registerExpr(t token.TokenType, e expr) {
	p.exprs[t] = e
}

func (p *Parser) registerOpExpr(t token.TokenType, e opExpr) {
	p.opExprs[t] = e
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	expr := p.exprs[p.curToken.Type]

	if expr == nil {
		p.noExpressionError(p.curToken.Type)
		return nil
	}
	leftExpr := expr()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		op := p.opExprs[p.peekToken.Type]
		if op == nil {
			return leftExpr
		}
		p.nextToken()
		leftExpr = op(leftExpr)
	}

	return leftExpr
}

func (p *Parser) noExpressionError(t token.TokenType) {
	msg := fmt.Sprintf("no expression for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{Token: p.curToken}
}

func (p *Parser) parseUndefined() ast.Expression {
	return &ast.Undefined{Token: p.curToken}
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseOpExpr(left ast.Expression) ast.Expression {
	e := &ast.OperationExpression{
		Token:    p.curToken,
		Left:     left,
		Operator: p.curToken.Literal,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	e.Right = p.parseExpression(precedence)

	return e
}

func (p *Parser) parsePrefixExpr() ast.Expression {
	e := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	e.Right = p.parseExpression(PREFIX)

	return e
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expression.Expression = p.parseClosureStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		expression.ElseExpression = p.parseClosureStatement()
	}

	return expression
}

func (p *Parser) parseWhileExpression() ast.Expression {
	expression := &ast.WhileExpression{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expression.Expression = p.parseClosureStatement()
	return expression
}

func (p *Parser) parseClosureStatement() *ast.ClosureStatement {
	block := &ast.ClosureStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}
	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseFunctionExpression() ast.Expression {
	fn := &ast.FunctionExpression{Token: p.curToken}

	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
	} else {
		if !p.expectPeek(token.IDENT) {
			return nil
		}

		fn.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		if !p.expectPeek(token.LPAREN) {
			return nil
		}
	}

	fn.Parameters = p.parseFunctionParameters()
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	fn.Expression = p.parseClosureStatement()
	return fn
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	params := make([]*ast.Identifier, 0)

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return params
	}

	p.nextToken()
	params = append(params, &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	})

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		params = append(params, &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		})
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return params
}

func (p *Parser) parseGoExpression() ast.Expression {
	exp := &ast.GoFunctionApplication{Token: p.curToken}
	p.nextToken()

	var fa interface{} = p.parseExpression(LOWEST)
	switch v := fa.(type) {
	case *ast.FunctionApplication:
		exp.FunctionApplication = v
		return exp
	default:
		p.errors = append(
			p.errors,
			fmt.Sprintf("Expected function application, got %T", v),
		)
	}

	return nil

}

func (p *Parser) parseApplicationExpression(function ast.Expression) ast.Expression {
	exp := &ast.FunctionApplication{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()

	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}
