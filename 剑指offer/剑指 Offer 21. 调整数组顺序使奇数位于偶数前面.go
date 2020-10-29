package 剑指offer

/* 	入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

提示：
1 <= nums.length <= 50000
1 <= nums[i] <= 10000
*/

// func exchange(nums []int) []int {
// 	odd := []int{}
// 	even := []int{}
// 	for _, value := range nums {
// 		if value%2 == 1 {
// 			odd = append(odd, value)
// 		} else {
// 			even = append(even, value)
// 		}
// 	}
// 	return append(odd, even...)
// }

func exchange(nums []int) []int {
	//  双指针，前后遍历
	start, end := 0, len(nums)-1
	for start < end {
		// 正向遍历nums直到nums[start]为偶数，或者start>end
		for start < end && nums[start]%2 == 1 {
			start++
		}
		// 逆向遍历nums直到nums[end]]为奇数，或者start>end
		for start < end && nums[end]%2 == 0 {
			end--
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
	return nums
}
