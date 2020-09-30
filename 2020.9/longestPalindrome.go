package main

//func main() {
//	s := "aaa"
//	fmt.Println(longestPalindrome(s))
//}
func longestPalindrome(s string) string {
	n := len(s)
	start, end := 0, 0
	if s == "" {
		return s
	}

	for i := 0; i < n; i++ {
		//中心为字符为1时
		left1, right1 := solve(s, i, i)
		//中心为字符为2时
		left2, right2 := solve(s, i, i+1)
		//奇数回文串时
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		//偶数回文串时
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func solve(s string, left, right int) (int, int) {
	// left, right = left-1, right+1扩大范围继续判断
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	//循环中对数据进行更改已不满足回文串条件
	return left + 1, right - 1
}
