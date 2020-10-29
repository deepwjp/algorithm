package main

import (
	"strconv"
	"strings"
)

func minNumber(nums []int) (res string) {
	var str []string
	for key, value := range nums {
		str = append(str, "")
		str[key] = string(strconv.Itoa(value))
	}
	//sort.Strings(str)
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str)-1; j++ {
			//var tmpStr []string
			//copy(tmpStr,str)
			tmp1 := strings.Join([]string{str[j], str[j+1]}, "")
			tmp2 := strings.Join([]string{str[j+1], str[j]}, "")
			if tmp1 > tmp2 {
				str[j], str[j+1] = str[j+1], str[j]
			}
		}
	}
	res = strings.Join(str, "")
	return
}
