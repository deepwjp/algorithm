package 剑指offer

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：
输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：
输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]

限制：
1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5
*/

type MaxQueue struct {
	q1  []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q1 = append(this.q1, value)
	// 删除max中所有比val小的值
	// 如果新插入的数据val比max中最后一个值last大，则将当前max中最后一个值抹去，因为在弹出last前，val始终在队列中，最大值始终不是last
	for len(this.max) != 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.q1) != 0 {
		n = this.q1[0]
		this.q1 = this.q1[1:]
		if this.max[0] == n {
			this.max = this.max[1:]
		}
	}
	return n
}
