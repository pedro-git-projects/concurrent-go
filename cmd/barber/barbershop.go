package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan Client
	Open            bool
}

func NewBarberShop() *BarberShop {
	clientChan := make(chan Client, NUMBER_OF_SEATS)
	doneChan := make(chan bool)

	shop := &BarberShop{
		ShopCapacity:    NUMBER_OF_SEATS,
		HairCutDuration: CUT_DURATION,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	return shop
}

func (b *BarberShop) AddBarber(name string) {
	brb := NewBarber(name)
	b.NumberOfBarbers++

	go func() {
		color.Yellow("%s goes to the waiting room to check for clients.", brb.Name)
		for {
			if len(b.ClientsChan) == 0 {
				brb.Sleeping = true
				color.Yellow("There is nothing to do so %s takes a nap.", brb.Name)
			}
			client, isOpen := <-b.ClientsChan

			if isOpen {
				if brb.Sleeping {
					brb.Sleeping = false
					color.Yellow("%s wakes %s up.", client, brb.Name)
					b.CutHair(*brb, client)
				} else {
					b.BarberGoHome(*brb)
					return
				}
			}
		}
	}()
}

func (b *BarberShop) CutHair(barber Barber, client Client) {
	color.Green("%s is cutting %s's hair.", barber.Name, client)
	time.Sleep(b.HairCutDuration)
	color.Green("%s has finished cutting %s's hair.", barber, client)
}

func (b *BarberShop) BarberGoHome(barber Barber) {
	color.Cyan("%s is going home", barber.Name)
	b.BarbersDoneChan <- true
}
