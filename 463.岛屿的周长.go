package main

/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	length := 0
	// m, n := len(grid), len(gird[0])
	// 遍历二位数组
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 判断四周
			if grid[i][j] == 1 {
				if i < 1 || grid[i-1][j] != 1 {
					length++
				}
				if j < 1 || grid[i][j-1] != 1 {
					length++
				}
				if i >= len(grid)-1 || grid[i+1][j] != 1 {
					length++
				}
				if j >= len(grid[0])-1 || grid[i][j+1] != 1 {
					length++
				}
			}
		}
	}
	return length
}

// @lc code=end
