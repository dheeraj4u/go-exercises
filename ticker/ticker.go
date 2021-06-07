package main

import (
	"fmt"
	"time"
)

// func main() {

// 	fmt.Println("Ticker starts here")

// 	ticker := time.NewTicker(1 * time.Second)

// 	for _ = range ticker.C {
// 		fmt.Println("Tock")
// 	}
// }

// ticker with goroutine
// func bg() {
// 	ticker := time.NewTicker(1 * time.Second)
// 	for _ = range ticker.C {
// 		fmt.Println("Tock")
// 	}
// }

// func main() {
// 	fmt.Println("Ticker started")

// 	go bg()

// 	//go other func here

// 	//select for continious loop
// 	select {}
// }

//Ticker with channels to stop

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tock", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	done <- true
	fmt.Println("Ticker stopped")

}
