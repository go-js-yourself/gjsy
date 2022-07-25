package interpreter

import (
	"os"

	"github.com/go-js-yourself/gjsy/pkg/evaluator"
	"github.com/go-js-yourself/gjsy/pkg/lexer"
	"github.com/go-js-yourself/gjsy/pkg/object"
	"github.com/go-js-yourself/gjsy/pkg/parser"
)

func Run(file string) {
	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(input))
	p := parser.New(l)
	program := p.ParseProgram()

	evaluator.Eval(program, object.NewEnvironment())
}
