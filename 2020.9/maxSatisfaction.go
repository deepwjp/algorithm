package main

import "sort"

//  1402. 做菜顺序
func maxSatisfaction(satisfaction []int) (max int) {
	n := len(satisfaction)
	sort.Ints(satisfaction)
	count := 1
	if satisfaction[n-1] <= 0 {
		return
	}
	max = satisfaction[n-1]
	for i := n - 1; i >= 0; i-- {
		curMax := 0
		for j := 1; j <= count; j++ {
			curMax += satisfaction[n-j] * (count - j + 1)
		}
		if curMax >= max {
			max = curMax
			count++
		}
	}
	return
}
