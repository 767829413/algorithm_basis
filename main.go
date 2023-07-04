package main

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var res = []int{}
	var stack = NewStack()
	for !stack.IsEmpty() || root != nil {
		for root != nil {
			// 放入根节点
			res = append(res, root.Val)
			stack.Push(root)
			// 一直找右
			root = root.Right
		}
		// 再寻找左节点
		if !stack.IsEmpty() {
			root = stack.Pop().Left
		}
	}
	// 得到 nrl 的结果
	tmp := []int{}
	for i := len(res) - 1; i >= 0; i-- {
		tmp = append(tmp, res[i])
	}
	return tmp
}

type Stack struct {
	data []*TreeNode
}

// 初始化栈
func NewStack() Stack {
	return Stack{
		data: []*TreeNode{},
	}
}

// 将元素压进栈
func (s *Stack) Push(x *TreeNode) {
	s.data = append(s.data, x)
}

// 将元素弹栈
func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 拿到栈顶元素
func (s *Stack) Top() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// 统计栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}
