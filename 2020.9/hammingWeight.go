package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	count := 0
	str := fmt.Sprintf("x=%b", num)
	for _, value := range str {
		if value == '1' {
			count++
		}
	}
	return count
}
