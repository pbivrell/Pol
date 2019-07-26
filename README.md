## Brief Description
POL is an interpreted procedural programming language designed
for an Independent Study at Colorado State University.

You can learn more about it a [program-pol.com](program-pol.com)

## Usage
Clone repo.

From repo directory execute: 

`go install Pol/main/pol.go`

*This will build a `pol` executable into the src directory*

To run code execute: 

`./pol filename`
                     
`./pol -src "source code"` *Note quotes in src need to be escaped*

POL allows you to print AST with the `-tree` option

Check out the examples directory to see examples of pol code

## Notes
* Error messages are a work in progress they are not guaranteed to be useful or correct.
* There are known bugs currently I do not have the time to fix them. I eventually plan to
  come back to this project.

## Dependencies
To use POL you need the go programming language. See docs here: [github.com/golang](https://github.com/golang).


### Hello world Lizzie
