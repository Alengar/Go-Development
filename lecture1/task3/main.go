package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{9, 8, 6, 4, 5}
	slice2 := []int{4, 5, 9, 6, 8}

	if sliceEqual(slice1, slice2) {
		fmt.Println("The slices are equal.")
	} else {
		fmt.Println("The slices are not equal.")
	}
}

func sliceEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	sort.Ints(slice1)
	sort.Ints(slice2)

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
