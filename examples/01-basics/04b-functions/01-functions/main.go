package main

import "log"

/**
 * A function is a block of code that can be called with a name.
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
 * Named return values can improve readability of the function and allow naked returns
 * In this case, result was declared as the return value but not initialized, then it was assigned the value of x/y and will be returned automatically
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
 * In this case, the function operate takes op as a parameter which is a function that takes two int parameters and returns an int, it then returns the result of applying op to x and y
 * @param x: first operand
 * @param y: second operand
 * @param op: function that defines the operation
 * @return: result of applying op to x and y
 */
func operate(x, y int, op func(int, int) int) int {
	return op(x, y)
}

/**
 * Function closure is a function that returns a function
 * In this case, the function counter returns a function that returns an int
 * The function counter has a count variable that is incremented each time the returned function is called
 * @return: a function that returns an int
 */
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

/**
 * Recursion is a function that calls itself
 * @param n: the number to calculate the factorial of
 * @return: the factorial of n
 */
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
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

	// Recursion function
	log.Println("\n8. Recursion")
	log.Printf("Factorial of 5: %d\n", factorial(5))
}
