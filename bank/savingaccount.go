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
func (s *SavingAccount) Accrue(rate float64) {
	s.interest += s.balance*rate
	s.balance += s.balance*rate
}