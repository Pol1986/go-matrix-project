package main

import "fmt"

\\Константы
const row = 3
const cols = 4

func sumMatrix(A [row][cols]int, B [row][cols]int) [row][cols]int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] += B[i][j]
		}
	}
	return A
}

func diagonalSum(A [row][cols]int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i][i]
	}
	return sum
}

func trans(A [row][cols]int) [cols][row]int {
	var transposed [cols][row]int
	for i := 0; i < len(A[0]); i++ {
		for j := 0; j < len(A); j++ {
			transposed[i][j] = A[j][i]
		}
	}
	return transposed
}

func main() {
	matrix := [row][cols]int{
		{10, 10, 10, 10},
		{10, 20, 10, 20},
		{-10, -20, -20, 10},
	}
	fmt.Println(sumMatrix(matrix, matrix))
	fmt.Println(diagonalSum(matrix))
	fmt.Println(trans(matrix))
}
