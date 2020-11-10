package main

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	n := len(nums)
	i, j := n-2, n-1
	// 两边逆向扫描
	// 先找到逆向第一个下降的数据
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 如果不是倒序slice
	if i >= 0 {
		// 第二遍逆向扫描，找到比nums[i]第一个数据
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 进行翻转
	reverse := func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	reverse(nums[i+1:])
}

// @lc code=end
