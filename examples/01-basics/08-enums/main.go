package main

import (
	"fmt"
	"log"
)

/**
 * Enum Examples demonstrates how to implement enum-like constants in Go.
 * Key concepts covered:
 * - Basic enums using iota
 * - Custom string representations
 * - Bitmask/flag enums
 * - Type safety with custom types
 */

/**
 * Direction represents compass directions
 * Shows basic enum pattern using iota
 */
type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
)

// String provides custom string representation for Direction
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

/**
 * DayOfWeek demonstrates enum with explicit values
 */
type DayOfWeek int

const (
	Sunday    DayOfWeek = 1
	Monday    DayOfWeek = 2
	Tuesday   DayOfWeek = 3
	Wednesday DayOfWeek = 4
	Thursday  DayOfWeek = 5
	Friday    DayOfWeek = 6
	Saturday  DayOfWeek = 7
)

/**
 * FileMode demonstrates bitwise flag enum pattern
 */
type FileMode uint

const (
	Read    FileMode = 1 << iota // 1
	Write                        // 2
	Execute                      // 4
)

func (f FileMode) String() string {
	var permissions []string
	if f&Read != 0 {
		permissions = append(permissions, "read")
	}
	if f&Write != 0 {
		permissions = append(permissions, "write")
	}
	if f&Execute != 0 {
		permissions = append(permissions, "execute")
	}
	return fmt.Sprintf("%v", permissions)
}

func main() {
	log.Println("=== Enum Examples ===")

	/**
	 * 1. Basic enum usage
	 * Shows how to use simple enums
	 */
	log.Println("\n1. Basic enum usage")
	var dir Direction = North
	log.Printf("Direction: %v (value: %d)\n", dir, dir)

	/**
	 * 2. Type safety
	 * Demonstrates type safety of enum types
	 */
	log.Println("\n2. Type safety")
	// This would cause a compile error:
	// dir = 5 // Cannot use 5 (type untyped int) as type Direction

	/**
	 * 3. Days of week
	 * Shows enum with explicit values
	 */
	log.Println("\n3. Days of week")
	today := Monday
	log.Printf("Today is day %d\n", today)

	/**
	 * 4. Bitwise flags
	 * Demonstrates using enums as flags
	 */
	log.Println("\n4. Bitwise flags")
	permissions := Read | Write // Combine flags using bitwise OR
	log.Printf("Permissions: %v\n", permissions)

	// Check if flag is set
	if permissions&Execute == 0 {
		log.Println("Execute permission not set")
	}

	// Add execute permission
	permissions |= Execute
	log.Printf("Updated permissions: %v\n", permissions)

	/**
	 * 5. Enum iteration
	 * Shows how to iterate over enum values
	 */
	log.Println("\n5. Enum iteration")
	for d := North; d <= West; d++ {
		log.Printf("Direction: %v\n", d)
	}
}
