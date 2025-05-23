package main

import "log"

/**
 * The range keyword provides a clean way to iterate over built-in types and custom collections.
 *
 * Key concepts:
 * - Range syntax and usage
 * - Different iteration patterns
 * - Value vs pointer semantics
 * - Performance considerations
 * - Common pitfalls
 *
 * Common use cases:
 * - Iterating over arrays and slices
 * - Processing map key-value pairs
 * - Character-by-character string processing
 * - Channel iteration
 * - Custom iterator implementations
 */

/**
 * User represents a basic user structure
 * Used to demonstrate ranging over structs
 */
type User struct {
	ID   int
	Name string
}

func main() {
	log.Println("=== Range Iterator Examples ===")

	/**
	 * 1. Range over slice
	 * Shows basic slice iteration with index and value
	 * Range returns the index and value of the slice
	 */
	log.Println("\n1. Range over slice")
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		log.Printf("Index: %d, Value: %d\n", i, num)
	}

	/**
	 * 2. Range over map
	 * Range returns the key and value of the map
	 */
	log.Println("\n2. Range over map")
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	for key, value := range colors {
		log.Printf("Key: %s, Value: %s\n", key, value)
	}

	/**
	 * 3. Range over string
	 * Range returns the position, character and unicode of the string
	 */
	log.Println("\n3. Range over string")
	text := "Hello, 世界"
	for i, char := range text {
		log.Printf("Position: %d, Character: %c, Unicode: %U\n", i, char, char)
	}

	/**
	 * 4. Range over channel
	 * Range returns the value of the channel until it is closed
	 */
	log.Println("\n4. Range over channel")
	ch := make(chan int, 3)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	for value := range ch {
		log.Printf("Received: %d\n", value)
	}

	/**
	 * 5. Range with blank identifier
	 * Range returns the value of the slice
	 */
	log.Println("\n5. Range with blank identifier")
	for _, num := range numbers {
		log.Printf("Value: %d\n", num)
	}

	/**
	 * 6. Range with struct slice
	 * Range returns the index and value of the struct slice
	 */
	log.Println("\n6. Range with struct slice")
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}
	for _, user := range users {
		log.Printf("User: %d - %s\n", user.ID, user.Name)
	}

	/**
	 * 7. Range over array
	 * Range returns the index and value of the array
	 */
	log.Println("\n7. Range over array")
	array := [3]string{"one", "two", "three"}
	for i, value := range array {
		log.Printf("Index: %d, Value: %s\n", i, value)
	}
}
