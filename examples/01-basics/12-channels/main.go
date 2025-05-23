package main

import (
	"log"
	"time"
)

/**
 * Channels in Go are typed conduits that enable safe communication between goroutines.
 * They provide a way to synchronize data transfer and coordinate concurrent operations.
 *
 * Key concepts:
 * - Channel creation and usage
 * - Buffered vs unbuffered channels
 * - Channel direction (send/receive)
 * - Channel synchronization
 * - Channel closing
 * - Select statements for multiple channel operations
 *
 * Common use cases:
 * - Data transfer between goroutines
 * - Synchronization and coordination
 * - Worker pools and pipelines
 * - Event handling
 * - Resource management
 * - Timeout handling with select
 */

/*
*
* producer generates data and sends it to a channel
* ch chan<- int: This parameter is a channel of type int that is used for sending data. The chan<- syntax indicates that this channel is only for sending data (write-only).
* count int: This is an integer parameter that specifies how many items the producer should generate and send to the channel.

defer close(ch): The defer keyword is used to ensure that the close(ch) function is called when the producer function completes. This closes the channel ch, signaling to any receiving goroutines that no more data will be sent on this channel.
*/
func producer(ch chan<- int, count int) {
	defer close(ch) // Close channel when done
	for i := 1; i <= count; i++ {
		ch <- i                            // ch <- i: Sends the current value of i to the channel ch.
		time.Sleep(100 * time.Millisecond) //Pauses the execution for 100 milliseconds before the next iteration. This simulates some delay in producing each item.
	}
}

/*
*

  - consumer receives data from a channel and processes it

  - @param ch: ch <-chan int: This parameter is a channel of type int that is used for receiving data. The <-chan syntax indicates that this channel is only for receiving data (read-only).

  - for value := range ch: This is a for loop that iterates over values received from the channel ch. The range keyword is used to receive values from the channel until it is closed. Each received value is assigned to the variable value.

  - @param id: consumer identifier
*/
func consumer(ch <-chan int, id int) {
	for value := range ch {
		log.Printf("Consumer %d received: %d\n", id, value)
	}
}

/**
 * ping demonstrates bidirectional channel usage
 * @param ping: channel for sending ping
 * @param pong: channel for receiving pong
 */
func ping(ping chan<- string, pong <-chan string) {
	for i := 0; i < 3; i++ {
		ping <- "ping"
		<-pong // Wait for pong
	}
}

/**
 * worker demonstrates using the select statement for handling multiple channel operations.
 * @param dataCh: channel for receiving data
 * @param quitCh: channel for receiving quit signal
 * select is used to handle multiple channels in a non-blocking way
 */
func worker(dataCh <-chan int, quitCh <-chan bool) {
	for {
		select {
		//The worker waits to receive an integer from the dataCh channel.
		case data := <-dataCh:
			log.Printf("Worker received: %d\n", data)

		//The worker waits for a signal on the quitCh channel.
		case <-quitCh:
			log.Println("Worker received quit signal")
			return
		case <-time.After(500 * time.Millisecond):
			log.Println("Worker timed out waiting for data")
			return
		}
	}
}

func main() {
	log.Println("=== Channel Examples ===")

	/**
	 * 1. Basic channel usage
	 * Shows simple send and receive operations
	 */
	log.Println("\n1. Basic channel usage")
	ch := make(chan int)
	go func() {
		ch <- 42 // Send value
	}()
	value := <-ch // Receive value
	log.Printf("Received: %d\n", value)

	/**
	 * 2. Buffered channel
	 * Demonstrates channel with capacity
	 */
	log.Println("\n2. Buffered channel")
	bufferedCh := make(chan int, 2)
	bufferedCh <- 1 // Won't block
	bufferedCh <- 2 // Won't block
	log.Printf("Buffered values: %d, %d\n", <-bufferedCh, <-bufferedCh)

	/**
	 * 3. Producer-Consumer pattern
	 * Shows channel usage in common pattern
	 */
	log.Println("\n3. Producer-Consumer pattern")
	dataCh := make(chan int, 3)
	go producer(dataCh, 5)
	go consumer(dataCh, 1)
	time.Sleep(time.Second)

	/**
	 * 4. Multiple consumers
	 * Demonstrates fan-out pattern
	 */
	log.Println("\n4. Multiple consumers")
	workCh := make(chan int, 5)
	go producer(workCh, 6)
	for i := 1; i <= 2; i++ {
		go consumer(workCh, i)
	}
	time.Sleep(time.Second)

	/**
	 * 5. Bidirectional communication
	 * Shows ping-pong pattern with channels
	 */
	log.Println("\n5. Bidirectional communication")
	pingCh := make(chan string)
	pongCh := make(chan string)
	go ping(pingCh, pongCh)
	for i := 0; i < 3; i++ {
		<-pingCh // Receive ping
		pongCh <- "pong"
	}

	/**
	 * 6. Select statement
	 * Demonstrates handling multiple channels
	 */
	log.Println("\n6. Select statement")
	dataCh = make(chan int)
	quitCh := make(chan bool)

	go worker(dataCh, quitCh)

	// Send some data
	dataCh <- 1
	dataCh <- 2

	// Send quit signal
	quitCh <- true
	time.Sleep(100 * time.Millisecond)

	/**
	 * 7. Select with default
	 * Shows non-blocking channel operations
	 */
	log.Println("\n7. Select with default")
	ch = make(chan int)
	select {
	case val := <-ch:
		log.Printf("Received: %d\n", val)
	default:
		log.Println("No value available")
	}

	/**
	 * 8. Select with timeout
	 * Demonstrates timeout pattern
	 */
	log.Println("\n8. Select with timeout")
	resultCh := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		resultCh <- "Done"
	}()

	select {
	case result := <-resultCh:
		log.Printf("Received result: %s\n", result)
	case <-time.After(1 * time.Second):
		log.Println("Operation timed out")
	}

	log.Println("Main: All done")
}
