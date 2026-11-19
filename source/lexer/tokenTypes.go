package lexer

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
	TokNot                               // !
)
