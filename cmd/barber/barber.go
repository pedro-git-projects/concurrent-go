package main

type Barber struct {
	Name     string
	Sleeping bool
}

func NewBarber(name string) *Barber {
	b := &Barber{
		Name:     name,
		Sleeping: false,
	}
	return b
}
