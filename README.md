## Brief Description
POL is an interpreted procedural programming language designed
for an Independent Study at Colorado State University.

## Reasons For Writing
I am building POL as part of an educational glimpse into the process
of writing a grammar, a lexer, a parser, and an interpreter. My goal is
to learn about the processes involved in language design
and implementation.

## Design Process
* Grammar Design

    The grammar for this language is a LL grammar. This choice was made because
    LL grammars can be easily implemented as a recursive descent
    parser (see below). The BNF for the grammar can be found in
    [Grammar.txt](doc/Grammar.txt)

* Lexer
	
    The lexer performs lexical analysis on .pol files and then tokenizes
    valid syntax for use by the parser. The lexer takes .pol files and
    produces .pol_lexed files, which are a white space delimited file of tokens.
    The code for the lexer is written in Go and can be found in [lexer.go](src/lexer/lexer.go)

* Parser

    The parser takes .pol_lexed files from the lexer and performs a
    predictive parsing method known as recursive descent. In the
    process of performing recursive descent, the parser creates an
    abstract syntax tree that will later be walked by the interpreter.
    When the parser encounters invalid syntax an error will be written out to the user.

* Interpreter

	TBW

## Resources
Dr. Wim Bohm - Colorado State University Professor who oversaw the project.

Essentials of Programming Languages 3rd Edition - Book used to aid interpreter development
