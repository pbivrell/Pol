## Brief Description
POL is an interpreted procedural programming language designed
for an Independent Study at Colorado State University.

## Reasons For Writing
I am building POL as part of an educational glimpse into the process
of writing a grammar, lexer, parser, and an interpreter. My goal is
to learn and take part in, the processes involved in language design 
and implementation.

## Design Process
* Grammar Design

    The grammar for this language is a LL grammar. This choice was made because
    LL grammars can be easily implemented as a recursive decent
    parser (see below). The BNF for the grammar can be found under
    [Docs/Grammar.txt](Docs/Grammar.txt)

* Lexer
	
    The Lexer preforms lexical analysis on .pol files and then tokenizes
    valid syntax for use by the parser. The lexer takes .pol files and
    produces .pol_lexed file which is a white space delimited file of tokens.
    The code for the lexer is written in C and can be found under [src/Lexer.c](src/Lexer.c)

* Parser

    The parser takes .pol_lexed files from the lexer and preforms a
    predictive parsing method known as recursive decent. In the
    processes of performing recursive decent the parser create an
    abstract syntax tree that will later be walked by the interpreter.
    When the parser encounters invalid syntax will be written out to
    user. The code for the parser is written in C and can be found under [src/Parser.c](src/Lexer.c)

* Interpreter

	TBW

## Resources
Dr. Wim Bohm - Colorado State University Professor who oversaw the project.

Essentials of Programming Languages 3rd Edition - Book used to aid interpreter development
