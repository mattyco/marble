package errors

import (
	"fmt"

	"github.com/mattyco/marble/source/position"
)

// Error is the template for a language error.
type Error interface {
	String() string
}

// SyntaxError contains the data that is returned via a syntax error.
type SyntaxError struct {
	Message  string
	Position *position.Position
}

func (s *SyntaxError) String() string {
	return fmt.Sprintf("Syntax Error: %s at %s", s.Message, s.Position.String())
}
