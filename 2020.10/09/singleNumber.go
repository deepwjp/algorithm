package main

import "sort"

//剑指 Offer 56 - II. 数组中数字出现的次数 II

//题也可用 map
func singleNumber(nums []int) (res int) {
	sort.Ints(nums)
	n := len(nums)
	if nums[0] != nums[1] && nums[0] != nums[2] {
		return nums[0]
	}
	if nums[n-1] != nums[n-2] && nums[n] != nums[n-3] {
		return nums[n-1]
	}
	for i := 1; i < n-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return
}
