package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/go-js-yourself/gjsy/pkg/evaluator"
	"github.com/go-js-yourself/gjsy/pkg/lexer"
	"github.com/go-js-yourself/gjsy/pkg/parser"
)

const PROMPT = "js> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	io.WriteString(out, "Welcome to Go JS Yourself!\n")

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			io.WriteString(out, "Execution had the following errors:")
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
