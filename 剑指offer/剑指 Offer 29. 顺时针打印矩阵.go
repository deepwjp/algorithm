package 剑指offer

/* 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

限制：
0 <= matrix.length <= 100
0 <= matrix[i].length <= 100
*/

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
