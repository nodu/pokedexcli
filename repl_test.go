package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input: "  hello  ",
			expected: []string{
				"hello",
			},
		},
		{
			input: "  hello world  ",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world     ",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lenghts are not equal: %v, vs %v",
				len(actual),
				len(cs.expected),
			)
			continue
		}
		// save indexing operation because we've already checked lenghts are equal
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", cs.input, actual, cs.expected)
			}

		}
	}
}
