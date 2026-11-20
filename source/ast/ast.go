package ast

import (
	"fmt"
	"strings"

	"github.com/mattyco/marble/source/lexer"
)

type Node interface {
	String() string
}

type Program struct {
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
	Value     *lexer.Token
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

func (p *Program) String() string {
	values := make([]string, 0)

	for _, n := range p.Nodes {
		values = append(values, n.String())
	}

	return strings.Join(values, "\n")
}

func (i *Integer) String() string {
	return fmt.Sprintf("integer(%s)", i.Value.Value)
}

func (f *Float) String() string {
	return fmt.Sprintf("float(%s)", f.Value.Value)
}

func (l *List) String() string {
	return fmt.Sprintf("list[ %s ]", func() string {
		out := make([]string, len(l.Elements))

		for i, e := range l.Elements {
			out[i] = e.String()
		}

		return strings.Join(out, ", ")
	}())
}

func (ie *IndexExpression) String() string {
	return fmt.Sprintf("%s[%s]", ie.Left.String(), ie.Index.String())
}

func (se *SliceExpression) String() string {
	return fmt.Sprintf("%s[%s:%s]", se.List.String(), se.Start.String(), se.End.String())
}

func (b *Bool) String() string {
	return fmt.Sprintf("boolean(%s)", b.Value.Value)
}

func (s *String) String() string {
	return fmt.Sprintf("string(%s)", s.Value.Value)
}

func (c *CallExpression) String() string {
	return fmt.Sprintf("identifier(%s(%s))", c.Value.Value, func() string {
		out := make([]string, len(c.Arguments))
		for i, arg := range c.Arguments {
			out[i] = arg.String()
		}
		return strings.Join(out, ", ")
	}())
}

func (i *Identifier) String() string {
	return fmt.Sprintf("identifier(%s)", i.Value.Value)
}

func (b *BinaryExpression) String() string {
	return fmt.Sprintf("%s %s %s", b.Left.String(), b.Operator.String(), b.Right.String())
}

func (v *VariableDeclaration) String() string {
	return fmt.Sprintf("identifier(%s) = %s", v.Identifier.Value, v.Value.String())
}

func (c *ConstantDeclaration) String() string {
	return fmt.Sprintf("identifier(%s) = %s", c.Identifier.Value, c.Value.String())
}
