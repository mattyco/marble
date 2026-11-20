package position

import "fmt"

// Position contains the x, and y coordinate inside of the code block.
type Position struct {
	X, Y int
}

// NewPosition creates a new position object.
func NewPosition(x, y int) *Position {
	return &Position{X: x, Y: y}
}

// Clone returns a duplicated position object.
func (p *Position) Clone() *Position {
	return &Position{
		X: p.X,
		Y: p.Y,
	}
}

// String returns the string varient of the position.
func (p *Position) String() string {
	return fmt.Sprintf("line %d, column %d", p.Y, p.X)
}
