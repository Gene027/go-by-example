package main

import (
	"log"
	"time"
)

/**
 * Switch is a control flow statement that allows you to match a value against a list of cases.
 * Key concepts covered:
 * - Basic switch syntax
 * - Multiple cases
 * - Expression switches
 * - Type switches
 * - Fallthrough behavior
 * - Case with multiple values
 */

func main() {
	log.Println("=== Switch Statement Examples ===")

	/**
	 * 1. Basic switch
	 * Simple value matching with single cases
	 */
	log.Println("\n1. Basic switch")
	day := "Wednesday"
	switch day {
	case "Monday":
		log.Println("Start of work week")
	case "Wednesday":
		log.Println("Mid-week")
	case "Friday":
		log.Println("TGIF!")
	default:
		log.Println("Another day")
	}

	/**
	 * 2. Switch with multiple cases
	 * Demonstrates grouping multiple values in a single case
	 */
	log.Println("\n2. Switch with multiple cases")
	num := 2
	switch num {
	case 1, 3, 5, 7, 9:
		log.Println("Odd number <= 9")
	case 2, 4, 6, 8:
		log.Println("Even number <= 8")
	default:
		log.Println("Number > 9 or <= 0")
	}

	// Switch with expression
	log.Println("\n3. Switch with expression")
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		log.Println("Good morning!")
	case hour < 17:
		log.Println("Good afternoon!")
	default:
		log.Println("Good evening!")
	}

	/**
	 * 4. Switch with fallthrough
	 * The fallthrough keyword is used to transfer control to the next case.
	 * It is used to execute the next case even if the current case matches.
	 */
	log.Println("\n4. Switch with fallthrough")
	switch num := 1; num {
	case 1:
		log.Println("One")
		fallthrough
	case 2:
		log.Println("Two")
		fallthrough
	case 3:
		log.Println("Three")
	}

	/**
	 * 5. Type switch
	 * The type switch is used to check the type of an interface.
	 */
	log.Println("\n5. Type switch")
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		log.Printf("Integer: %d\n", v)
	case string:
		log.Printf("String: %s\n", v)
	case bool:
		log.Printf("Boolean: %t\n", v)
	default:
		log.Printf("Unknown type: %T\n", v)
	}
}
