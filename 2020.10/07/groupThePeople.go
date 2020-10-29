package main

//1282. 用户分组
func groupThePeople(groupSizes []int) (res [][]int) {
	maps := make(map[int][]int)

	for key, value := range groupSizes {
		tmp := maps[value]
		tmp = append(tmp, key)
		maps[value] = tmp
	}
	//fmt.Println(maps)
	for key, value := range maps {

		n := len(value)
		if n > key {
			count := n / key
			if n%key == 0 {
				for i := 0; i < count; i++ {
					res = append(res, nil)
					res[len(res)-1] = append(res[len(res)-1], value[key*i:key*(i+1)]...)
				}
			} else {
				for i := 0; i <= count; i++ {
					res = append(res, nil)
					if i == count {
						res[len(res)-1] = append(res[len(res)-1], value[key*i:]...)
					} else {
						res[len(res)-1] = append(res[len(res)-1], value[key*i:key*(i+1)]...)
					}
				}
			}
		} else {
			res = append(res, nil)
			res[len(res)-1] = value
		}
	}
	return
}
