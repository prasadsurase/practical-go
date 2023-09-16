package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1 []int
	fmt.Println("Length of empty slice: ", len(s1))
	if s1 == nil {
		fmt.Println("Slice is empty")
	}
	// s2 := make([]int, 15)
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%#v\n", s2)

	s3 := s2[1:4]
	fmt.Printf("%#v\n", s3)

	// fmt.Printf("%#v\n", s3[:100]) // panic
	s3 = append(s3, 100)
	fmt.Printf("S3 append 100: %#v\n", s3)
	s3 = append(s3, 101)
	fmt.Printf("S3 append 101: %#v\n", s3)
	fmt.Printf("S2: %#v\n", s2)

	fmt.Printf("S2 len=%d, capacity=%d\n", len(s2), cap(s2))
	fmt.Printf("S3 len=%d, capacity=%d\n", len(s3), cap(s3))

	var s4 []int
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Printf("S4 len=%d, capacity=%d\n", len(s4), cap(s4))

	fmt.Println(concatStrs([]string{"A", "B"}, []string{"C", "D", "E"})) // [A B C D E]

	fmt.Println("Median of [1,2,3]", median([]float64{1, 2, 3}))
	fmt.Println("Median of [1,2,3,4]", median([]float64{1, 2, 3, 4}))
	fmt.Println("Median of [1,2,3,5,9]", median([]float64{1, 2, 3, 5, 9}))
}

func median(values []float64) float64 {
	sort.Float64s(values)
	i := len(values) / 2
	if len(values)%2 == 1 {
		return values[i]
	}
	v := (values[i-1] + values[i]) / 2
	return v
}

func concatStrs(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func appendInt(s []int, v int) []int {
	l := len(s)
	if len(s) < cap(s) { //enough space in current slice.
		s = s[:len(s)+1]
	} else {
		fmt.Printf("Reallocate to new slice with length %d -> %d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}
	s[l] = v
	return s
}
