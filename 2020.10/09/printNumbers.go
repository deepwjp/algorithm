package main

import "math"

//剑指 Offer 17. 打印从1到最大的n位数
func printNumbers(n int) (res []int) {
	max := int(math.Pow(10, float64(n)))
	for i := 1; i < max; i++ {
		res = append(res, i)
	}
	return
}
