package parser

import (
	"fmt"
	"function_parser/ast"
	"function_parser/lexer"
	"function_parser/token"
	"strconv"
)

const (
	// predefined precedence levels for operators
	_ int = iota
	LOWEST
	SUM      // sum or subtraction (+, -)
	PRODUCT  // product or division (*, /)
	POWER    // power or squar root (^, sqrt)
	PREFIX   // negative numbers (- or + unary)
	CALL     // call functions(sin, cos, ln, etc.)
	GROUPING // expression in parentheses ()
)

var precedences = map[token.TokenType]int{
	token.PLUS:   SUM,
	token.MINUS:  SUM,
	token.SLASH:  PRODUCT,
	token.TIMES:  PRODUCT,
	token.POWER:  POWER,
	token.SQRT:   POWER,
	token.LPAREN: CALL,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.NUMBER, p.parserNumberLiteral)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.TIMES, p.parseInfixExpression)
	p.registerInfix(token.POWER, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)

	p.nextToken() // populate currToken and peekToken
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseMathExpression() *ast.Function {
	mathExpression := new(ast.Function)

	mathExpression.Expression = p.parseExpression(LOWEST)

	return mathExpression
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.EOF) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) parserNumberLiteral() ast.Expression {
	value, err := strconv.ParseFloat(p.currToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	return &ast.NumberLiteral{Token: p.currToken, Value: value}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
		Left:     left,
	}

	precedence := p.currPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.FunctionCall{Token: p.currToken, Function: function}
	exp.Argument = p.parseExpressionParam(token.RPAREN)
	return exp
}

func (p *Parser) parseExpressionParam(end token.TokenType) ast.Expression {
	if p.peekTokenIs(end) {
		p.nextToken()
		return nil
	}

	p.nextToken()
	return p.parseExpression(LOWEST)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekErrors(t)
	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool { return p.peekToken.Type == t }
func (p *Parser) currTokenIs(t token.TokenType) bool { return p.currToken.Type == t }

func (p *Parser) Errors() []string { return p.errors }

func (p *Parser) peekErrors(t token.TokenType) {
	msg := fmt.Sprintf(
		"Expected next token to be %s, got %s instead",
		t,
		p.peekToken.Type,
	)

	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	p.errors = append(p.errors,
		fmt.Sprintf("no prefix parse function for %s found", t),
	)
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
