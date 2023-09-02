package main

import (
	"fmt"
)

func main() {

	fmt.Println("Length of Go:", len("Go"))
	fmt.Println("Length of Go 😄:", len("Go 😄"))
	banner("Go", 8)
	banner("Go 😄", 8)
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println("")
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println("")
}
