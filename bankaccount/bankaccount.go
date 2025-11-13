package bankaccount

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	mu      sync.Mutex
	balance int
}

func NewAccount() *BankAccount {
	return &BankAccount{
		balance: 0,
	}
}

func (a *BankAccount) Balance() int {
	return a.balance
}

func (a *BankAccount) Withdraw(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amount <= 0 {
		return fmt.Errorf("not allowed to withdraw amount equal to or lower than zero")
	}

	if a.balance < amount {
		return fmt.Errorf("insufficient funds: balance: %d attempted %d", a.balance, amount)
	}

	a.balance = a.balance - amount

	return nil
}

func (a *BankAccount) Deposit(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amount <= 0 {
		return fmt.Errorf("not allowed to deposit amount equal or lower than zero")
	}

	a.balance += amount
	return nil
}
