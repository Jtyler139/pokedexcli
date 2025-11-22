package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	}{
		{
			input:		"  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "The CHEESE gROmit",
			expected: []string{"the", "cheese", "gromit"},
		},
		{
			input: "	wow 	that's	a	lot	of	blank	space	",
			expected: []string{"wow", "that's", "a", "lot", "of", "blank", "space"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The length of actual input (%d) and the expected input (%d) do not match", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("The expected word (%s) and the given word (%s) do not match", expectedWord, word)
			}
		}
	}
}