package main

import (
	"fmt"
	"sync"
)

type SampleMap struct {
	mu   sync.Mutex
	data map[string]SampleStruct
}

type SampleStruct struct {
	Name  string
	Value int
}

func NewSampleMap() *SampleMap {
	return &SampleMap{
		data: make(map[string]SampleStruct),
	}
}

func (m *SampleMap) Store(key string, value SampleStruct) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *SampleMap) Load(key string) (SampleStruct, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.data[key]
	return val, ok
}

func (m *SampleMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func main() {
	sampleMap := NewSampleMap()

	sampleMap.Store("car", SampleStruct{Name: "Car", Value: 10})
	sampleMap.Store("bus", SampleStruct{Name: "Bus", Value: 5})

	if val, ok := sampleMap.Load("car"); ok {
		fmt.Printf("Loaded: %+v\n", val)
	} else {
		fmt.Println("Key 'car' not found.")
	}

	if val, ok := sampleMap.Load("bus"); ok {
		fmt.Printf("Loaded: %+v\n", val)
	} else {
		fmt.Println("Key 'bus' not found.")
	}

	sampleMap.Delete("car")

	if val, ok := sampleMap.Load("car"); ok {
		fmt.Printf("Loaded after deletion: %+v\n", val)
	} else {
		fmt.Println("Key 'car' not found after deletion.")
	}
}
