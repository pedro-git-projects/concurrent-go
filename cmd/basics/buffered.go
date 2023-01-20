package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Printf("Got %d from channel\n", i)
		time.Sleep(1 * time.Second)
	}
}

func buffChan() {
	ch := make(chan int, 100) // Nothing is suspended
	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Printf("Sending %d to channel...\n", i)
		ch <- i
		fmt.Println("sent ", i, " to channel")
	}
	fmt.Println("Done")
	close(ch)
}
