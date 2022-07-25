package evaluator

import (
	"fmt"
	"strings"

	"github.com/go-js-yourself/gjsy/pkg/object"
)

var builtins = map[string]*object.BuiltinObj{
	"console": &object.BuiltinObj{
		Funcs: map[string]*object.BuiltinFunc{
			"log": &object.BuiltinFunc{
				Func: func(args ...object.Object) object.Object {
					out := make([]string, len(args))
					for i, a := range args {
						out[i] = a.Inspect()
					}
					fmt.Println(strings.Join(out, " "))
					return &object.Undefined{}
				},
			},
		},
	},
}
