package main

import (
	"errors"
	"fmt"
	"sync"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
	mutex   sync.Mutex
}

func main() {
	wallet := Wallet{}
	wallet.Deposit(10)
	fmt.Println("Balance after deposit:", wallet.Balance())

	err := wallet.Withdraw(5)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Balance after withdrawal:", wallet.Balance())
}

//put some amount
func (w *Wallet) Deposit(amount Bitcoin) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.balance += amount
}

// Takes some amount from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if amount > w.balance {
		return errors.New("insufficient funds")
	}

	w.balance -= amount
	return nil
}

// Returns current balance
func (w *Wallet) Balance() Bitcoin {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.balance
}