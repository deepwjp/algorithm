package main

//剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
