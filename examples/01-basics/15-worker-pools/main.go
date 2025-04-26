package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * Worker Pools, WaitGroups, and Rate Limiting in Go demonstrate concurrent task processing patterns.
 *
 * Key concepts:
 * - Worker Pool: Group of goroutines processing tasks concurrently
 * - WaitGroup: Synchronization for multiple goroutines
 * - Rate Limiting: Controlling the speed of operations
 * - Task Distribution: Fan-out pattern
 * - Result Collection: Fan-in pattern
 *
 * Common use cases:
 * - Parallel processing
 * - Batch operations
 * - API rate limiting
 * - Resource management
 * - Load balancing
 */

/**
 * Task represents a unit of work
 */
type Task struct {
	ID     int
	Result int
}

/**
 * WorkerStats tracks worker pool statistics
 */
type WorkerStats struct {
	tasksProcessed uint64 // Atomic counter for processed tasks
	totalTime      int64  // Atomic counter for total processing time (nanoseconds)
}

/**
 * worker processes tasks from the task queue
 * @param id: worker identifier
 * @param tasks: channel for receiving tasks
 * @param results: channel for sending results
 * @param stats: pointer to WorkerStats for tracking statistics
 * @param wg: WaitGroup for synchronization
 */
func worker(id int, tasks <-chan Task, results chan<- Task, stats *WorkerStats, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		start := time.Now()

		// Simulate processing time
		time.Sleep(100 * time.Millisecond)
		task.Result = task.ID * 2 // Simple computation

		// Update atomic counters
		atomic.AddUint64(&stats.tasksProcessed, 1)
		atomic.AddInt64(&stats.totalTime, time.Since(start).Nanoseconds())

		log.Printf("Worker %d processed task %d (Total: %d)\n",
			id,
			task.ID,
			atomic.LoadUint64(&stats.tasksProcessed))

		results <- task
	}
}

/**
 * rateLimitedWorker demonstrates rate limiting with time.Ticker
 * @param id: worker identifier
 * @param tasks: channel for receiving tasks
 * @param results: channel for sending results
 * @param rate: maximum operations per second
 * @param stats: pointer to WorkerStats for tracking statistics
 * @param wg: WaitGroup for synchronization
 */
func rateLimitedWorker(id int, tasks <-chan Task, results chan<- Task, rate int, stats *WorkerStats, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create rate limiter
	limiter := time.NewTicker(time.Second / time.Duration(rate))
	defer limiter.Stop()

	for task := range tasks {
		<-limiter.C // Wait for rate limit
		start := time.Now()

		// Process task
		time.Sleep(50 * time.Millisecond)
		task.Result = task.ID * 2

		// Update atomic counters
		atomic.AddUint64(&stats.tasksProcessed, 1)
		atomic.AddInt64(&stats.totalTime, time.Since(start).Nanoseconds())

		log.Printf("Rate-limited worker %d processed task %d (Total: %d)\n",
			id,
			task.ID,
			atomic.LoadUint64(&stats.tasksProcessed))

		results <- task
	}
}

func main() {
	log.Println("=== Worker Pools, WaitGroups, and Rate Limiting Examples ===")

	/**
	 * 1. Basic Worker Pool with Atomic Counters
	 */
	log.Println("\n1. Basic Worker Pool with Atomic Counters")
	numWorkers := 3
	numTasks := 10

	tasks := make(chan Task, numTasks)
	results := make(chan Task, numTasks)
	var wg sync.WaitGroup
	stats := &WorkerStats{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, stats, &wg)
	}

	// Send tasks
	go func() {
		for i := 1; i <= numTasks; i++ {
			tasks <- Task{ID: i}
		}
		close(tasks)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results and print statistics
	for result := range results {
		log.Printf("Got result for task %d: %d\n", result.ID, result.Result)
	}

	avgTime := float64(atomic.LoadInt64(&stats.totalTime)) / float64(atomic.LoadUint64(&stats.tasksProcessed))
	log.Printf("Statistics - Tasks Processed: %d, Average Time: %.2f ms\n",
		atomic.LoadUint64(&stats.tasksProcessed),
		avgTime/1e6)

	/**
	 * 2. Rate Limited Worker Pool
	 * Demonstrates controlled processing speed
	 */
	log.Println("\n2. Rate Limited Worker Pool")
	rateLimitedTasks := make(chan Task, numTasks)
	rateLimitedResults := make(chan Task, numTasks)
	var rateLimitedWg sync.WaitGroup

	// Start rate-limited workers
	operationsPerSecond := 2
	for i := 1; i <= numWorkers; i++ {
		rateLimitedWg.Add(1)
		go rateLimitedWorker(i, rateLimitedTasks, rateLimitedResults, operationsPerSecond, stats, &rateLimitedWg)
	}

	// Send tasks
	go func() {
		for i := 1; i <= numTasks; i++ {
			rateLimitedTasks <- Task{ID: i}
		}
		close(rateLimitedTasks)
	}()

	// Collect results
	go func() {
		rateLimitedWg.Wait()
		close(rateLimitedResults)
	}()

	// Process results
	for result := range rateLimitedResults {
		log.Printf("Got rate-limited result for task %d: %d\n", result.ID, result.Result)
	}

	/**
	 * 3. Dynamic Worker Pool
	 * Shows how to adjust pool size based on load
	 */
	log.Println("\n3. Dynamic Worker Pool")
	dynamicTasks := make(chan Task, numTasks)
	dynamicResults := make(chan Task, numTasks)
	var dynamicWg sync.WaitGroup

	// Start initial workers
	initialWorkers := 2
	for i := 1; i <= initialWorkers; i++ {
		dynamicWg.Add(1)
		go worker(i, dynamicTasks, dynamicResults, stats, &dynamicWg)
	}

	// Add more workers based on load
	go func() {
		if len(dynamicTasks) > 5 { // High load threshold
			dynamicWg.Add(1)
			go worker(initialWorkers+1, dynamicTasks, dynamicResults, stats, &dynamicWg)
		}
	}()

	// Send and process tasks
	for i := 1; i <= 5; i++ {
		dynamicTasks <- Task{ID: i}
	}
	close(dynamicTasks)
	dynamicWg.Wait()
	close(dynamicResults)

	for result := range dynamicResults {
		log.Printf("Got dynamic result for task %d: %d\n", result.ID, result.Result)
	}

	log.Println("Main: All done")
}
