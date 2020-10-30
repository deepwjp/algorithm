package main

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	start, end := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			start = i
			break
		}
	}
	if start == -1 {
		return []int{start, end}
	}
	for i := len(nums) - 1; i >= start; i-- {
		if nums[i] == target {
			end = i
			break
		}

	}
	return []int{start, end}
}

// @lc code=end
