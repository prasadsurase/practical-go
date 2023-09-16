package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer file.Close()

	data, err := wordFrequency(file)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("%#v", data)
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	wordsCount := make(map[string]int) // count of words
	for s.Scan() {
		words := strings.Split(s.Text(), " ") //current line's text
		if len(words) != 0 {
			for _, word := range words {
				word = strings.ToLower(word)
				// TODO: remove special characters
				// TODO: return most common N words
				if _, ok := wordsCount[word]; ok {
					wordsCount[word]++
				} else {
					wordsCount[word] = 1
				}
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("Error: %s", err)
	}
	return wordsCount, nil
}
