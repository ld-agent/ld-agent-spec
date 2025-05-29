// =============================================================================
// CALCULATOR PLUGIN EXAMPLE (Go)
// =============================================================================

package main

import "errors"

var ModuleInfo = map[string]interface{}{
	"name":        "Simple Calculator",
	"description": "Basic arithmetic operations for demonstration",
	"author":      "ld-agent Team",
	"version":     "1.0.0",
	"platform":    "any",
	"go_version":  ">=1.21",
	"dependencies": []string{},
	"environment_variables": map[string]interface{}{},
}

func AddNumbers(a, b float64) float64 {
	return a + b
}

func SubtractNumbers(a, b float64) float64 {
	return a - b
}

func MultiplyNumbers(a, b float64) float64 {
	return a * b
}

func DivideNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

var ModuleExports = map[string]interface{}{
	"tools": []interface{}{
		map[string]interface{}{
			"name":        "add_numbers",
			"function":    AddNumbers,
			"description": "Add two numbers together",
		},
		map[string]interface{}{
			"name":        "subtract_numbers",
			"function":    SubtractNumbers,
			"description": "Subtract the second number from the first",
		},
		map[string]interface{}{
			"name":        "multiply_numbers",
			"function":    MultiplyNumbers,
			"description": "Multiply two numbers together",
		},
		map[string]interface{}{
			"name":        "divide_numbers",
			"function":    DivideNumbers,
			"description": "Divide the first number by the second",
		},
	},
} 