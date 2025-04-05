package lexer

import "function_parser/token"

// Lexer represents a lexical analyzer for tokenizing input strings.
type Lexer struct {
	input        string // The input string to be tokenized
	position     int    // Current position in input (points to current char)
	readPosition int    // Current reading position in input (after current char)
	ch           byte   // Current character under examination
}

// New creates a new Lexer instance with the given input string.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Initialize the lexer by reading the first character
	return l
}

// peekChar returns the next character in the input without advancing the position.
// If the end of the input is reached, it returns 0.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// readChar advances the lexer to the next character in the input.
// If the end of the input is reached, it sets the current character to 0.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// skipWhitespace skips over whitespace characters in the input.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// NextToken retrieves the next token from the input and advances the lexer.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.TIMES, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '^':
		tok = newToken(token.POWER, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.IDENT
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()

	return tok
}

// isDigit checks if the given character is a numeric digit (0-9).
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// isLetter checks if the given character is a letter (a-z, A-Z) or an underscore (_).
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// readNumber reads a sequence of numeric digits (and at most one dot for decimals) from the input.
// It returns the number as a string.
func (l *Lexer) readNumber() string {
	position := l.position
	seenDot := false

	for isDigit(l.ch) || (l.ch == '.' && !seenDot) {
		if l.ch == '.' {
			seenDot = true
		}
		l.readChar()
	}

	return l.input[position:l.position]
}

// readIdentifier reads a sequence of letters or underscores from the input and returns it as a string.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// newToken creates a new token with the given type and character.
func newToken(tokenTypen token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenTypen, Literal: string(ch)}
}
