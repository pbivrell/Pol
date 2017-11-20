## Brief Description
POL is an interpreted procedural programming language designed
for an Independent Study at Colorado State University.

## Usage
Clone repo.

From repo directory execute: `go build src/pol/main.go`

*This will build a `main` executable*

To run code execute: `./main filename`

*Try `./main examples/first.pol` for a glimpse at most everything pol has to offer*

## Notes
* Current running the executable above only produces parse trees.
* Error messages are a work in progress they are not guaranteed to be useful or correct.
* You can add optional numeric arguments to see debug output for parser and lexer. The following example will produce debug output for the parser but none for the lexer. 
`./main examples/first.pol 10 0`

## Dependencies
To use POL you need the go programming language. See docs here: [github.com/golang](https://github.com/golang).
