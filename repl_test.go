package main

import (
	"strings"
	"testing"
)

func TestCase(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("The words in the slices do not match")
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}

func cleanInput(s string) []string {
	convert := []rune(s)
	var slice []string
	var tmp []rune
	for index, value := range convert {
		if index == 0 {
			if value != 32 {
				tmp = append(tmp, value)
			}
			continue
		}
		if value == 32 {
			if tmp != nil {
				conversion := string(tmp)
				slice = append(slice, strings.ToLower(conversion))
				tmp = nil
			}
			continue
		}
		tmp = append(tmp, value)
	}
	return slice
}
