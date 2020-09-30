package main

//func main() {
//	arr := []int{-2, -1}
//	fmt.Print(maxSubArray(arr))
//}

// 贪心
func maxSubArray(nums []int) int {
	sumPrefix := nums[0] //之前和
	maxSum := nums[0]    //最大和
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		sumPrefix = max(nums[i], sumPrefix+nums[i])
		maxSum = max(maxSum, sumPrefix)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//暴力破解 超时
//func maxSubArray(nums []int) int {
//	var sumMax int = nums[0]
//
//	for i := 0; i < len(nums); i++ {
//
//		for j := i; j < len(nums); j++ {
//			temp := 0
//			for k := i; k <= j; k++ {
//				temp += nums[k]
//				if temp > sumMax {
//					sumMax = temp
//				}
//			}
//
//		}
//	}
//	return sumMax
//}
