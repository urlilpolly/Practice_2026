package main

import (
	"fmt"
	"math/rand"
)

func generateMatrix(m int, n int) [][]int {
	matrix := make([][]int, m)

	randomNumbers := rand.Perm(20)
	counter := 0

	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = randomNumbers[counter]
			counter++
		}
	}

	return matrix
}

func main() {
	m, n := 3, 4
	matrix := generateMatrix(m, n)

	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%4d ", val) // Выравнивание по правому краю
		}
		fmt.Println()
	}
}
