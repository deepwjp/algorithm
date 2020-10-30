package 剑指offer

/* 统计一个数字在排序数组中出现的次数。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

限制：
0 <= 数组长度 <= 50000
*/
func search(nums []int, target int) int {
	// 边界条件判断
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 1
		}
		return 0
	}
	// 统计当前片段内次数
	times := 0
	i := len(nums) / 2
	// 二分查找递归
	if nums[i] > target {
		times = times + search(nums[:i], target)
	} else if nums[i] < target {
		times = times + search(nums[i:], target)
	} else {
		for _, value := range nums {
			if value == target {
				times++
			}
		}
	}
	return times
}
