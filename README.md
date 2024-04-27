# Collectify: Laravel Collections-like Functionality for Go

Collectify is a Go package designed to bring the convenience and power of Laravel's collection to the Go programming language. With Collectify, you can manipulate arrays, slices, maps, and other data structures with ease, using expressive methods and fluent syntax.

## Features

- **Map, Filter, Reduce**: Easily transform and manipulate data using familiar functional programming methods.
- **Chaining**: Chain multiple operations together fluently for concise and readable code.
- **Powerful Iteration**: Iterate over collections using closures or predefined functions.
- **Lazy Evaluation**: Operations are lazily evaluated, improving performance and memory efficiency.
- **Supports Various Data Types**: Works seamlessly with arrays, slices, maps, and custom data types.
- **Method Chaining**: Chain multiple operations together in a single statement for clean, readable code.

## Installation

To install Collectify, use `go get github.com/nopetimizer/collectify`:

```bash
go get github.com/nopetimizer/collectify
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/nopetimizer/collectify"
)

func main() {
    // Create a new collection
    col := collectify.NewIndexed[int]([]int{1, 2, 3, 4, 5})

    // Map each element to its square
    squares := col.Map(func(i interface{}) interface{} {
        return i.(int) * i.(int)
    })

    // Filter elements greater than 10
    filtered := squares.Filter(func(i interface{}) bool {
        return i.(int) > 10
    })

    // Reduce to find the sum
    sum := filtered.Reduce(0, func(accumulator interface{}, currentValue interface{}) interface{} {
        return accumulator.(int) + currentValue.(int)
    })

    fmt.Println("Sum of squares greater than 10:", sum) // Output: Sum of squares greater than 10: 54
}
```

## Contributing

Contributions are welcome! If you have ideas for new features, improvements, or find any bugs, please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
