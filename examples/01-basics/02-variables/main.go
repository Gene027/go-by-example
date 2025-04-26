package main

import (
	"log"
)

/**
 * Variables Examples demonstrates different ways to declare and use variables in Go.
 * Key concepts covered:
 * - Variable declaration syntax
 * - Short variable declaration
 * - Multiple variable declaration
 * - Zero values
 * - Variable scope
 * - Type conversion
 */

// Package level variables
var globalVar = "I'm a global variable"

// Multiple variable declaration
var (
	name    string = "John"
	age     int    = 30
	isValid bool   = true
)

func main() {
	log.Println("=== Variables Examples ===")

	/**
	 * 1. Basic variable declaration
	 * Shows different ways to declare variables
	 */
	log.Println("\n1. Basic variable declaration")
	var name string = "Alice"
	var age = 25       // Type inferred
	job := "Developer" // Short declaration
	log.Printf("Name: %s, Age: %d, Job: %s\n", name, age, job)

	/**
	 * 2. Multiple variable declaration
	 * Demonstrates declaring multiple variables at once
	 */
	log.Println("\n2. Multiple variable declaration")
	var x, y int = 10, 20
	a, b := 30, "hello"
	log.Printf("x: %d, y: %d, a: %d, b: %s\n", x, y, a, b)

	/**
	 * 3. Zero values
	 * Shows default values for uninitialized variables
	 */
	log.Println("\n3. Zero values")
	var (
		intZero    int
		floatZero  float64
		boolZero   bool
		stringZero string
	)
	log.Printf("Zero values - Int: %d, Float: %f, Bool: %v, String: %q\n",
		intZero, floatZero, boolZero, stringZero)

	/**
	 * 4. Type conversion
	 * Demonstrates explicit type conversion between compatible types
	 */
	log.Println("\n4. Type conversion")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	log.Printf("Conversions - Int: %d, Float: %f, Uint: %d\n", i, f, u)

	// Short declaration syntax (only inside functions)
	shortVar := 42
	log.Printf("Short declaration: %v (type: %T)\n", shortVar, shortVar)

	// Regular declaration
	var regularVar int = 100
	log.Printf("Regular declaration: %v (type: %T)\n", regularVar, regularVar)

	// Multiple assignments
	log.Printf("\n=== Multiple Assignments ===\n")
	var c = true
	log.Printf("c: %v\n", c)

	// Variable shadowing (demonstrating scope)
	log.Printf("\n=== Variable Shadowing ===\n")
	// x := 5
	log.Printf("Outer x: %d\n", x)
	{
		x := 10 // This creates a new variable x in this scope
		log.Printf("Inner x: %d\n", x)
	}
	log.Printf("Outer x after block: %d\n", x)

	// Using global variables
	log.Printf("\n=== Global Variables ===\n")
	log.Printf("Global var: %s\n", globalVar)
	log.Printf("Name: %s, Age: %d, IsValid: %t\n", name, age, isValid)
}
