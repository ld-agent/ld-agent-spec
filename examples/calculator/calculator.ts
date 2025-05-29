// =============================================================================
// CALCULATOR PLUGIN EXAMPLE (TypeScript)
// =============================================================================

export const moduleInfo = {
    name: "Simple Calculator",
    description: "Basic arithmetic operations for demonstration",
    author: "ld-agent Team",
    version: "1.0.0",
    platform: "any",
    node_version: ">=18.0.0",
    dependencies: [],
    environment_variables: {}
};

export function addNumbers(a: number, b: number): number {
    return a + b;
}

export function subtractNumbers(a: number, b: number): number {
    return a - b;
}

export function multiplyNumbers(a: number, b: number): number {
    return a * b;
}

export function divideNumbers(a: number, b: number): number {
    if (b === 0) {
        throw new Error("Cannot divide by zero");
    }
    return a / b;
}

export const moduleExports = {
    tools: [addNumbers, subtractNumbers, multiplyNumbers, divideNumbers]
}; 