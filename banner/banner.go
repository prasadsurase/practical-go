package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	text := "Go"
	fmt.Println("Length of Go:", len(text))
	banner(text, 8)

	for i, r := range text {
		fmt.Println(i, r)
		fmt.Printf("%c of type %T\n", r, r)
	}

	text = "G😄"
	fmt.Println("Length of G😄:", len(text))
	banner(text, 8)

	for i, r := range text {
		fmt.Println(i, r)
		fmt.Printf("%c of type %T\n", r, r)
	}
}

func banner(text string, width int) {
	// Doesnt work for unicode.
	// padding := (width - len(text)) / 2
	padding := (width - utf8.RuneCountInString(text)) / 2
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
