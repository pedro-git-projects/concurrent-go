package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "Message from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 2"
	}
}

func channelSelect() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	chan1 := make(chan string)
	chan2 := make(chan string)

	go server1(chan1)
	go server2(chan2)

	for {
		select { // select chooses a case at random if more than one matches
		case s1 := <-chan1:
			fmt.Println("Case 1", s1)
		case s2 := <-chan1:
			fmt.Println("Case 2", s2)
		case s3 := <-chan2:
			fmt.Println("Case 3", s3)
		case s4 := <-chan2:
			fmt.Println("Case 4", s4)
		}
	}
}
