package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		v := rand.Int()
		if v == 0 {
			v = 1
		}
		data[i] = v
	}

	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxValue {
			maxValue = data[i]
		}
	}

	return maxValue
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		chunkSize = 1
	}

	results := make([]int, CHUNKS)
	var wg sync.WaitGroup
	actualChunks := 0

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		if start >= len(data) {
			break
		}

		end := start + chunkSize
		if i == CHUNKS-1 || end > len(data) {
			end = len(data)
		}

		index := actualChunks
		actualChunks++

		wg.Add(1)
		go func(idx int, chunk []int) {
			defer wg.Done()
			results[idx] = maximum(chunk)
		}(index, data[start:end])
	}

	wg.Wait()

	return maximum(results[:actualChunks])
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
