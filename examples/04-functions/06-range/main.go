package main

import "log"

/**
 * Range Examples demonstrates different ways to iterate over various data structures in Go.
 * Range is commonly used with:
 * - Arrays and Slices: Provides index and value
 * - Maps: Provides key and value
 * - Strings: Provides index and rune (Unicode code point)
 * - Channels: Provides values received from channel
 */

/**
 * Person represents a basic user structure
 * Used to demonstrate ranging over custom types
 */
type Person struct {
	Name string
	Age  int
}

/**
 * processItems demonstrates range over a slice with custom processing
 * @param items: slice of strings to process
 * @return: map of string to int containing item counts
 */
func processItems(items []string) map[string]int {
	counts := make(map[string]int)
	for _, item := range items {
		counts[item]++
	}
	return counts
}

func main() {
	log.Println("=== Range Examples ===")

	/**
	 * 1. Range over slice
	 * Demonstrates basic slice iteration providing both index and value
	 */
	log.Println("\n1. Range over slice")
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		log.Printf("Index: %d, Value: %d\n", i, num)
	}

	/**
	 * 2. Range over map
	 * Shows how to iterate over map entries with key-value pairs
	 */
	log.Println("\n2. Range over map")
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}
	for name, score := range scores {
		log.Printf("Name: %s, Score: %d\n", name, score)
	}

	/**
	 * 3. Range over string
	 * Demonstrates iteration over string as Unicode code points (runes)
	 */
	log.Println("\n3. Range over string")
	text := "Hello, 世界"
	for i, char := range text {
		log.Printf("Position: %d, Character: %c, Unicode: %U\n",
			i, char, char)
	}

	/**
	 * 4. Range with blank identifier
	 * Shows how to skip index or value when not needed
	 */
	log.Println("\n4. Range with blank identifier")
	for _, num := range numbers {
		log.Printf("Value only: %d\n", num)
	}

	/**
	 * 5. Range over custom types
	 * Demonstrates ranging over slice of custom structs
	 */
	log.Println("\n5. Range over custom types")
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Carol", 28},
	}
	for _, person := range people {
		log.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}

	/**
	 * 6. Range with map manipulation
	 * Shows how to safely modify map while ranging
	 */
	log.Println("\n6. Range with map manipulation")
	items := []string{"apple", "banana", "apple", "cherry"}
	counts := processItems(items)
	for item, count := range counts {
		log.Printf("Item: %s, Count: %d\n", item, count)
	}

	/**
	 * 7. Range over channel
	 * Demonstrates how to range over a channel until it's closed
	 */
	log.Println("\n7. Range over channel")
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for num := range ch {
		log.Printf("Received: %d\n", num)
	}
}
