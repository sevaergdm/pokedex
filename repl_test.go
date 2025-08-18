package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  This is a string  ",
			expected: []string{"this", "is", "a", "string"},
		},
		{
			input:    " I AM SHOUTING ",
			expected: []string{"i", "am", "shouting"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected output of length %d, but got: %d", len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %s, but got %s", expectedWord, word)
			}
		}
	}
}
