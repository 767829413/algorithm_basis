package main

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	println(mySqrt(8))
	println(mySqrt1(8))
}

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
// 输入：x = 4 输出：2
// 输入：x = 8 输出：2 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
// 因为是向下取整,所以这个平方根肯定在左边区域,如果向上取整难么平方根一定是在右边区域
func mySqrt(x int) int {
	var left, right = 0, x
	for left < right {
		mid := (left + right + 1) / 2
		if mid <= x/mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func mySqrt1(x int) int {
	var left, right = 0, x
	for left < right {
		mid := (left + right) / 2
		if mid <= x/mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
