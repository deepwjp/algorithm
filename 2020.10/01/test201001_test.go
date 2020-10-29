package main

import (
	"fmt"
	"testing"
)

func Test201001(t *testing.T) {
	i := []int{5, 6, 7}
	hello(1)
	fmt.Println(i[0])
}

func hello(num ...int) {
	num[0] = 18
}
