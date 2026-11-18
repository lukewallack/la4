// Bank.go: resembles Bank.java, defines the Bank "class".
package bank

import "fmt"

type BankInterface interface {
	String() string
	Accrue()
}

type Bank struct {
	accounts map[AccountInterface]struct{} // Contains our set of accounts
}

// Constructor: returns a new bank that contains an empty 
// set of AccountInterface values.
func NewBank() *Bank {
	return &Bank {
		accounts: make(map[AccountInterface]struct{}),
	}
}

// Add an account to a Bank.
func (b *Bank) Add(account AccountInterface) {
	b.accounts[account] = struct{}{}
}

// Accrue interest on all accounts.
func (b *Bank) Accrue(rate float64) {
	// Set up a channel for accrue goroutines
	accrueChan := make(chan float64)

	// Loop through each account, ignoring value since it's an empty struct.
	for a, _ := range b.accounts {
		go a.Accrue(rate, accrueChan)
	}

	// Compute total interest
	total := 0.0
	for range b.accounts {
		total += <-accrueChan
	}

	// Print out total interest
	fmt.Printf("Total interest: %.2f\n", total)
}

// Tostring: print out all of a bank's accounts.
func (b *Bank) String() string {
	r := ""
	for a, _ := range b.accounts { // Ignore value, we only care about key
		r += a.String() + "\n"
	}
	return r
}