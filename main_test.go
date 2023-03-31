package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestGroup(t *testing.T) {
	matrixTests := []struct {
		matrix [][]block
		color  int
		len    int
	}{
		{
			matrix: [][]block{
				{{0}, {0}, {0}, {0}},
				{{1}, {0}, {1}, {0}},
				{{0}, {2}, {1}, {1}},
				{{0}, {0}, {1}, {0}},
				{{0}, {2}, {1}, {1}},
				{{1}, {0}, {1}, {0}},
				{{1}, {0}, {0}, {0}},
			},
			color: 1,
			len:   7,
		},
		{
			matrix: [][]block{
				{{2}, {0}, {0}, {0}},
				{{2}, {0}, {2}, {0}},
				{{2}, {2}, {2}, {2}},
				{{0}, {0}, {4}, {2}},
				{{0}, {2}, {3}, {2}},
				{{1}, {1}, {1}, {1}},
				{{1}, {1}, {1}, {1}},
			},
			color: 2,
			len:   9,
		},
	}
	for _, test := range matrixTests {
		longestGroup := findLongestGroup(test.matrix)

		assert := assert.New(t)
		assert.Equal(test.color, longestGroup[0].color)
		assert.Equal(test.len, len(longestGroup))

	}
}

func BenchmarkFindLongestGroup(b *testing.B) {
	height, width, colorsCount := 100, 100, 10
	matrix := newMatrix(height, width, colorsCount)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findLongestGroup(matrix)
	}
}
