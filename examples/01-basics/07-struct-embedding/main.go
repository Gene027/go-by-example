package main

import (
	"fmt"
	"log"
)

/**
 * Struct Embedding Examples demonstrates composition through embedding in Go.
 * Key concepts covered:
 * - Basic struct embedding
 * - Method promotion
 * - Interface satisfaction through embedding
 * - Multiple embedding
 * - Embedding and field name conflicts
 */

/**
 * Address represents a basic location structure
 * Used to demonstrate embedding in other structs
 */
type Address struct {
	Street  string
	City    string
	Country string
}

func (a Address) Format() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

/**
 * Person demonstrates basic struct embedding
 * Embeds Address to inherit its fields and methods
 */
type Person struct {
	Name    string
	Age     int
	Address // Embedded struct
}

/**
 * Employee demonstrates multiple embedding
 * Shows how to embed multiple structs and handle conflicts
 */
type Employee struct {
	Person         // Embedded Person struct
	CompanyName    string
	CompanyAddress Address // Regular field, not embedded
}

/**
 * Logger interface for demonstration
 * Shows interface satisfaction through embedding
 */
type Logger interface {
	Log(message string)
}

/**
 * BaseLogger provides basic logging functionality
 * Will be embedded in other types
 */
type BaseLogger struct {
	prefix string
}

func (b BaseLogger) Log(message string) {
	log.Printf("%s: %s\n", b.prefix, message)
}

/**
 * Service demonstrates interface satisfaction through embedding
 * Automatically satisfies Logger interface through BaseLogger
 */
type Service struct {
	BaseLogger
	name string
}

func main() {
	log.Println("=== Struct Embedding Examples ===")

	/**
	 * 1. Basic struct embedding
	 * Shows how embedded fields can be accessed directly
	 */
	log.Println("\n1. Basic struct embedding")
	person := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			Country: "USA",
		},
	}
	log.Printf("Person: %s, Address: %s\n", person.Name, person.Format())

	/**
	 * 2. Multiple embedding
	 * Demonstrates embedding multiple structs
	 */
	log.Println("\n2. Multiple embedding")
	employee := Employee{
		Person: Person{
			Name: "Bob",
			Age:  35,
			Address: Address{
				Street:  "456 Work St",
				City:    "Chicago",
				Country: "USA",
			},
		},
		CompanyName: "Tech Corp",
		CompanyAddress: Address{
			Street:  "789 Corp Ave",
			City:    "New York",
			Country: "USA",
		},
	}
	log.Printf("Employee: %s, Home: %s, Work: %s\n",
		employee.Name,
		employee.Address.Format(),
		employee.CompanyAddress.Format())

	/**
	 * 3. Interface satisfaction through embedding
	 * Shows how embedded types can satisfy interfaces
	 */
	log.Println("\n3. Interface satisfaction through embedding")
	service := Service{
		BaseLogger: BaseLogger{prefix: "SERVICE"},
		name:       "MyService",
	}
	service.Log("Service started")

	/**
	 * 4. Method promotion
	 * Demonstrates how methods are promoted from embedded types
	 */
	log.Println("\n4. Method promotion")
	log.Printf("Person address: %s\n", person.Format())
}
