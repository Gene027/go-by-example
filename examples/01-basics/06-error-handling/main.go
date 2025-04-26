package main

import (
	"errors"
	"fmt"
	"log"
)

/**
 * Error Handling Examples demonstrates common error handling patterns in Go.
 * Key concepts covered:
 * - Error interface
 * - Creating custom errors
 * - Multiple error handling patterns
 * - Error wrapping and unwrapping
 * - Panic and recover
 */

/**
 * CustomError demonstrates a custom error type
 * Shows how to implement the error interface
 */
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

/**
 * divide performs division and handles division by zero
 * @param a: numerator
 * @param b: denominator
 * @return: result and error if division by zero
 */
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

/**
 * validateAge checks if age is within valid range
 * @param age: age to validate
 * @return: custom error with specific error code
 */
func validateAge(age int) error {
	if age < 0 {
		return &CustomError{Code: 1, Message: "age cannot be negative"}
	}
	if age > 150 {
		return &CustomError{Code: 2, Message: "age is unreasonably high"}
	}
	return nil
}

/**
 * processWithPanic demonstrates panic and recover
 * Shows how to handle unexpected errors
 */
func processWithPanic() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n", r)
		}
	}()

	panic("something went wrong")
}

func main() {
	log.Println("=== Error Handling Examples ===")

	/**
	 * 1. Basic error handling
	 * Shows standard error checking pattern
	 */
	log.Println("\n1. Basic error handling")
	result, err := divide(10, 2)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Printf("Result: %.2f\n", result)
	}

	/**
	 * 2. Custom error types
	 * Demonstrates using custom error structures
	 */
	log.Println("\n2. Custom error types")
	if err := validateAge(200); err != nil {
		log.Printf("Validation error: %v\n", err)
	}

	/**
	 * 3. Multiple error handling
	 * Shows handling different types of errors
	 */
	log.Println("\n3. Multiple error handling")
	ages := []int{-5, 25, 200}
	for _, age := range ages {
		if err := validateAge(age); err != nil {
			switch e := err.(type) {
			case *CustomError:
				log.Printf("Custom error with code %d: %s\n", e.Code, e.Message)
			default:
				log.Printf("Unknown error: %v\n", err)
			}
		} else {
			log.Printf("Age %d is valid\n", age)
		}
	}

	/**
	 * 4. Error wrapping
	 * Demonstrates wrapping and unwrapping errors
	 */
	log.Println("\n4. Error wrapping")
	_, err = divide(1, 0)
	if err != nil {
		wrappedErr := fmt.Errorf("calculation error: %w", err)
		log.Printf("Wrapped error: %v\n", wrappedErr)
		if errors.Is(wrappedErr, err) {
			log.Println("Original error found in wrapped error")
		}
	}

	/**
	 * 5. Panic and recover
	 * Shows how to handle panic situations
	 */
	log.Println("\n5. Panic and recover")
	processWithPanic()
	log.Println("Continued after panic")
}
