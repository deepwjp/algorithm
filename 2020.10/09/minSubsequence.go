package main

//1403. 非递增顺序的最小子序列
import (
	"sort"
)

func minSubsequence(nums []int) (res []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sumMax := 0
	curMax := 0
	for _, value := range nums {
		sumMax += value
	}
	sumMax = sumMax / 2
	for _, value := range nums {
		curMax += value
		res = append(res, value)
		if curMax > sumMax {
			return res
		}
	}
	return
}
