//Explanation:
//Item Struct: Represents individual items in a store.
//NewItem Function: Serves as a factory function, but its responsibility could reside in Cart to closely tie the creation of an item to its usage context.
//Cart Struct: Acts as a Creator, responsible for creating Item instances through the AddItemToCart method and managing a collection of items.
//Client Code: Demonstrated in the main function, where items are added directly to the cart, showcasing how responsibility for item creation is contained within the Cart class.
//Summary:
//The Creator pattern helps in putting responsibility for the creation of related objects within a class that aggregates or uses them, thereby guiding object creation in a way that enhances cohesion and reduces unnecessary dependencies. By following this principle, you can create maintainable and scalable systems due to well-defined class responsibilities.

package main

import (
	"fmt"
)

// Item represents a product item in a store
type Item struct {
	Name  string
	Price float64
}

// NewItem is a factory method in the Cart class to create Items
func NewItem(name string, price float64) *Item {
	return &Item{Name: name, Price: price}
}

// Cart is the class that creates Items and holds them in a collection
type Cart struct {
	items []*Item
}

// AddItemToCart adds a new item to the cart
func (c *Cart) AddItemToCart(name string, price float64) {
	item := NewItem(name, price)
	c.items = append(c.items, item)
}

// TotalPrice calculates the total price of items in the cart
func (c *Cart) TotalPrice() float64 {
	total := 0.0
	for _, item := range c.items {
		total += item.Price
	}
	return total
}

// PrintCartDetails prints the details of the cart
func (c *Cart) PrintCartDetails() {
	fmt.Println("Cart Details:")
	for _, item := range c.items {
		fmt.Printf("- %s: $%.2f\n", item.Name, item.Price)
	}
	fmt.Printf("Total: $%.2f\n", c.TotalPrice())
}

func main() {
	cart := &Cart{}
	cart.AddItemToCart("Laptop", 1500.00)
	cart.AddItemToCart("Smartphone", 699.99)

	// Print cart details
	cart.PrintCartDetails()
}

//OUTPUT
//
//Cart Details:
//- Laptop: $1500.00
//- Smartphone: $699.99
//Total: $2199.99
