package lexer

import (
	"slices"
	"unicode"

	"github.com/mattyco/marble/source/errors"
)

// TokenType represents an individual lexer token.
type TokenType int

// Keywords is the dictionary of all recognised key language words.
var Keywords []string = []string{
	"if", "elif", "else",
	"for", "while", "continue", "break",
	"int", "float", "string", "bool", "list",
	"var", "const", "function", "return",
}

const (
	TokInteger      TokenType = iota + 1 // 1,2,3...
	TokFloat                             // 1.0, 2.0, 3.5...
	TokString                            // "Hello, world!"
	TokBoolean                           // true/false
	TokIdentifier                        // exampleName
	TokKeyword                           // function
	TokPlus                              // +
	TokMinus                             // -
	TokDivide                            // /
	TokMultiply                          // *
	TokMod                               // %
	TokDiv                               // //
	TokGreaterThn                        // >
	TokGreaterThnEq                      // >=
	TokLessThn                           // <
	TokLessThnEq                         // <=
	TokAsign                             // =
	TokEquals                            // ==
	TokLeftParen                         // (
	TokRightParen                        // )
	TokLeftSquare                        // [
	TokRightSquare                       // ]
	TokLeftBrace                         // {
	TokRightBrace                        // }
	TokBang                              // !
	TokNotEq                             // !=
	TokDot                               // .
	TokComma                             // ,
	TokAnd                               // &&
	TokOr                                // ||
	TokColon                             // :
)

// buildString builds a string token.
func (l *Lexer) buildString() (*Token, errors.Error) {
	tok := NewToken(TokString, "", l.position)

	var value []rune

	l.nextCharacter()

	for l.current != nil && *l.current != '"' {
		value = append(value, *l.current)
		l.nextCharacter()
	}

	if l.current == nil || *l.current != '"' {
		return nil, &errors.SyntaxError{Message: "string declared but never closed", Position: tok.Position.Clone()}
	}

	tok.Value = string(value)
	return tok, nil
}

// buildNumber builds a number token and also defines if it is a float or an integer.
func (l *Lexer) buildNumber() (*Token, errors.Error) {
	tok := NewToken(TokInteger, "", l.position)

	var value []rune
	points := 0

	for l.current != nil && (unicode.IsDigit(*l.current) || *l.current == '.') {
		if *l.current == '.' {
			// Checks if a floating point has already been included.
			// 1.0000 would pass
			// 1.000.000 would fail etc
			if points != 0 {
				return nil, &errors.SyntaxError{Message: "invalid floating point defined", Position: tok.Position.Clone()}
			}

			points++
			tok.Type = TokFloat
		}

		value = append(value, *l.current)
		l.nextCharacter()
	}

	// This ensures that values of 123. or 1... etc cannot be tokenised.
	if l.current != nil && *l.current == '.' {
		return nil, &errors.SyntaxError{Message: "invalid floating point defined", Position: tok.Position.Clone()}
	}

	tok.Value = string(value)
	return tok, nil
}

// buildIdentifier builds an identifier token, then also compares it against the keywords dictionary.
func (l *Lexer) buildIdentifier() *Token {
	tok := NewToken(TokIdentifier, "", l.position)

	var value []rune

	for l.current != nil && (unicode.IsLetter(*l.current) || unicode.IsDigit(*l.current) || *l.current == '_') {
		value = append(value, *l.current)
		l.nextCharacter()
	}

	tok.Value = string(value)

	if slices.Contains(Keywords, tok.Value) {
		tok.Type = TokKeyword
	}

	if tok.Value == "true" || tok.Value == "false" {
		tok.Type = TokBoolean
	}

	return tok
}
