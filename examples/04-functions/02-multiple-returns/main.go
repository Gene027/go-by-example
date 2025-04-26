package main

import (
	"errors"
	"log"
)

/**
 * Multiple Return Values Examples
 * Demonstrates different patterns for functions that return multiple values:
 * - Error handling pattern
 * - Named return values
 * - Multiple values of different types
 * - Function returns with error handling
 */

/**
 * divide demonstrates the common error handling pattern in Go
 * Returns both the result and an error if the operation fails
 * @param x: dividend
 * @param y: divisor
 * @return: quotient and error if division by zero
 */
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

/**
 * getCoordinates shows named return values
 * Variables x and y are declared in the return signature
 * @return x: x coordinate
 * @return y: y coordinate
 */
func getCoordinates() (x, y int) {
	x = 10
	y = 20
	return // naked return
}

/**
 * getUserInfo demonstrates returning multiple values of different types
 * Common pattern for returning entity data
 * @return: name (string), age (int), active status (bool)
 */
func getUserInfo() (string, int, bool) {
	return "Alice", 30, true
}

// Function returning function and error
func getOperation(op string) (func(int, int) int, error) {
	switch op {
	case "add":
		return func(x, y int) int { return x + y }, nil
	case "multiply":
		return func(x, y int) int { return x * y }, nil
	default:
		return nil, errors.New("unknown operation")
	}
}

func main() {
	log.Println("=== Multiple Return Values Examples ===")

	// Basic multiple returns
	log.Println("\n1. Basic multiple returns")
	result, err := divide(10, 2)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Printf("Division result: %.2f\n", result)
	}

	// Error handling
	log.Println("\n2. Error handling")
	result, err = divide(10, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Named return values
	log.Println("\n3. Named return values")
	x, y := getCoordinates()
	log.Printf("Coordinates: (%d, %d)\n", x, y)

	// Multiple values with different types
	log.Println("\n4. Different types")
	name, age, active := getUserInfo()
	log.Printf("User: %s, Age: %d, Active: %t\n", name, age, active)

	// Ignoring return values
	log.Println("\n5. Ignoring values")
	name, _, _ = getUserInfo() // Ignore age and active status
	log.Printf("Only name: %s\n", name)

	// Function and error return
	log.Println("\n6. Function return")
	operation, err := getOperation("add")
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		result := operation(5, 3)
		log.Printf("Operation result: %d\n", result)
	}
}
