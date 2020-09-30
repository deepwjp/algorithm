package main

//func main() {
//	strs := []string{"10", "0", "1"}
//	fmt.Println(findMaxForm(strs, 1, 1))
//}

//m->0  n->1
func findMaxForm(strs []string, m int, n int) int {
	maps := make([]map[string]int, 100)
	//falg := []int{}
	//无用
	//strs = sortLen(strs)
	maps = count(strs, maps)
	max := 0
	tempMax := 0
	for key := 0; key < len(strs); key++ {
		if m >= maps[key]["0"] && n >= maps[key]["1"] {
			tempMax++
			m -= maps[key]["0"]
			n -= maps[key]["1"]
			//flag[key] = 1
			if tempMax > max {
				max = tempMax
			}
		} else {
			tempMax--
			//flag[key-1] = 0
		}
	}
	return max
}

func sortLen(strs []string) []string {
	for i := 0; i < len(strs)-1; i++ {
		for j := i + 1; j < len(strs); j++ {
			if len(strs[i]) > len(strs[j]) {
				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}
	return strs
}

func count(str []string, maps []map[string]int) []map[string]int {
	for key, value := range str {
		temp := map[string]int{"0": 0, "1": 0}
		for i := 0; i < len(value); i++ {
			if value[i] == '0' {
				temp["0"]++
			}
		}
		temp["1"] = len(value) - temp["0"]
		maps[key] = temp
	}
	return maps
}
