// A savings account is just a regular account with
// interest that accrues.
package bank

import "gobank/customer"

type SavingInterface interface {
	AccountInterface
}

type SavingAccount struct {
	Account 			// All account members
	interest float64 	// Plus interest
}

// Constructor: create a new saving account
func NewSavingAccount(number string, customer *customer.Customer, balance float64) *SavingAccount {
	return &SavingAccount {
		Account: Account {
			number: 	number,
			customer: 	customer,
			balance: 	balance,
		},
		interest: 0,
	}
}

// Saving accounts accrue interest
func (s *SavingAccount) Accrue(rate float64, accrueChan chan<- float64) {
	amount := s.balance*rate
	s.interest += amount
	s.balance += amount
	accrueChan <- amount
}