package main

import (
	"fmt"
	"log"
)

/**
 * Struct and Struct Embedding Examples demonstrates structs and composition through embedding in Go.
 * Structs are a collection of fields that are used to represent a single entity.
 * Embedding is a way to inherit fields and methods from other structs.
 * Key concepts covered:
 * - Basic struct definition and usage
 * - Struct fields and methods
 * - Struct embedding for composition
 * - Method promotion
 * - Interface satisfaction through embedding
 * - Multiple embedding
 * - Embedding and field name conflicts
 */

/**
 * Address represents a basic location structure
 * Shows basic struct definition with fields
 */
type Address struct {
	Street  string // Each field has a name and type
	City    string
	Country string
}

// Method defined on Address struct
func (a Address) Format() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

/**
 * Person demonstrates basic struct definition and embedding
 * Shows both regular fields and embedded structs
 */
type Person struct {
	Name    string // Regular struct fields
	Age     int
	Address // Embedded struct - inherits all fields and methods
}

/**
 * Employee demonstrates struct composition through multiple embedding
 * Shows both embedded structs and regular struct fields
 */
type Employee struct {
	Person         // Embedded Person struct - inherits Name, Age, and Address
	CompanyName    string
	CompanyAddress Address // Regular field - normal struct composition
}

/**
 * Logger interface for demonstration
 * Shows how structs can implement interfaces
 */
type Logger interface {
	Log(message string)
}

/**
 * BaseLogger provides basic logging functionality
 * Shows a simple struct with a method
 */
type BaseLogger struct {
	prefix string
}

// Method that implements Logger interface
func (b BaseLogger) Log(message string) {
	log.Printf("%s: %s\n", b.prefix, message)
}

/**
 * Service demonstrates interface implementation through embedding
 * Shows how embedded structs can help satisfy interfaces
 */
type Service struct {
	BaseLogger // Embedded struct provides Log method
	name       string
}

func main() {
	log.Println("=== Struct and Struct Embedding Examples ===")

	/**
	 * 1. Basic struct usage and embedding
	 * Shows struct initialization and field access
	 */
	log.Println("\n1. Basic struct usage and embedding")
	person := Person{ // Initialize struct with field values
		Name: "Alice",
		Age:  30,
		Address: Address{ // Initialize embedded struct
			Street:  "123 Main St",
			City:    "Boston",
			Country: "USA",
		},
	}
	// Access fields and methods directly due to embedding
	log.Printf("Person: %s, Address: %s\n", person.Name, person.Format())

	/**
	 * 2. Multiple embedding and composition
	 * Shows complex struct relationships
	 */
	log.Println("\n2. Multiple embedding and composition")
	employee := Employee{
		Person: Person{ // Initialize embedded Person struct
			Name: "Bob",
			Age:  35,
			Address: Address{
				Street:  "456 Work St",
				City:    "Chicago",
				Country: "USA",
			},
		},
		CompanyName: "Tech Corp",
		CompanyAddress: Address{ // Initialize regular struct field
			Street:  "789 Corp Ave",
			City:    "New York",
			Country: "USA",
		},
	}
	log.Printf("Employee: %s, Home: %s, Work: %s\n",
		employee.Name, // Promoted field from Person
		employee.Address.Format(),
		employee.CompanyAddress.Format())

	/**
	 * 3. Interface satisfaction through embedding
	 * Shows how embedded types can provide interface implementation
	 */
	log.Println("\n3. Interface satisfaction through embedding")
	service := Service{
		BaseLogger: BaseLogger{prefix: "SERVICE"},
		name:       "MyService",
	}
	service.Log("Service started") // Method from embedded BaseLogger

	/**
	 * 4. Method promotion from embedded types
	 * Shows how embedded struct methods become available
	 */
	log.Println("\n4. Method promotion")
	log.Printf("Person address: %s\n", person.Format()) // Format method from embedded Address
}
