package main

import "log"

/**
 * Generics in Go (introduced in Go 1.18) provide a way to write functions and data
 * structures that can work with multiple types while maintaining type safety.
 *
 * Key benefits of generics:
 * - Code reusability: Write functions that work with different types
 * - Type safety: Catch type errors at compile time
 * - Reduced code duplication: Avoid writing similar code for different types
 *
 * Common use cases:
 * - Generic data structures (maps, slices, trees)
 * - Generic algorithms (sorting, searching)
 * - Type-safe operations across multiple types
 *
 * Important concepts:
 * - Type parameters: Placeholders for concrete types
 * - Constraints: Interface types that specify what operations are allowed
 * - Type inference: Compiler's ability to deduce type arguments
 * - Type sets: Groups of types that satisfy certain constraints
 */

/**
 * Stack demonstrates a generic data structure
 * Type parameter T can be any type
 */
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

/**
 * Number constraint allows any numeric type
 */
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

/**
 * Sum demonstrates generic function with constraints
 * Works with any numeric type
 */
func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

/**
 * Pair demonstrates generic struct with multiple type parameters
 */
type Pair[T, U any] struct {
	First  T
	Second U
}

func main() {
	log.Println("=== Generics Examples ===")

	/**
	 * 1. Generic data structure
	 * Shows how to use generic Stack
	 */
	log.Println("\n1. Generic data structure")
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	if val, ok := intStack.Pop(); ok {
		log.Printf("Popped: %d\n", val)
	}

	strStack := NewStack[string]()
	strStack.Push("hello")
	strStack.Push("world")
	if val, ok := strStack.Pop(); ok {
		log.Printf("Popped: %s\n", val)
	}

	/**
	 * 2. Generic function with constraints
	 * Demonstrates type-constrained operations
	 */
	log.Println("\n2. Generic function with constraints")
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3}
	log.Printf("Sum of ints: %d\n", Sum(ints))
	log.Printf("Sum of floats: %.2f\n", Sum(floats))

	/**
	 * 3. Multiple type parameters
	 * Shows struct with different types
	 */
	log.Println("\n3. Multiple type parameters")
	pair := Pair[string, int]{
		First:  "Age",
		Second: 25,
	}
	log.Printf("Pair: %s = %d\n", pair.First, pair.Second)

	/**
	 * 4. Type inference
	 * Demonstrates compiler's type deduction
	 */
	log.Println("\n4. Type inference")
	// Type parameters can often be inferred
	numbers := []int{1, 2, 3}
	sum := Sum(numbers) // Type parameter [int] is inferred
	log.Printf("Inferred sum: %d\n", sum)
}
