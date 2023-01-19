package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHello() {
	defer wg.Done()
	fmt.Printf("Hello!\n")
}

func main() {

	wg.Add(4)

	salutation := "Hello"
	go func() {
		defer wg.Done()
		salutation = "Welcome"
	}() // closure variables can be accessed in goroutines

	for _, salutation := range []string{"Olá", "Hallo", "Konnichwa"} {
		wg.Add(1)
		go func(salutation string) { // ranged variables need to be passed
			// as parameters to closures
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}

	go sayHello() // named

	go func() {
		defer wg.Done()
		fmt.Printf("Hello!\n")
	}() // anonymous

	helloAnon := func() {
		defer wg.Done()
		fmt.Printf("Hello!\n")
	} // anonymous saved on var

	go helloAnon()
	wg.Wait()
	fmt.Println(salutation)
}
