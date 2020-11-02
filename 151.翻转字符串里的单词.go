package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	strList := strings.Split(s, " ")
	var res []string
	for i := len(strList) - 1; i >= 0; i-- {
		str := strings.TrimSpace(strList[i])
		if len(str) > 0 {
			res = append(res, strList[i])
		}
	}
	return strings.Join(res, " ")
}

// @lc code=end
