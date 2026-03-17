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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size == 0 {
		return nil
	}
	slice := make([]int, 0, size)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for range size {
		randomNumber := rnd.Int() % size
		slice = append(slice, randomNumber)
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	if len(data) == 0 {
		return 0
	}
	slice := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	for i := range CHUNKS {
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		// Последний чанк забирает остаток
		if i == CHUNKS-1 {
			endIndex = len(data)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			slice = append(slice, maximum(data[startIndex:endIndex]))
			mu.Unlock()
		}()
	}
	wg.Wait()
	return maximum(slice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(slice)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
