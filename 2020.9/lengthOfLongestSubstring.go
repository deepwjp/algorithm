package main

//  leetcode.3 最长字串
import (
	"strings"
)

//func main() {
//	s := "qwerqwq"
//	fmt.Println(lengthOfLongestSubstring(s))
//}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start, end := 0, 1
	for i := 1; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}

	return end - start
}
