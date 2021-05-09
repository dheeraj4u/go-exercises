package main

// Swap the values of two variables without a temporary variable

import "fmt"

func main() {
	fmt.Println(swap())
}

func swap() []int {
	a, b := 15, 10
	b, a = a, b
	return []int{a, b}
}
