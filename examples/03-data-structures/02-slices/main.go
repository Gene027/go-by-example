package main

import "log"

/**
 * Slice Examples demonstrates the flexibility and power of Go slices.
 * Key concepts covered:
 * - Slice is a dynamic array that can grow and shrink
 * - Slice creation and initialization
 * - Slice operations (append, copy)
 * - Slice capacity and growth
 * - Slice manipulation (subslicing)
 * - Common slice patterns
 */

func main() {
	log.Println("=== Slice Examples ===")

	/**
	 * 1. Creating slices
	 * Shows different ways to create and initialize slices
	 */
	log.Println("\n1. Creating slices")
	numbers := []int{1, 2, 3, 4, 5}
	log.Printf("Numbers: %v\n", numbers)

	/**
	 * 2. Make function
	 * Demonstrates slice creation with specified length and capacity
	 */
	log.Println("\n2. Using make")
	scores := make([]int, 3, 5) // length 3, capacity 5
	log.Printf("Scores: %v, len: %d, cap: %d\n",
		scores, len(scores), cap(scores))

	/**
	 * 3. Append operation
	 * Shows how to grow slices dynamically
	 */
	log.Println("\n3. Appending to slices")
	scores = append(scores, 100)
	log.Printf("After append: %v, len: %d, cap: %d\n",
		scores, len(scores), cap(scores))

	// Slicing operations
	log.Println("\n4. Slice operations")
	log.Printf("First two elements: %v\n", numbers[:2])
	log.Printf("Last two elements: %v\n", numbers[len(numbers)-2:])
	log.Printf("Middle elements: %v\n", numbers[1:4])

	// Slice of slices
	log.Println("\n5. Slice of slices")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	log.Printf("Matrix:\n")
	for i, row := range matrix {
		log.Printf("Row %d: %v\n", i, row)
	}

	// Copy slices
	log.Println("\n6. Copying slices")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	log.Printf("Source: %v\n", src)
	log.Printf("Destination: %v\n", dst)
	log.Printf("Number of elements copied: %d\n", copied)

	// Slice capacity growth
	log.Println("\n7. Slice capacity growth")
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		log.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
	}

	// Nil slices
	log.Println("\n8. Nil slices")
	var nilSlice []int
	log.Printf("Nil slice: %v, Length: %d, Is nil? %v\n",
		nilSlice, len(nilSlice), nilSlice == nil)
}
