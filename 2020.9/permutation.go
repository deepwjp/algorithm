package main

//func main() {
//	str := "12345678"
//	fmt.Println(permutation(str))
//}

func permutation(S string) []string {
	var res []string
	dfsPermutation(S, &res, "")
	return res
}

func dfsPermutation(str string, res *[]string, path string) {
	if str == "" {
		*res = append(*res, path)
	}

	for i := 0; i < len(str); i++ {
		temp := str[:]
		cur := str[i]
		if i != len(str)-1 {
			temp = temp[:i] + temp[i+1:]
		} else {
			temp = temp[:len(temp)-1]
		}
		dfsPermutation(temp, res, path+string(cur))
	}
}
