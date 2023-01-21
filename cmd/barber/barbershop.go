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
	barber := NewBarber(name)
	b.NumberOfBarbers++

	go func() {
		color.Yellow("%s goes to the waiting room to check for clients.", barber.Name)
		for {
			if len(b.ClientsChan) == 0 {
				barber.Sleeping = true
				color.Yellow("There is nothing to do so %s takes a nap.", barber.Name)
			}

			client, isOpen := <-b.ClientsChan

			if isOpen {
				if barber.Sleeping {
					barber.Sleeping = false
					color.Yellow("%s wakes %s up.", client.Name, barber.Name)
				}
				b.CutHair(*barber, client)
			} else {
				b.BarberGoHome(*barber)
				return
			}
		}
	}()
}

func (b *BarberShop) CutHair(barber Barber, client Client) {
	color.Green("%s is cutting %s's hair.", barber.Name, client.Name)
	time.Sleep(b.HairCutDuration)
	color.Green("%s has finished cutting %s's hair.", barber.Name, client.Name)
}

func (b *BarberShop) BarberGoHome(barber Barber) {
	color.Cyan("%s is going home.", barber.Name)
	b.BarbersDoneChan <- true
}

func (b *BarberShop) CloseShop() {
	color.Cyan("Closing shop for the day.")
	close(b.ClientsChan)
	b.Open = false

	for a := 1; a <= b.NumberOfBarbers; a++ {
		<-b.BarbersDoneChan
	}
	close(b.BarbersDoneChan)

	color.Green("The barber shop is now closed for the day, and everyone has gone home.")
	color.Green("---------------------------------------------------------------------")
}

func (b *BarberShop) addClient(client Client) {
	color.Green("*** %s has arrived.", client.Name)

	if b.Open {
		select {
		case b.ClientsChan <- client:
			color.Yellow("%s seats in the waiting room.", client.Name)
		default:
			color.Red("The waiting room is full, so %s goes home.", client.Name)
		}
	} else {
		color.Red("The shop is already closed, so %s goes home.", client.Name)
	}
}
