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

## 按位与

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

## 全排列

![1.png](https://s2.loli.net/2023/06/25/vnq85Sb6B4pLXIo.png)

```go
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
```

## 全排列 II

![2.png](https://s2.loli.net/2023/06/25/PE6iUVOBQ4ZTNnw.png)

```go
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
```
