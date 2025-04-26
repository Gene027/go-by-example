package main

import (
	"log"
	"regexp"
	"strings"
	"text/template"
)

/**
 * String Manipulation in Go demonstrates various string operations, formatting,
 * template processing, and regular expressions.
 *
 * Key concepts:
 * - String functions and manipulation
 * - String formatting with fmt
 * - Text template processing
 * - Regular expression patterns
 *
 * Common use cases:
 * - Text processing
 * - Data formatting
 * - Template rendering
 * - Pattern matching
 * - Input validation
 */

/**
 * User represents a data structure for template demonstration
 */
type User struct {
	Name    string
	Age     int
	Hobbies []string
}

/**
 * stringFunctions demonstrates common string operations
 */
func stringFunctions() {
	text := "  Hello, World!  "
	log.Printf("Original: %q\n", text)
	log.Printf("Trimmed: %q\n", strings.TrimSpace(text))
	log.Printf("Upper: %s\n", strings.ToUpper(text))
	log.Printf("Lower: %s\n", strings.ToLower(text))
	log.Printf("Contains 'World': %v\n", strings.Contains(text, "World"))
	log.Printf("Replace: %s\n", strings.Replace(text, "World", "Go", 1)) // The int argument is the number of replacements to make

	parts := strings.Split("a,b,c", ",")
	log.Printf("Split: %v\n", parts)
	log.Printf("Join: %s\n", strings.Join(parts, "-"))
}

/**
 * stringFormatting shows various formatting options
 */
func stringFormatting() {
	name := "Alice"
	age := 30
	height := 1.75
	items := []string{"apple", "banana", "orange"}

	// Basic formatting
	log.Printf("String: %s\n", name)
	log.Printf("Integer: %d\n", age)
	log.Printf("Float: %.2f\n", height)
	log.Printf("Array: %v\n", items)

	// Advanced formatting
	log.Printf("Padded number: %05d\n", age)
	log.Printf("Scientific: %e\n", height)
	log.Printf("Type information: %#v\n", items)
	log.Printf("Binary: %b\n", age)
}

/**
 * textTemplates demonstrates text template processing
 */
func textTemplates() {
	user := User{
		Name:    "Bob",
		Age:     25,
		Hobbies: []string{"reading", "gaming", "coding"},
	}

	// Simple template
	tmpl := `User Profile:
Name: {{.Name}}
Age: {{.Age}}
Hobbies: {{range .Hobbies}}
  - {{.}}{{end}}
`
	t := template.Must(template.New("profile").Parse(tmpl))
	t.Execute(log.Writer(), user)

	// Template with functions
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"join":  strings.Join,
	}

	advTmpl := `
{{upper .Name}}'s Profile
Hobbies: {{join .Hobbies ", "}}
`
	t = template.Must(template.New("advanced").Funcs(funcMap).Parse(advTmpl))
	t.Execute(log.Writer(), user)
}

/**
 * regularExpressions shows pattern matching and validation
 */
func regularExpressions() {
	// Email validation pattern
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailRegex := regexp.MustCompile(emailPattern)

	emails := []string{
		"user@example.com",
		"invalid.email@",
		"another.user@domain.co.uk",
	}

	for _, email := range emails {
		log.Printf("Email %q valid: %v\n", email, emailRegex.MatchString(email))
	}

	// Find and replace
	text := "Contact us at support@example.com or sales@example.com"
	pattern := regexp.MustCompile(`\w+@\w+\.\w+`)

	// Find all matches
	matches := pattern.FindAllString(text, -1)
	log.Printf("Found emails: %v\n", matches)

	// Replace matches
	masked := pattern.ReplaceAllString(text, "[EMAIL]")
	log.Printf("Masked text: %s\n", masked)

	// Extract information
	log.Println("\nExtracting information:")
	datePattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "Date: 2024-03-15"
	if parts := datePattern.FindStringSubmatch(date); len(parts) > 0 {
		log.Printf("Year: %s, Month: %s, Day: %s\n", parts[1], parts[2], parts[3])
	}
}

func main() {
	log.Println("=== String Manipulation Examples ===")

	/**
	 * 1. String Functions
	 * Shows common string operations
	 */
	log.Println("\n1. String Functions")
	stringFunctions()

	/**
	 * 2. String Formatting
	 * Demonstrates various formatting options
	 */
	log.Println("\n2. String Formatting")
	stringFormatting()

	/**
	 * 3. Text Templates
	 * Shows template processing
	 */
	log.Println("\n3. Text Templates")
	textTemplates()

	/**
	 * 4. Regular Expressions
	 * Demonstrates pattern matching
	 */
	log.Println("\n4. Regular Expressions")
	regularExpressions()

	log.Println("Main: All done")
}
