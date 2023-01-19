package main

import "fmt"

func main() {
	stringStream := make(chan string) // bidirectional channel
	readStream := make(chan<- interface{})
	// writeStream := make(<-chan interface{})

	go func() {
		// sending
		stringStream <- "Hello channels"
		readStream <- "Sending"
		// writeStream <- "Sending" will error out
	}()
	fmt.Println(<-stringStream) // reading

}
