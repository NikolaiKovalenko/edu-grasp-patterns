//Explanation:
//PaymentService Interface: Defines a common interface for all payment service implementations, allowing for loose coupling.
//Concrete Implementations (PaypalService, StripeService): Each implements the PaymentService interface, providing specific payment processing logic.
//PaymentProcessor: Acts as the intermediary layer that decouples the client from the payment services. It uses the PaymentService interface to process payments, allowing easy switching between different payment services.
//Client Code: Demonstrates how to use PaymentProcessor to handle payments through different services by selecting the service at runtime, showcasing flexibility and modularity.

package main

import "fmt"

// PaymentService interface defines the contract for payment services
type PaymentService interface {
	ProcessPayment(amount float64) string
}

// PaypalService is a concrete implementation of PaymentService
type PaypalService struct{}

func (p *PaypalService) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of $%.2f through PayPal.", amount)
}

// StripeService is another concrete implementation of PaymentService
type StripeService struct{}

func (s *StripeService) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of $%.2f through Stripe.", amount)
}

// PaymentProcessor provides an indirection layer between the client and payment services
type PaymentProcessor struct {
	paymentService PaymentService
}

// NewPaymentProcessor creates a new PaymentProcessor with the specified PaymentService
func NewPaymentProcessor(service PaymentService) *PaymentProcessor {
	return &PaymentProcessor{paymentService: service}
}

// Pay processes the payment using the specified PaymentService
func (pp *PaymentProcessor) Pay(amount float64) {
	result := pp.paymentService.ProcessPayment(amount)
	fmt.Println(result)
}

func main() {
	// Client choosing PayPal service
	paypal := &PaypalService{}
	paymentProcessor := NewPaymentProcessor(paypal)
	paymentProcessor.Pay(100.0)

	// Client switching to Stripe service
	stripe := &StripeService{}
	paymentProcessor = NewPaymentProcessor(stripe)
	paymentProcessor.Pay(200.0)
}

// OUTPUT
//
//Processed payment of $100.00 through PayPal.
//Processed payment of $200.00 through Stripe.
