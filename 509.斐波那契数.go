package main

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	f0, f1 := 0, 1

	for i := 2; i <= N; i = i + 2 {
		f0 = f0 + f1
		f1 = f0 + f1
	}
	if N%2 == 0 {
		return f1
	}
	return f0
}

// @lc code=end
