package main

import (
	"fmt"
	"time"

	"github.com/pedro-git-projects/concurrent-go/pkg"
)

func SleepingBarber(b *BarberShop) {
	shopClosing := make(chan bool)
	closed := make(chan bool)
	rng := pkg.RNG()

	go func() {
		<-time.After(TIME_OPEN)
		shopClosing <- true
		b.CloseShop()
		closed <- true
	}()

	i := 1
	go func() {
		for {
			randomArrivalMs := rng.Int() % (2 * ARRIVAL_RATE)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomArrivalMs)):
				c := NewClient(fmt.Sprintf("Client #%d", i))
				b.addClient(*c)
				i++
			}
		}
	}()

	<-closed
}
