package main

import "fmt"

type Number interface {
	int | float64
}

func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	maxx := nums[0]
	for _, n := range nums[1:] {
		if n > maxx {
			maxx = n
		}
	}
	return maxx
}

func max2[T int | float64](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	maxx := nums[0]
	for _, n := range nums[1:] {
		if n > maxx {
			maxx = n
		}
	}
	return maxx
}

func main() {
	fmt.Println(max([]int{2, 1, 3}))
	fmt.Println(max([]float64{2, 1, 3}))
	fmt.Println(max2([]int{2, 1, 3}))
	fmt.Println(max2([]float64{2, 1, 3}))
}
