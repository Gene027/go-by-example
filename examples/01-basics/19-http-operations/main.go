package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/**
 * HTTP Operations and Context in Go demonstrates web service patterns.
 *
 * Key concepts:
 * - HTTP client operations
 * - HTTP server implementation
 * - Context for cancellation
 * - Middleware patterns
 * - Request/Response handling
 *
 * Common use cases:
 * - RESTful APIs
 * - Web services
 * - API integration
 * - Timeout handling
 * - Request tracing
 */

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// middleware demonstrates a basic middleware pattern
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: %s %s", r.Method, r.URL.Path)

		next(w, r)

		log.Printf("Duration: %v", time.Since(start))
	}
}

func httpServerExample() {
	// Basic handler
	http.HandleFunc("/", middleware(func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Status:  "success",
			Message: "Welcome to the API",
		}
		json.NewEncoder(w).Encode(response)
	}))

	// JSON handler with context
	http.HandleFunc("/api/data", middleware(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		select {
		case <-ctx.Done():
			http.Error(w, "Request timeout", http.StatusGatewayTimeout)
			return
		case <-time.After(1 * time.Second):
			response := Response{
				Status:  "success",
				Message: "Data retrieved",
				Data:    map[string]string{"key": "value"},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}))

	// Start server in goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)
}

func httpClientExample() {
	// Create client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/api/data", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse response
	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %+v\n", result)
}

func contextExample() {
	// Context with value
	ctx := context.WithValue(context.Background(), "userID", "123")

	// Context with timeout
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Use context in goroutine
	go func(ctx context.Context) {
		if userID, ok := ctx.Value("userID").(string); ok {
			log.Printf("Processing for user: %s\n", userID)
		}

		select {
		case <-ctx.Done():
			log.Println("Operation cancelled:", ctx.Err())
		case <-time.After(1 * time.Second):
			log.Println("Operation completed")
		}
	}(ctx)

	// Wait for operation
	time.Sleep(1500 * time.Millisecond)
}

func main() {
	log.Println("=== HTTP and Context Examples ===")

	log.Println("\n1. Starting HTTP Server")
	httpServerExample()

	log.Println("\n2. HTTP Client Operations")
	httpClientExample()

	log.Println("\n3. Context Handling")
	contextExample()

	log.Println("Main: All done")
}
