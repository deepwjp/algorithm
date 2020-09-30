package main

import (
	"sort"
)

//func main() {
//	nums1:=[]int{1,2,3}
//	nums2:=[]int{3,4,4}
//	fmt.Println(findMedianSortedArrays(nums1,nums2))
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	all := make([]int,cap(nums1)+cap(nums2))
	all= append(nums1,nums2...)

	sort.Ints(all[:])
	leng:= len(all)
	if leng%2 == 0{

		num:=float64(all[leng/2-1]+all[leng/2])/2.0
		return  float64(num)
	}
	return float64(all[leng/2])
}

