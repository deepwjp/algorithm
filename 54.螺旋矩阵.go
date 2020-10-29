package main

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var ret []int
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		// 向右
		for j := left; j <= right; j++ {
			ret = append(ret, matrix[up][j])
		}
		if up++; up > down {
			break
		}

		// 向下
		for i := up; i <= down; i++ {
			ret = append(ret, matrix[i][right])
		}
		if right--; right < left {
			break
		}

		// 向左
		for j := right; j >= left; j-- {
			ret = append(ret, matrix[down][j])
		}
		if down--; down < up {
			break
		}

		// 向上
		for i := down; i >= up; i-- {
			ret = append(ret, matrix[i][left])
		}
		if left++; left > right {
			break
		}
	}
	return ret
}

// @lc code=end
