package brushbook

import (
	"fmt"
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
