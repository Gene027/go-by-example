package main

import "log"

/**
 * Interface is a type that defines a set of methods while struct is a type that defines a set of fields (data structure). structs can implement interfaces and can be used to group data of different types together.
 * In Go, methods are functions that are associated with a type.
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

/*
*

  - Area is a method that returns the area of a rectangle

  - The syntax for a method is func (receiver) methodName(parameters) returnType. The receiver is the type that the method is associated with. The method can be called on an instance of the type used to create the receiver eg r.Area() where r is an instance of Rectangle

  - Rectangle will implement Shape implicitly by implementing the Area() method and Perimeter() method declared in the Shape interface
*/
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

/*
* Rectangle implements the Printer interface by implementing the Print() method declared in the Printer interface
* Shows how structs can implement multiple interfaces
 */
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
	 * 3. Empty interface is an interface that has no methods. It is the most general interface type in Go. Any type implements the empty interface. Just like the any type in TypeScript.
	 * Shows how empty interface can hold values of any type
	 */
	log.Println("\n3. Empty interface")
	var i interface{}
	i = 42
	log.Printf("Empty interface with int: %v\n", i)
	i = "hello"
	log.Printf("Empty interface with string: %v\n", i)

	/**
	 * 4. Type assertions is used to check the type of an interface and convert it to a specific type.
	 * The syntax is variable.(Type) and it returns the value and a boolean that is true if the type assertion is successful.
	 */
	log.Println("\n4. Type assertions")
	var shape Shape = Circle{Radius: 2}

	if circle, ok := shape.(Circle); ok {
		log.Printf("Shape is a circle with radius %.2f\n", circle.Radius)
	}

	/**
	 * 5. Type switches is used to check the type of an interface and execute different code blocks based on the type.
	 * The syntax is switch v := variable.(type) {}.
	 * The switch statement is used to check the type of the variable and execute different code blocks based on the type.
	 * The type switch is useful when you have a variable that implements multiple interfaces and you want to handle each type differently.
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
	 * 6. Interface composition is the ability to combine multiple interfaces into a single interface.
	 * The syntax is type InterfaceName interface {}.
	 * The interface composition is useful when you want to create a new interface that is a combination of multiple interfaces.
	 */
	log.Println("\n6. Interface composition")
	printers := []Printer{rect, circ}
	for _, p := range printers {
		p.Print()
	}
}
