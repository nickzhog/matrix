package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestGroup(t *testing.T) {
	matrix := [][]block{
		{{0}, {0}, {0}, {0}},
		{{1}, {0}, {1}, {0}},
		{{0}, {2}, {1}, {1}},
		{{0}, {0}, {1}, {0}},
		{{0}, {2}, {1}, {1}},
		{{1}, {0}, {1}, {0}},
		{{1}, {0}, {0}, {0}},
	}
	longestGroup, gtype := findLongestGroup(matrix)

	assert := assert.New(t)
	assert.Equal(TypeVertical, gtype)
	assert.Equal(1, longestGroup[0].color)
	assert.Equal(5, len(longestGroup))
}

func BenchmarkFindLongestGroup(b *testing.B) {
	height, width, colorsCount := 100, 100, 10
	matrix := newMatrix(height, width, colorsCount)

	for n := 0; n < b.N; n++ {
		findLongestGroup(matrix)
	}
}
