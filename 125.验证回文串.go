package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	reg := regexp.MustCompile(`[a-zA-Z0-9]*`)
	//根据规则提取关键信息
	strs := reg.FindAllStringSubmatch(s, -1)
	str := ""
	for _, value := range strs {
		str += value[0]
	}
	str = strings.ToLower(str)
	fmt.Println(str)
	for i := 0; i < len(str)/2; i++ {
		fmt.Println(str[i], str[len(str)-1-i])
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
