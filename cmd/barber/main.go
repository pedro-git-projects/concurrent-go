package main

func main() {
	welcome()
	bs := NewBarberShop()
	bs.AddBarber("Douglas")
	SleepingBarber(bs)
}
