package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello    world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "qual lumpur WOWOWW Incroyable",
			expected: []string{"qual", "lumpur", "wowoww", "incroyable"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lenghth does not match: %v vs %v", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("'%v' does not match expected word '%v'", word, expectedWord)
			}
		}

	}
}
