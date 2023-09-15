package main

import (
	"testing"
)

func TestSliceEqual(t *testing.T) {
	tests := []struct {
		slice1   []int
		slice2   []int
		expected bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1, 2, 3}, false},
		{[]int{}, []int{}, true},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := sliceEqual(test.slice1, test.slice2)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v for %v and %v", test.expected, result, test.slice1, test.slice2)
			}
		})
	}
}
