package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattyco/marble/source/lexer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Marble ~ https://github.com/mattyco/marble")
		fmt.Printf("%s <file>\r\n", filepath.Base(os.Args[0]))
		return
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read file '%s', error: %v\r\n", os.Args[1], err)
		return
	}

	tokens, error := lexer.NewLexer([]rune(string(bytes))).Tokenise()
	if error != nil {
		fmt.Println(error.String())
		return
	}

	for _, tok := range tokens {
		fmt.Println(tok.String())
	}
}
