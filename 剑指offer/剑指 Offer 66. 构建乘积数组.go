package 剑指offer

/* 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。



示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]


提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000 */

/*
初始化：数组 BB ，其中 B[0] = 1B[0]=1 ；辅助变量 tmp = 1tmp=1 ；
计算 B[i]B[i] 的 下三角 各元素的乘积，直接乘入 B[i]B[i] ；
计算 B[i]B[i] 的 上三角 各元素的乘积，记为 tmptmp ，并乘入 B[i]B[i] ；
返回 BB 。
*/
func constructArr(a []int) []int {
	if len(a) == 0 {
		return a
	}
	tmp := 1
	b := make([]int, len(a))
	b[0] = 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1] // 下三角
	}
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1] // 上三角
		b[i] *= tmp   // 下三角 * 上三角
	}
	return b
}
