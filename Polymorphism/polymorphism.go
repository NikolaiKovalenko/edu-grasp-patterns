//Explanation:
//Drawer Interface: Defines a Draw method that must be implemented by any shape.
//Concrete Types (Circle, Rectangle, Triangle): Implement the Drawer interface, providing their own implementation of the Draw method.
//Client Code: Demonstrates polymorphism by iterating over a collection of Drawer interfaces and calling Draw, without needing to know the specific type of each shape.

package main

import "fmt"

// Drawer interface defines a method for drawing shapes
type Drawer interface {
	Draw()
}

// Circle represents a circle and implements the Drawer interface
type Circle struct {
	Radius float64
}

// Draw outputs the draw action for a circle
func (c *Circle) Draw() {
	fmt.Printf("Drawing a circle with radius %.2f\n", c.Radius)
}

// Rectangle represents a rectangle and implements the Drawer interface
type Rectangle struct {
	Width, Height float64
}

// Draw outputs the draw action for a rectangle
func (r *Rectangle) Draw() {
	fmt.Printf("Drawing a rectangle with width %.2f and height %.2f\n", r.Width, r.Height)
}

// Triangle represents a triangle and implements the Drawer interface
type Triangle struct {
	Base, Height float64
}

// Draw outputs the draw action for a triangle
func (t *Triangle) Draw() {
	fmt.Printf("Drawing a triangle with base %.2f and height %.2f\n", t.Base, t.Height)
}

func main() {
	// Create a slice of Drawer interface
	shapes := []Drawer{
		&Circle{Radius: 5},
		&Rectangle{Width: 3, Height: 4},
		&Triangle{Base: 6, Height: 4},
	}

	// Iterate over shapes and draw each one
	for _, shape := range shapes {
		shape.Draw()
	}
}

//OUTPUT
//
//Drawing a circle with radius 5.00
//Drawing a rectangle with width 3.00 and height 4.00
//Drawing a triangle with base 6.00 and height 4.00
