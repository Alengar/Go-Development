package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4}, 7, []int{2, 3}},
		{[]int{}, 5, nil}, // пустой слайс
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := twoSum(test.nums, test.target)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v for nums=%v, target=%d, but got %v", test.expected, test.nums, test.target, result)
			}
		})
	}
}
