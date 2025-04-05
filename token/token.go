package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	PLUS  TokenType = "+"
	MINUS TokenType = "-"
	TIMES TokenType = "*"
	SLASH TokenType = "/"
	POWER TokenType = "^"

	IDENT  TokenType = "IDENT" // functions and variables
	NUMBER TokenType = "NUMBER"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
)
