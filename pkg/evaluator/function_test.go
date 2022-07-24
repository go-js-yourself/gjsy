package evaluator

import (
	"testing"

	"github.com/go-js-yourself/gjsy/pkg/object"
)

func TestFunctionObject(t *testing.T) {
	input := "function(x) { x + 2 };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "{\n(x + 2);\n}\n"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = function(x) {x;}; identity(5);", 5},
		{"let identity = function(x) { return x; }; identity(5);", 5},
		{"let double = function(x) { x * 2; }; double(5);", 10},
		{"let add = function(x, y) { x + y; }; add(2, 3);", 5},
		{"let add = function(x, y) { x + y; }; add(add(2,3), 5);", 10},
		{"function(x) { x; }(5)", 5},
		{"function foo(x) { x; }; foo(5)", 5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestNestedClosures(t *testing.T) {
	input := `
	let newAdder = function(x) {
		function(y) { return x + y };
	};
	
	let addTwo = newAdder(2);
	addTwo(2);`

	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 4)
}
