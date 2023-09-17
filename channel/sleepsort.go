package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{9, 3, 5, 3, 2, 5, 1, 4, 6, 2, 1, 8}
	fmt.Printf("Input: %#v\n", values)
	data := sleepSort(values)
	fmt.Printf("Input: %#v\n", data)
}

func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		// fmt.Println("Value: ", n)
		go func(j int) {
			time.Sleep(time.Duration(j) * time.Millisecond) // this helps sort.
			// fmt.Println("J Value: ", j)
			ch <- j
		}(n)
	}
	// close(ch)

	// out := make([]int, 10, 10)
	var out []int
	for range values {
		out = append(out, <-ch)
	}
	return out
}
