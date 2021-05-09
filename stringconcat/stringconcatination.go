package main

import "fmt"

func main() {

	// Creating and initializing strings
	// using var keyword
	var str1 string
	str1 = "Hello "

	var str2 string
	str2 = "Reader!"

	// Concatenating strings
	// Using + operator
	fmt.Println("New string 1: ", str1+str2)

	// Creating and initializing strings
	// Using shorthand declaration
	str3 := "Welcome"
	str4 := "Educative.io"

	// Concatenating strings
	// Using + operator
	result := str3 + " to " + str4

	fmt.Println("New string 2: ", result)

}
