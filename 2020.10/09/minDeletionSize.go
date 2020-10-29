package main

//944. 删列造序
func minDeletionSize(A []string) (res int) {
	lenA, lenS := len(A), len(A[0])
	for j := 0; j < lenS; j++ {
		for i := 1; i < lenA; i++ {
			if A[i-1][j] > A[i][j] {
				res++
				break
			}
		}
	}

	return res
}
