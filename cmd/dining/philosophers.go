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
	ate       bool
	lock      *sync.Mutex
}

type Fork map[int]*sync.Mutex

type Philosophers struct {
	List       []Philosopher
	FinishList []Philosopher
	Size       int
	Forks      Fork
	Done       *sync.WaitGroup
	Seated     *sync.WaitGroup
}

func NewPhilosopher(name string, rightFork, leftFork int) *Philosopher {
	mu := &sync.Mutex{}
	p := &Philosopher{
		name:      name,
		rightFork: rightFork,
		leftFork:  leftFork,
		ate:       false,
		lock:      mu,
	}
	return p
}

func NewPhilosophers(list []Philosopher) *Philosophers {
	p := &Philosophers{}

	p.List = list
	p.Size = len(p.List)

	p.Forks = make(Fork)
	for i := 0; i < p.Size; i++ {
		p.Forks[i] = &sync.Mutex{}
	}

	done := &sync.WaitGroup{}
	seated := &sync.WaitGroup{}
	p.Done = done
	p.Seated = seated

	return p
}

func (p *Philosophers) dine() {
	p.Done.Add(p.Size)
	p.Seated.Add(p.Size)

	for i := 0; i < p.Size; i++ {
		go eat(p.List[i], p.Done, p.Forks, p.Seated, &p.FinishList)
	}

	p.Done.Wait()
}

func (p Philosophers) printFinishOrder() {
	for i := range p.FinishList {
		fmt.Printf("%s finished in position #%d\n", p.FinishList[i].name, i+1)
	}
}

func eat(philosopher Philosopher, wg *sync.WaitGroup, forks Fork, seated *sync.WaitGroup, finishedList *[]Philosopher) {
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

	philosopher.lock.Lock()
	philosopher.ate = true
	*finishedList = append(*finishedList, philosopher)
	philosopher.lock.Unlock()
}
