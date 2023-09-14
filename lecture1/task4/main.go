package main

import "fmt"

func main() {
	nums := IntSlice{9, 8, 9, 3, 11, 5}

	fmt.Println("Original Slice:", nums)

	nums.Sort()

	fmt.Println("Sorted Slice:", nums)
}

type IntSlice []int

func (s IntSlice) Sort() { // Используя Bubble sort
	n := len(s)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if s[i-1] > s[i] {
				s[i-1], s[i] = s[i], s[i-1]
				swapped = true
			}
		}
		n--
	}
}
