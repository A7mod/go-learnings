package interfacee

import "fmt"

// Interface : BankAccount
type BankAccount interface {
	GetBalance() float64
	Deposit(amount float64)
}

// Savings Account struct
type SavingsAccount struct {
	balance float64
}

// Checking Account struct
type CheckingAccount struct {
	balance float64
}

// SavingsAccount ka GetBalance() Method
func (s SavingsAccount) GetBalance() float64 {
	return s.balance
}

// CheckingAccount ka GetBalance() Method
func (c CheckingAccount) GetBalance() float64 {
	return c.balance
}

// SavingsAccount ka Deposit Method
func (s *SavingsAccount) Deposit(amount float64) {
	s.balance += amount
	fmt.Println("Savings Account: Deposited", amount)
}

// CheckingAccount ka Deposit Method
func (c *CheckingAccount) Deposit(amount float64) {
	c.balance += amount
	fmt.Println("Checking Account: Deposited", amount)
}

func ShowAccounts() {
	// Polymorphism: Interface ka reference dono structs ko hold kar skta h.
	var account1 BankAccount = &SavingsAccount{balance: 15000}
	var account2 BankAccount = &CheckingAccount{balance: 20000}

	account1.Deposit(2500)
	account2.Deposit(1500)

	fmt.Println("Savings Account Balance:", account1.GetBalance())
	fmt.Println("Checkimg Account Balance:", account2.GetBalance())

	fmt.Println("----------------------------")
}
