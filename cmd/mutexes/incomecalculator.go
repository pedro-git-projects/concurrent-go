package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func calculateYearlyIncome() {
	wg := sync.WaitGroup{}
	balance := new(int)
	balanceMutex := sync.Mutex{}

	fmt.Printf("Initial balance: $%d.00\n", *balance)

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balanceMutex.Lock()
				temp := *balance
				temp += income.Amount
				*balance = temp
				balanceMutex.Unlock()

				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()
	fmt.Printf("Final balance: $%d.00\n", *balance)
}
