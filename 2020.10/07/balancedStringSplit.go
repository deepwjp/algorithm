package main

//1221. 分割平衡字符串
func balancedStringSplit(s string) (res int) {
	countR, countL := 0, 0

	for _, value := range s {
		if value == 'R' {
			countR++
		} else {
			countL++
		}
		if countR == countL {
			res++
			countR, countL = 0, 0
		}
	}
	return
}
