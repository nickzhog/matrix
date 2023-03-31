package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var height, width, colorsCount int
	fmt.Println("Введите ширину матрицы:")
	fmt.Scan(&width)
	if width < 1 {
		log.Fatal("не может быть меньше 1")
	}

	fmt.Println("Введите высоту матрицы:")
	fmt.Scan(&height)
	if height < 1 {
		log.Fatal("не может быть меньше 1")
	}

	fmt.Println("Введите количество цветов:")
	fmt.Scan(&colorsCount)
	if colorsCount < 1 {
		log.Fatal("не может быть меньше 1")
	}

	matrix := newMatrix(height, width, colorsCount)

	fmt.Println("Матрица:")
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%d ", matrix[i][j].color)
		}
		fmt.Println()
	}

	longestGroup := findLongestGroup(matrix)
	fmt.Printf("Наибольшее количество одинаковых цветов подряд: %d, цвет: %d \n",
		len(longestGroup), longestGroup[0].color)
}

type block struct {
	color int
}

func newMatrix(height, width, colorsCount int) [][]block {
	matrix := make([][]block, height)

	rand.Seed(time.Now().UnixNano())

	for i := range matrix {
		matrix[i] = make([]block, width)
		for j := range matrix[i] {
			matrix[i][j] = block{color: rand.Intn(colorsCount)}
		}
	}

	return matrix
}

type point struct {
	row int
	col int
}

func findLongestGroup(matrix [][]block) []block {
	visited := make(map[point]bool)
	longestGroup := make([]block, 0)

	for i := range matrix {
		for j := range matrix[i] {
			if !visited[point{i, j}] {
				color := matrix[i][j].color
				group := bfs(matrix, visited, color, i, j)

				if len(group) > len(longestGroup) {
					longestGroup = group
				}
			}
		}
	}

	return longestGroup
}

func bfs(matrix [][]block, visited map[point]bool, color, row, col int) []block {
	queue := []point{{row, col}}
	group := make([]block, 0)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if visited[p] {
			continue
		}

		if matrix[p.row][p.col].color != color {
			continue
		}

		group = append(group, matrix[p.row][p.col])
		visited[p] = true

		if p.row > 0 {
			queue = append(queue, point{p.row - 1, p.col})
		}
		if p.row < len(matrix)-1 {
			queue = append(queue, point{p.row + 1, p.col})
		}
		if p.col > 0 {
			queue = append(queue, point{p.row, p.col - 1})
		}
		if p.col < len(matrix[0])-1 {
			queue = append(queue, point{p.row, p.col + 1})
		}
	}

	return group
}
