package practice

import "fmt"

func rotate(matrix [][]int)  {
	length := len(matrix)
	//主对角线反转i, j -> j, i
	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	////列反转 j,i -> j, n-1-i
	for j := 0; j < length /2; j++ {
		for i := 0; i < length; i++ {
			matrix[i][j], matrix[i][length - 1- j] = matrix[i][length - 1-j], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}
