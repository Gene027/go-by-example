package main

import (
	"fmt"
	"log"
)

/**
 * Variadic Functions Examples
 * Demonstrates different ways to use variadic parameters in Go:
 * - Basic variadic functions
 * - Mixed type parameters
 * - Forwarding variadic arguments
 * - Common use cases and patterns
 */

/**
 * sum is a basic variadic function that adds numbers
 * @param numbers: variable number of integers
 * @return: sum of all provided numbers
 */
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

/**
 * printValues demonstrates variadic function with interface{} type
 * Allows printing values of any type with a common prefix
 * @param prefix: string to prepend to each value
 * @param values: variable number of values of any type
 */
func printValues(prefix string, values ...interface{}) {
	for _, value := range values {
		log.Printf("%s: %v (type: %T)\n", prefix, value, value)
	}
}

// Variadic function with other parameters
func joinStrings(separator string, parts ...string) string {
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += separator
		}
		result += part
	}
	return result
}

// Forwarding variadic arguments
func wrapper(numbers ...int) int {
	log.Printf("Wrapping call to sum with numbers: %v\n", numbers)
	return sum(numbers...) // Note the ... to forward variadic args
}

func main() {
	log.Println("=== Variadic Functions Examples ===")

	// Basic usage
	log.Println("\n1. Basic variadic function")
	log.Printf("Sum: %d\n", sum(1, 2, 3, 4, 5))

	// Empty variadic call
	log.Println("\n2. Empty variadic call")
	log.Printf("Sum of nothing: %d\n", sum())

	// Slice as variadic argument
	log.Println("\n3. Slice as variadic argument")
	numbers := []int{1, 2, 3, 4, 5}
	log.Printf("Sum from slice: %d\n", sum(numbers...))

	// Mixed types
	log.Println("\n4. Mixed types")
	printValues("Item", 42, "hello", true, 3.14)

	// String joining
	log.Println("\n5. String joining")
	result := joinStrings(", ", "apple", "banana", "orange")
	log.Printf("Joined string: %s\n", result)

	// Forwarding variadic arguments
	log.Println("\n6. Forwarding arguments")
	log.Printf("Wrapped sum: %d\n", wrapper(1, 2, 3))

	// fmt.Printf as a variadic function example
	log.Println("\n7. fmt.Printf example")
	fmt.Printf("Formatting multiple values: %d %s %.2f\n", 42, "test", 3.14)
}
