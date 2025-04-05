package ast_test

import (
	"math"
	"testing"

	"github.com/ArtroxGabriel/sigma-parser/ast"
	"github.com/ArtroxGabriel/sigma-parser/token"
)

func TestString(t *testing.T) {
	tests := []struct {
		mathExpression ast.Function
		expectedString string
	}{
		{
			expectedString: "(5 + 3)",
			mathExpression: ast.Function{
				Expression: &ast.InfixExpression{
					Token:    token.Token{Type: token.PLUS, Literal: "+"},
					Left:     &ast.NumberLiteral{Token: token.Token{Type: token.NUMBER, Literal: "5"}, Value: 5},
					Operator: "+",
					Right:    &ast.NumberLiteral{Token: token.Token{Type: token.NUMBER, Literal: "3"}, Value: 3},
				},
			},
		},
		{
			expectedString: "((-3.14) / e)",
			mathExpression: ast.Function{
				Expression: &ast.InfixExpression{
					Token: token.Token{Type: token.SLASH, Literal: "/"},
					Left: &ast.PrefixExpression{
						Token:    token.Token{Type: token.MINUS, Literal: "-"},
						Operator: "-",
						Right:    &ast.NumberLiteral{Token: token.Token{Type: token.NUMBER, Literal: "3.14"}, Value: 3.14},
					},
					Operator: "/",
					Right:    &ast.Constant{Token: token.Token{Type: token.IDENT, Literal: "e"}, Value: math.E, Name: "e"},
				},
			},
		},
		{
			expectedString: "sin(90)",
			mathExpression: ast.Function{
				Expression: &ast.FunctionCall{
					Token:    token.Token{Type: token.IDENT, Literal: "sin"},
					Function: &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "sin"}, Value: "sin"},
					Argument: &ast.NumberLiteral{Token: token.Token{Type: token.NUMBER, Literal: "90"}, Value: 90},
				},
			},
		},
		{
			expectedString: "((x ^ 2) + (y ^ 2))",
			mathExpression: ast.Function{
				Expression: &ast.InfixExpression{
					Token:    token.Token{Type: token.PLUS, Literal: "+"},
					Operator: "+",
					Left: &ast.InfixExpression{
						Token:    token.Token{Type: token.POWER, Literal: "^"},
						Operator: "^",
						Left:     &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
						Right:    &ast.NumberLiteral{Token: token.Token{Type: token.NUMBER, Literal: "2"}, Value: 2},
					},
					Right: &ast.InfixExpression{
						Token:    token.Token{Type: token.POWER, Literal: "^"},
						Operator: "^",
						Left:     &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "y"}, Value: "y"},
						Right:    &ast.NumberLiteral{Token: token.Token{Type: token.NUMBER, Literal: "2"}, Value: 2},
					},
				},
			},
		},
		{
			expectedString: "(sqrt(16) * pi)",
			mathExpression: ast.Function{
				Expression: &ast.InfixExpression{
					Token:    token.Token{Type: token.TIMES, Literal: "*"},
					Operator: "*",
					Left: &ast.FunctionCall{
						Token:    token.Token{Type: token.IDENT, Literal: "sqrt"},
						Function: &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "sqrt"}, Value: "sqrt"},
						Argument: &ast.NumberLiteral{Token: token.Token{Type: token.NUMBER, Literal: "16"}, Value: 16},
					},
					Right: &ast.Constant{Token: token.Token{Type: token.IDENT, Literal: "pi"}, Value: math.Pi, Name: "pi"},
				},
			},
		},
	}

	for _, tt := range tests {
		if tt.mathExpression.String() != tt.expectedString {
			t.Errorf("expected %s, got %s", tt.expectedString, tt.mathExpression.String())
		}
	}
}
