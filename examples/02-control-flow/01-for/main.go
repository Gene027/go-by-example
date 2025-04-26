package main

import "log"

/**
 * For Loop Examples demonstrates different ways to use for loops in Go.
 * Key concepts covered:
 * - Basic for loop syntax
 * - For as while loop
 * - Infinite loops with break
 * - Continue statement
 * - Nested loops with labels
 */

func main() {
	log.Println("=== For Loop Examples ===")

	/**
	 * 1. Basic for loop
	 * Traditional three-component loop: initialization, condition, post
	 */
	log.Println("\n1. Basic for loop")
	for i := 0; i < 3; i++ {
		log.Printf("Count: %d\n", i)
	}

	/**
	 * 2. For as while loop
	 * Demonstrates using for with only a condition
	 */
	log.Println("\n2. For as while loop")
	sum := 1
	for sum < 50 {
		sum *= 2
		log.Printf("Current sum: %d\n", sum)
	}

	/**
	 * 3. Infinite loop with break
	 * Shows how to create and control infinite loops
	 */
	log.Println("\n3. Infinite loop with break")
	count := 0
	for {
		count++
		if count > 3 {
			break
		}
		log.Printf("Infinite loop iteration: %d\n", count)
	}

	/**
	 * 4. For loop with continue
	 * Demonstrates skipping iterations
	 */
	log.Println("\n4. For loop with continue (printing odd numbers)")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		log.Printf("Odd number: %d\n", i)
	}

	/**
	 * 5. Nested loops with labels
	 * Shows how to use labels for breaking out of nested loops
	 */
	log.Println("\n5. Nested loops with labels")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i*j > 3 {
				log.Printf("Breaking outer loop at i=%d, j=%d\n", i, j)
				break outer
			}
			log.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
