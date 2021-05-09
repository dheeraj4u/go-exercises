// Implement the SumOfSquares function which takes an integer, c and returns the sum of all squares between 1 and c. You’ll need to use select statements, goroutines, and channels.

// For example, entering 5 would return 55 because 1^2 + 2^2 + 3^2 + 4^2 + 5^2 = 55

// You can use the following code as a starting point:

/*
package main

import "fmt"

func SumOfSquares(c, quit chan int) {

// your code here

}

func main() {

  mychannel := make(chan int)

  quitchannel:= make(chan int)

  sum:= 0

  go func() {
    for i := 0; i < 6; i++ {
      sum += <-mychannel
    }
    fmt.Println(sum)
  }()
  SumOfSquares(mychannel, quitchannel)
}
*/

// Solution
package main

import "fmt"

func SumOfSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}

func main() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 0
	}()
	SumOfSquares(mychannel, quitchannel)
}
