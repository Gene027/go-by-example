package main

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/**
 * Data Formats and Time Handling in Go demonstrates common data processing tasks.
 *
 * Key concepts:
 * - JSON/XML encoding and decoding
 * - Time operations and formatting
 * - Random number generation
 * - Number parsing and conversion
 * - Base64 encoding/decoding
 */

// Person represents a data structure for serialization examples
type Person struct {
	Name      string    `json:"name" xml:"name"`
	Age       int       `json:"age" xml:"age"`
	Birthday  time.Time `json:"birthday" xml:"birthday"`
	Addresses []Address `json:"addresses" xml:"address"`
}

// Address represents a nested structure
type Address struct {
	Street string `json:"street" xml:"street"`
	City   string `json:"city" xml:"city"`
}

func jsonExample() {
	person := Person{
		Name:     "Alice",
		Age:      30,
		Birthday: time.Date(1993, time.April, 15, 0, 0, 0, 0, time.UTC),
		Addresses: []Address{
			{Street: "123 Main St", City: "Boston"},
			{Street: "456 Oak Rd", City: "New York"},
		},
	}

	// Marshal (encode) to JSON
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("JSON:\n%s\n", jsonData)

	// Unmarshal (decode) from JSON
	var decodedPerson Person
	if err := json.Unmarshal(jsonData, &decodedPerson); err != nil {
		log.Fatal(err)
	}
	log.Printf("Decoded: %+v\n", decodedPerson)
}

func xmlExample() {
	person := Person{
		Name:     "Bob",
		Age:      25,
		Birthday: time.Date(1998, time.July, 10, 0, 0, 0, 0, time.UTC),
		Addresses: []Address{
			{Street: "789 Pine St", City: "Chicago"},
		},
	}

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("XML:\n%s\n", xmlData)

	// Unmarshal from XML
	var decodedPerson Person
	if err := xml.Unmarshal(xmlData, &decodedPerson); err != nil {
		log.Fatal(err)
	}
	log.Printf("Decoded: %+v\n", decodedPerson)
}

func timeExample() {
	now := time.Now()
	log.Printf("Current time: %v\n", now)
	log.Printf("Unix epoch: %d\n", now.Unix())
	log.Printf("Formatted (RFC3339): %s\n", now.Format(time.RFC3339))
	log.Printf("Formatted (custom): %s\n", now.Format("2006-01-02 15:04:05"))

	// Time parsing
	timeStr := "2024-03-15 14:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Parsed time: %v\n", parsedTime)

	// Time calculations
	future := now.Add(24 * time.Hour)
	duration := future.Sub(now)
	log.Printf("Tomorrow: %v (in %v)\n", future, duration)
}

func randomExample() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random numbers
	log.Printf("Random int: %d\n", rand.Int())
	log.Printf("Random float: %f\n", rand.Float64())
	log.Printf("Random range (1-100): %d\n", rand.Intn(100)+1)

	// Random string generation
	bytes := make([]byte, 16)
	rand.Read(bytes)
	log.Printf("Random bytes (base64): %s\n", base64.StdEncoding.EncodeToString(bytes))
}

func numberParsingExample() {
	// String to number conversion
	intStr := "42"
	num, err := strconv.Atoi(intStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Parsed int: %d\n", num)

	floatStr := "3.14159"
	pi, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Parsed float: %f\n", pi)

	// Number to string conversion
	log.Printf("Number as string: %s\n", strconv.Itoa(num))
	log.Printf("Float as string: %s\n", strconv.FormatFloat(pi, 'f', 2, 64))
}

func base64Example() {
	data := []byte("Hello, World!")

	// Encode
	encoded := base64.StdEncoding.EncodeToString(data)
	log.Printf("Base64 encoded: %s\n", encoded)

	// Decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Base64 decoded: %s\n", decoded)
}

func main() {
	log.Println("=== Data Formats and Time Examples ===")

	log.Println("\n1. JSON Processing")
	jsonExample()

	log.Println("\n2. XML Processing")
	xmlExample()

	log.Println("\n3. Time Operations")
	timeExample()

	log.Println("\n4. Random Numbers")
	randomExample()

	log.Println("\n5. Number Parsing")
	numberParsingExample()

	log.Println("\n6. Base64 Encoding")
	base64Example()

	log.Println("Main: All done")
}
