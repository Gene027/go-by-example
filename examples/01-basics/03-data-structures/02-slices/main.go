package main

import "log"

/**
 * Slices are dynamic arrays that can grow and shrink.
 * They are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
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
	 * The slice is created using the [] operator and the size is inferred from the number of elements in the array not like arrays which have a fixed size.
	 */
	log.Println("\n1. Creating slices")
	numbers := []int{1, 2, 3, 4, 5}
	log.Printf("Numbers: %v\n", numbers)

	/**
	 * 2. Make function
	 * The make function is used to create a slice with a specified length and capacity.
	 * The length is the number of elements in the slice.
	 * The capacity is the maximum number of elements that the slice can hold.
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
	scores = append(scores, 100) // can add multiple elements at once eg. scores = append(scores, 100, 200, 300)
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

	//Append One Slice to Another
	log.Println("\n9. Append one slice to another")
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice1 = append(slice1, slice2...) // here we are appending slice2 to slice1 by spreading the elements of slice2
	log.Printf("Slice1: %v\n", slice1)
}
