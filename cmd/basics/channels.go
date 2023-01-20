package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping                                      // recieves value in ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s)) // recieves formatted s
	}
}

func readAndRespond() {
	ping := make(chan string)
	pong := make(chan string)
	go shout(ping, pong)

	fmt.Println("Enter a value and press RETURN (enter Q to quit)")

	for {
		fmt.Print("-> ")
		input := new(string)
		fmt.Scanf("%s", input)

		if *input == strings.ToLower("q") {
			break
		}

		ping <- *input     // sends to ping
		response := <-pong // recieves from pong
		fmt.Printf("Response: %s\n", response)
	}
	fmt.Println("All done! Closing channels")
	close(ping)
	close(pong)
}
