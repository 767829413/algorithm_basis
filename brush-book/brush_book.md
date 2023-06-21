# 练习册

## 螺旋矩阵

![1](https://pic.imgdb.cn/item/648bfcd61ddac507ccf4b5e5.png)

```go
func SpiralMatrix(matrix [][]int) []int {
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
```

## 

![2.png](https://s2.loli.net/2023/06/20/bADp8oI6HtmL7zN.png)

```go
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
```