![logo](https://github.com/user-attachments/assets/b711049a-3873-4a1c-8cbc-9e69e5b5f104)

# Resulto


[![Go Reference](https://pkg.go.dev/badge/github.com/yourusername/resulto.svg)](https://pkg.go.dev/github.com/yourusername/resulto)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/resulto)](https://goreportcard.com/report/github.com/yourusername/resulto)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Rust-inspired Result type for Go, enabling elegant error handling with generics.

## Features

- Generic Result type supporting any value and error types
- Chainable operations for cleaner error handling
- Panic-free error propagation
- Fully tested and production-ready

## Installation

```bash
go get github.com/yourusername/resulto
```

## Usage

### Basic Usage

```go
package main

import (
    "errors"
    "fmt"
    "github.com/yourusername/resulto"
)

func divide(a, b int) resulto.Result[int] {
    if b == 0 {
        return resulto.Failure[int](errors.New("division by zero"))
    }
    return resulto.Success(a / b)
}

func main() {
    // Successful result
    result1 := divide(10, 2)
    if result1.IsOk() {
        fmt.Printf("Result: %d\n", result1.Unwrap())
    }
    
    // Failed result with safe unwrap
    result2 := divide(5, 0)
    value := result2.UnwrapOr(0)
    fmt.Printf("Safe value: %d\n", value)
}
```

### Advanced Usage

```go
package main

import (
    "errors"
    "fmt"
    "github.com/yourusername/resulto"
    "strconv"
)

func parseAndDouble(input string) resulto.Result[int] {
    return parseString(input).Map(func(n int) int {
        return n * 2
    })
}

func parseString(input string) resulto.Result[int] {
    n, err := strconv.Atoi(input)
    if err != nil {
        return resulto.Failure[int](err)
    }
    return resulto.Success(n)
}

func main() {
    // Chaining operations
    parseAndDouble("5").Match(
        func(n int) { fmt.Printf("Doubled: %d\n", n) },
        func(err error) { fmt.Printf("Error: %s\n", err.Error()) },
    )
    
    // Error handling
    parseAndDouble("invalid").Match(
        func(n int) { fmt.Printf("Doubled: %d\n", n) },
        func(err error) { fmt.Printf("Error: %s\n", err.Error()) },
    )
}
```

## API Reference

### Creating Results

- `Success[T](value T) Result[T]`: Create a successful result
- `Failure[T](err error) Result[T]`: Create a failed result

### Methods

- `IsOk() bool`: Check if result is successful
- `Unwrap() T`: Get value or panic if error
- `UnwrapOr(def T) T`: Get value or default if error
- `UnwrapErr() error`: Get error or panic if successful

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
