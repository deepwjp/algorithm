package main

import (
	"fmt"
)

func main() {
	res := []int{1, 2}
	fmt.Println(len(res), cap(res))
	maps := make(map[int]int)
	res = append(res, []int{1, 2, 3}...)
	fmt.Println(len(res), cap(res), maps)

}
