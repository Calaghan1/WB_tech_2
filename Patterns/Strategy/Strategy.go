// Паттерн "Стратегия" (Strategy) - это поведенческий паттерн проектирования, который позволяет определить семейство алгоритмов, 
// инкапсулировать каждый из них и обеспечивать их взаимозаменяемость. 
// Он позволяет клиентам выбирать подходящий алгоритм из семейства и использовать его независимо от контекста.


package main

import "fmt"

// Strategy interface
type PaymentStrategy interface {
	Pay(amount int)
}

// ConcreteStrategyA
type CreditCardPayment struct{}

func (cc *CreditCardPayment) Pay(amount int) {
	fmt.Printf("Paid %d via credit card\n", amount)
}

// ConcreteStrategyB
type PayPalPayment struct{}

func (pp *PayPalPayment) Pay(amount int) {
	fmt.Printf("Paid %d via PayPal\n", amount)
}

// Context
type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func (cart *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	cart.paymentStrategy = strategy
}

func (cart *ShoppingCart) Checkout(amount int) {
	cart.paymentStrategy.Pay(amount)
}

func main() {
	cart := &ShoppingCart{}

	creditCard := &CreditCardPayment{}
	payPal := &PayPalPayment{}

	cart.SetPaymentStrategy(creditCard)
	cart.Checkout(100)

	cart.SetPaymentStrategy(payPal)
	cart.Checkout(50)
}