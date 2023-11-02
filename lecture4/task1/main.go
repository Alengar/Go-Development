package main

import (
	"fmt"
	"sync"
)

var (
	numbers []int
	wg      sync.WaitGroup
)

func appendNumber(n int) {
	numbers = append(numbers, n)
	wg.Done()
}

func main() {
	for i := 0; i < 985; i++ {
		wg.Add(1)
		go appendNumber(i)
	}

	wg.Wait()

	// Длина мапы 'numbers' может не быть равной 985 из-за race condition
	fmt.Printf("Конечная Длина мапы 'numbers': %d\n", len(numbers))
}
