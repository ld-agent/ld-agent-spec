# ld-agent Architecture

## Overview

**ld-agent** is a dynamic linker for agentic systems, inspired by traditional dynamic linkers like `ld-so`. Just as `ld-so` discovers, loads, and links shared libraries at runtime to create executable programs, ld-agent discovers, loads, and links agentic capabilities at runtime to create composable AI systems.

## Core Concepts

### Dynamic Linking for AI

Traditional dynamic linking involves:
1. **Discovery** - Finding shared libraries in standard locations
2. **Loading** - Loading libraries into memory
3. **Symbol Resolution** - Finding exported functions and variables
4. **Relocation** - Adjusting addresses for runtime execution
5. **Execution** - Calling linked functions

ld-agent applies the same concepts to AI capabilities:
1. **Discovery** - Finding plugin modules in standard locations
2. **Loading** - Loading plugins into the runtime environment
3. **Symbol Resolution** - Finding exported AI tools and functions
4. **Registration** - Building a registry of available capabilities
5. **Execution** - Calling AI tools with proper argument mapping

### Plugin as Shared Object

In traditional systems, shared objects (`.so`, `.dll`, `.dylib`) export symbols that can be linked at runtime. In ld-agent, plugins export standardized symbols:

- **Module Metadata** (`_module_info`) - Like ELF headers, describes the module
- **Module Exports** (`_module_exports`) - Like symbol tables, lists available functions

## Architecture Components

### 1. Plugin Loader

The core component responsible for discovery and loading:

```
PluginLoader
├── Discovery Engine
│   ├── File System Scanner
│   ├── Platform Filter
│   └── Dependency Resolver
├── Loading Engine
│   ├── Language-Specific Loaders
│   ├── Validation Engine
│   └── Error Handler
└── Registry
    ├── Symbol Table
    ├── Metadata Store
    └── Runtime Cache
```

### 2. Discovery Engine

**File System Scanner**
- Scans standard plugin directories (`plugins/`)
- Identifies plugin files by extension (`.py`, `.so`, `.js`)
- Handles both single-file and package plugins

**Platform Filter**
- Checks plugin platform compatibility
- Filters plugins based on current OS
- Validates runtime requirements

**Dependency Resolver**
- Analyzes plugin dependencies
- Checks for conflicts
- Validates version compatibility

### 3. Loading Engine

**Language-Specific Loaders**
- Python: Uses `importlib` for dynamic imports
- Go: Uses `plugin` package for shared library loading
- TypeScript/JavaScript: Uses `require()` or `import()`

**Validation Engine**
- Validates plugin structure against specification
- Checks required symbols (`_module_info`, `_module_exports`)
- Validates function signatures and metadata

**Error Handler**
- Graceful error handling for plugin failures
- Detailed error reporting and logging
- Isolation to prevent system crashes

### 4. Registry

**Symbol Table**
- Maps tool names to callable functions
- Maintains function metadata and signatures
- Enables fast lookup and execution

**Metadata Store**
- Stores plugin information and capabilities
- Tracks dependencies and requirements
- Manages environment variable definitions

**Runtime Cache**
- Caches loaded plugins for performance
- Manages plugin lifecycle
- Handles hot-reloading (future feature)

## Language Implementations

### Python Implementation

```
ld-agent-python/
├── ldagent/
│   ├── loader.py          # Core plugin loader
│   ├── env_utils.py       # Environment management
│   ├── depcheck.py        # Dependency checking
│   └── cli.py             # Command-line interface
├── plugins/               # Plugin directory
├── examples/              # Usage examples
└── setup.py              # Package configuration
```

**Key Features:**
- Uses `importlib` for dynamic module loading
- Pydantic for validation and type safety
- Automatic environment variable management
- CLI tools for plugin management

### Go Implementation

```
ld-agent-go/
├── types.go               # Core type definitions
├── loader.go              # Plugin loader implementation
├── errors.go              # Error definitions
├── loader_test.go         # Unit tests
├── plugins/               # Plugin directory
├── example/               # Usage examples
└── Makefile              # Build configuration
```

**Key Features:**
- Uses Go's `plugin` package for shared library loading
- Reflection for runtime type checking
- Strong typing with compile-time safety
- High-performance execution

### TypeScript Implementation

```
ld-agent-ts/
├── src/
│   ├── types.ts           # Type definitions
│   ├── loader.ts          # Plugin loader
│   ├── index.ts           # Main exports
│   ├── plugins/           # Plugin directory
│   └── example/           # Usage examples
├── package.json           # npm configuration
└── tsconfig.json          # TypeScript configuration
```

