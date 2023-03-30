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

	longestGroup, gtype := findLongestGroup(matrix)
	fmt.Printf("Наибольшее количество одинаковых цветов подряд: %s %d, цвет: %d \n",
		gtype, len(longestGroup), longestGroup[0].color)
}

type block struct {
	color int
}

const (
	TypeHorizontal = "горизонтально"
	TypeVertical   = "вертикально"
)

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

func findLongestGroup(matrix [][]block) ([]block, string) {
	var longestGroup []block
	var groupType string

	// Проверка горизонтальных групп
	for i := range matrix {
		currentGroup := make([]block, 0)
		for j := range matrix[i] {
			if len(currentGroup) == 0 || currentGroup[len(currentGroup)-1].color == matrix[i][j].color {
				currentGroup = append(currentGroup, matrix[i][j])
			} else {
				if len(currentGroup) > len(longestGroup) {
					longestGroup = currentGroup
					groupType = TypeHorizontal
				}
				currentGroup = []block{matrix[i][j]}
			}
		}
		if len(currentGroup) > len(longestGroup) {
			longestGroup = currentGroup
			groupType = TypeHorizontal
		}
	}

	// Проверка вертикальных групп
	for j := range matrix[0] {
		currentGroup := make([]block, 0)
		for i := range matrix {
			if len(currentGroup) == 0 || currentGroup[len(currentGroup)-1].color == matrix[i][j].color {
				currentGroup = append(currentGroup, matrix[i][j])
			} else {
				if len(currentGroup) > len(longestGroup) {
					longestGroup = currentGroup
					groupType = TypeVertical
				}
				currentGroup = []block{matrix[i][j]}
			}
		}
		if len(currentGroup) > len(longestGroup) {
			longestGroup = currentGroup
			groupType = TypeVertical
		}
	}

	return longestGroup, groupType
}
