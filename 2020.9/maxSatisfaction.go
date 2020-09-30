package main

//func main() {
//	arr := []int{-1, -8, 0, 5, -7}
//	fmt.Println(maxSatisfaction(arr))
//}

func maxSatisfaction(satisfaction []int) int {
	max := 0
	cur := 0
	for key, value := range satisfaction {
		cur += value * (key + 1)
		if cur > max {
			max = cur
		}
	}
	return max
}