**Key Features:**
- Uses Node.js module system for loading
- Zod for runtime validation
- TypeScript for compile-time type safety
- Async/await support for modern JavaScript

## Plugin Lifecycle

### 1. Discovery Phase

```
Scan plugins/ directory
├── Find plugin files (.py, .so, .js)
├── Check platform compatibility
├── Validate file structure
└── Build discovery manifest
```

### 2. Loading Phase

```
For each discovered plugin:
├── Load module using language-specific loader
├── Extract _module_info metadata
├── Validate plugin structure
├── Check dependencies
└── Register in symbol table
```

### 3. Registration Phase

```
For each loaded plugin:
├── Extract _module_exports
├── Validate function signatures
├── Register tools in symbol table
├── Store metadata in registry
└── Setup environment variables
```

### 4. Execution Phase

```
When tool is called:
├── Lookup function in symbol table
├── Validate arguments
├── Call function with proper mapping
├── Handle errors gracefully
└── Return results
```

## Environment Management

ld-agent provides comprehensive environment variable management:

### Variable Declaration
Plugins declare their environment variables in metadata:
```json
{
  "environment_variables": {
    "API_KEY": {
      "description": "Service API key",
      "required": true
    },
    "TIMEOUT": {
      "description": "Request timeout",
      "required": false,
      "default": "30"
    }
  }
}
```

### Template Generation
ld-agent generates `.env.template` files with:
- All declared variables from all plugins
- Descriptions and default values
- Required/optional indicators
- Conflict detection

### Validation
Runtime validation ensures:
- Required variables are set
- No naming conflicts between plugins
- Proper value formats

## Dependency Management

### Declaration
Plugins declare dependencies in metadata:
```json
{
  "dependencies": [
    "requests>=2.28.0",
    "pydantic>=2.0.0"
  ]
}
```

### Resolution
ld-agent checks:
- Package availability
- Version compatibility
- Conflict resolution
- Installation status

### Requirements Generation
Automatic generation of:
- `requirements.txt` (Python)
- `go.mod` entries (Go)
- `package.json` dependencies (TypeScript)

## Error Handling and Isolation

### Plugin Isolation
- Plugin failures don't crash the system
- Graceful degradation when plugins fail
- Detailed error reporting and logging

### Error Categories
1. **Discovery Errors** - File not found, permission issues
2. **Loading Errors** - Import failures, syntax errors
3. **Validation Errors** - Missing symbols, invalid structure
4. **Runtime Errors** - Function call failures, exceptions

### Recovery Strategies
- Skip failed plugins and continue loading others
- Provide fallback mechanisms where possible
- Clear error messages for debugging

## Performance Considerations

### Lazy Loading
- Plugins loaded on-demand when first accessed
- Reduces startup time and memory usage
- Enables selective loading based on requirements

### Caching
- Loaded plugins cached in memory
- Metadata cached for fast lookups
- Symbol table optimized for frequent access

### Parallel Loading
- Multiple plugins loaded concurrently
- Dependency-aware loading order
- Efficient resource utilization

## Security Considerations

### Plugin Validation
- Strict validation of plugin structure
- Signature verification (future feature)
- Sandboxing capabilities (future feature)

### Environment Isolation
- Environment variables scoped to plugins
- No cross-plugin variable access
- Secure credential management

### Code Execution
- Controlled execution environment
- Error boundary protection
- Resource usage monitoring (future feature)

## Future Enhancements

### Hot Reloading
- Dynamic plugin reloading without restart
- Development-time productivity improvements
- Configuration change detection

### Plugin Marketplace
- Centralized plugin repository
- Version management and updates
- Community-contributed plugins

### Advanced Dependency Management
- Automatic dependency installation
- Version conflict resolution
- Dependency graph visualization

### Monitoring and Observability
- Plugin performance metrics
- Usage analytics
- Health monitoring

## Comparison with Traditional Linkers

| Feature | Traditional Linker | ld-agent |
|---------|-------------------|----------|
| **Discovery** | `/lib`, `/usr/lib` | `plugins/` directory |
| **File Format** | ELF, PE, Mach-O | Language modules |
| **Symbol Table** | Function addresses | Function references |
| **Loading** | Memory mapping | Language imports |
| **Linking** | Address resolution | Registry lookup |
| **Execution** | Direct calls | Wrapper functions |
| **Dependencies** | Library dependencies | Package dependencies |
| **Versioning** | SONAME versioning | Semantic versioning |

This architecture enables ld-agent to provide the same composability and modularity for AI systems that dynamic linking provides for traditional software systems. 
