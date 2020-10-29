package main

//921. 使括号有效的最少添加
func minAddToMakeValid(S string) (res int) {
	count := 0
	for _, value := range S {
		// 保持左右括号平衡
		if value == '(' {
			count++
		}
		if value == ')' {
			count--
		}
		// 如果右括号比左括号多，需要添加的数量+1,更新count
		if count == -1 {
			res++
			count++
		}
	}
	//优括号多余左括号的数量加上末尾左括号的数量
	return res + count
}
