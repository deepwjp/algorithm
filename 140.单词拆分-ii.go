package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	// 边界条件判断
	if s == "" || len(s) == 0 || wordDict == nil || len(wordDict) == 0 {
		return nil
	}
	// 保存结果
	res := make([]string, 0)
	set := [26]bool{}

	// 将 wordDict 转化为 string 将每个元素放到set中进行去重
	for _, v := range strings.Join(wordDict, "") {
		set[v-'a'] = true
	}

	// 进行判断 如果 wordDict中有元素不在 s中一定是无法构成词
	for _, v := range s {
		if !set[v-'a'] {
			return res
		}
	}

	var dfs func(nowString string, index int)
	max := len(s)
	dfs = func(nowString string, index int) {
		// 递归终止条件
		if index >= max {
			res = append(res, nowString[1:])
			return
		}
		// 遍历 wordDict 拿到每个词
		for _, word := range wordDict {
			lenWord := len(word)
			// 如果那个词跟当前 s可以匹配的上 则进入下一层
			if index+lenWord <= len(s) && s[index:index+lenWord] == word {
				dfs(nowString+" "+word, index+lenWord)
			}
		}
	}
	dfs("", 0)
	return res
}

// @lc code=end
