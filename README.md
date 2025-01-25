# Unexpected Zero Value on Missing Map Keys in Go

This repository demonstrates an uncommon yet potentially problematic behavior in Go: accessing a non-existent key in a map returns the zero value of the value type instead of signaling an error.  This can lead to subtle bugs if not carefully considered.

The `bug.go` file contains code showcasing this behavior. Accessing keys that haven't been set will result in a `0` being printed, which might be mistaken for an actual value.

The `bugSolution.go` offers a few strategies for mitigating this issue, such as checking for key existence using the comma ok idiom.

## Contributing

Feel free to contribute improvements or alternative solutions!