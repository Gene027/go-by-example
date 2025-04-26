package main

import "log"

/**
 * Array Examples demonstrates the basic usage of arrays in Go.
 * Key concepts covered:
 * - Array declaration and initialization
 * - Fixed-size nature of arrays
 * - Array operations and indexing
 * - Multi-dimensional arrays
 * - Array comparisons
 */

func main() {
	log.Println("=== Array Examples ===")

	/**
	 * 1. Basic array declaration
	 * Shows different ways to declare and initialize arrays
	 */
	log.Println("\n1. Basic array declaration")
	var numbers [5]int // Zero-initialized array
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	log.Printf("Numbers: %v\n", numbers)

	/**
	 * 2. Array initialization
	 * Demonstrates array literal syntax and partial initialization
	 */
	log.Println("\n2. Array initialization")
	fruits := [3]string{"apple", "banana", "orange"}
	log.Printf("Fruits: %v\n", fruits)

	/**
	 * 3. Array with implicit size
	 * Shows how Go can infer array size from initialization
	 */
	log.Println("\n3. Array with implicit size")
	colors := [...]string{"red", "green", "blue"}
	log.Printf("Colors (length %d): %v\n", len(colors), colors)

	/**
	 * 4. Multi-dimensional arrays
	 * Demonstrates arrays of arrays (matrices)
	 */
	log.Println("\n4. Multi-dimensional arrays")
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i, row := range matrix {
		log.Printf("Row %d: %v\n", i, row)
	}

	// Array bounds
	log.Println("\n5. Array bounds and length")
	log.Printf("Length of fruits: %d\n", len(fruits))

	/**
	 * 5. Array comparisons
	 * Shows how arrays can be compared directly
	 */
	log.Println("\n5. Array comparisons")
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}
	log.Printf("arr1 == arr2: %v\n", arr1 == arr2)
	log.Printf("arr1 == arr3: %v\n", arr1 == arr3)

	/**
	 * 6. Array copying
	 * Demonstrates how arrays are copied by value
	 */
	log.Println("\n6. Array copying")
	original := [3]int{1, 2, 3}
	copy := original
	copy[0] = 100
	log.Printf("Original: %v\n", original)
	log.Printf("Copy: %v\n", copy)
}
