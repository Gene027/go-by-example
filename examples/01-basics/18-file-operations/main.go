package main

import (
	"bufio"
	"crypto/sha256"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

/**
 * File Operations and System Utilities in Go demonstrates common file handling,
 * path manipulation, and cryptographic operations.
 *
 * Key concepts:
 * - File reading and writing
 * - Directory operations
 * - Path manipulation
 * - URL parsing
 * - Cryptographic hashing
 *
 * Common use cases:
 * - Configuration file handling
 * - Log file processing
 * - Data import/export
 * - Web URL validation
 * - Data integrity verification
 */

func readFileExample() {
	// Reading entire file
	content, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("File contents: %s\n", content)

	// Reading file line by line
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Printf("Line: %s\n", scanner.Text())
	}
}

func writeFileExample() {
	// Writing string to file
	data := []byte("Hello, Go!\n")
	err := os.WriteFile("testdata/output.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Append to file
	file, err := os.OpenFile("testdata/output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString("More data\n"); err != nil {
		log.Fatal(err)
	}
}

func directoryExample() {
	// Create directory
	if err := os.MkdirAll("testdata/subdir", 0755); err != nil {
		log.Fatal(err)
	}

	// List directory contents
	entries, err := os.ReadDir("testdata")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Printf("Error getting info for %s: %v\n", entry.Name(), err)
			continue
		}
		log.Printf("Name: %s, Size: %d, IsDir: %v\n",
			entry.Name(), info.Size(), entry.IsDir())
	}

	// Walk directory tree
	err = filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		log.Printf("Path: %s, IsDir: %v\n", path, info.IsDir())
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func urlParsingExample() {
	urlStr := "https://user:pass@example.com:8080/path?key=value#fragment"
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Scheme: %s\n", parsedURL.Scheme)
	log.Printf("Host: %s\n", parsedURL.Host)
	log.Printf("Path: %s\n", parsedURL.Path)
	log.Printf("Query: %v\n", parsedURL.Query())
	log.Printf("Fragment: %s\n", parsedURL.Fragment)

	// Build URL
	u := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/api",
		RawQuery: "sort=desc&page=1",
	}
	log.Printf("Built URL: %s\n", u.String())
}

func hashingExample() {
	// Calculate SHA256 hash of string
	data := []byte("Hello, World!")
	hash := sha256.Sum256(data)
	log.Printf("SHA256 hash: %x\n", hash)

	// Hash file content
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		log.Fatal(err)
	}
	log.Printf("File hash: %x\n", hasher.Sum(nil))
}

func tempFileExample() {
	// Create temporary file
	tempFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	log.Printf("Temp file created: %s\n", tempFile.Name())

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "example-*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	log.Printf("Temp directory created: %s\n", tempDir)
}

func main() {
	log.Println("=== File Operations and System Utilities Examples ===")

	// Create testdata directory if it doesn't exist
	if err := os.MkdirAll("testdata", 0755); err != nil {
		log.Fatal(err)
	}

	// Create sample input file
	sampleData := []byte("Line 1\nLine 2\nLine 3\n")
	if err := os.WriteFile("testdata/input.txt", sampleData, 0644); err != nil {
		log.Fatal(err)
	}

	log.Println("\n1. File Reading")
	readFileExample()

	log.Println("\n2. File Writing")
	writeFileExample()

	log.Println("\n3. Directory Operations")
	directoryExample()

	log.Println("\n4. URL Parsing")
	urlParsingExample()

	log.Println("\n5. SHA256 Hashing")
	hashingExample()

	log.Println("\n6. Temporary Files and Directories")
	tempFileExample()

	log.Println("Main: All done")
}
