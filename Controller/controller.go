//Explanation:
//Order Struct: Represents the data regarding a customer's order.
//OrderProcessor Class: Acts as a service class that contains the business logic for processing an order.
//OrderController Class: Serves as the controller, receiving orders from the user interface layer and delegating processing to the OrderProcessor.
//Client Code: Represents how a client interacts with the system through the OrderController, which manages the order processing workflow.

package main

import "fmt"

// Order represents an order in the system
type Order struct {
	ID    int
	Items []string
}

// NewOrder creates a new order
func NewOrder(id int, items []string) *Order {
	return &Order{ID: id, Items: items}
}

// OrderProcessor handles the processing of orders
type OrderProcessor struct {
}

// ProcessOrder handles the logic for processing an order
func (op *OrderProcessor) ProcessOrder(order *Order) {
	fmt.Printf("Processing order #%d: %v\n", order.ID, order.Items)
	// Here, additional logic can be added, such as payment processing or inventory checking
}

// OrderController is the controller class managing the order processing
type OrderController struct {
	processor *OrderProcessor
}

// NewOrderController initializes the OrderController with a processor
func NewOrderController(processor *OrderProcessor) *OrderController {
	return &OrderController{processor: processor}
}

// SubmitOrder is the method for handling a user request to order
func (oc *OrderController) SubmitOrder(order *Order) {
	fmt.Println("Controller: Received order submission.")
	oc.processor.ProcessOrder(order)
}

func main() {
	processor := &OrderProcessor{}
	controller := NewOrderController(processor)

	order := NewOrder(1, []string{"Laptop", "Smartphone"})
	controller.SubmitOrder(order)
}

//OUTPUT
//
//Controller: Received order submission.
//Processing order #1: [Laptop Smartphone]
