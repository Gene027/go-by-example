package main

import "log"

/**
 * Arrays are fixed-size collections of elements of the same type. Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value
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
	 * This example shows how to declare an array with an explicit size.
	 * The size is specified in the declaration.
	 */
	log.Println("\n1. Basic array declaration")
	var numbers [5]int // Zero-initialized array
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	log.Printf("Numbers: %v\n", numbers)

	/**
	 * 2. Array initialization
	 * This example shows how to declare an array with an explicit size.
	 * The size is specified in the declaration.
	 */
	log.Println("\n2. Array initialization")
	fruits := [3]string{"apple", "banana", "orange"}
	log.Printf("Fruits: %v\n", fruits)

	/**
	 * 3. Array with implicit size
	 * This example shows how to declare an array with an implicit size.
	 * The size is inferred from the number of elements in the array.
	 */
	log.Println("\n3. Array with implicit size")
	colors := [...]string{"red", "green", "blue"}
	log.Printf("Colors (length %d): %v\n", len(colors), colors) //len() is used to get the length of the array

	/**
	 * 4. Multi-dimensional arrays 2D
	 * This example shows how to declare a 2D array.
	 * The size is specified in the declaration.
	 */
	log.Println("\n4. Multi-dimensional arrays")
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	// This is a loop that iterates over the rows of the matrix
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
	 * The copy has a different memory address than the original.
	 */
	log.Println("\n6. Array copying")
	original := [3]int{1, 2, 3}
	copy := original
	copy[0] = 100
	log.Printf("Original: %v\n", original)
	log.Printf("Copy: %v\n", copy)
}
