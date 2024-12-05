
## Why need to export
In Go, exporting a field or a function means making it accessible from other packages. This is done by capitalizing the first letter of the field or function name. Exporting is necessary for several reasons:

1. Package Encapsulation:
   - Go uses packages to organize code. By default, identifiers (variables, functions, types, etc.) are unexported and only accessible within the same package. Exporting allows other packages to access these identifiers.

2. Inter-package Communication:
   - When you want to share data structures or functions between different packages, you need to export them. This is crucial for modularity and code reuse.

3. Standard Library Usage:
   - Many functions in Go's standard library, such as those in the encoding/json package, rely on reflection to access struct fields. For reflection to work, the fields must be exported.