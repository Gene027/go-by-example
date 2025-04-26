package main

import (
	"log"
	"time"
)

/**
 * Timers, Tickers, and Timeouts in Go demonstrate time-based operations and scheduling.
 *
 * Key concepts:
 * - Timer: Fires once after a specified duration
 * - Ticker: Fires repeatedly at specified intervals
 * - Timeout patterns: Limiting operation duration
 * - Timer cancellation and reset
 * - Ticker cleanup
 *
 * Common use cases:
 * - Delayed operations
 * - Periodic tasks
 * - Operation timeouts
 * - Rate limiting
 * - Heartbeat signals
 */

/**
 * delayedOperation demonstrates basic timer usage
 * Shows how to execute code after a delay
 */
func delayedOperation() {
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	log.Println("Starting delayed operation")
	<-timer.C
	log.Println("Delayed operation executed")
}

/**
 * periodicTask demonstrates ticker usage
 * Shows how to perform recurring tasks
 * @param duration: interval between ticks
 * @param count: number of ticks to process
 */
func periodicTask(duration time.Duration, count int) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for i := 1; i <= count; i++ {
		<-ticker.C
		log.Printf("Tick %d: Executing periodic task\n", i)
	}
}

/**
 * timeoutOperation demonstrates timeout pattern
 * Shows how to limit operation duration
 * @param timeout: maximum duration to wait
 */
func timeoutOperation(timeout time.Duration) {
	done := make(chan bool)

	go func() {
		// Simulate long operation
		time.Sleep(2 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		log.Println("Operation completed successfully")
	case <-time.After(timeout):
		log.Println("Operation timed out")
	}
}

/**
 * timerReset demonstrates timer reset functionality
 * Shows how to reuse timers
 */
func timerReset() {
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	log.Println("Starting timer")
	timer.Reset(500 * time.Millisecond) // Reset to shorter duration
	<-timer.C
	log.Println("Timer fired after reset")
}

func main() {
	log.Println("=== Timers, Tickers, and Timeouts Examples ===")

	/**
	 * 1. Basic timer usage
	 * Shows one-time delayed execution
	 */
	log.Println("\n1. Basic timer usage")
	delayedOperation()

	/**
	 * 2. Periodic ticker
	 * Demonstrates recurring tasks
	 */
	log.Println("\n2. Periodic ticker")
	periodicTask(500*time.Millisecond, 3)

	/**
	 * 3. Operation timeout
	 * Shows how to limit operation duration
	 */
	log.Println("\n3. Operation timeout")
	timeoutOperation(1 * time.Second)

	/**
	 * 4. Timer reset
	 * Demonstrates timer reuse
	 */
	log.Println("\n4. Timer reset")
	timerReset()

	/**
	 * 5. Multiple tickers
	 * Shows handling multiple time sources
	 */
	log.Println("\n5. Multiple tickers")
	ticker1 := time.NewTicker(500 * time.Millisecond)
	ticker2 := time.NewTicker(800 * time.Millisecond)
	defer ticker1.Stop()
	defer ticker2.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	count := 0
	for {
		select {
		case <-ticker1.C:
			log.Println("Ticker 1 fired")
			count++
		case <-ticker2.C:
			log.Println("Ticker 2 fired")
			count++
		case <-done:
			log.Printf("Processed %d ticks\n", count)
			return
		}
	}
}
