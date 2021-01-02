package main

import (
	"fmt"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

// Calculate returns x * 3.
func CalculateMultiply(x int) (result int) {
	result = x * 3
	return result
}

func main() {
	fmt.Println("Hello World")
}
