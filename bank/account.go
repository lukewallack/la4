package bank // Accounts belong in banks I think

import (
	"strconv"
	"gobank/customer"
)

type AccountInterface interface {
	Accrue(rate float64, accrueChan chan<- float64)
	Deposit(amount float64)
	Withdraw(amount float64)
	Balance() 	float64
	String() 	string
}

type Account struct {
	number 		string
	customer 	*customer.Customer
	balance 	float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	a.balance -= amount
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) String() string {
	return a.number + ":" + a.customer.String() + ":" + strconv.FormatFloat(a.balance, 'f', 2, 64)
}