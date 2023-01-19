package main

import (
	"fmt"
	"sync"
	"time"
)

func sleeping(n int, wg *sync.WaitGroup) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine #%d is sleeping...\n", i+1)
			time.Sleep(1)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines completed")
}

// come back to this problem after channels
func spinner(delay time.Duration, wg *sync.WaitGroup) {
	wg.Add(4)
	for _, r := range `-\|/` {
		go func(r rune) {
			defer wg.Done()
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}(r)
	}
	wg.Wait()
	fmt.Println()
}
