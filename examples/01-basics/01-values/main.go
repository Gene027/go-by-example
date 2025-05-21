package main

import (
	"log"
)

/**
 * Values are the basic building blocks of Go programs.
 * They are the smallest unit of data in Go.
 * They can be of different types, such as:
 * - Numeric types (integers, floats)
 * - Strings and characters
 * - Booleans
 * - Constants
 * - Type inference
 */

func main() {
	log.Println("=== Basic Values Examples ===")

	/**
	 * 1. Integer values
	 * Demonstrates different integer types and operations
	 */
	log.Println("\n1. Integer values")
	intValue := 42
	int64Value := int64(42)
	log.Printf("Integer: %d, Int64: %d\n", intValue, int64Value) // %d is the format specifier for an integer

	/**
	 * 2. Floating-point values
	 * Shows float32 and float64 types
	 */
	log.Println("\n2. Floating-point values")
	float32Value := float32(3.14)
	float64Value := 3.14159
	log.Printf("Float32: %f, Float64: %f\n", float32Value, float64Value) // %f is the format specifier for a float

	/**
	 * 3. String values
	 * Demonstrates string literals and operations
	 */
	log.Println("\n3. String values")
	str1 := "Hello"
	str2 := "World"
	log.Printf("Strings: %s %s\n", str1, str2)
	log.Printf("Concatenated: %s\n", str1+" "+str2) // %s is the format specifier for a string

	/**
	 * 4. Boolean values
	 * Shows boolean operations and comparisons
	 */
	log.Println("\n4. Boolean values")
	isTrue := true
	isFalse := false
	log.Printf("Boolean AND: %v\n", isTrue && isFalse) // %v is the format specifier for a boolean
	log.Printf("Boolean OR: %v\n", isTrue || isFalse)

	/**
	 * 5. Type inference
	 * Demonstrates Go's type inference capabilities
	 */
	log.Println("\n5. Type inference")
	inferredInt := 42                                                      // Type inferred as int
	inferredFloat := 3.14                                                  // Type inferred as float64
	log.Printf("Types - Int: %T, Float: %T\n", inferredInt, inferredFloat) // %T is the format specifier for a type

	log.Println("\n6. Printf")
	log.Printf("Integer: %d, Float: %f\n", inferredInt, inferredFloat)

}
