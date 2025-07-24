package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "CHARMANDER Bulbasaur PIkAchu",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	// Loop over cases here
	for _, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("cleanInput(%q) = %v, want %v", c.input, actual, c.expected)
		}

	}
}
