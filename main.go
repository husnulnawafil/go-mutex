package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var balance sync.Mutex

type Income struct {
	Source string
	Amount int
}

func main() {
	// variables for bank balance
	var bankBalance int

	// print out starting values
	fmt.Printf("initial account balance : %d.00\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifs", Amount: 10},
		{Source: "Part Time Job", Amount: 50},
		{Source: "Investment", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d you have earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	// print out final balance
	fmt.Printf("Final Bank Balance : $%d.00\n", bankBalance)

}
