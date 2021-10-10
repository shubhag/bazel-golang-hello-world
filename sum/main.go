package main

import (
	"flag"
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

// Main method takes 2 numbers as arguments and returns their sum.
func main() {
	var num1, num2 int
	flag.IntVar(&num1, "first", 1, "First number")
	flag.IntVar(&num2, "second", 2, "Second number")
	flag.Parse()

	fmt.Printf("Sum of provided numbers: %d\n", sum(num1, num2))
}
