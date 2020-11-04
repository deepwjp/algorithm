package main

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	ans := nums[0]
	cur := nums[0]
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(nums); i++ {
		// 判断 cur是否大于0
		// 如果小于0,就舍弃之前的值，否则更新值
		cur = max(nums[i], cur+nums[i])
		ans = max(cur, ans)
	}
	return ans
}

// @lc code=end
