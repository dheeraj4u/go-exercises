/*Implement a concurrent Merge Sort solution using goroutines and channels.

You can use this sequential Merge Sort implementation as a starting point:



package main
import "fmt"

func Merge(left, right [] int) [] int{
  merged := make([] int, 0, len(left) + len(right))

  for len(left) > 0 || len(right) > 0{

    if len(left) == 0 {
      return append(merged,right...)

    }else if len(right) == 0 {
      return append(merged,left...)

    }else if left[0] < right[0] {
      merged = append(merged, left[0])
      left = left[1:]

    }else{
      merged = append(merged, right [0])
      right = right[1:]
    }
  }
  return merged
}

func MergeSort(data [] int) [] int {
  if len(data) <= 1 {
    return data
  }
  mid := len(data)/2
  left := MergeSort(data[:mid])
  right := MergeSort(data[mid:])
  return Merge(left,right)
}

func main(){
  data := [] int{9,4,3,6,1,2,10,5,7,8}
  fmt.Printf("%v\n%v\n", data, MergeSort(data))
}*/

// Solution

package main

import "fmt"

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	done := make(chan bool)
	mid := len(data) / 2
	var left []int
	go func() {
		left = MergeSort(data[:mid])
		done <- true
	}()
	right := MergeSort(data[mid:])
	<-done
	return Merge(left, right)
}

func main() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n%v\n", data, MergeSort(data))
}

// Firstly, in merge sort, we keep dividing our array recursively into the right side and the left side and call the MergeSort function on both sides from line 30 to line 34.

// Now we have to make sure that Merge(left,right) is executed after we get return values from both the recursive calls, i.e. both the left and right must be updated before Merge(left,right) can be executable. Hence, we introduce a channel of type bool on line 26 and send true on it as soon as left = MergeSort(data[:mid]) is executed (line 32).

// The <-done operation blocks the code on line 35 before the statement Merge(left,right) so that it does not proceed until our goroutine has finished. After the goroutine has finished and we receive true on the done channel, the code proceeds forward to Merge(left,right) statement on line 36.
