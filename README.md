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

## Credits

Miles Possing, Ivan Valdes Castillo
