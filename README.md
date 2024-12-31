# Combin

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/combin.svg)](https://pkg.go.dev/github.com/akramarenkov/combin)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/combin)](https://goreportcard.com/report/github.com/akramarenkov/combin)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/combin/badge.svg)](https://coveralls.io/github/akramarenkov/combin)

## Purpose

Library that allows you to drawn combinations

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/combin"
)

func main() {
    for combination := range combin.Every([]int{4, 3, 2, 1}) {
        fmt.Println(combination)
    }
    // Output:
    // [4]
    // [4 3]
    // [4 3 2]
    // [4 3 2 1]
    // [4 3 1]
    // [4 2]
    // [4 2 1]
    // [4 1]
    // [3]
    // [3 2]
    // [3 2 1]
    // [3 1]
    // [2]
    // [2 1]
    // [1]
}
```
