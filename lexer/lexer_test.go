package lexer_test

import (
	"testing"

	"github.com/ArtroxGabriel/sigma-parser/lexer"
	"github.com/ArtroxGabriel/sigma-parser/token"
)

func TestNextToken_ReadOneToken(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  token.Token
	}{
		{name: "NUMBER integer token", input: "69", want: token.Token{Type: token.NUMBER, Literal: "69"}},
		{name: "NUMBER float token", input: "3.14", want: token.Token{Type: token.NUMBER, Literal: "3.14"}},
		{name: "PLUS token", input: "+", want: token.Token{Type: token.PLUS, Literal: "+"}},
		{name: "MINUS token", input: "-", want: token.Token{Type: token.MINUS, Literal: "-"}},
		{name: "TIMES token", input: "*", want: token.Token{Type: token.TIMES, Literal: "*"}},
		{name: "SLASH token", input: "/", want: token.Token{Type: token.SLASH, Literal: "/"}},
		{name: "POWER token", input: "^", want: token.Token{Type: token.POWER, Literal: "^"}},
		{name: "SIN token", input: "sin", want: token.Token{Type: token.IDENT, Literal: "sin"}},
		{name: "COS token", input: "cos", want: token.Token{Type: token.IDENT, Literal: "cos"}},
		{name: "TAN token", input: "tan", want: token.Token{Type: token.IDENT, Literal: "tan"}},
		{name: "SQRT token", input: "sqrt", want: token.Token{Type: token.IDENT, Literal: "sqrt"}},
		{name: "LN token", input: "ln", want: token.Token{Type: token.IDENT, Literal: "ln"}},
		{name: "LOG token", input: "log", want: token.Token{Type: token.IDENT, Literal: "log"}},
		{name: "LOG2 token", input: "log2", want: token.Token{Type: token.IDENT, Literal: "log2"}},
		{name: "ABS token", input: "abs", want: token.Token{Type: token.IDENT, Literal: "abs"}},
		{name: "EXP token", input: "exp", want: token.Token{Type: token.IDENT, Literal: "exp"}},
		{name: "PI token", input: "pi", want: token.Token{Type: token.IDENT, Literal: "pi"}},
		{name: "E token", input: "e", want: token.Token{Type: token.IDENT, Literal: "e"}},
		{name: "LPAREN token", input: "(", want: token.Token{Type: token.LPAREN, Literal: "("}},
		{name: "RPAREN token", input: ")", want: token.Token{Type: token.RPAREN, Literal: ")"}},
		{name: "ACOS token", input: "acos", want: token.Token{Type: token.IDENT, Literal: "acos"}},
		{name: "ATAN token", input: "atan", want: token.Token{Type: token.IDENT, Literal: "atan"}},
		{name: "NUMBER token", input: "123", want: token.Token{Type: token.NUMBER, Literal: "123"}},
		{name: "ILLEGAL token", input: "@", want: token.Token{Type: token.ILLEGAL, Literal: "@"}},
		{name: "EOF token", input: "", want: token.Token{Type: token.EOF, Literal: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			got := l.NextToken()
			if got.Type != tt.want.Type {
				t.Errorf("NextToken().Type = %v, want %v", got.Type, tt.want.Type)
			}
			if got.Literal != tt.want.Literal {
				t.Errorf("NextToken().Literal = %v, want %v", got.Literal, tt.want.Literal)
			}
		})
	}
}

func TestNextToken_ReadExpression(t *testing.T) {
	tests := []struct {
		input string
		want  []token.Token
	}{
		{
			input: "10 + 5",
			want: []token.Token{
				{Type: token.NUMBER, Literal: "10"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: "5"},
			},
		},
		{
			input: "3.14 * pi",
			want: []token.Token{
				{Type: token.NUMBER, Literal: "3.14"},
				{Type: token.TIMES, Literal: "*"},
				{Type: token.IDENT, Literal: "pi"},
			},
		},
		{
			input: "3.14.0 + 5",
			want: []token.Token{
				{Type: token.NUMBER, Literal: "3.14"},
				{Type: token.ILLEGAL, Literal: "."},
				{Type: token.NUMBER, Literal: "0"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: "5"},
			},
		},
		{
			input: "sin(90) + cos(0)",
			want: []token.Token{
				{Type: token.IDENT, Literal: "sin"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.NUMBER, Literal: "90"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.IDENT, Literal: "cos"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.NUMBER, Literal: "0"},
				{Type: token.RPAREN, Literal: ")"},
			},
		},
		{
			input: "log2(8) / 2",
			want: []token.Token{
				{Type: token.IDENT, Literal: "log2"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.NUMBER, Literal: "8"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.NUMBER, Literal: "2"},
			},
		},
		{
			input: "sqrt(4) ^ 2",
			want: []token.Token{
				{Type: token.IDENT, Literal: "sqrt"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.NUMBER, Literal: "4"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.POWER, Literal: "^"},
				{Type: token.NUMBER, Literal: "2"},
			},
		},
		{
			input: "(1 + 2) * (3 - 4) / 5",
			want: []token.Token{
				{Type: token.LPAREN, Literal: "("},
				{Type: token.NUMBER, Literal: "1"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.NUMBER, Literal: "2"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.TIMES, Literal: "*"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.NUMBER, Literal: "3"},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.NUMBER, Literal: "4"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.NUMBER, Literal: "5"},
			},
		},
		{
			input: "tan(e) + log(100) @",
			want: []token.Token{
				{Type: token.IDENT, Literal: "tan"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "e"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.IDENT, Literal: "log"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.NUMBER, Literal: "100"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.ILLEGAL, Literal: "@"},
			},
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		for _, want := range tt.want {
			got := l.NextToken()
			if got.Type != want.Type {
				t.Errorf("NextToken().Type = %v, want %v", got.Type, want.Type)
			}
			if got.Literal != want.Literal {
				t.Errorf("NextToken().Literal = %v, want %v", got.Literal, want.Literal)
			}
		}
	}
}
