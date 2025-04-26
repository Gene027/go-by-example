package main

import "log"

/**
 * Interface Examples demonstrates the usage and power of interfaces in Go.
 * Key concepts covered:
 * - Interface declaration and implementation
 * - Implicit interface satisfaction
 * - Empty interface
 * - Type assertions and type switches
 * - Interface composition
 */

/**
 * Shape defines a common interface for geometric shapes
 * Any type that implements Area() and Perimeter() is a Shape
 */
type Shape interface {
	Area() float64
	Perimeter() float64
}

/**
 * Printer defines an interface for types that can be printed
 * Used to demonstrate interface composition
 */
type Printer interface {
	Print()
}

/**
 * Rectangle implements the Shape interface
 * Shows how structs can implement interfaces
 */
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Print() {
	log.Printf("Rectangle: %.2f x %.2f\n", r.Width, r.Height)
}

/**
 * Circle implements the Shape interface
 * Another example of interface implementation
 */
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func (c Circle) Print() {
	log.Printf("Circle: radius %.2f\n", c.Radius)
}

/**
 * PrintShape demonstrates using interfaces as function parameters
 * @param s: any type that implements the Shape interface
 */
func PrintShape(s Shape) {
	log.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	log.Println("=== Interface Examples ===")

	/**
	 * 1. Basic interface usage
	 * Shows how different types can implement the same interface
	 */
	log.Println("\n1. Basic interface usage")
	rect := Rectangle{Width: 5, Height: 3}
	circ := Circle{Radius: 2}
	PrintShape(rect)
	PrintShape(circ)

	/**
	 * 2. Interface slices
	 * Demonstrates storing different types in a slice through interfaces
	 */
	log.Println("\n2. Interface slices")
	shapes := []Shape{rect, circ}
	for _, shape := range shapes {
		PrintShape(shape)
	}

	/**
	 * 3. Empty interface
	 * Shows how empty interface can hold values of any type
	 */
	log.Println("\n3. Empty interface")
	var i interface{}
	i = 42
	log.Printf("Empty interface with int: %v\n", i)
	i = "hello"
	log.Printf("Empty interface with string: %v\n", i)

	/**
	 * 4. Type assertions
	 * Demonstrates safely checking and converting interface types
	 */
	log.Println("\n4. Type assertions")
	var shape Shape = Circle{Radius: 2}
	if circle, ok := shape.(Circle); ok {
		log.Printf("Shape is a circle with radius %.2f\n", circle.Radius)
	}

	/**
	 * 5. Type switches
	 * Shows how to handle multiple possible types
	 */
	log.Println("\n5. Type switches")
	shapes = []Shape{Rectangle{Width: 3, Height: 4}, Circle{Radius: 2}}
	for _, shape := range shapes {
		switch v := shape.(type) {
		case Rectangle:
			log.Printf("Rectangle with width %.2f\n", v.Width)
		case Circle:
			log.Printf("Circle with radius %.2f\n", v.Radius)
		default:
			log.Printf("Unknown shape type\n")
		}
	}

	/**
	 * 6. Interface composition
	 * Demonstrates combining multiple interfaces
	 */
	log.Println("\n6. Interface composition")
	printers := []Printer{rect, circ}
	for _, p := range printers {
		p.Print()
	}
}
