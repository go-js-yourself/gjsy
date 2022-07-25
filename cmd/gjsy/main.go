package main

import (
	"fmt"
	"os"

	"github.com/go-js-yourself/gjsy/pkg/interpreter"
	"github.com/go-js-yourself/gjsy/pkg/repl"
)

const USAGE = `Go JS yourself!
A JavaScript interpreter in Go.

Usage:
	gjsy <command>

Commands:
	repl		Starts a REPL.
	<file name>	Process a JavaScript file.`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(USAGE)
		return
	}

	if os.Args[1] == "repl" {
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	if _, err := os.Stat(os.Args[1]); err != nil {
		fmt.Println(USAGE)
		return
	}

	interpreter.Run(os.Args[1])
}
