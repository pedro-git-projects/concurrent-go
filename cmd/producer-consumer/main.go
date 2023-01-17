package main

import "fmt"

func main() {
	fmt.Println("The pizzaria is open for business")
	fmt.Println("---------------------------------")

	pizzaria := NewPizzaria(10)
	producer := NewProducer()

	go ProducePizzas(producer, pizzaria)

	consumer(producer, pizzaria)

	fmt.Println("Done for the day")
	fmt.Println("----------------")
	fmt.Printf("We made %d pizzas, but failed to make %d, with %d attempts total\n", pizzaria.PiazzasMade, pizzaria.PiazzasFailed, pizzaria.Total)

}
