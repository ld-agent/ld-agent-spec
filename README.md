# ld-agent Plugin Specification

**Version 1.0** | [Changelog](CHANGELOG.md) | [Examples](examples/) | [Schemas](schemas/)

## Overview

This repository contains the official plugin specification for **ld-agent** - a dynamic linker for agentic AI systems. Just as `ld-so` discovers, loads, and links shared objects at runtime, ld-agent discovers, loads, and links AI capabilities at runtime to create composable AI systems.

## What's in this Repository

- 📋 **[Plugin Specification](PLUGIN_SPECIFICATION.md)** - The formal specification that all ld-agent plugins must follow
- 🏗️ **[Architecture Guide](ARCHITECTURE.md)** - Core design principles and implementation patterns
- 📊 **[JSON Schemas](schemas/)** - Machine-readable validation schemas for each language
- 💡 **[Examples](examples/)** - Reference implementations showing best practices
- 📚 **[Documentation](docs/)** - Migration guides, best practices, and FAQ

## Specification Versioning

| Version | Status | Languages | Notable Changes |
|---------|---------|-----------|-----------------|
| 1.0 | Current | Python, Go, TypeScript | Initial specification |
| 0.9 | Legacy | Python | Beta version |

## Language Implementations

The specification is implemented across multiple languages:

- 🐍 **Python**: [`ld-agent-python`](https://github.com/your-org/ld-agent-python)
- 🚀 **Go**: [`ld-agent-go`](https://github.com/your-org/ld-agent-go)  
- 📦 **TypeScript**: [`ld-agent-typescript`](https://github.com/your-org/ld-agent-typescript)

## Quick Start

### Writing Your First Plugin

Choose your language to get started:

#### Python
```python
_module_info = {
    "name": "My Plugin",
    "description": "What my plugin does",
    "version": "1.0.0"
}

def my_function(input: str) -> str:
    """My function description."""
    return f"Processed: {input}"

_module_exports = {"tools": [my_function]}
```

#### Go
```go
var ModuleInfo = map[string]interface{}{
    "name": "My Plugin",
    "description": "What my plugin does", 
    "version": "1.0.0",
}

func MyFunction(input string) string {
    return "Processed: " + input
}

var ModuleExports = map[string]interface{}{
    "tools": []interface{}{MyFunction},
}
```

#### TypeScript
```typescript
export const moduleInfo = {
    name: "My Plugin",
    description: "What my plugin does",
    version: "1.0.0"
};

export function myFunction(input: string): string {
    return `Processed: ${input}`;
}

export const moduleExports = {tools: [myFunction]};
```

## Validation

Use our validation tools to ensure your plugins conform to the specification:

```bash
# Validate Python plugin
python -m ldagent.validate my_plugin.py

# Validate Go plugin  
go run github.com/your-org/ld-agent-go/cmd/validate my_plugin.so

# Validate TypeScript plugin
npx ld-agent-typescript validate my_plugin.ts
```

## Contributing to the Specification

1. 📖 Read the [Contributing Guide](CONTRIBUTING.md)
2. 💬 Discuss changes in [GitHub Discussions](https://github.com/your-org/ld-agent-spec/discussions)
3. 📝 Submit an RFC for major changes
4. 🔧 Open a PR for minor improvements

## Specification Compliance

To be considered compliant, plugin implementations must:

- ✅ Export required metadata (`_module_info`/`ModuleInfo`/`moduleInfo`)
- ✅ Export capability definitions (`_module_exports`/`ModuleExports`/`moduleExports`)
- ✅ Follow language-specific naming conventions
- ✅ Pass validation with official schemas
- ✅ Include comprehensive documentation

## Community

- 🐛 [Report Issues](https://github.com/your-org/ld-agent-spec/issues)
- 💬 [Discussions](https://github.com/your-org/ld-agent-spec/discussions)  
- 🔧 [Plugin Registry](https://github.com/your-org/ld-agent-plugins)
- 🌟 [Awesome Plugins](https://github.com/your-org/awesome-ld-agent)

## License

MIT License - see [LICENSE](LICENSE) file for details. 