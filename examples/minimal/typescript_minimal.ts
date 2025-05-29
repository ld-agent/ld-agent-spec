// =============================================================================
// MINIMAL LD-AGENT PLUGIN EXAMPLE (TypeScript)
// =============================================================================

export const moduleInfo = {
    name: "Minimal Example",
    description: "The simplest possible ld-agent plugin",
    version: "1.0.0"
};

export function helloWorld(): string {
    return "Hello from ld-agent!";
}

export const moduleExports = {
    tools: [helloWorld]
}; 