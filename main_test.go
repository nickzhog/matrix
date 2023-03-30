package main

import (
	"testing"
)

func BenchmarkFindLongestGroup(b *testing.B) {
	height, width, colorsCount := 100, 100, 10
	matrix := newMatrix(height, width, colorsCount)

	for n := 0; n < b.N; n++ {
		findLongestGroup(matrix)
	}
}
