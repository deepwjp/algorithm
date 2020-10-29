package main

/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) (res []int) {

	for key, value1 := range nums {
		tmp := 0
		for key2, value2 := range nums {
			if key == key2 {
				continue
			}
			if value1 > value2 {
				tmp++
			}
		}
		res = append(res, tmp)
	}
	return
}

// @lc code=end
