package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go sendData(ch1, 5, 7, 4)
	go sendData(ch2, 10, 11, 8)
	go sendData(ch3, 3, 2, 9)

	merged := mergeChannels(ch1, ch2, ch3)

	for num := range merged {
		fmt.Println(num)
	}
}

func sendData(ch chan int, nums ...int) {
	for _, num := range nums {
		ch <- num
	}
	close(ch)
}

func mergeChannels(channels ...chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	output := func(ch <-chan int) {
		defer wg.Done()

		for num := range ch {
			merged <- num
		}
	}

	wg.Add(len(channels))

	for _, ch := range channels {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
