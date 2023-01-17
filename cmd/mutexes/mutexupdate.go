package main

import (
	"fmt"
	"sync"
)

// thread safe
func mutexUpdate(wg *sync.WaitGroup) {
	msg := "Hello, world!"
	mutex := sync.Mutex{}

	wg.Add(2)
	go updateMessageMutex("Hello, universe!", &msg, &mutex, wg)
	go updateMessageMutex("Hello, cosmos!", &msg, &mutex, wg)
	wg.Wait()

	fmt.Println(msg)
}
