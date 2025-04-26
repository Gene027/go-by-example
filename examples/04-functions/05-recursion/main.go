package main

import "log"

/**
 * Recursion Examples demonstrates different types of recursive functions in Go.
 * Types of recursion covered:
 * - Basic recursion
 * - Tree recursion
 * - Tail recursion
 * - Mutual recursion
 * Common patterns and best practices are shown for each type.
 */

/**
 * factorial calculates n! using basic recursion
 * Base case: n <= 1 returns 1
 * Recursive case: n * factorial(n-1)
 * @param n: number to calculate factorial for
 * @return: factorial of n
 */
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

/**
 * fibonacci calculates the nth Fibonacci number using recursion
 * Shows tree recursion where each call spawns multiple recursive calls
 * @param n: position in Fibonacci sequence
 * @return: nth Fibonacci number
 */
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

/**
 * TreeNode represents a binary tree structure
 * Used to demonstrate recursive tree traversal
 */
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * TraverseInOrder performs an in-order traversal of the binary tree
 * Demonstrates recursive traversal of a tree structure
 * Order: left subtree -> current node -> right subtree
 */
func (t *TreeNode) TraverseInOrder() {
	if t == nil {
		return
	}
	t.Left.TraverseInOrder()
	log.Printf("%d ", t.Value)
	t.Right.TraverseInOrder()
}

/**
 * factorialTail demonstrates tail recursion optimization
 * Uses an accumulator parameter to maintain state between recursive calls
 * @param n: number to calculate factorial for
 * @param acc: accumulator that holds intermediate results
 * @return: factorial calculated using tail recursion
 */
func factorialTail(n int, acc int) int {
	if n <= 1 {
		return acc
	}
	return factorialTail(n-1, n*acc)
}

/**
 * isEven and isOdd demonstrate mutual recursion
 * Two functions that call each other to determine if a number is even/odd
 * @param n: number to check
 * @return: true if n is even, false otherwise
 */
func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return isOdd(n - 1)
}

/**
 * isOdd is part of mutual recursion with isEven
 * Determines if a number is odd by checking if n-1 is even
 * @param n: number to check
 * @return: true if n is odd, false otherwise
 */
func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

/**
 * main demonstrates various recursion patterns:
 * - Basic recursion with factorial
 * - Tree recursion with Fibonacci
 * - Binary tree traversal
 * - Tail recursion optimization
 * - Mutual recursion
 */
func main() {
	log.Println("=== Recursion Examples ===")

	// Factorial calculation
	log.Println("\n1. Basic recursion - Factorial")
	log.Printf("Factorial of 5: %d\n", factorial(5))

	// Fibonacci sequence
	log.Println("\n2. Fibonacci sequence")
	for i := 0; i < 10; i++ {
		log.Printf("Fibonacci(%d): %d\n", i, fibonacci(i))
	}

	// Binary tree traversal
	log.Println("\n3. Tree traversal")
	root := &TreeNode{
		Value: 1,
		Left:  &TreeNode{Value: 2},
		Right: &TreeNode{Value: 3},
	}
	log.Println("In-order traversal:")
	root.TraverseInOrder()

	// Tail recursion
	log.Println("\n4. Tail recursion")
	log.Printf("Factorial(5) with tail recursion: %d\n", factorialTail(5, 1))

	// Mutual recursion
	log.Println("\n5. Mutual recursion")
	log.Printf("Is 4 even? %v\n", isEven(4))
	log.Printf("Is 5 even? %v\n", isEven(5))
}
