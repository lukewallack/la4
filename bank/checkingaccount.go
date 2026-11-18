// A checking account is just a regular account. 
// We do not accrue interest here.
package bank

import "gobank/customer"

type CheckingInterface interface {
	AccountInterface
}

type CheckingAccount struct {
	Account // Checking only has Account members.
}

func NewCheckingAccount(number string, customer *customer.Customer, balance float64) *CheckingAccount {
	return &CheckingAccount {
		Account: Account {
			number: 	number,
			customer: 	customer,
			balance: 	balance,
		},
	}
}

// We do not accrue interest in this checking account.
func (s *CheckingAccount) Accrue(rate float64, accrueChan chan<- float64) { accrueChan <- 0 }