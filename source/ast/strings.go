package ast

import (
	"fmt"
	"strings"
)

func (p *Program) String() string {
	values := make([]string, 0)

	for _, n := range p.Nodes {
		values = append(values, n.String())
	}

	return strings.Join(values, "\n")
}

func (b *Body) String() string {
	values := make([]string, 0)

	for _, n := range b.Nodes {
		values = append(values, n.String())
	}

	return "Body\n" + strings.Join(values, "\n")
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
	return fmt.Sprintf("identifier(%s(%s))", c.Value.String(), func() string {
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

func (c *Conditional) String() string {
	var output strings.Builder

	fmt.Fprintf(&output, "Conditional\n")
	fmt.Fprintf(&output, "  Condition: %s\n", c.Match.Condition.String())
	fmt.Fprintf(&output, "  Consequence: %s\n", c.Match.Consequence.String())

	for _, branch := range c.Other {
		fmt.Fprintf(&output, "  Elif:\n%s", branch.Condition.String())
		fmt.Fprintf(&output, "  Consequence: %s\n", branch.Consequence.String())
	}

	if c.NoMatch != nil {
		fmt.Fprintf(&output, "  Else:\n%s", c.NoMatch.String())
	}

	return output.String()
}

func (f *FunctionDeclare) String() string {
	var output strings.Builder

	fmt.Fprintf(&output, "Function\n")
	fmt.Fprintf(&output, "  Name: %s\n", f.Identifier.Value)

	for _, arg := range f.Arguments {
		fmt.Fprintf(&output, "  Argument: %s (%s)\n", arg.Identifier.Value, arg.Type.Value)
	}

	fmt.Fprintf(&output, "  Body: %s\n", f.Body.String())

	return output.String()
}
