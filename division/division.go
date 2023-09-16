package main

import (
	"fmt"
)

func saveDiv(a, b int) (q int, err error) {
	defer func() {
		if e := recover(); e != nil {
			// log.Println("Error: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
}

func main() {
	fmt.Println(saveDiv(4, 2))
	fmt.Println(saveDiv(4, 0))
}
