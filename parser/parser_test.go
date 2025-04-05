package parser_test

import (
	"testing"

	"github.com/ArtroxGabriel/sigma-parser/ast"
	"github.com/ArtroxGabriel/sigma-parser/lexer"
	"github.com/ArtroxGabriel/sigma-parser/parser"
)

func TestIdentifierExpression(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"x"},
		{"sqrt"},
		{"log2"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		mathExpression := p.ParseFunction()

		checkParserErrors(t, p)
		ident, ok := mathExpression.Expression.(*ast.Identifier)
		if !ok {
			t.Fatalf(
				"exp is not *ast.Identifier. got=%T",
				mathExpression.Expression)
		}
		if ident.Value != tt.input {
			t.Errorf("ident.Value not %s. got=%s", tt.input, ident.Value)
		}
		if ident.TokenLiteral() != tt.input {
			t.Errorf("ident.TokenLiteral not %s. got=%s", tt.input,
				ident.TokenLiteral())
		}

	}
}

func TestNumberLiteralExpression(t *testing.T) {
	tests := []struct {
		input string
		want  float64
	}{
		{input: "5", want: 5},
		{input: "3.14", want: 3.14},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		mathExpression := p.ParseFunction()

		checkParserErrors(t, p)
		ident, ok := mathExpression.Expression.(*ast.NumberLiteral)
		if !ok {
			t.Fatalf(
				"exp is not *ast.Identifier. got=%T",
				mathExpression.Expression)
		}
		if ident.Value != tt.want {
			t.Errorf("ident.Value not %f. got=%f", tt.want, ident.Value)
		}
		if ident.TokenLiteral() != tt.input {
			t.Errorf("ident.TokenLiteral not %f. got=%s", tt.want,
				ident.TokenLiteral())
		}

	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTest := []struct {
		input        string
		operator     string
		value        float64
		valueLiteral string
	}{
		{"-15", "-", 15, "15"},
		{"-2.18", "-", 2.18, "2.18"},
	}

	for _, tt := range prefixTest {
		l := lexer.New(tt.input)
		p := parser.New(l)
		function := p.ParseFunction()
		checkParserErrors(t, p)

		exp, ok := function.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf(
				"exp is not *ast.PrefixExpression. got=%T",
				function.Expression)
		}
		if exp.Operator != tt.operator {
			t.Errorf("exp.Operator not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.value, tt.valueLiteral) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTest := []struct {
		input        string
		leftValue    any
		leftLiteral  string
		operator     string
		rightValue   any
		rightLiteral string
	}{
		{"5 + 5;", 5, "5", "+", 5, "5"},
		{"5 - 5;", 5, "5", "-", 5, "5"},
		{"5 * 5;", 5, "5", "*", 5, "5"},
		{"5 / 5;", 5, "5", "/", 5, "5"},
		{"5 ^ 5;", 5, "5", "^", 5, "5"},
		{"x + PI", "x", "x", "+", "PI", "PI"},
	}

	for _, tt := range infixTest {
		l := lexer.New(tt.input)
		p := parser.New(l)
		function := p.ParseFunction()
		checkParserErrors(t, p)

		if !testInfixExpression(
			t,
			function.Expression,
			tt.leftValue,
			tt.leftLiteral,
			tt.operator,
			tt.rightValue,
			tt.rightLiteral,
		) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"a + log2(b * c) + d",
			"((a + log2((b * c))) + d)",
		},
		{
			"sqrt(a + b + c * d / f + g)",
			"sqrt((((a + b) + ((c * d) / f)) + g))",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			checkParserErrors(t, p)
			function := p.ParseFunction()

			actual := function.String()
			if actual != tt.expected {
				t.Errorf("expected %q. got=%q", tt.expected, actual)
			}
		})
	}
}

func testInfixExpression(
	t *testing.T,
	exp ast.Expression,
	left any,
	leftLiteral string,
	operator string,
	right any,
	rightLiteral string,
) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left, leftLiteral) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("Exp.Operator is not '%s'. got=%q", operator,
			opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right, rightLiteral) {
		return false
	}

	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected any,
	expectedLiteral string,
) bool {
	switch v := expected.(type) {
	case float64:
		return testNumberLiteral(t, exp, v, expectedLiteral)
	case int:
		return testNumberLiteral(t, exp, float64(v), expectedLiteral)
	case string:
		return testIdentifier(t, exp, v)
	default:
		t.Errorf("type of exp not handled. got=%T", exp)
		return false
	}
}

func testNumberLiteral(t *testing.T, il ast.Expression, value float64, valueLiteral string) bool {
	integ, ok := il.(*ast.NumberLiteral)
	if !ok {
		t.Errorf(
			"il not *ast.IntegerLiterat. go=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf(
			"integ.Value not %f. got=%f",
			value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != valueLiteral {
		t.Errorf(
			"integ.TokenLiteral() not %s. got=%s",
			valueLiteral, integ.TokenLiteral())
		return false
	}
	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("iden.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("iden.TokenLiteral() not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
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
