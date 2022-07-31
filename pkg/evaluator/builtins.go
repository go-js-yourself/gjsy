package evaluator

import (
	"fmt"
	"strings"
	"sync"

	"github.com/go-js-yourself/gjsy/pkg/object"
)

var wg sync.WaitGroup

var builtins = map[string]*object.BuiltinObj{
	"console": {
		Funcs: map[string]*object.BuiltinFunc{
			"log": {
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
	"wg": {
		Funcs: map[string]*object.BuiltinFunc{
			"add": {
				Func: func(args ...object.Object) object.Object {
					if len(args) != 1 {
						return newError("wrong number of arguments. got=%d, want=1",
							len(args))
					}
					switch arg := args[0].(type) {
					case *object.Integer:
						wg.Add(int(arg.Value))
						return &object.Undefined{}
					default:
						return newError("argument to `add` not supported, got %s",
							args[0].Type())
					}
				},
			},
			"done": {
				Func: func(args ...object.Object) object.Object {
					if len(args) != 0 {
						return newError("wrong number of arguments. got=%d, want=0",
							len(args))
					}
					wg.Done()
					return &object.Undefined{}
				},
			},
			"wait": {
				Func: func(args ...object.Object) object.Object {
					if len(args) != 0 {
						return newError("wrong number of arguments. got=%d, want=0",
							len(args))
					}
					wg.Wait()
					return &object.Undefined{}
				},
			},
		},
	},
}
