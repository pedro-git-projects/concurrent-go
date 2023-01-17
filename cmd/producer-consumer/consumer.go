package main

import "fmt"

func consumer(producer *Producer, pizzaria *Pizzaria) {
	for pizza := range producer.data {
		if pizza.pizzaNumber <= pizzaria.PizzaTarget {
			if pizza.success {
				fmt.Println(pizza.message)
				fmt.Printf("Order #%d is out for delivery!\n", pizza.pizzaNumber)
			} else {
				fmt.Println(pizza.message)
				fmt.Println("The customer is really mad!")

			}
		} else {
			fmt.Println("Done making pizzas...")
			err := producer.Close()
			if err != nil {
				fmt.Printf("*** Error closing channel: %v\n", err)
			}
		}
	}
}
