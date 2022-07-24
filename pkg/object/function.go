package object

import (
	"fmt"
	"strings"

	"github.com/go-js-yourself/gjsy/pkg/ast"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.ClosureStatement
	Env        *Environment
	Name       *ast.Identifier
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f *Function) Inspect() string {
	out := "function"

	params := make([]string, len(f.Parameters))
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	if f.Name != nil {
		out += " " + f.Name.String()
	}

	return fmt.Sprintf("%s(%s)%s",
		out,
		strings.Join(params, ", "),
		f.Body.String(),
	)
}
