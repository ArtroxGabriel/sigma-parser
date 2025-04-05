package ast

import (
	"bytes"
	"function_parser/token"
)

// Node is the base interface for all AST nodes
type Node interface {
	TokenLiteral() string
	String() string
}

// Expression represents any mathematical expression
type Expression interface {
	Node
	expressionNode()
}

// NumberLiteral represents a literal number (integer or decimal)
type NumberLiteral struct {
	Token token.Token
	Value float64 // Using float64 to support decimal numbers
}

func (*NumberLiteral) expressionNode()         {}
func (nl *NumberLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NumberLiteral) String() string       { return nl.Token.Literal }

// PrefixExpression represents unary prefix operations (e.g., -5)
type PrefixExpression struct {
	Token    token.Token // The prefix token (-, +)
	Operator string      // The operator itself
	Right    Expression  // The expression to the right
}

func (*PrefixExpression) expressionNode()         {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression represents binary operations (e.g., 2+3, 4*5)
type InfixExpression struct {
	Token    token.Token // The operator token
	Left     Expression  // The expression to the left
	Operator string      // The operator (+, -, *, /, ^)
	Right    Expression  // The expression to the right
}

func (*InfixExpression) expressionNode()         {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// FunctionCall represents a mathematical function call (e.g., sin(x))
type FunctionCall struct {
	Token    token.Token // The function token
	Function Expression  // Function name (sin, cos, etc.)
	Argument Expression  // The function argument
}

func (*FunctionCall) expressionNode()         {}
func (fc *FunctionCall) TokenLiteral() string { return fc.Token.Literal }
func (fc *FunctionCall) String() string {
	var out bytes.Buffer

	out.WriteString(fc.Function.String())
	out.WriteString("(")
	out.WriteString(fc.Argument.String())
	out.WriteString(")")

	return out.String()
}

// Identifier represents a variable like x, y, z
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (*Identifier) expressionNode()        {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// Constant represents mathematical constants like PI and E
type Constant struct {
	Token token.Token
	Name  string  // Constant name (PI, E)
	Value float64 // Constant value
}

func (*Constant) expressionNode()        {}
func (c *Constant) TokenLiteral() string { return c.Token.Literal }
func (c *Constant) String() string       { return c.Name }

// Function is the root node containing the complete mathematical expression
type Function struct {
	Expression Expression // The complete mathematical expression
}

func (me *Function) TokenLiteral() string {
	if me.Expression != nil {
		return me.Expression.TokenLiteral()
	}
	return ""
}

func (me *Function) String() string {
	if me.Expression != nil {
		return me.Expression.String()
	}
	return ""
}
