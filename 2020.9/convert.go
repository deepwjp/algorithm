package main

import (
	"fmt"
	"strings"
)

//func main() {
//	fmt.Println(convert("LEETCODEISHIRING", 4))
//}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	b := numRows - 1
	strArr := make([]string, numRows)
	for k, v := range s {
		if (k/b)%2 == 0 {
			fmt.Println("a", k, b, strArr[k%b], string(v))
			strArr[k%b] += string(v)
		} else {
			strArr[b-(k%b)] += string(v)
			fmt.Println("b", k, b, strArr[k%b], string(v))
		}
	}
	return strings.Join(strArr, "")
}
