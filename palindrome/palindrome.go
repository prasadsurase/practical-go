package main

import (
	"fmt"
)

func main() {
	text := "g"
	fmt.Printf("%v is palindrome: %v\n", text, isPalindrome(text))
	text = "go"
	fmt.Printf("%v is palindrome: %v\n", text, isPalindrome(text))
	text = "gog"
	fmt.Printf("%v is palindrome: %v\n", text, isPalindrome(text))
	text = "gogo"
	fmt.Printf("%v is palindrome: %v\n", text, isPalindrome(text))
	text = "gogog"
	fmt.Printf("%v is palindrome: %v\n", text, isPalindrome(text))
	text = "gogogo"
	fmt.Printf("%v is palindrome: %v\n", text, isPalindrome(text))
	text = "gogogog"
	fmt.Printf("%v is palindrome: %v\n", text, isPalindrome(text))
	text = "g😄g😄g"
	fmt.Printf("%v is palindrome: %v\n", text, isPalindrome(text))
}

// Fails
// func isPalindrome(text string) bool {
// 	fmt.Println("=======================================")
// 	is_palindrome := true
// 	text_len := utf8.RuneCountInString(text)
// 	for i := 0; i < (text_len / 2); i++ {
// 		fmt.Printf("%v %v\n", text[i], text[text_len-i-1])
// 		if text[i] != text[text_len-1] {
// 			return false
// 		} else {
// 			continue
// 		}
// 	}
// 	return is_palindrome
// }

func isPalindrome(text string) bool {
	rs := []rune(text)
	for i := 0; i < len(rs)/2; i++ {
		// fmt.Printf("%v %v\n", rs[i], rs[len(rs)-i-1])
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
}
