# gjsy (Go JS Yourself!)

A JS interpreter written in GO. Now with multi-thread support!

## Pre-requisites

This project uses Go 1.18, but the `Makefile` commands will run on Docker to
help with the project set up, prerequisites, etc. So, the only two
pre-requisites are:

* make
* [Docker](https://docs.docker.com/get-docker/)

## Testing

If using the provided `Makefile`, run:

```
make test
# or with Go
go test ./...
```

## Building

After building, a single binary `gjsy` can be found in the `build/bin`
directory, if running with go, it is usually installed in `$GOPATH/bin`

```
make build
# or
go install ./...
```

## REPL

This project includes a REPL, you can run it with:

```
make repl
# or
go run cmd/gjsy repl
```

## File interpreter

There are several provided examples in the `examples` directory, to interpret
a single file, run

```
make run FILE=examples/hello_world.js
# or
go run cmd/gjsy examples/hello_world.js
```

## Writeup

### Parsing and Tokenization

TODO

### Environment

The environment represents the evaluation context for any given expression.
It is defined in [engironment.go](./pkg/object/environment.go) and allows
the persistance of state during the lifetime of the program.

### Evaluation

The [evaluator](./pkg/evaluator/evaluator.go) is responsible for the execution
of the program. It does this by recursively evaluating the program, its
statements, and expressions. In the above file, there is a single switch
statement which selects on the token type to evaluate. Further evaluation
is conducted in subordonate files. For example the
[identifier.go](./pkg/evaluator/identifier.go) file evaluates an identifier
given the environment.

### Parallelism and Go

Perhaps the only unique facet of this project is the application of
go routines to evaluate function expressions in a
[go routine](https://golangbot.com/goroutines/). This can be found
in [function.go](./pkg/evaluator/function.go) in the
`applyGoFunction` function. This function spawns a new go routine
to run the `Eval` function given the current environment. The
routines generated by using the `go` application syntax can then
be waited on by using the built in `wg` object
[builtins.go](./pkg/evaluator/builtins.go). There is an example
of this in [multi_thread.js](./examples/multi_thread.js) which
can be run by using the following command:

```shell
FILE=examples/multi_thread.js make run
```

## Credits

Miles Possing, Ivan Valdes Castillo

## References

This effort was heavily influenced by the book
Thorsten Ball's book "Writing an Interpreter in Go"
