# ld-agent Plugin Specification

**Version 1.0**

This specification defines the standard interface that all ld-agent plugins must implement across all supported languages (Python, Go, TypeScript/JavaScript).

## Plugin Structure

All ld-agent plugins must export two standardized symbols:

1. **Module Metadata** (`_module_info`) - Describes the plugin
2. **Module Exports** (`_module_exports`) - Defines available capabilities

### Module Metadata (`_module_info`)

```json
{
  "name": "Plugin Display Name",
  "description": "What this plugin does",
  "author": "Plugin Author",
  "version": "1.0.0",
  "platform": "any|darwin|linux|windows",
  "python_requires": ">=3.10",
  "dependencies": ["package>=version"],
  "environment_variables": {
    "VAR_NAME": {
      "description": "Variable description",
      "required": true|false,
      "default": "default_value"
    }
  }
}
```

### Module Exports (`_module_exports`)

```json
{
  "tools": [function1, function2, ...]
}
```

## Language-Specific Implementations

### Python

```python
# =============================================================================
# START OF MODULE METADATA
# =============================================================================
_module_info = {
    "name": "Example Plugin",
    "description": "An example plugin",
    "author": "Your Name",
    "version": "1.0.0",
    "platform": "any",
    "python_requires": ">=3.10",
    "dependencies": ["pydantic>=2.0.0"],
    "environment_variables": {
        "EXAMPLE_API_KEY": {
            "description": "API key for example service",
            "required": True
        }
    }
}
# =============================================================================
# END OF MODULE METADATA
# =============================================================================

from typing import Annotated
from pydantic import Field

def example_function(
    message: Annotated[str, Field(description="Input message")]
) -> str:
    """Example function that processes a message."""
    return f"Processed: {message}"

# =============================================================================
# START OF EXPORTS
# =============================================================================
_module_exports = {
    "tools": [example_function]
}
# =============================================================================
# END OF EXPORTS
# =============================================================================
```

### Go

```go
package main

import (
    "fmt"
    "reflect"
)

// ModuleInfo contains plugin metadata
var ModuleInfo = map[string]interface{}{
    "name":        "Example Plugin",
    "description": "An example plugin",
    "author":      "Your Name",
    "version":     "1.0.0",
    "platform":    "any",
}

// ExampleFunction processes a message
func ExampleFunction(message string) string {
    return fmt.Sprintf("Processed: %s", message)
}

// ModuleExports defines available tools
var ModuleExports = map[string]interface{}{
    "tools": []interface{}{
        map[string]interface{}{
            "name":        "example_function",
            "function":    ExampleFunction,
            "description": "Example function that processes a message",
            "parameters": map[string]interface{}{
                "message": map[string]interface{}{
                    "type":        "string",
                    "description": "Input message",
                },
            },
        },
    },
}
```

### TypeScript/JavaScript

```typescript
// Module metadata
export const moduleInfo = {
    name: "Example Plugin",
    description: "An example plugin",
    author: "Your Name",
    version: "1.0.0",
    platform: "any",
    dependencies: [],
    environmentVariables: {
        EXAMPLE_API_KEY: {
            description: "API key for example service",
            required: true
        }
    }
};

// Example function
export function exampleFunction(message: string): string {
    return `Processed: ${message}`;
}

// Module exports
export const moduleExports = {
    tools: [exampleFunction]
};
```

## Discovery and Loading

### Python
- Scans `plugins/` directory for `.py` files and packages
- Uses `importlib` to dynamically load modules
- Validates structure using Pydantic

### Go
- Scans `plugins/` directory for `.so` shared libraries
- Uses Go's `plugin` package to load symbols
- Validates structure using reflection

### TypeScript/JavaScript
- Scans `plugins/` directory for `.js`/`.ts` files
- Uses `require()` or `import()` to load modules
- Validates structure using Zod schemas

## Function Signatures

All plugin functions should:
1. Have clear, descriptive names
2. Include comprehensive docstrings/comments
3. Use typed parameters with descriptions
4. Return typed results
5. Handle errors gracefully

### Parameter Annotations

- **Python**: Use `typing.Annotated` with Pydantic `Field`
- **Go**: Use struct tags for parameter metadata
- **TypeScript**: Use JSDoc comments or schema definitions

## Environment Variables

Plugins can declare environment variables they need:

```json
{
  "environment_variables": {
    "PLUGIN_API_KEY": {
      "description": "API key for the service",
      "required": true
    },
    "PLUGIN_TIMEOUT": {
      "description": "Request timeout in seconds",
      "required": false,
      "default": "30"
    }
  }
}
```

ld-agent will:
- Generate `.env.template` files with all variables
- Validate required variables are set
- Check for naming conflicts between plugins

## Dependencies

Plugins can declare their dependencies:

```json
{
  "dependencies": [
    "requests>=2.28.0",
    "pydantic>=2.0.0"
  ]
}
```

ld-agent will:
- Check if dependencies are installed
- Generate requirements files
- Validate version compatibility

## Platform Compatibility

Plugins can specify platform requirements:

- `"any"` - Works on all platforms
- `"darwin"` - macOS only
- `"linux"` - Linux only  
- `"windows"` - Windows only

ld-agent will only load compatible plugins for the current platform.

## Error Handling

Plugins should handle errors gracefully and return meaningful error messages. ld-agent will catch and log plugin errors without crashing the entire system.

## Best Practices

1. **Single Responsibility**: Each plugin should have a focused purpose
2. **Clear Naming**: Use descriptive names for plugins and functions
3. **Documentation**: Include comprehensive docstrings and metadata
4. **Error Handling**: Handle edge cases and provide helpful error messages
5. **Testing**: Include tests for plugin functionality
6. **Dependencies**: Minimize external dependencies when possible
7. **Environment**: Use environment variables for configuration
8. **Versioning**: Follow semantic versioning for plugin releases

## Examples

See the `examples/` directory in each language implementation for complete working examples of plugins following this specification. 
