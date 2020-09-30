package main

//全排列
func permute(nums []int) [][]int {
	res := [][]int{}
	dfsPermute(nums, &res, []int{})
	return res
}

func dfsPermute(nums []int, res *[][]int, path []int) {
	if len(nums) == 0 {
		*res = append(*res, path)
	}
	for i := 0; i < len(nums); i++ {
		var temp []int
		temp = append(temp, nums...)
		cur := nums[i]
		if i != len(nums)-1 {
			temp = append(temp[:i], temp[i+1:]...)
		} else {
			temp = temp[:len(temp)-1]
		}
		dfsPermute(temp, res, append(path, cur))
	}
}

//
//func permute(nums []int) (res [][]int) {
//	if len(nums) == 1 {
//		return [][]int{nums}
//	}
//	for i, num := range nums {
//		// 把num从 nums 拿出去 得到tmp
//		tmp := make([]int, len(nums)-1)
//		copy(tmp[0:], nums[0:i])
//		copy(tmp[i:], nums[i+1:])
//		// sub 是把num 拿出去后，数组中剩余数据的全排列
//		sub := permute(tmp)
//		for _, s := range sub {
//			res = append(res, append(s, num))
//		}
//	}
//	return
//}
