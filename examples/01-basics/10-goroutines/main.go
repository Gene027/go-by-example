package main

import (
	"log"
	"time"
)

/**
 * Goroutines in Go are lightweight threads of execution that enable concurrent programming.
 * They are managed by the Go runtime rather than the operating system, making them very efficient.
 *
 * Key concepts:
 * - Concurrent execution: Multiple goroutines run independently
 * - Lightweight: Very small memory footprint (a few KB)
 * - Scalability: Can run thousands of goroutines simultaneously
 * - Communication: Use channels to communicate between goroutines
 * - Scheduling: Go runtime handles scheduling automatically
 *
 * Common use cases:
 * - Parallel processing
 * - Background tasks
 * - Server request handling
 * - Asynchronous operations
 * - Event handling
 */

/**
 * worker simulates a task that takes some time to complete
 * @param id: identifier for the worker
 * @param duration: how long the task takes
 */
func worker(id int, duration time.Duration) {
	log.Printf("Worker %d: Starting\n", id)
	time.Sleep(duration)
	log.Printf("Worker %d: Done\n", id)
}

/**
 * counter demonstrates a simple concurrent counter
 * @param name: identifier for the counter
 * @param n: number of iterations
 */
func counter(name string, n int) {
	for i := 1; i <= n; i++ {
		log.Printf("Counter %s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	log.Println("=== Goroutines Examples ===")

	/**
	 * 1. Basic goroutine usage
	 * Shows how to start concurrent execution
	 */
	log.Println("\n1. Basic goroutine usage")
	// Start two goroutines which will run concurrently more like async
	go worker(1, 2*time.Second)
	go worker(2, 1*time.Second)
	time.Sleep(3 * time.Second) // Wait for workers to finish

	/**
	 * 2. Multiple goroutines
	 * func(id int) { ... }: This is an anonymous function that takes an integer parameter id.
	 * (i): This is the value of i that will be passed to the anonymous function.
	 * go func(id int) { ... }(i): This is the actual goroutine invocation, function call with i as the argument.
	 * The anonymous function is executed in a new goroutine, and the value of i is captured at the time the goroutine is created.
	 */
	log.Println("\n2. Multiple goroutines")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			log.Printf("Goroutine %d executing\n", id)
		}(i)
	}
	time.Sleep(time.Second)

	/**
	 * 3. Concurrent counters
	 * Shows independent execution of goroutines
	 */
	log.Println("\n3. Concurrent counters")
	go counter("A", 3)
	go counter("B", 3)
	time.Sleep(time.Second)

	/**
	 * 4. Anonymous goroutine
	 * Demonstrates using anonymous functions as goroutines
	 */
	log.Println("\n4. Anonymous goroutine")
	go func() {
		log.Println("Executing anonymous goroutine")
	}()
	time.Sleep(time.Second)

	/**
	 * 5. Goroutine with closure
	 * Shows how goroutines can access variables from their closure
	 */
	log.Println("\n5. Goroutine with closure")
	message := "Hello from closure"
	go func() {
		log.Println(message)
	}()
	time.Sleep(time.Second)

	log.Println("Main: All done")
}
