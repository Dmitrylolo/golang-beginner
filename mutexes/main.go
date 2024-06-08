package main

import (
	"fmt"
	"sync"
)

type Account struct {
	Balance int
	Mutex   *sync.Mutex
}

func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance -= amount
	a.Mutex.Unlock()

	wg.Done()
}

func (a *Account) Deposit(amoount int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance += amoount
	a.Mutex.Unlock()

	wg.Done()
}

func main() {
	fmt.Println("Mutexes")
	var m sync.Mutex

	account := Account{
		Balance: 1000,
		Mutex:   &m,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go account.Withdraw(700, &wg)
	go account.Deposit(500, &wg)
	wg.Wait()
	fmt.Println("Balance updated")
	fmt.Println("Balance:", account.Balance)
}
