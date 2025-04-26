package main

import "log"

/**
 * If/Else Examples demonstrates conditional execution patterns in Go.
 * Key concepts covered:
 * - Basic if-else syntax
 * - Initialization statements
 * - Multiple conditions (else-if)
 * - Nested conditionals
 * - Logical operators
 * - Error checking patterns
 */

/**
 * checkNumber determines if a number is positive, negative, or zero
 * @param n: number to check
 * @return: string describing the number
 */
func checkNumber(n int) string {
	if n < 0 {
		return "negative"
	} else if n == 0 {
		return "zero"
	} else {
		return "positive"
	}
}

func main() {
	log.Println("=== If/Else Examples ===")

	// Basic if-else
	log.Println("\n1. Basic if-else")
	age := 18
	if age >= 18 {
		log.Println("Adult")
	} else {
		log.Println("Minor")
	}

	// If with initialization statement
	log.Println("\n2. If with initialization")
	if score := 85; score >= 70 {
		log.Printf("Passed with score: %d\n", score)
	} else {
		log.Printf("Failed with score: %d\n", score)
	}
	// Note: score is not accessible here

	// Multiple else-if conditions
	log.Println("\n3. Multiple else-if conditions")
	numbers := []int{-1, 0, 1}
	for _, num := range numbers {
		result := checkNumber(num)
		log.Printf("Number %d is %s\n", num, result)
	}

	// Nested if statements
	log.Println("\n4. Nested if statements")
	num := 15
	if num > 0 {
		if num%2 == 0 {
			log.Printf("%d is positive and even\n", num)
		} else {
			log.Printf("%d is positive and odd\n", num)
		}
	}

	// Logical operators in conditions
	log.Println("\n5. Logical operators in conditions")
	temperature := 25
	humidity := 60
	if temperature > 20 && humidity < 70 {
		log.Println("Weather is comfortable")
	}
	if temperature < 0 || temperature > 35 {
		log.Println("Extreme temperature!")
	}

	// Error checking pattern (common in Go)
	log.Println("\n6. Error checking pattern")
	if _, err := someFunction(); err != nil {
		log.Printf("Error occurred: %v\n", err)
	}
}

// Helper function for error checking example
func someFunction() (int, error) {
	return 0, nil // Simplified example
}
