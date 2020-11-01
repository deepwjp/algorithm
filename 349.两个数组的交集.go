package main

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) (res []int) {
	maps := make(map[int]bool)
	for _, value := range nums1 {
		maps[value] = false
	}
	for _, value := range nums2 {
		if _, ok := maps[value]; ok {
			maps[value] = true
			// res = append(res, value)
		}
	}
	for key, value := range maps {
		if value == true {
			res = append(res, key)
		}
	}
	return
}

// @lc code=end
