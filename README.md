## Brief Description
POL is an interpreted procedural programming language designed
for an Independent Study at Colorado State University.

## Reasons For Writing
I am building POL as part of an educational glimpse into the process
of writing a grammer, lexer, parser, and an interpreter. My goal is 
to learn and take part in, the processes involved in language design 
and implementation.

## Design Process
* Grammer Design 

    The grammer for this langauge is a LL grammer. This choice was made because
    LL grammers can be easily implemented as a recursive decent
    parser (see below). The BNF for the grammer can be found under
    [Docs/Grammer.txt](Docs/Grammer.txt)

* Lexer
	
    The Lexer preforms lexical anylisis on .pol files and then tokenizes
    valid syntax for use by the parser. The lexer takes .pol files and
    produces .pol_lexed file which is a white space delimited file of tokens.
    The code for the lexer is written in C and can be found under [src/Lexer.c](src/Lexer.c)

* Parser

    The parser takes .pol_lexed files from the lexer and preforms a
    predicitive parsing method known as recursive decent. In the
    processes of proforming recursive decent the parser create an
    abstract syntax tree that will later be walked by the interpreter.
    When the parser encounters invalid syntax will be written out to
    user. The code for the parser is written in C and can be found under [src/Parser.c](src/Lexer.c)

* Interpreter

	TBW

## Resources
Dr. Wim Bohm - Colorado State University Professor who oversaw the project.

Essentials of Programming Languages 3rd Edition - Book used to aid interpreter development

