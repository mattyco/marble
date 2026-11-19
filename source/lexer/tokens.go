package lexer

import "fmt"

// TokenType represents an individual lexer token.
type TokenType int

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

// Token contains the type, string value, and the position in the code block.
type Token struct {
	Type     TokenType
	value    string
	Position *Position
}

// NewToken creates a new token object.
func NewToken(ttype TokenType, value string, position *Position) *Token {
	return &Token{
		Type:     ttype,
		value:    value,
		Position: position.Clone(),
	}
}

// String returns the string representation of the token.
func (t *Token) String() string {
	return fmt.Sprintf("type: %d position: (%s)", t.Type, t.Position.String())
}
