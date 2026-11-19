package lexer

import (
	"fmt"
	"slices"
	"unicode"
)

// TokenType represents an individual lexer token.
type TokenType int

// Keywords is the dictionary of all recognised key language words.
var Keywords []string = []string{
	"if", "elif", "else",
	"for", "while", "continue", "break",
	"int", "float", "string", "bool", "list",
	"var", "const", "function", "class", "return",
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
	TokNot                               // !
)

// buildString builds a string token.
func (l *Lexer) buildString() (*Token, error) {
	tok := NewToken(TokString, "", l.position)

	var value []rune

	l.nextCharacter()

	for l.current != nil && *l.current != '"' {
		value = append(value, *l.current)
		l.nextCharacter()
	}

	if *l.current != '"' {
		return nil, fmt.Errorf("string declared at %s but never closed", tok.Position.String())
	}

	tok.value = string(value)
	return tok, nil
}

// buildNumber builds a number token and also defines if it is a float or an integer.
func (l *Lexer) buildNumber() (*Token, error) {
	tok := NewToken(TokInteger, "", l.position)

	var value []rune
	points := 0

	for l.current != nil && unicode.IsDigit(*l.current) || *l.current == '.' {
		if *l.current == '.' {
			// Checks if a floating point has already been included.
			// 1.0000 would pass
			// 1.000.000 would fail etc
			if points != 0 {
				return nil, fmt.Errorf("invalid floating point defined at %s", tok.Position.String())
			}

			points++
			tok.Type = TokFloat
		}

		value = append(value, *l.current)
		l.nextCharacter()
	}

	tok.value = string(value)
	return tok, nil
}

// buildIdentifier builds an identifier token, then also compares it against the keywords dictionary.
func (l *Lexer) buildIdentifier() *Token {
	tok := NewToken(TokIdentifier, "", l.position)

	var value []rune

	for l.current != nil && unicode.IsLetter(*l.current) || unicode.IsDigit(*l.current) || *l.current == '_' {
		value = append(value, *l.current)
		l.nextCharacter()
	}

	tok.value = string(value)

	if slices.Contains(Keywords, tok.value) {
		tok.Type = TokKeyword
	}

	if tok.value == "true" || tok.value == "false" {
		tok.Type = TokBoolean
	}

	return tok
}
