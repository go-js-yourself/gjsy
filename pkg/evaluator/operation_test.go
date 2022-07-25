package evaluator

import "testing"

func TestEvalIngtegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5+5+5+5", 20},
		{"-50 * 2", -100},
		{"20 / 10", 2},
		{"71 - 2", 69}, // nice
		{"20 * 2 / 10", 4},
		{"20 + 2 * 10", 40},
		// {"(20 + 2) * 10", 220},
		// {"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalIntegerComparison(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"10 > 1", true},
		{"10 > 20", false},
		{"10 < 1", false},
		{"10 < 20", true},
		{"10 <= 10", true},
		{"10 <= 11", true},
		{"10 <= 9", false},
		{"10 >= 10", true},
		{"10 >= 11", false},
		{"10 >= 9", true},
		{"10 == 10", true},
		{"10 == 11", false},
		{"10 != 11", true},
		{"10 != 10", false},
		{"true && true", true},
		{"false && true", false},
		{"true || false", true},
		{"false || false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}
