package lexer

import "fmt"

// Position contains the x, and y coordinate inside of the code block.
type Position struct {
	x, y int
}

// NewPosition creates a new position object.
func NewPosition(x, y int) *Position {
	return new(Position)
}

// Clone returns a duplicated position object.
func (p *Position) Clone() *Position {
	return &Position{
		x: p.x,
		y: p.y,
	}
}

// String returns the string varient of the position.
func (p *Position) String() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}
