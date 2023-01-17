package main

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func NewProducer() *Producer {
	p := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	return p
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func ProducePizzas(producer *Producer, pizzaria *Pizzaria) {
	i := 0
	for { // breaks when quit chan recieves smth
		currentPizza := pizzaria.MakePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case producer.data <- *currentPizza:

			case quitChan := <-producer.quit:
				close(producer.data)
				close(quitChan)
				return
			}
		}
	}
}
