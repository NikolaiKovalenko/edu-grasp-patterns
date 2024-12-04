//Explanation:
//Product Class: Represents a product with its name and price. It encapsulates data relevant to the product.
//DiscountCalculator Class: An information expert dedicated to calculating discounts. It has the necessary information (product price and discount percentage) to perform its calculations.
//Checkout Class: Manages a collection of products and calculates the total cost. It also has the logic to display the checkout summary, demonstrating that it has responsibility over products and their total cost.
//Client Code: In the main function, products are added to the checkout, and a summary is printed, showing how the responsibility for calculating the discount is appropriately assigned to the DiscountCalculator.

package main

import "fmt"

// Product represents a product in an online shop
type Product struct {
	Name  string
	Price float64
}

// DiscountCalculator is responsible for calculating discounts
type DiscountCalculator struct{}

// Calculate calculates the discount amount for a product based on a discount percentage
func (d *DiscountCalculator) Calculate(product Product, discountPercent float64) float64 {
	return product.Price * (discountPercent / 100)
}

// Checkout is responsible for processing the checkout
type Checkout struct {
	products []Product
	total    float64
}

// AddProduct adds a product to the checkout
func (c *Checkout) AddProduct(product Product) {
	c.products = append(c.products, product)
	c.total += product.Price
}

// GetTotal returns the total price of products in the checkout
func (c *Checkout) GetTotal() float64 {
	return c.total
}

// PrintSummary prints the checkout summary along with any applicable discounts
func (c *Checkout) PrintSummary(discountPercent float64) {
	fmt.Println("Checkout Summary:")
	for _, product := range c.products {
		fmt.Printf("- %s: $%.2f\n", product.Name, product.Price)
	}

	if discountPercent > 0 {
		discountCalculator := DiscountCalculator{}
		discountAmount := discountCalculator.Calculate(c.products[0], discountPercent) // Assuming discount applies to the first product for example
		finalPrice := c.total - discountAmount
		fmt.Printf("Discount: -$%.2f\n", discountAmount)
		fmt.Printf("Total after discount: $%.2f\n", finalPrice)
	} else {
		fmt.Printf("Total: $%.2f\n", c.total)
	}
}

func main() {
	// Create products
	product1 := Product{Name: "Laptop", Price: 1200.00}
	product2 := Product{Name: "Mouse", Price: 25.00}

	// Create checkout and add products
	checkout := &Checkout{}
	checkout.AddProduct(product1)
	checkout.AddProduct(product2)

	// Print checkout summary with a discount
	checkout.PrintSummary(10) // Applying a 10% discount
}

//OUTPUT
//
//Checkout Summary:
//- Laptop: $1200.00
//- Mouse: $25.00
//Discount: -$120.00
//Total after discount: $1105.00
