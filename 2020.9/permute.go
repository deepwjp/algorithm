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
