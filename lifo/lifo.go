package main

import "fmt"

func main() {
	// Implement a Stack (LIFO)

	// Create

	var stack []string

	// Push

	stack = append(stack, "world!")

	stack = append(stack, "Hello ")

	for len(stack) > 0 {

		// Print top

		n := len(stack) - 1

		fmt.Print(stack[n])

		// Pop

		stack = stack[:n]

	}

	// Output: Hello world!

}
