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
func (l *Lexer) Tokenise() ([]*Token, error) {
	if l.position == nil {
		// This ensures that the position of the token can be stored as intended, and also so that
		// the nextCharacter function can process the position instead of dereferencing a nilpointer.
		l.position = NewPosition(-1, -1)
	}

	if l.idx != -1 {
		// this is so it starts properly from the beginning of the code block.
		l.idx = -1
	}

	// start at index 0, with the initial character.
	l.nextCharacter()

	var tokens []*Token = make([]*Token, 0)

	for l.current != nil {
		switch *l.current {
		// Recognised whitespace characters that the lexer ignores.
		case '\r', '\t', ' ':
			l.nextCharacter()
			continue
		
		
		}
	}

	return tokens, nil
}
