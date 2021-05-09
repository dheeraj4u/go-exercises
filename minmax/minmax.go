package main

import "fmt"

// Implement min and max behavior

// Min returns the smaller of x or y.

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Max returns the larger of x or y.

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(Min(5, 10))
	fmt.Println(Max(5, 10))
}
