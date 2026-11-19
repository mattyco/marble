package lexer

// Lexer reads through the code block and tokenises it into parseable tokens.
type Lexer struct {
	code []rune
	idx  int

	position *Position
	current  *rune
}

// NewLexer creates a new lexer object.
func NewLexer(code []rune) *Lexer {
	return &Lexer{
		code: code,

		idx:      -1,
		position: NewPosition(-1, -1),
		current:  nil,
	}
}

// nextCharacter shifts to the next character in the sequence.
func (l *Lexer) nextCharacter() {
	l.idx++
	l.position.x++

	if l.idx >= len(l.code) {
		l.current = nil
		return
	}

	l.current = &l.code[l.idx]

	if *l.current == '\n' {
		l.position.x = 0
		l.position.y++
	}
}

// peekNextCharacter returns the next character in the sequence without enforcing the change.
func (l *Lexer) peekNextCharacter() *rune {
	if l.idx+1 < len(l.code) {
		return &l.code[l.idx+1]
	}

	return nil
}

// Tokenise converts the plan-text into parseable tokens.
func (t *Token) Tokenise() ([]*Token, error) {
	var tokens []*Token = make([]*Token, 0)

	return tokens, nil
}
