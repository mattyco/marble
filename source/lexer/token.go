package lexer

import "fmt"

// Token contains the type, string value, and the position in the code block.
type Token struct {
	Type     TokenType
	value    string
	Position *Position
}

// NewToken creates a new token object.
func NewToken(ttype TokenType, value string, position *Position) *Token {
	return &Token{
		Type:     ttype,
		value:    value,
		Position: position.Clone(),
	}
}

// String returns the string representation of the token.
func (t *Token) String() string {
	return fmt.Sprintf("type: %d \"%s\" position: (%s)", t.Type, t.value, t.Position.String())
}
