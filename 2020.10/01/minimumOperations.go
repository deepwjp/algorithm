package main

import "math"

//LCP 19. 秋叶收藏集
/*
   状态数组，dp[i][j]中：i表示终止下标,j表示：0为左半边 r，1为中间部分 y，2为右半边 r
   dp[i][j] 表示从0到i需要调整的叶子数
*/
func minimumOperations(leaves string) int {
	const inf = math.MaxInt32
	n := len(leaves)
	dp := make([][3]int, n)
	//如果当前叶子颜色为y则需要更改，更改数量+1
	isYellow := func(i int) int {
		if leaves[i] == 'y' {
			return 1
		}
		return 0
	}
	//比较更改次数多少，返回次数较少的
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 第一片叶子只能为 第一部分r
	dp[0][0] = isYellow(0)
	// 第一片叶子不可能为中间y 部分和后面r 部分，因为取最小值，所以设置最大inf
	dp[0][1] = inf
	dp[0][2] = inf
	// 第二片叶子不能为后面r 部分，因为取最小值，所以设置最大inf
	dp[1][2] = inf
	//因为第一篇叶子只能是r，所有从第二片叶子开始遍历字符串，
	for i := 1; i < n; i++ {
		// 判断第i片叶子是否是
		//isRed := isRed(i)
		// 判断第i片叶子是否是y
		isYellow := isYellow(i)
		//如果第i片属于第一部分，那么第i-1片也一定属于第一部分，前i-1片需要修改的数量加上当前修改情况为前i项需要修改的数量
		dp[i][0] = dp[i-1][0] + isYellow
		/*如果第i片叶子属于第二部分，那么第i-1片可能有两种情况，一种是第i-1片属于第一部分表示为dp[i-1][0]，一种是第i-1片属于第二部分dp[i-1][1]，这时需要选择较小的作为结果*/
		dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1 - isYellow
		if i >= 2 {
			/*如果第i片叶子属于第二部分，那么第i-1片可能有两种情况，一种是第i-1片属于第二部分表示为dp[i-1][1]，一种是第i-1片属于第三部分dp[i-1][2]，这时需要选择较小的作为结果*/
			dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + isYellow
		}
	}
	return dp[n-1][2]
}
