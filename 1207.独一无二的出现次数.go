package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) (res bool) {
	maps := make(map[int]int)
	tmp := []int{}
	for _, value := range arr {
		maps[value]++
	}
	for _, value := range maps {
		tmp = append(tmp, value)
	}
	sort.Ints(tmp)
	for i := 1; i < len(tmp); i++ {
		if tmp[i] == tmp[i-1] {
			return false
		}
	}
	return true
}

// @lc code=end
