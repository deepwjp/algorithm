package main

import (
	"fmt"
	"sort"
)

// 面试题 08.08. 有重复字符串的排列组合
func main() {
	//nums := []int{-1, 2, -1, 2, 1, -1, 2, 1}
	//nums := []int{1, 2, 1}
	str := "qqe"
	fmt.Print(permutation2(str))
}

func permutation2(str string) []string {
	if len(str) == 0 {
		return nil
	}
	newStr := []string{}
	for _, value := range str {
		newStr = append(newStr, string(value))
	}
	result := []string{}
	sort.Strings(newStr)

	dfsPermutation2(newStr, 0, &result)
	return result
}

// 回溯函数实现
// i表示本次函数需要放置的元素位置
func dfsPermutation2(str []string, i int, result *[]string) {
	n := len(str)
	if i == n-1 {
		var tmp string
		//tmp := make([]string, n)
		//copy(tmp, str)
		//tmp = append(tmp, value)
		for _, value := range str {
			tmp += value
		}
		*result = append(*result, tmp)
		return
	}
	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
	for k := i; k < n; k++ {
		// 跳过重复数字
		if k != i && str[k] == str[i] {
			continue
		}
		str[k], str[i] = str[i], str[k]
		dfsPermutation2(str, i+1, result)
	}
	// 还原状态
	for k := n - 1; k > i; k-- {
		str[i], str[k] = str[k], str[i]
	}
}
