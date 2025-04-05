package parser_test

import (
	"testing"

	"github.com/ArtroxGabriel/sigma-parser/ast"
	"github.com/ArtroxGabriel/sigma-parser/lexer"
	"github.com/ArtroxGabriel/sigma-parser/parser"
)

func TestIdentifierExpression(t *testing.T) {
	input := "sqrt"

	l := lexer.New(input)
	p := parser.New(l)
	mathExpression := p.ParseMathExpression()
	checkParserErrors(t, p)

	ident, ok := mathExpression.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf(
			"exp is not *ast.Identifier. got=%T",
			mathExpression.Expression)
	}
	if ident.Value != input {
		t.Errorf("ident.Value not %s. got=%s", "sqrt", ident.Value)
	}
	if ident.TokenLiteral() != input {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "sqrt",
			ident.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d erros", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}
