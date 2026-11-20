package parser

import (
	"github.com/mattyco/marble/source/errors"
	"github.com/mattyco/marble/source/lexer"
	"github.com/mattyco/marble/source/parser/ast"
)

// Parser goes through each lexer token and converts them into valid AST nodes.
type Parser struct {
	tokens   []*lexer.Token
	current  *lexer.Token
	position int
}

// NewParser creates a.new parser instance.
func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{
		tokens:   tokens,
		current:  nil,
		position: -1,
	}
}

// nextToken shifts to the next token in the sequence.
func (p *Parser) nextToken() {
	p.position++
	if len(p.tokens) <= p.position {
		p.current = nil
		return
	}

	p.current = p.tokens[p.position]
}

// peekNextToken returns the next token in the sequence.
func (p *Parser) peekNextToken() *lexer.Token {
	if p.position+1 >= len(p.tokens) {
		return nil
	}

	return p.tokens[p.position+1]
}

// Parse runs the language parser.
func (p *Parser) Parse() (ast.Node, errors.Error) {
	p.nextToken()

	nodes := make([]ast.Node, 0)

	

	return &ast.Program{Nodes: nodes}, nil
}
