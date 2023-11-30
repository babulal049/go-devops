package main

import "fmt"

// PaymentProcessor is an interface for processing payments
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// CreditCardProcessor is a struct that will implement PaymentProcessor
type CreditCardProcessor struct{}

func (p CreditCardProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing a credit card payment of $%.2f\n", amount)
}

// PayPalProcessor is another struct that will implement PaymentProcessor
type PayPalProcessor struct{}

func (p PayPalProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing a PayPal payment of $%.2f\n", amount)
}

// MakePayment takes a PaymentProcessor and an amount, and processes the payment
func MakePayment(processor PaymentProcessor, amount float64) {
	processor.ProcessPayment(amount)
}

func main() {
	creditCardProcessor := CreditCardProcessor{}
	paypalProcessor := PayPalProcessor{}

	MakePayment(creditCardProcessor, 50.25) // Pay using credit card
	MakePayment(paypalProcessor, 75.00)     // Pay using PayPal
}
