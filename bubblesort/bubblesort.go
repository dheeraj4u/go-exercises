package main

import "fmt"

func main() {
	data := []int{4, 2, 1, 5, 3, 4}
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]

			}
		}
	}
	fmt.Println(data)
}
