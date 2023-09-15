package main

import (
	"fmt"
	"testing"
)

func TestIntSliceSort(t *testing.T) {
	tests := []struct {
		input    IntSlice
		expected IntSlice
	}{
		{IntSlice{9, 8, 9, 3, 11, 5}, IntSlice{3, 5, 8, 9, 9, 11}},
		{IntSlice{5, 4, 3, 2, 1}, IntSlice{1, 2, 3, 4, 5}},
		{IntSlice{1}, IntSlice{1}},
		{IntSlice{}, IntSlice{}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Sorting %v", test.input), func(t *testing.T) {
			test.input.Sort()
			if !test.input.Equal(test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, test.input)
			}
		})
	}
}

func (s IntSlice) Equal(other IntSlice) bool {
	if len(s) != len(other) {
		return false
	}
	for i, val := range s {
		if val != other[i] {
			return false
		}
	}
	return true
}
