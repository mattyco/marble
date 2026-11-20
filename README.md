# Marble

An interpreted, static typed programming language written in Go.

[![License](https://img.shields.io/github/license/mattyco/marble)](LICENSE)

# Lexer

The lexer takes the raw code block, or plain text - and converts it into parseable tokens. In the example provided below, we can see the output that the lexer gives for a function declaration.

The lexer knows basic rules, such as how strings, integers/floats, and identifiers are formed. It can also detect boolean values or keyword values. 

```
function cool(cool string) {
    var helloworld string = "Hello, world!"
}
```

```
type: 6 value: "function" position: (line 1, column 1)
type: 5 value: "cool" position: (line 1, column 10)
type: 5 value: "cool" position: (line 1, column 15)
type: 6 value: "string" position: (line 1, column 20)
type: 23 value: "{" position: (line 1, column 28)
type: 6 value: "var" position: (line 2, column 6)
type: 5 value: "helloworld" position: (line 2, column 10)
type: 6 value: "string" position: (line 2, column 21)
type: 17 value: "=" position: (line 2, column 28)
type: 3 value: "Hello, world!" position: (line 2, column 30)
type: 24 value: "}" position: (line 3, column 2)
```

Once tokenisation is complete, the resulting tokens are passed to the parser, which analyses their structure and converts them into an AST (Abstract Syntax Tree).

# Parser

The parser takes the tokens produced by the lexer and converts them into an AST (Abstract Syntax Tree).

Unlike the lexer, the parser understands the syntax and grammatical rules of Marble. It determines how tokens relate to one another and whether they form valid expressions, statements and declarations.

Following from the example function declaration above, we can see how the AST develops.

```
Function
  Name: cool
  Argument: cool (string)
  Body: Variable:
        identifier(helloworld) = string(Hello, world!)
```

If the tokens do not follow Marble's grammatical rules, the parser produces a syntax error.

Once parsing is complete, the resulting AST is passed into the program evaluator.