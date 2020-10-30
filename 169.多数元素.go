package main

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
/*
解法
1. hash 统计
2. sort 排序，取中点
3. 摩尔投票法
*/
func majorityElement(nums []int) int {
	// 票数
	votes := 0
	// 众数
	var x int
	for _, value := range nums {
		// 根据区间判断，如果之前的票数和为0，那么后半段众数票数一定大于0
		// 所以当票数为 0 时，设置当前数为众数，然后统计票数
		if votes == 0 {
			x = value
		}
		// 如果当前数是众数，票数增加，否则票数减少
		if value == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}

// @lc code=end
