# =============================================================================
# MINIMAL LD-AGENT PLUGIN EXAMPLE (Python)
# =============================================================================

_module_info = {
    "name": "Minimal Example",
    "description": "The simplest possible ld-agent plugin",
    "version": "1.0.0"
}

def hello_world() -> str:
    """Returns a simple greeting."""
    return "Hello from ld-agent!"

_module_exports = {
    "tools": [hello_world]
} 