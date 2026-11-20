package ast

import (
	"github.com/mattyco/marble/source/lexer"
)

type Node interface {
	String() string
}

type Program struct {
	Nodes []Node
}

type Body struct {
	Nodes []Node
}

type Integer struct {
	Value *lexer.Token
}

type Bool struct {
	Value *lexer.Token
}

type String struct {
	Value *lexer.Token
}

type Float struct {
	Value *lexer.Token
}

type List struct {
	Elements []Node
}

type IndexExpression struct {
	Left  Node // list
	Index Node // should be integer.
}

type SliceExpression struct {
	List  Node // list
	Start Node // can be empty so do keep in mind.
	End   Node // can be empty so do keep in mind.
}

type Identifier struct {
	Value *lexer.Token
}

type CallExpression struct {
	Value     Node
	Arguments []Node
}

type BinaryExpression struct {
	Left     Node
	Right    Node
	Operator *lexer.Token
}

type VariableDeclaration struct {
	Identifier *lexer.Token
	Value      Node
}

type ConstantDeclaration struct {
	Identifier *lexer.Token
	Value      Node
}

type Condition struct {
	Condition   Node
	Consequence Node
}

type Conditional struct {
	Match   *Condition   // if
	Other   []*Condition // elif branches
	NoMatch Node         // else
}

type Argument struct {
	Identifier *lexer.Token // foo
	Type       *lexer.Token // int, string, float, bool, list
}

type FunctionDeclare struct {
	Identifier *lexer.Token
	Arguments  []*Argument
	Body       Node
}

type WhileCondition struct {
	Condition Node
	Body      Node
}

type Return struct {
	Token *lexer.Token
	Value Node
}

type Continue struct {
	Token *lexer.Token
}

type Break struct {
	Token *lexer.Token
}
