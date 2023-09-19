package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(time.Millisecond * time.Duration(10))
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Millisecond * time.Duration(20))
		ch2 <- 2
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	select {
	case val := <-ch1:
		fmt.Printf("Channel 1 received %d\n", val)
	case val := <-ch2:
		fmt.Printf("Channel 2 received %d\n", val)

	case <-ctx.Done(): // case <-time.After(5 * time.Millisecond):
		fmt.Println("timeout")
	}

	fmt.Println("Hello world")
}
