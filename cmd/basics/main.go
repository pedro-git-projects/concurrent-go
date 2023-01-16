package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	// use wait groups to stop the main go routine from returning before the others
	var wg sync.WaitGroup

	words := []string{
		"s0",
		"s1",
		"s2",
		"s3",
		"s4",
		"s5",
		"s6",
	}

	wg.Add(len(words))

	for i := 0; i < len(words); i++ {
		// go routines are unordered
		go printSomething(fmt.Sprintf("%d: %s", i, words[i]), &wg)
	}

	wg.Add(3) // else the delta becomes negative
	go printSomething("This is the first thing to be printed", &wg)
	go printSomething("This is the second thing to be printed", &wg)
	go printSomething("This is the third thing to be printed", &wg)
	wg.Wait()
}
