package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	// Проверяем на нулевой срез
	size := 0
	slice := generateRandomElements(size)

	assert.Nil(t, slice)
	// Проверяем что создается заданного размера
	size = 10
	slice = generateRandomElements(size)

	assert.Equal(t, size, len(slice))
	// Проверяем что создаются random срезы
	size = 20
	slice1 := generateRandomElements(size)
	time.Sleep(1 * time.Second)
	slice2 := generateRandomElements(size)

	assert.NotEqual(t, slice1, slice2)
}
func TestMaximum(t *testing.T) {
	// Проверяем если нулевой срез
	slice := []int{}
	max := maximum(slice)

	assert.Equal(t, 0, max)
	// Проверяем если 1 элемент в срезе
	slice = []int{3}
	max = maximum(slice)

	assert.Equal(t, slice[0], max)
	// Проверяем максимум в срезе
	slice = []int{3, 2, 0, 1}
	max = maximum(slice)

	assert.Equal(t, 3, max)
}
func TestMaxChunks(t *testing.T) {
	// Проверяем если нулевой срез
	slice := []int{}
	max := maxChunks(slice)

	assert.Equal(t, 0, max)
	// Проверяем если 1 элемент в срезе
	slice = []int{3}
	max = maxChunks(slice)

	assert.Equal(t, slice[0], max)
	// Проверяем максимум в срезе c len(slice)>CHUNKS
	slice = []int{3, 2, 0, 1, 5, 6, 7, 8, 9}
	max = maxChunks(slice)

	assert.Equal(t, 9, max)
	// Проверяем максимум в срезе c len(slice)<CHUNKS
	slice = []int{3, 2, 0, 1, 5}
	max = maxChunks(slice)

	assert.Equal(t, 5, max)
}
