package main

//1111. 有效括号的嵌套深度
func maxDepthAfterSplit(seq string) (res []int) {
	for key, value := range seq {
		if value == '(' {
			res = append(res, key%2)
		} else {
			res = append(res, 1-key%2)
		}
	}

	return
}
