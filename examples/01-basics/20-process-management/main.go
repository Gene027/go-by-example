package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

/**
 * Process Management and System Operations in Go demonstrates handling of
 * system processes, signals, and environment variables.
 *
 * Key concepts:
 * - Process spawning and management
 * - Signal handling
 * - Environment variables
 * - Command execution
 * - Process cleanup
 *
 * Common use cases:
 * - System automation
 * - Service management
 * - Graceful shutdown
 * - Configuration
 * - Process supervision
 */

func commandLineExample() {
	// Get command-line arguments
	log.Printf("Program name: %s\n", os.Args[0])
	log.Printf("Arguments: %v\n", os.Args[1:])

	// Get environment variables
	log.Printf("HOME: %s\n", os.Getenv("HOME"))
	log.Printf("PATH: %s\n", os.Getenv("PATH"))

	// Set environment variable
	os.Setenv("MY_VAR", "my_value")
	log.Printf("MY_VAR: %s\n", os.Getenv("MY_VAR"))
}

func processSpawningExample() {
	// Execute command and wait for completion
	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Command output:\n%s\n", output)

	// Start long-running process
	longCmd := exec.Command("sleep", "5")
	if err := longCmd.Start(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Process started with PID: %d\n", longCmd.Process.Pid)

	// Wait for completion in goroutine
	go func() {
		if err := longCmd.Wait(); err != nil {
			log.Printf("Process error: %v\n", err)
		}
		log.Println("Long process completed")
	}()
}

func signalHandlingExample() {
	// Create channel for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create channel for cleanup completion
	done := make(chan bool)

	// Handle signals in goroutine
	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %v\n", sig)

		// Perform cleanup
		cleanup()

		done <- true
	}()

	log.Println("Process running. Press Ctrl+C to exit...")
	<-done
}

func cleanup() {
	log.Println("Performing cleanup...")
	// Simulate cleanup tasks
	time.Sleep(500 * time.Millisecond)
	log.Println("Cleanup completed")
}

func execExample() {
	// Replace current process with new one
	// Note: This will only execute in demonstration
	log.Println("This would replace the current process:")
	log.Printf("exec.Command(\"ls\", \"-l\").Run()")
}

func main() {
	log.Println("=== Process Management Examples ===")

	log.Println("\n1. Command Line and Environment")
	commandLineExample()

	log.Println("\n2. Process Spawning")
	processSpawningExample()

	log.Println("\n3. Signal Handling")
	signalHandlingExample()

	log.Println("\n4. Exec Operations")
	execExample()

	log.Println("Main: All done")
}
