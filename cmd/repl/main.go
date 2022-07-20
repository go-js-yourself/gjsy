package main

import (
	"os"

	"github.com/go-js-yourself/gjsy/pkg/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
