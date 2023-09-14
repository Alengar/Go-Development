package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "car", "racecar"}, ""},
		{[]string{"hello", "hell", "hellfire"}, "hell"},
		{[]string{"", "abc", "def"}, ""},
		{[]string{}, ""},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := longestCommonPrefix(test.input)
			if result != test.expected {
				t.Errorf("For input %v, expected %s, but got %s", test.input, test.expected, result)
			}
		})
	}
}
