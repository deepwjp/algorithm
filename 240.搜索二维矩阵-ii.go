package main

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n := len(matrix)
	m := len(matrix[0])
	i, j := n-1, 0
	for i >= 0 && j < m {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		}

	}
	return false
}

// @lc code=end
