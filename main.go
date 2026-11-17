// main.go: This file is the program entry point and
// represents the main method in Bank.java.
package main

import (
	"fmt"
	"gobank/bank"
	"gobank/customer"
)

func main() {
	var b *bank.Bank = bank.NewBank()
	var c *customer.Customer = customer.NewCustomer("Ann")
	b.Add(bank.NewCheckingAccount("01001", c, 100.00))
	b.Add(bank.NewSavingAccount("01002", c, 200.00))
	b.Accrue(0.02)

	fmt.Printf(b.String())
}