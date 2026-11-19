package lexer

import (
	"fmt"

	"github.com/mattyco/marble/source/position"
)

// Token contains the type, string value, and the position in the code block.
type Token struct {
	Type     TokenType
	Value    string
	Position *position.Position
}

// NewToken creates a new token object.
func NewToken(ttype TokenType, value string, pos *position.Position) *Token {
	if pos == nil {
		pos = position.NewPosition(-1, -1)
	}

	return &Token{
		Type:     ttype,
		Value:    value,
		Position: pos.Clone(),
	}
}

// String returns the string representation of the token.
func (t *Token) String() string {
	return fmt.Sprintf("type: %d value: \"%s\" position: (%s)", t.Type, t.Value, t.Position.String())
}
