package main

import (
	"fmt"
	"time"
)

type Pizzaria struct {
	PizzaTarget   int
	PiazzasFailed int
	PiazzasMade   int
	Total         int
}

func NewPizzaria(target int) *Pizzaria {
	p := &Pizzaria{
		PizzaTarget: target,
	}
	return p
}

func (p *Pizzaria) MakePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if p.PizzaTarget >= pizzaNumber {
		delay := getRandomNumber(5)
		fmt.Printf("Recieved order #%d!\n", pizzaNumber)

		rnd := getRandomNumber(12)
		msg := ""
		success := false

		if rnd >= 5 {
			p.PiazzasMade++
		} else {
			p.PiazzasFailed++
		}
		p.Total++

		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The coock quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return p
	} else {
		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
		}
	}
}
