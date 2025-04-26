package main

import "log"

/**
 * Pointer Examples demonstrates the usage and behavior of pointers in Go.
 * Key concepts covered:
 * - Pointer declaration and initialization
 * - Dereferencing pointers
 * - Pointer arithmetic (not supported in Go)
 * - Pointer to structs
 * - Nil pointers
 */

/**
 * Person represents a basic struct to demonstrate struct pointers
 * Used to show how pointers work with composite types
 */
type Person struct {
	Name string
	Age  int
}

/**
 * modifyValue demonstrates pointer usage to modify values
 * @param ptr: pointer to an integer
 * Shows how to modify a value through its pointer
 */
func modifyValue(ptr *int) {
	*ptr = 42
}

/**
 * updatePerson shows how to work with struct pointers
 * @param p: pointer to Person struct
 * Demonstrates updating struct fields through a pointer
 */
func updatePerson(p *Person) {
	p.Name = "Updated " + p.Name // No need to dereference for struct fields
	p.Age++
}

func main() {
	log.Println("=== Pointer Examples ===")

	/**
	 * 1. Basic pointer usage
	 * Shows declaration and dereferencing of pointers
	 */
	log.Println("\n1. Basic pointer usage")
	x := 10
	ptr := &x // & operator takes the address of x
	log.Printf("Value: %d, Pointer: %p\n", x, ptr)
	*ptr = 20 // * operator dereferences ptr
	log.Printf("Modified value: %d\n", x)

	/**
	 * 2. Zero value of pointers
	 * Demonstrates nil pointers and initialization
	 */
	log.Println("\n2. Zero value of pointers")
	var nilPtr *int
	log.Printf("Nil pointer: %v\n", nilPtr)
	log.Printf("Is nil? %v\n", nilPtr == nil)

	/**
	 * 3. Pointers to structs
	 * Shows how to work with struct pointers
	 */
	log.Println("\n3. Pointers to structs")
	person := &Person{
		Name: "Alice",
		Age:  25,
	}
	log.Printf("Before: %+v\n", *person)
	updatePerson(person)
	log.Printf("After: %+v\n", *person)

	/**
	 * 4. Function with pointer parameter
	 * Demonstrates passing pointers to functions
	 */
	log.Println("\n4. Function with pointer parameter")
	value := 10
	log.Printf("Before: %d\n", value)
	modifyValue(&value)
	log.Printf("After: %d\n", value)

	/**
	 * 5. New function
	 * Shows how to create pointers using new()
	 */
	log.Println("\n5. New function")
	ptr2 := new(int)
	*ptr2 = 100
	log.Printf("Value through new pointer: %d\n", *ptr2)
}
