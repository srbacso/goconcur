package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Main job2", Amount: 1500},
		{Source: "Main job3", Amount: 520},
		{Source: "Main job4", Amount: 540},
		{Source: "Main job5", Amount: 500},
		{Source: "Main job6", Amount: 300},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income, m *sync.Mutex) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				m.Lock()
				bankBalance += income.Amount
				m.Unlock()
			}
		}(i, income, &balance)
	}

	wg.Wait()

	fmt.Printf("Final account balance: $%d.00", bankBalance)
	fmt.Println()
}
