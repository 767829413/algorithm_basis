package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// s, _ := reader.ReadString('\n')
	// n, _ := strconv.Atoi(strings.TrimSpace(s))
	// print(100)
	// print1(1)
	// println(result)

	// println(fib(n))
	// println(loopFib(n))
	// steps = n
	// stairClimbing(0)
	// println(methodNumbers)
	// fmt.Println(permute([]int{0, 1}))
	fmt.Println(permuteUnique([]int{1, 1, 2, 3, 3, 2, 2}))
}

var result = 0

// 递的过程解
func print(n int) {
	if n == 1 {
		println(n)
		result += n
		return
	}
	print(n - 1)
	println(n)
	result += n
}

// 归的过程解
func print1(n int) {
	if n == 100 {
		println(n)
		result += n
		return
	}
	println(n)
	result += n
	print1(n + 1)
}

// 递归求斐波那契数
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 迭代求斐波那契数
func loopFib(n int) int {

	if n <= 2 {
		return 1
	}
	var a, b = 1, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

// 爬楼梯问题:
// 一个楼梯总共有n级台阶，问从第0级台阶走到第n级台阶一共有多少种方案? 你每次可以走1个台阶或者两个台阶
var steps int
var methodNumbers = 0

func stairClimbing(start int) {
	if start == steps {
		methodNumbers++
		return
	}

	if start < steps {
		stairClimbing(start + 1)
		stairClimbing(start + 2)
	}
}

// 全排列问题I: 给定一个不含重复数字的数组 nums，返回其所有可能的全排列。你可以按任意顺序返回答案
// 输入: nums =[1,2,3] 输出:[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 输入: nums = [0,1] 输出:[[0,1],[1,0]]
// 输入: nums = [1] 输出:[[1]]
func permute(nums []int) [][]int {
	var result [][]int
	var tmp = []int{}
	var has = make([]int, len(nums))
	var l = len(nums)
	if l == 0 {
		return result
	}
	var fullyAligned func(nums []int)
	fullyAligned = func(nums []int) {
		if len(tmp) == l {
			tmpCopy := make([]int, len(tmp))
			copy(tmpCopy, tmp)
			result = append(result, tmpCopy)
			return
		}
		for i := 0; i < l; i++ {
			if has[i] == 0 {
				tmp = append(tmp, nums[i])
				has[i] = 1
				fullyAligned(nums)
				tmp = tmp[:len(tmp)-1]
				has[i] = 0
			}
		}
	}
	fullyAligned(nums)
	return result
}

// 全排列问题II: 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列
// 输入：nums = [1,1,2] 输出：[[1,1,2],[1,2,1],[2,1,1]]
// 输入：nums = [1,2,3] 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var tmp = []int{}
	var has = make([]int, len(nums))
	var l = len(nums)
	if l == 0 {
		return result
	}
	var fullyAligned func(deep int)
	fullyAligned = func(deep int) {
		if l == deep {
			tmpCopy := make([]int, len(tmp))
			copy(tmpCopy, tmp)
			result = append(result, tmpCopy)
			return
		}
		for i := 0; i < l; i++ {
			if i >= 1 && has[i-1] == 0 && nums[i-1] == nums[i] {
				continue
			}
			if has[i] != 0 {
				continue
			}
			tmp = append(tmp, nums[i])
			has[i] = 1
			fullyAligned(deep + 1)
			tmp = tmp[:len(tmp)-1]
			has[i] = 0
		}
	}
	fullyAligned(0)
	return result
}
