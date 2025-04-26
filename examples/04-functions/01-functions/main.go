package main

import "log"

/**
 * Basic function examples demonstrating different ways to define and use functions in Go.
 * This file covers:
 * - Basic function declaration and calls
 * - Multiple parameters
 * - Named return values
 * - Functions as values
 * - Anonymous functions
 * - Closures
 * - Functions as struct fields
 */

/**
 * greet creates a greeting message for the given name
 * @param name: the name to greet
 * @return: formatted greeting string
 */
func greet(name string) string {
	return "Hello, " + name
}

/**
 * add demonstrates a function with multiple parameters of the same type
 * @param x: first number to add
 * @param y: second number to add
 * @return: sum of x and y
 */
func add(x, y int) int {
	return x + y
}

/**
 * divide shows named return value usage
 * Named return values can improve readability and allow naked returns
 * @param x: dividend
 * @param y: divisor
 * @return result: quotient of x/y
 */
func divide(x, y float64) (result float64) {
	result = x / y
	return // naked return uses named return value
}

/**
 * operate demonstrates function as a parameter
 * Higher-order function that takes a function as an argument
 * @param x: first operand
 * @param y: second operand
 * @param op: function that defines the operation
 * @return: result of applying op to x and y
 */
func operate(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// Function closure
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	log.Println("=== Function Examples ===")

	// Basic function call
	log.Println("\n1. Basic function")
	message := greet("Alice")
	log.Printf("Greeting: %s\n", message)

	// Multiple parameter function
	log.Println("\n2. Multiple parameters")
	sum := add(5, 3)
	log.Printf("Sum: %d\n", sum)

	// Named return value
	log.Println("\n3. Named return value")
	result := divide(10.0, 2.0)
	log.Printf("Division result: %.2f\n", result)

	// Function as value
	log.Println("\n4. Function as value")
	multiply := func(x, y int) int {
		return x * y
	}
	log.Printf("Multiplication: %d\n", operate(4, 5, multiply))

	// Anonymous function
	log.Println("\n5. Anonymous function")
	func(x int) {
		log.Printf("Anonymous function with value: %d\n", x)
	}(42)

	// Closure example
	log.Println("\n6. Closure")
	count := counter()
	for i := 0; i < 3; i++ {
		log.Printf("Count: %d\n", count())
	}

	// Function as field
	log.Println("\n7. Function as field")
	type Calculator struct {
		operation func(int, int) int
	}
	calc := Calculator{operation: add}
	log.Printf("Calculator result: %d\n", calc.operation(10, 5))
}
