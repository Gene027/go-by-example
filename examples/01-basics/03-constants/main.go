package main

import (
	"log"
)

/**
 * Constants are used to store values that cannot be changed.
 * They can be of different types, such as:
 * - Numeric types (integers, floats)
 * - Strings and characters
 * - Booleans
 * - Arrays and slices
 * - Maps and structs
 * - Functions and methods
 * - Pointers and interfaces

 * Constants are declared using the const keyword.
 * They can be typed or untyped.
 * They can be enumerated using the iota keyword.
 * They can be constant expressions.
 * They can be constant rules and limitations
 */

// Untyped constants
const (
	Pi    = 3.14159
	Hello = "Hello, World!"
)

/**
 * Define named constants for days of the week
 * Using iota for automatic enumeration. The first constant is 0, the second is 1, and so on.
 */
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// Typed constants
const (
	MaxConnections int     = 100
	Timeout        float64 = 30.0
)

/**
 * Define typed constants for size units
 * Here we are using a constant expression to calculate the value of the other constants.
 */
const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
)

func main() {
	log.Println("=== Constants Examples ===")

	/**
	 * 1. Basic constants
	 * Shows declaration and usage of simple constants
	 */
	log.Println("\n1. Basic constants")
	log.Printf("Pi: %f, Greeting: %s\n", Pi, Hello)

	/**
	 * 2. Typed constants
	 * Demonstrates constants with explicit types
	 */
	log.Println("\n2. Typed constants")
	log.Printf("Max Connections: %d\n", MaxConnections)
	log.Printf("Timeout: %.1f seconds\n", Timeout)

	/**
	 * 3. Enumerated constants
	 * Shows usage of iota-based enumeration
	 */
	log.Println("\n3. Enumerated constants")
	log.Printf("Days: Sunday=%d, Monday=%d, Saturday=%d\n",
		Sunday, Monday, Saturday)

	/**
	 * 4. Size constants
	 * Demonstrates constant expressions and calculations
	 */
	log.Println("\n4. Size constants")
	log.Printf("Sizes: KB=%d, MB=%d, GB=%d\n", KB, MB, GB)

	// Constant expressions
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)
	log.Printf("\n=== Memory Units ===\n")
	log.Printf("Kilobyte: %d bytes\n", KB)
	log.Printf("Megabyte: %d bytes\n", MB)
	log.Printf("Gigabyte: %d bytes\n", GB)

	// Demonstrating type inference with constants
	const untypedNum = 42
	var intNum int = untypedNum          // Works fine
	var float64Num = float64(untypedNum) // Works fine
	log.Printf("\n=== Type Inference ===\n")
	log.Printf("Untyped constant as int: %d\n", intNum)
	log.Printf("Untyped constant as float64: %f\n", float64Num)
}
