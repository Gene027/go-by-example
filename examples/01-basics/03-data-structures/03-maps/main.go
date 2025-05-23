package main

import "log"

/**
 * Maps are key-value pairs that are unordered and unique.
 * Key concepts covered:
 * - Map creation and initialization
 * - Key-value operations
 * - Map iteration
 * - Nested maps
 * - Zero values and nil maps
 */

func main() {
	log.Println("=== Map Examples ===")

	/**
	 * 1. Basic map creation
	 * Syntax: b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	 */
	log.Println("\n1. Basic map creation")
	scores := map[string]int{
		"Alice": 95,
		"Bob":   89,
		"Carol": 92,
	}
	log.Printf("Scores: %v\n", scores)

	/**
	 * 2. Creating map with make
	 * Demonstrates dynamic map creation using make function
	 */
	log.Println("\n2. Creating map with make")
	ages := make(map[string]int) // ages != nil at this point
	ages["Tom"] = 25
	ages["Jerry"] = 30
	log.Printf("Ages: %v\n", ages)

	// Zero value of maps
	log.Println("\n9. Zero value of maps")
	var nilMap map[string]int
	log.Printf("Nil map: %v, Is nil? %v\n", nilMap, nilMap == nil)

	/**
	 * 3. Map operations
	 * Shows common map operations: access, modify, check existence
	 */
	log.Println("\n3. Accessing and modifying maps")
	log.Printf("Bob's score: %d\n", scores["Bob"])
	scores["Bob"] = 91
	log.Printf("Bob's new score: %d\n", scores["Bob"])

	/**
	 * 4. Key existence check
	 * Demonstrates the comma-ok idiom for checking key existence
	 */
	log.Println("\n4. Checking key existence")
	score, exists := scores["David"]
	if exists {
		log.Printf("David's score: %d\n", score)
	} else {
		log.Printf("David's score not found\n")
	}

	// Deleting entries
	log.Println("\n5. Deleting from maps")
	delete(scores, "Carol")
	log.Printf("After deleting Carol: %v\n", scores)

	// Map length
	log.Println("\n6. Map length")
	log.Printf("Number of scores: %d\n", len(scores))

	// Iterating over maps
	log.Println("\n7. Iterating over maps")
	for name, score := range scores {
		log.Printf("%s: %d\n", name, score)
	}

	// Maps of maps
	log.Println("\n8. Nested maps")
	users := map[string]map[string]string{
		"user1": {
			"name":  "John Doe",
			"email": "john@example.com",
		},
		"user2": {
			"name":  "Jane Doe",
			"email": "jane@example.com",
		},
	}
	log.Printf("Users: %v\n", users)

	// Map doesnt implement copy
	log.Println("\n9. Map doesnt implement copy")
	originalMap := map[string]int{"a": 1, "b": 2}
	copiedMap := originalMap
	copiedMap["a"] = 100 // This will change the original map
	log.Printf("Original map: %v\n", originalMap)
	log.Printf("Copied map: %v\n", copiedMap)

	// How to copy maps
	log.Println("\n10. How to copy maps")
	originalMap2 := map[string]int{"a": 1, "b": 2}
	copiedMap2 := make(map[string]int)
	for key, value := range originalMap2 {
		copiedMap2[key] = value
	}
}
