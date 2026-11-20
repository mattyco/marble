package parser

import (
	"fmt"

	"github.com/mattyco/marble/source/errors"
	"github.com/mattyco/marble/source/lexer"
	"github.com/mattyco/marble/source/parser/ast"
)

// parseVariableDeclare parses the declaration of a variable node.
func (p *Parser) parseVariableDeclare() (ast.Node, errors.Error) {
	base := p.current.Position.Clone()

	p.nextToken()

	node := &ast.VariableDeclaration{}

	if p.current == nil {
		return nil, &errors.SyntaxError{Message: "improper declaration of variable", Position: base.Clone()}
	}

	if p.current.Type == lexer.TokIdentifier {
		node.Identifier = p.current

		p.nextToken()

		if p.current == nil {
			return nil, &errors.SyntaxError{Message: "improper declaration of variable", Position: node.Identifier.Position.Clone()}
		}

		if p.current.Type != lexer.TokKeyword {
			return nil, &errors.SyntaxError{Message: fmt.Sprintf("expected type definition got '%s'", p.current.Value), Position: node.Identifier.Position.Clone()}
		}

		node.Type = p.current

		p.nextToken()

		if p.current == nil {
			return nil, &errors.SyntaxError{Message: "improper declaration of variable", Position: node.Type.Position.Clone()}
		}

		if p.current.Type != lexer.TokAsign {
			return nil, &errors.SyntaxError{Message: fmt.Sprintf("expected assign but got '%s'", p.current.Value), Position: node.Type.Position.Clone()}
		}

		p.nextToken()

		expression, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		node.Value = expression
		return node, nil
	}

	return nil, &errors.SyntaxError{Message: fmt.Sprintf("expected identifier got '%s'", p.current.Value), Position: p.current.Position.Clone()}
}

// parseConstantDeclare parses the declaration of a variable node.
func (p *Parser) parseConstantDeclare() (ast.Node, errors.Error) {
	base := p.current.Position.Clone()

	p.nextToken()

	node := &ast.ConstantDeclaration{}

	if p.current == nil {
		return nil, &errors.SyntaxError{Message: "improper declaration of constant", Position: base.Clone()}
	}

	if p.current.Type == lexer.TokIdentifier {
		node.Identifier = p.current

		p.nextToken()

		if p.current == nil {
			return nil, &errors.SyntaxError{Message: "improper declaration of constant", Position: node.Identifier.Position.Clone()}
		}

		if p.current.Type != lexer.TokKeyword {
			return nil, &errors.SyntaxError{Message: fmt.Sprintf("expected type definition got '%s'", p.current.Value), Position: node.Identifier.Position.Clone()}
		}

		node.Type = p.current

		p.nextToken()

		if p.current == nil {
			return nil, &errors.SyntaxError{Message: "improper declaration of constant", Position: node.Type.Position.Clone()}
		}

		if p.current.Type != lexer.TokAsign {
			return nil, &errors.SyntaxError{Message: fmt.Sprintf("expected assign but got '%s'", p.current.Value), Position: node.Type.Position.Clone()}
		}

		p.nextToken()

		expression, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		node.Value = expression
		return node, nil
	}

	return nil, &errors.SyntaxError{Message: fmt.Sprintf("expected identifier got '%s'", p.current.Value), Position: p.current.Position.Clone()}
}
