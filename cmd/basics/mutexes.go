package main

import (
	"fmt"
	"sync"
)

func commonIncrement() {
	count := new(int)
	lock := &sync.Mutex{}

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		*count++
		fmt.Printf("Incrementing: %d\n", *count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		*count--
		fmt.Printf("Decrementing: %d\n", *count)
	}

	// Increment
	arithmetic := &sync.WaitGroup{}
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic completed")
}
