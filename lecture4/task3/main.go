package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}
	mux := &sync.RWMutex{}

	// Запускаем горутину для записи в мапу
	go writeLoop(m, mux)

	// Запускаем несколько горутин для чтения мапы
	for i := 0; i < 4; i++ {
		go readLoop(m, mux)
	}

	// Блокируем выполнение программы, чтобы она не завершилась сразу
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			// Блокируем мьютекс для записи
			mux.Lock()
			m[i] = i
			// Разблокируем мьютекс после записи
			mux.Unlock()
		}
	}
}

func readLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		// Блокируем мьютекс для чтения
		mux.RLock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		// Разблокируем мьютекс после чтения
		mux.RUnlock()
	}
}
