package lexer

import (
	"fmt"
	"unicode"

	"github.com/mattyco/marble/source/errors"
	"github.com/mattyco/marble/source/position"
)

// Lexer reads through the code block and tokenises it into parseable tokens.
type Lexer struct {
	code []rune
	idx  int

	position *position.Position
	current  *rune
}

// NewLexer creates a new lexer object.
func NewLexer(code []rune) *Lexer {
	return &Lexer{
		code:     code,
		idx:      -1,
		position: position.NewPosition(1,1),
		current:  nil,
	}
}

// nextCharacter shifts to the next character in the sequence.
func (l *Lexer) nextCharacter() {
	l.idx++

	if l.idx >= len(l.code) {
		l.current = nil
		return
	}

	l.current = &l.code[l.idx]

	if *l.current == '\n' {
		l.position.X = 0
		l.position.Y++
	}

	l.position.X++
}

// peekNextCharacter returns the next character in the sequence without enforcing the change.
func (l *Lexer) peekNextCharacter() *rune {
	if l.idx+1 < len(l.code) {
		return &l.code[l.idx+1]
	}

	return nil
}

// matchNextCharacter matches the current character with a required next character.
func (l *Lexer) matchNextCharacter(want rune) (*rune, bool) {
	next := l.peekNextCharacter()

	if next != nil && *next == want {
		return next, true
	}

	return nil, false
}

// Tokenise converts the plan-text into parseable tokens.
func (l *Lexer) Tokenise() ([]*Token, errors.Error) {
	if l.position == nil {
		// This ensures that the position of the token can be stored as intended, and also so that
		// the nextCharacter function can process the position instead of dereferencing a nilpointer.
		l.position = position.NewPosition(1,1)
	}

	if l.idx != -1 {
		// this is so it starts properly from the beginning of the code block.
		l.idx = -1
		l.position = position.NewPosition(1,1)
	}

	// start at index 0, with the initial character.
	l.nextCharacter()

	var tokens []*Token = make([]*Token, 0)

	for l.current != nil {
		switch *l.current {
		// Recognised whitespace characters that the lexer ignores.
		case '\r', '\t', ' ', '\n':
			l.nextCharacter()
			continue

		case '=':
			if next, ok := l.matchNextCharacter('='); ok {
				tokens = append(tokens, NewToken(TokEquals, string(*l.current)+string(*next), l.position))
				l.nextCharacter()
				l.nextCharacter()
				continue
			}

			tokens = append(tokens, NewToken(TokAsign, string(*l.current), l.position))
			l.nextCharacter()
		case '/':
			if next, ok := l.matchNextCharacter('/'); ok {
				tokens = append(tokens, NewToken(TokDiv, string(*l.current)+string(*next), l.position))
				l.nextCharacter()
				l.nextCharacter()
				continue
			}

			tokens = append(tokens, NewToken(TokDivide, string(*l.current), l.position))
			l.nextCharacter()
		case '>':
			if next, ok := l.matchNextCharacter('='); ok {
				tokens = append(tokens, NewToken(TokGreaterThnEq, string(*l.current)+string(*next), l.position))
				l.nextCharacter()
				l.nextCharacter()
				continue
			}

			tokens = append(tokens, NewToken(TokGreaterThn, string(*l.current), l.position))
			l.nextCharacter()
		case '<':
			if next, ok := l.matchNextCharacter('='); ok {
				tokens = append(tokens, NewToken(TokLessThnEq, string(*l.current)+string(*next), l.position))
				l.nextCharacter()
				l.nextCharacter()
				continue
			}

			tokens = append(tokens, NewToken(TokLessThn, string(*l.current), l.position))
			l.nextCharacter()
		case '!':
			if next, ok := l.matchNextCharacter('='); ok {
				tokens = append(tokens, NewToken(TokNotEq, string(*l.current)+string(*next), l.position))
				l.nextCharacter()
				l.nextCharacter()
				continue
			}

			tokens = append(tokens, NewToken(TokBang, string(*l.current), l.position))
			l.nextCharacter()
		case '&':
			if next, ok := l.matchNextCharacter('&'); ok {
				tokens = append(tokens, NewToken(TokAnd, string(*l.current)+string(*next), l.position))
				l.nextCharacter()
				l.nextCharacter()
				continue
			}

			return nil, &errors.SyntaxError{Message: fmt.Sprintf("unexpected token '%s'", string(*l.current)), Position: l.position.Clone()}
		case '|':
			if next, ok := l.matchNextCharacter('|'); ok {
				tokens = append(tokens, NewToken(TokOr, string(*l.current)+string(*next), l.position))
				l.nextCharacter()
				l.nextCharacter()
				continue
			}

			return nil, &errors.SyntaxError{Message: fmt.Sprintf("unexpected token '%s'", string(*l.current)), Position: l.position.Clone()}
		case '%':
			tokens = append(tokens, NewToken(TokMod, string(*l.current), l.position))
			l.nextCharacter()
		case '+':
			tokens = append(tokens, NewToken(TokPlus, string(*l.current), l.position))
			l.nextCharacter()
		case '-':
			tokens = append(tokens, NewToken(TokMinus, string(*l.current), l.position))
			l.nextCharacter()
		case '*':
			tokens = append(tokens, NewToken(TokMultiply, string(*l.current), l.position))
			l.nextCharacter()
		case '(':
			tokens = append(tokens, NewToken(TokLeftParen, string(*l.current), l.position))
			l.nextCharacter()
		case ')':
			tokens = append(tokens, NewToken(TokRightParen, string(*l.current), l.position))
			l.nextCharacter()
		case '[':
			tokens = append(tokens, NewToken(TokLeftSquare, string(*l.current), l.position))
			l.nextCharacter()
		case ']':
			tokens = append(tokens, NewToken(TokRightSquare, string(*l.current), l.position))
			l.nextCharacter()
		case '{':
			tokens = append(tokens, NewToken(TokLeftBrace, string(*l.current), l.position))
			l.nextCharacter()
		case '}':
			tokens = append(tokens, NewToken(TokRightBrace, string(*l.current), l.position))
			l.nextCharacter()
		case '.':
			tokens = append(tokens, NewToken(TokDot, string(*l.current), l.position))
			l.nextCharacter()
		case ',':
			tokens = append(tokens, NewToken(TokComma, string(*l.current), l.position))
			l.nextCharacter()
		case ':':
			tokens = append(tokens, NewToken(TokColon, string(*l.current), l.position))
			l.nextCharacter()
		default:
			if *l.current == '"' {
				tok, err := l.buildString()
				if err != nil {
					return nil, err
				}

				tokens = append(tokens, tok)
				l.nextCharacter()
				continue
			}

			if unicode.IsDigit(*l.current) {
				tok, err := l.buildNumber()
				if err != nil {
					return nil, err
				}

				tokens = append(tokens, tok)
				l.nextCharacter()
				continue
			}

			if *l.current == '_' || unicode.IsLetter(*l.current) {
				tokens = append(tokens, l.buildIdentifier())
				l.nextCharacter()
				continue
			}

			return nil, &errors.SyntaxError{Message: fmt.Sprintf("unexpected token '%s'", string(*l.current)), Position: l.position.Clone()}
		}
	}

	return tokens, nil
}
