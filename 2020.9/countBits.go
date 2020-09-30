package main

//func main() {
//	fmt.Println(countBits(2))
//}

func countBits(num int) []int {
	arr := make([]int, num+1)

	arr[0] = 0
	if num == 0 {
		return arr
	}
	arr[1] = 1
	if num == 1 {
		return arr
	}

	for i := 2; i <= num; i++ {
		//当前数与前一个数相与，的数的个数+1
		temp := i & (i - 1)
		arr[i] = arr[temp] + 1
	}
	return arr
}
