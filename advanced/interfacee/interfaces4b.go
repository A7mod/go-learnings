package interfacee

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64) string
}

type CreditCard struct{}
type PayPal struct{}
type PhonePe struct{}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card.", amount)
}

func (p PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

func (p PhonePe) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PhonePe UPI.", amount)
}

func ProcessPayment(p PaymentProcessor, amount float64) {
	fmt.Println(p.Pay(amount))
}

func ShowPayments() {
	cc := CreditCard{}
	pay := PayPal{}
	pe := PhonePe{}

	ProcessPayment(cc, 1200.00)
	ProcessPayment(pay, 2200.75)
	ProcessPayment(pe, 20000.25)

	fmt.Println("----------------------------")
}
