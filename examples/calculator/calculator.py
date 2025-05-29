# =============================================================================
# START OF MODULE METADATA
# =============================================================================
_module_info = {
    "name": "Simple Calculator",
    "description": "Basic arithmetic operations for demonstration",
    "author": "ld-agent Team",
    "version": "1.0.0",
    "platform": "any",
    "python_requires": ">=3.10",
    "dependencies": [],
    "environment_variables": {}
}
# =============================================================================
# END OF MODULE METADATA
# =============================================================================

from typing import Annotated
from pydantic import Field


def add_numbers(
    a: Annotated[float, Field(description="First number to add")],
    b: Annotated[float, Field(description="Second number to add")]
) -> float:
    """Add two numbers together and return the result."""
    return a + b


def subtract_numbers(
    a: Annotated[float, Field(description="Number to subtract from")],
    b: Annotated[float, Field(description="Number to subtract")]
) -> float:
    """Subtract the second number from the first."""
    return a - b


def multiply_numbers(
    a: Annotated[float, Field(description="First number to multiply")],
    b: Annotated[float, Field(description="Second number to multiply")]
) -> float:
    """Multiply two numbers together."""
    return a * b


def divide_numbers(
    a: Annotated[float, Field(description="Number to divide")],
    b: Annotated[float, Field(description="Number to divide by")]
) -> float:
    """Divide the first number by the second."""
    if b == 0:
        raise ValueError("Cannot divide by zero")
    return a / b


# =============================================================================
# START OF EXPORTS
# =============================================================================
_module_exports = {
    "tools": [add_numbers, subtract_numbers, multiply_numbers, divide_numbers]
}
# =============================================================================
# END OF EXPORTS
# ============================================================================= 
