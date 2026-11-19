package errors

import (
	"fmt"

	"github.com/mattyco/marble/source/position"
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

func (b *BaseError) String() string {
	return fmt.Sprintf("%s: %s at %s", b.Name, b.Message, b.Position.String())
}

type SyntaxError struct {
	BaseError
}
