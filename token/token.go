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

	IDENT  TokenType = "IDENT"
	NUMBER TokenType = "NUMBER"

	SIN  TokenType = "SIN"
	COS  TokenType = "COS"
	TAN  TokenType = "TAN"
	ASIN TokenType = "ASIN"
	ACOS TokenType = "ACOS"
	ATAN TokenType = "ATAN"
	SQRT TokenType = "SQRT"

	LN   TokenType = "LN"
	LOG  TokenType = "LOG"
	LOG2 TokenType = "LOG2"

	ABS TokenType = "ABS"
	EXP TokenType = "EXP"

	PI TokenType = "PI"
	E  TokenType = "E"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
)
