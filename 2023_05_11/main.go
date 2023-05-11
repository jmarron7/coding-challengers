package main

import "fmt"

func pascals(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i][i] = 1
	}
	return result
}

func pascalsSymmetry(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j <= i/2; j++ {
			val := result[i-1][j-1] + result[i-1][j]
			row[j], row[i-j] = val, val
		}
		result[i] = row
	}
	return result
}

func main() {
	fmt.Println(pascals(1)) // [[1]]
	fmt.Println(pascals(2)) // [[1], [1, 1]]
	fmt.Println(pascals(3)) // [[1], [1, 1], [1, 2, 1]]
	fmt.Println(pascals(4)) // [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1]]
	fmt.Println(pascals(5)) // [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]

	fmt.Println(pascalsSymmetry(1)) // [[1]]
	fmt.Println(pascalsSymmetry(2)) // [[1], [1, 1]]
	fmt.Println(pascalsSymmetry(3)) // [[1], [1, 1], [1, 2, 1]]
	fmt.Println(pascalsSymmetry(4)) // [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1]]
	fmt.Println(pascalsSymmetry(5)) // [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
}
