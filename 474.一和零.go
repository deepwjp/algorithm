package main

/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {

	if len(strs) == 0 {
		return 0
	}

	var dp [][]int
	dp = make([][]int, m+1, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1, n+1)
	}

	for _, str := range strs {
		mp := make(map[int32]int)
		// 获取当前字符串中0和1的个数
		for _, c := range str {
			mp[c] += 1
		}
		// 二维简化的背包
		for i := m; i >= mp['0']; i-- {
			for j := n; j >= mp['1']; j-- {
				dp[i][j] = max(1+dp[i-mp['0']][j-mp['1']], dp[i][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
