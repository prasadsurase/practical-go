package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/*
		count := 0
			// using mutex
			var mu sync.Mutex
			var wg sync.WaitGroup

			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := 0; j < 1000; j++ {
						// using mutex
						mu.Lock()
						count++
						mu.Unlock()
					}
				}()
			}
			wg.Wait()
	*/

	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
