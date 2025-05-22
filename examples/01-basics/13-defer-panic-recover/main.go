package main

import "log"

/**
 * Defer, Panic, and Recover in Go demonstrates error handling and cleanup patterns.
 *
 * Key concepts:
 * - Defer statement execution order (LIFO)
 * - Resource cleanup with defer
 * - Panic mechanism
 * - Recover from panics
 * - Common use cases and patterns
 *
 * Common use cases:
 * - File handling (closing files)
 * - Database connections
 * - Mutex unlocking
 * - HTTP response body closing
 * - Cleanup operations
 */

/**
 * readFile simulates file operations with defer
 * Shows typical resource cleanup pattern
 */
func readFile(filename string) error {
	log.Printf("Opening file: %s\n", filename)
	defer log.Printf("Closing file: %s\n", filename)

	// Simulate file operations
	log.Printf("Reading file: %s\n", filename)
	return nil
}

/**
 * multipleDefers demonstrates defer execution order
 * Shows how multiple defers are handled
 */
func multipleDefers() {
	defer log.Println("First defer")
	defer log.Println("Second defer")
	defer log.Println("Third defer")
	log.Println("Function body")
}

/**
 * panicAndRecover demonstrates panic recovery
 * Shows how to handle unexpected errors
 */
func panicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n", r)
		}
	}()

	log.Println("Before panic")
	panic("something went wrong")
	log.Println("After panic") // This line never executes
}

/**
 * deferWithArguments shows argument evaluation timing
 * Demonstrates when defer arguments are evaluated
 */
func deferWithArguments(x int) {
	defer log.Printf("Deferred value: %d\n", x) // x here is evaluated at the time the defer is registered not when the function return therefore x is what is passed in
	x = 2
	log.Printf("Current value: %d\n", x)
}

func main() {
	log.Println("=== Defer, Panic, and Recover Examples ===")

	/**
	 * 1. Basic defer usage
	 * Shows simple cleanup pattern similar to finally in other languages which is executed after the function returns
	 */
	log.Println("\n1. Basic defer usage")
	readFile("example.txt")

	/**
	 * 2. Multiple defers
	 * Demonstrates LIFO execution order
	 */
	log.Println("\n2. Multiple defers")
	multipleDefers()

	/**
	 * 3. Defer with arguments
	 * Shows argument evaluation timing
	 */
	log.Println("\n3. Defer with arguments")
	deferWithArguments(1)

	/**
	 * 4. Panic and recover
	 * Demonstrates error recovery
	 */
	log.Println("\n4. Panic and recover")
	panicAndRecover()

	/**
	 * 5. Defer in loops
	 * Shows proper defer usage in loops
	 */
	log.Println("\n5. Defer in loops")
	for i := 1; i <= 3; i++ {
		func() {
			defer log.Printf("Cleanup iteration %d\n", i)
			log.Printf("Processing iteration %d\n", i)
		}()
	}

	log.Println("Main: All done")
}
