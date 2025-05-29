# ld-agent Plugin Examples

This directory contains reference implementations that demonstrate how to write ld-agent plugins across different languages and complexity levels.

## Examples by Complexity

### 1. Minimal Examples (`minimal/`)

The simplest possible plugins showing basic structure:
- **Python**: `python_minimal.py` - Hello world function
- **Go**: `go_minimal.go` - Hello world function  
- **TypeScript**: `typescript_minimal.ts` - Hello world function

### 2. Calculator Examples (`calculator/`)

Complete arithmetic plugin showing best practices:
- **Python**: `calculator.py` - Type annotations with Pydantic
- **Go**: `calculator.go` - Error handling and type safety
- **TypeScript**: `calculator.ts` - Modern TypeScript patterns

### 3. Complex Examples (`complex/`)

Advanced patterns for real-world plugins:
- Environment variable handling
- External API integration
- Async operations
- Error handling
- Dependency management

## Testing Your Plugins

Each language implementation provides validation tools:

```bash
# Python
python -m ldagent.validate your_plugin.py

# Go  
go run github.com/ld-agent/ld-agent-go/cmd/validate your_plugin.so

# TypeScript
npx ld-agent-typescript validate your_plugin.ts
```

## Plugin Development Guidelines

1. **Start Simple** - Begin with minimal examples
2. **Follow Patterns** - Use the calculator examples as templates
3. **Validate Early** - Test plugins with validation tools
4. **Document Thoroughly** - Include comprehensive docstrings/comments
5. **Handle Errors** - Implement proper error handling patterns

## Language-Specific Notes

### Python
- Use Pydantic `Field` for parameter descriptions
- Follow PEP 8 naming conventions
- Include type hints for all functions

### Go
- Use `map[string]interface{}` for flexible metadata
- Handle errors as return values
- Follow Go naming conventions (PascalCase)

### TypeScript
- Export `moduleInfo` and `moduleExports` constants
- Use proper TypeScript types
- Handle async operations with Promises

## Contributing Examples

To add new examples:

1. Create appropriately named files in relevant directories
2. Follow the specification format exactly
3. Include comprehensive comments
4. Test with validation tools
5. Update this README

## Reference

- [Plugin Specification](../PLUGIN_SPECIFICATION.md)
- [Architecture Guide](../ARCHITECTURE.md)
- [JSON Schemas](../schemas/)

---

**Need help?** Check the [FAQ](../docs/faq.md) or open a [discussion](https://github.com/your-org/ld-agent-spec/discussions). 
