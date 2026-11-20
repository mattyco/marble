package errors

import (
	"fmt"

	"github.com/mattyco/marble/source/lexer/position"
)

// Error is the template for a language error.
type Error interface {
	String() string
}

// BaseError is the error base for each individual error.
type BaseError struct {
	Name     string
	Message  string
	Position *position.Position
}

// String returns the error in a string format; "Name: message at line 1, column 1".
func (b *BaseError) String() string {
	return fmt.Sprintf("%s: %s at %s", b.Name, b.Message, b.Position.String())
}

type SyntaxError struct {
	BaseError
}

// String returns the error in a string format; "Name: message at line 1, column 1".
func (b *SyntaxError) String() string {
	return fmt.Sprintf("SyntaxError: %s at %s", b.Message, b.Position.String())
}
