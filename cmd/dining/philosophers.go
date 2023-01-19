package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var Philosophers = []Philosopher{
	{name: "Platão", rightFork: 0, leftFork: 4},
	{name: "Sócrates", rightFork: 0, leftFork: 1},
	{name: "Aristóteles", rightFork: 1, leftFork: 2},
	{name: "Pitágoras", rightFork: 2, leftFork: 3},
	{name: "Epicuro", rightFork: 3, leftFork: 4},
}

func dine() {
	wg := &sync.WaitGroup{} // when it zeroes, everyone has done eating
	wg.Add(len(Philosophers))

	seated := &sync.WaitGroup{} // when it zeroes, everyone is seated
	seated.Add(len(Philosophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(Philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(Philosophers); i++ {
		go diningProblem(Philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat
	seated.Done()
	fmt.Printf("%s is seated at the table.\n", philosopher.name)

	// eat n times
	for i := HUNGER; i > 0; i-- {
		// get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(EAT_TIME)

		fmt.Printf("\t%s has both forks and is thinking.\n", philosopher.name)
		time.Sleep(EAT_TIME)

		forks[philosopher.rightFork].Unlock()
		forks[philosopher.leftFork].Unlock()

		fmt.Printf("\t%s put down the forks\n", philosopher.name)
	}

	fmt.Printf("%s is satisfied\n", philosopher.name)
	fmt.Printf("%s has left the table\n", philosopher.name)
}
