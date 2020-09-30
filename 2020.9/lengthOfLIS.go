package main

import (
	"sort"
)

//func main() {
//	arr := []int{2, 2}
//	fmt.Println(lengthOfLIS(arr))
//}

//动态规划
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lenNum := make([]int, len(nums))
	// 初始化长度slice
	for i := 0; i < len(nums); i++ {
		lenNum[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		temp := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && temp <= lenNum[i]+lenNum[j] {
				temp = lenNum[i] + lenNum[j]
			}
		}
		if temp > lenNum[i] {
			lenNum[i] = temp
		}
	}
	sort.Ints(lenNum[:])
	return lenNum[len(lenNum)-1]
}
