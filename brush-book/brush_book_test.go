package brushbook

import (
	"fmt"
	"math"
	"testing"
)

func SpiralMatrix1(matrix [][]int) []int {
	res := []int{}
	n := len(matrix)
	if n == 0 {
		return res
	}
	m := len(matrix[0])
	if m == 0 {
		return res
	}
	up, down, left, right := 0, n-1, 0, m-1
	for {
		// 向右走
		for walk := left; walk <= right; walk++ {
			res = append(res, matrix[up][walk])
		}
		up++
		if up > down {
			break
		}
		// 向下走
		for walk := up; walk <= down; walk++ {
			res = append(res, matrix[walk][right])
		}
		right--
		if left > right {
			break
		}
		// 向左走
		for walk := right; walk >= left; walk-- {
			res = append(res, matrix[down][walk])
		}
		down--
		if up > down {
			break
		}
		// 向上走
		for walk := down; walk >= up; walk-- {
			res = append(res, matrix[walk][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

func SpiralMatrix2(matrix [][]int) []int {
	res := []int{}
	n := len(matrix)
	if n == 0 {
		return res
	}
	m := len(matrix[0])
	if m == 0 {
		return res
	}
	// 确定操作 顺时针 右 下 左 上
	dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	// 确定是否取过值
	has := make([][]int, n)
	for i := range has {
		has[i] = make([]int, m)
	}
	total := m * n
	for index, tar, x, y := 0, 0, 0, 0; index < total; index++ {
		res = append(res, matrix[x][y])
		has[x][y] = 1

		a, b := x+dx[tar], y+dy[tar]

		for i := 0; i < 4; i++ {
			if a >= n || a < 0 || b >= m || b < 0 || has[a][b] == 1 {
				tar = (tar + 1) % 4
				a, b = x+dx[tar], y+dy[tar]
			}
		}
		x, y = a, b
	}
	return res
}

func TestSpiralMatrix(t *testing.T) {
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	res := SpiralMatrix2(matrix)
	fmt.Println(res)
}

func rangeBitwiseAnd(left int, right int) int {
	res := 0
	// 0 <= left <= right <= 2^31 - 1
	// 总共31个1,从高位开始取
	for i := 30; i >= 0; i-- {
		if ((left >> i) & 1) != ((right >> i) & 1) {
			break
		}

		if ((left >> i) & 1) == 1 {
			res += 1 << i
		}
	}
	return res
}

func TestBitwiseAndRange(t *testing.T) {
	println(rangeBitwiseAnd(5, 7))             // == 4)
	println(rangeBitwiseAnd(0, 0))             // == 0)
	println(rangeBitwiseAnd(1, math.MaxInt32)) // == 0)
}

// 采用单调栈
func trap1(height []int) int {
	stack := []int{}
	res := 0
	for i := 0; i < len(height); i++ {
		last := 0
		for len(stack) != 0 && height[stack[len(stack)-1]] <= height[i] {
			res += (height[stack[len(stack)-1]] - last) * (i - stack[len(stack)-1] - 1)
			last = height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			res += (height[i] - last) * (i - stack[len(stack)-1] - 1)
		}

		stack = append(stack, i)
	}
	return res
}

// 采用双指针
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	left, right, lm, rm, res := 0, len(height)-1, 0, 0, 0

	for left < right {
		if height[left] > height[right] {
			if rm > height[right] {
				res += rm - height[right]
			} else {
				rm = height[right]
			}
			right--
		} else {
			if lm > height[left] {
				res += lm - height[left]
			} else {
				lm = height[left]
			}
			left++
		}
	}
	return res
}
func TestTrap(t *testing.T) {
	println(trap1([]int{0,1,0,2,1,0,1,3,2,1,2,1}) == 6)
}
