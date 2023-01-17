package main

import "fmt"

func race() {
	msg := "Hello wrold!"

	wg.Add(2)
	/* Race condition  to access msg */
	go updateMessage("Hello, universe!", msg, &wg)
	go updateMessage("Hello, cosmos!", msg, &wg)
	wg.Wait()

	fmt.Println(msg)
}
