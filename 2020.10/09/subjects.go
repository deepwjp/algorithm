package mains

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			//mask右移i位，如果数字i的某一个位置是1，就把数组中对应的数字添加到集合
			if (mask>>i)&1 == 1 {
				set = append(set, v)
			}
		}
		tmp := append([]int(nil), set...)
		ans = append(ans, tmp)
	}
	return
}
