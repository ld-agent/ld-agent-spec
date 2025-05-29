// =============================================================================
// MINIMAL LD-AGENT PLUGIN EXAMPLE (Go)
// =============================================================================

package main

var ModuleInfo = map[string]interface{}{
	"name":        "Minimal Example",
	"description": "The simplest possible ld-agent plugin",
	"version":     "1.0.0",
}

func HelloWorld() string {
	return "Hello from ld-agent!"
}

var ModuleExports = map[string]interface{}{
	"tools": []interface{}{
		map[string]interface{}{
			"name":        "hello_world",
			"function":    HelloWorld,
			"description": "Returns a simple greeting",
		},
	},
} 