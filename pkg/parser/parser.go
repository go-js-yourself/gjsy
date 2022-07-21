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
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL    // func(X)
)

var precedences = map[token.TokenType]int{
	token.EQ:    EQ,
	token.NEQ:   EQ,
	token.LT:    LTGT,
	token.GT:    LTGT,
	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.MOD:   PRODUCT,
	token.DIV:   PRODUCT,
	token.TIMES: PRODUCT,
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

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
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
