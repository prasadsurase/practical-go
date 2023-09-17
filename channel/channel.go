package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("hello world")
	fmt.Println("hola")

	/*
		for i := 0; i < 3; i++ {
			func() {
				go fmt.Println(i)
			}()
		}
	*/

	/*
		for i := 0; i < 3; i++ {
			i := i
			go func() {
				fmt.Println(i)
			}()
		}
	*/

	/*
		for i := 0; i < 3; i++ {
			go func(n int) {
				fmt.Println(n)
			}(i)
		}
	*/

	/*
		for i := 0; i < 3; i++ {
			go func() {
				fmt.Println(i)
			}()
		}
	*/
	time.Sleep(time.Second * 1)

	ch := make(chan string)
	go func() {
		ch <- "hola amigo"
	}()
	msg := <-ch
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("Message #%d", i+1)
		}
		time.Sleep(time.Second * 1)
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("Another Message #%d", i+1)
		}
		close(ch)
		// close(ch) // panic
	}()

	for msg := range ch {
		fmt.Println("got: ", msg)
	}
}
