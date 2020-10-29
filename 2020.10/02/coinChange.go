package main

import "math"

// 322. 零钱兑换
func coinChange(coins []int, amount int) int {
	// 定义一个足够大的数
	inf := math.MaxInt32
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = inf
	}
	// 完全背包模板
	for _, value := range coins {
		for i := value; i <= amount; i++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-value]+1)))
		}
	}
	//如果 dp[amount]更新后的结果与最开始相同，表示无法组成amount数值
	if dp[amount] == inf {
		return -1
	}
	return dp[amount]

}
