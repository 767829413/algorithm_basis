package main

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// q := NewQueue()
	// q.Push(1)
	// q.Push(2)
	// println(q.Pop())
	// println(q.Pop())

	s := Constructor()
	s.Push(1)
	s.Push(2)
	println(s.Top())
	println(s.Pop())
	println(s.Empty())

}

type MyStack struct {
	data *Queue
	tmp  *Queue
}

func Constructor() MyStack {
	return MyStack{
		data: NewQueue(),
		tmp:  NewQueue(),
	}
}

func (this *MyStack) Push(x int) {
	this.data.Push(x)
}

func (this *MyStack) Pop() int {
	var v = -1
	for !this.data.IsEmpty() {
		v = this.data.Pop()
		this.tmp.Push(v)
	}
	for !this.tmp.IsEmpty() {
		t := this.tmp.Pop()
		if v != t {
			this.data.Push(t)
		}
	}
	return v
}

func (this *MyStack) Top() int {
	var v = -1
	for !this.data.IsEmpty() {
		v = this.data.Pop()
		this.tmp.Push(v)
	}
	for !this.tmp.IsEmpty() {
		this.data.Push(this.tmp.Pop())
	}
	return v
}

func (this *MyStack) Empty() bool {
	return this.data.IsEmpty()
}

type Queue struct {
	data []int
}

// 初始化队列
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// 入队
func (q *Queue) Push(x int) {
	q.data = append([]int{x}, q.data...)
}

// 从队列中取出队头元素
func (q *Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[len(q.data)-1]
}

// 从队列中弹出队头元素
func (q *Queue) Pop() int {
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
