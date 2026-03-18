package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	testCases := map[string]int{
		"empty slice":                     0,
		"creates slice of specified size": 10,
	}
	for name, size := range testCases {
		t.Run(name, func(t *testing.T) {
			slice := generateRandomElements(size)
			assert.Equal(t, size, len(slice))
		})
	}
	t.Run("generates different slices", func(t *testing.T) {
		size := 20
		slice1 := generateRandomElements(size)
		time.Sleep(1 * time.Second)
		slice2 := generateRandomElements(size)
		assert.NotEqual(t, slice1, slice2)
	})
}
func TestMaximum(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		expected int
	}{
		{name: "check max empty slice",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "check max in single slice",
			slice:    []int{3},
			expected: 3,
		},
		{
			name:     "check max in slice",
			slice:    []int{3, 2, 0, 1},
			expected: 3,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			max := maximum(test.slice)
			assert.Equal(t, test.expected, max)
		})
	}
}
func TestMaxChunks(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		expected int
	}{
		{name: "check max empty slice",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "check max in single slice",
			slice:    []int{3},
			expected: 3,
		},
		{
			name:     "check max in slice",
			slice:    []int{3, 2, 0, 1, 5, 6, 7, 8, 9},
			expected: 9,
		},
		{
			name:     "check max in slice",
			slice:    []int{3, 2, 0, 1, 5},
			expected: 5,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			max := maxChunks(test.slice)
			assert.Equal(t, test.expected, max)
		})
	}
}
