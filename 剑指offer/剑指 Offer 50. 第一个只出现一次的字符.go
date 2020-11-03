package 剑指offer

/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:
s = "abaccdeff"
返回 "b"
s = ""
返回 " "

限制：
0 <= s 的长度 <= 50000
*/
func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	list := make([]int, 26)
	for _, value := range s {
		list[value-'a']++
	}
	for i := 0; i < len(s); i++ {
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
