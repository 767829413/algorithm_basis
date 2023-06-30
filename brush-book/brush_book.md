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

## 接雨水

![2.png](https://s2.loli.net/2023/06/29/tlEcuh3veasgZPM.png)

```go
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
```

## 滑动窗口最大值

![1.png](https://s2.loli.net/2023/06/29/IA6vzrYRchXBGs7.png)

```go
func maxSlidingWindow(nums []int, k int) []int {
	q := NewQueue()
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if !q.IsEmpty() && i-k+1 > q.PeekFrist() {
			q.RemoveFrist()
		}
		for !q.IsEmpty() && nums[i] > nums[q.PeekTail()] {
			q.RemoveTail()
		}
		q.PushTail(i)
		if i >= k-1 {
			res = append(res, nums[q.PeekFrist()])
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

// 获取队尾元素
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
```