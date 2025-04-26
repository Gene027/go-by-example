package main

import (
	"testing"
)

func TestLineFilter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		filters  []string
		expected string
	}{
		{
			name:     "No filters",
			input:    "line1\nline2\nline3",
			filters:  []string{},
			expected: "line1\nline2\nline3",
		},
		{
			name:     "Single filter",
			input:    "line1\nline2\nline3",
			filters:  []string{"line2"},
			expected: "line1\nline3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lineFilter(tt.input, tt.filters)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func BenchmarkLineFilter(b *testing.B) {
	input := "line1\nline2\nline3\nline4\nline5"
	filters := []string{"line2", "line4"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lineFilter(input, filters)
	}
}
