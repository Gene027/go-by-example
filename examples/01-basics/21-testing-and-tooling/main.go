package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
)

/**
 * Testing, Tooling, and Advanced Features in Go demonstrates development tools
 * and additional language features.
 *
 * Key concepts:
 * - Embed directive
 * - Command-line flag parsing
 * - Subcommand handling
 * - Line filtering
 * - Testing patterns
 *
 * Common use cases:
 * - CLI applications
 * - Resource embedding
 * - Text processing
 * - Application testing
 * - Performance benchmarking
 */

//go:embed testdata/sample.txt
var embeddedContent string

//go:embed testdata/config.json
var configData []byte

// Command-line flags
var (
	verbose = flag.Bool("verbose", false, "Enable verbose output")
	format  = flag.String("format", "text", "Output format (text/json)")
	count   = flag.Int("count", 1, "Number of iterations")
)

// Subcommands
type SubCommand struct {
	Name        string
	Description string
	Execute     func([]string) error
}

func lineFilter(input string, filters []string) string {
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		include := true
		for _, filter := range filters {
			if strings.Contains(line, filter) {
				include = false
				break
			}
		}
		if include {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

func main() {
	log.Println("=== Testing and Tooling Examples ===")

	// Define subcommands
	commands := map[string]SubCommand{
		"filter": {
			Name:        "filter",
			Description: "Filter input lines",
			Execute: func(args []string) error {
				if len(args) < 1 {
					return fmt.Errorf("filter requires at least one argument")
				}
				filtered := lineFilter(embeddedContent, args)
				fmt.Println(filtered)
				return nil
			},
		},
		"version": {
			Name:        "version",
			Description: "Show version information",
			Execute: func([]string) error {
				fmt.Println("Version 1.0.0")
				return nil
			},
		},
	}

	// Parse flags
	flag.Parse()

	// Handle subcommands
	if flag.NArg() > 0 {
		if cmd, ok := commands[flag.Arg(0)]; ok {
			if err := cmd.Execute(flag.Args()[1:]); err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	// Main program logic
	log.Printf("Embedded content length: %d bytes\n", len(embeddedContent))
	log.Printf("Config data length: %d bytes\n", len(configData))

	if *verbose {
		log.Printf("Verbose mode enabled")
		log.Printf("Format: %s\n", *format)
		log.Printf("Count: %d\n", *count)
	}
}
