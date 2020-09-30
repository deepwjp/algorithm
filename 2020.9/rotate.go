package main

import "fmt"

//func main() {
//	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
//	rotate(matrix)
//}

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 1 {
		return
	}

	// mirror by diagonal
	for i := 0; i < n; i++ {
		for j := n - i; j < n; j++ {
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}
	fmt.Println(matrix)
	// reverse the element order in each row
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
}

//func rotate(matrix [][]int) [][]int {
//	// 将旋转后的矩阵添加到现有矩阵之后
//	leng := len(matrix)
//	for i := leng - 1; i >= 0; i-- {
//		temp := matrix[i]
//		for j := 0; j < leng; j++ {
//			matrix[j] = append(matrix[j][:], temp[j])
//		}
//	}
//	// 更新新矩阵
//	for i := 0; i < leng; i++ {
//		matrix[i] = matrix[i][leng:]
//	}
//	return matrix
//}
