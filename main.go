package main

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

}

func maxSlidingWindow(nums []int, k int) []int {
	q := NewQueue()
	res := []int{}
	for i, idx := 0, 0; i < len(nums); i++ {
		if !q.IsEmpty() && i-k+1 > q.PeekFrist() {
			q.RemoveFrist()
		}
		
	}
	return res
}

type Queue struct {
	data []int
}

// 初始化队列
func NewQueue() Queue {
	return Queue{
		data: []int{},
	}
}

// 从左边入队
func (q *Queue) PushFrist(x int) {
	q.data = append([]int{x}, q.data...)
}

// 从右边入队
func (q *Queue) PushTail(x int) {
	q.data = append(q.data, x)
}

// 获取队头元素
func (q *Queue) PeekFrist() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[0]
}

// 从队头弹出元素
func (q *Queue) RemoveFrist() int {
	if q.IsEmpty() {
		return -1
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

// 获取队头元素
func (q *Queue) PeekTail() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[len(q.data)-1]
}

// 从队尾弹出元素
func (q *Queue) RemoveTail() int {
	if q.IsEmpty() {
		return -1
	}
	v := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	return v
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool { return len(q.data) == 0 }

// 统计队列的大小
func (q *Queue) Size() int { return len(q.data) }
