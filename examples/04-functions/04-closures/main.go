package main

import "log"

/**
 * Closure Examples demonstrates the use of closures in Go.
 * A closure is a function value that references variables from outside its body.
 * Key concepts covered:
 * - Function factories
 * - State encapsulation
 * - Variable capture
 * - Mutable state in closures
 */

/**
 * makeCounter creates a counter function that maintains its own state
 * Each returned function has its own independent count variable
 * @return: function that returns incrementing integers
 */
func makeCounter() func() int {
	count := 0 // This variable is "captured" by the closure
	return func() int {
		count++
		return count
	}
}

/**
 * makeAdder creates a function that adds a fixed base value
 * Demonstrates closure capturing parameters
 * @param base: the fixed value to add
 * @return: function that adds base to its argument
 */
func makeAdder(base int) func(int) int {
	return func(x int) int {
		return base + x // 'base' is captured from the outer scope
	}
}

// Logger factory - demonstrates closure with multiple variables
func makeLogger(prefix string) func(string) {
	count := 0 // Private counter for this logger
	return func(msg string) {
		count++
		log.Printf("[%s] (%d) %s\n", prefix, count, msg)
	}
}

func main() {
	log.Println("=== Closure Examples ===")

	// Basic counter closure
	log.Println("\n1. Counter closure")
	counter1 := makeCounter()
	counter2 := makeCounter()
	log.Printf("Counter1: %d\n", counter1()) // 1
	log.Printf("Counter1: %d\n", counter1()) // 2
	log.Printf("Counter2: %d\n", counter2()) // 1 (separate state)

	// Adder closure
	log.Println("\n2. Adder closure")
	add5 := makeAdder(5)
	add10 := makeAdder(10)
	log.Printf("Adding 5 to 3: %d\n", add5(3))   // 8
	log.Printf("Adding 10 to 3: %d\n", add10(3)) // 13

	// Logger closure
	log.Println("\n3. Logger closure")
	debugLog := makeLogger("DEBUG")
	errorLog := makeLogger("ERROR")
	debugLog("First debug message")
	debugLog("Second debug message")
	errorLog("First error message")

	// Closure capturing loop variable
	log.Println("\n4. Loop variable capture")
	funcs := make([]func(), 3)
	for i := 0; i < 3; i++ {
		i := i // Create new variable in each iteration
		funcs[i] = func() {
			log.Printf("Captured value: %d\n", i)
		}
	}
	for _, f := range funcs {
		f()
	}

	// Closure with mutable state
	log.Println("\n5. Mutable state in closure")
	accumulator := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}()
	log.Printf("Accumulator: %d\n", accumulator(1))
	log.Printf("Accumulator: %d\n", accumulator(2))
	log.Printf("Accumulator: %d\n", accumulator(3))
}
