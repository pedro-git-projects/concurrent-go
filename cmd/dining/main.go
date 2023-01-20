package main

func main() {
	phi0 := NewPhilosopher("Platão", 0, 4)
	phi1 := NewPhilosopher("Sócrates", 0, 1)
	phi2 := NewPhilosopher("Aristóteles", 1, 2)
	phi3 := NewPhilosopher("Pitágoras", 2, 3)
	phi4 := NewPhilosopher("Epicuro", 3, 4)

	var philosophers = []Philosopher{
		*phi0,
		*phi1,
		*phi2,
		*phi3,
		*phi4,
	}

	initMsg()
	p := NewPhilosophers(philosophers)
	p.dine()
	endMsg()
	p.printFinishOrder()
}
