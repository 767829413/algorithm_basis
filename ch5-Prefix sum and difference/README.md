# 前缀和与差分

## 一维前缀和

`定义和基本思路`

```text
假设当前我们有一个数组 有元素 a1 a2 a3 .... ai
那么 si 就表示从 a1 一直加到 ai 的和.
```

 **前缀和中下标一定从1开始**
    
![1.jpg](https://s2.loli.net/2023/07/19/R58yrx3aPbJvGqM.png)

1. 前缀和的作用就是可以快速的求出来数组中一段数的和。比如想求一下原数组中 [l,r] 也就是从 al 一直加到 ar
2. 如果不用前缀和的话，我们可以设 si 有: Si= a1 + a2 + ... + ai

假设我们想求L~R此区间的和。R > L

那么，我们的SL-1则有: SL-1 = a1 + a2 + ... + aL-1

SR则有: SR = a1 + a2 + a3 + ... + aL-1 + aL + ... + aR

那么 SR - SL-1 则为: SR - SL-1 = a1 + a2 + a3 + ... aL-1 + aL + ... + aR - (a1 + a2 + a3 + ... aL-1)

最后的结果不难发现就是我们的 aL + aR 的和。 sum[L,R] = aL + aL+1 + aL+2 + .... +aR = SR-SL-1

我们为了方便起见，将s0设置为0，下标从1开始,那么如何快速的求出sl 和 sr呢? 不难看出我们有如下的递推公式?

 数组下表从1开始: si = s[i - 1] + a[i]  
 数组下标从0开始: s[i]= s[i - 1] + a[i-1]

`模板和代码实现`

```go
var reader = bufio.NewReader(os.Stdin)
func main() {
	fmt.Print("请输入数组值：")
	inputArr, _ := reader.ReadString('\n')
	arrStr := strings.Split(strings.TrimSpace(inputArr), " ")
	fmt.Print("请输入左右边界：")
	inputLimit, _ := reader.ReadString('\n')
	limitStr := strings.Split(strings.TrimSpace(inputLimit), " ")
	var allInt = func(arr []string) []int {
		result := make([]int, len(arr))
		for i := 0; i < len(arr); i++ {
			result[i], _ = strconv.Atoi(arr[i])
		}
		return result
	}
	arr := allInt(arrStr)
	limit := allInt(limitStr)

	if limit[0] > limit[1] {
		limit[0], limit[1] = limit[1], limit[0]
	}
	if limit[0] == 0 {
		limit[0] = 1
	}
	s := make([]int, len(arr))
	s[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		s[i] = s[i-1] + arr[i]
	}
	fmt.Println(limit)
	fmt.Println(arr)
	fmt.Println(s)
	fmt.Printf("L: %d -> R: %d => SUM: %d", limit[0], limit[1], s[limit[1]]-s[limit[0]-1])
}
```

`相关题目`

 1. ![1](https://pic.imgdb.cn/item/64b8a3a31ddac507cc6416db.png)

```go
func subarraySum(nums []int, k int) int {
	s, m, result := make([]int, len(nums)+1), make(map[int]int), 0
	m[0] = 1
	for i := 1; i <= len(nums); i++ {
		s[i] = s[i-1] + nums[i-1]
	}
	for i := 1; i <= len(nums); i++ {
		result += m[s[i]-k]
		m[s[i]]++
	}
	return result
}
```

## 二维前缀和

`定义和基本思路`

 ```text
 二维前缀和就是二维数组里一小块区域的和.
 二维前缀和可以快速的求出二维数组中一块范围数组元素的和。
 ```

 ![1.png](https://s2.loli.net/2023/07/25/nSa2mPwRlTbAt41.png)

 **s[x2][y2] - s[x1-1][y2] - s[x2][y1-1] + s[x1-1][y1-1]**

首先我们知道求一维前缀和S[i]的公式为S[i] = S[i - 1] + a[i]。

二维数组其实也是一样的，公式为: 是S[i,j] = a[i,j] + S[i - 1, j] + S[i, j - 1] - S[i - 1]\[j - 1]

递推的思路，就可以求出从下标 (1，1)到(i，j) 的矩阵和了。

蓝色区域的元素和求解思路:

```text
S[蓝色区域] = S[x2, y2] - S[x1 - 1][y2] - S[x2][y1 - 1] + s[x1 - 1][y1 - 1]
首先: 每个S代表左上角到右下角的和
S[x2,y2] 就是左上角(x2, y2)这个坐标的和.
S[x1-1,y2] 就是从左上角到(x1-1,y2)这个坐标的和
S[x2,y1-1] 就是从左上角到(x2,y1-1)这个坐标的和
S[x1-1,y1-1] 就是从左上角到(x1-1,y1-1)这个坐标的和
因为 (x1 - 1, y1 - 1) 被减了两次,所以加回来一次,所得为二维前缀和
```

求S数组的思路:

 ![2.png](https://s2.loli.net/2023/07/25/b5gzmWIOt4Sw6XK.png)

 S[i,j] = S[i-1,j] + S[i,j-1] - S[i-1,j-1] + Value(i,j)

`模板和代码实现`

```go
func main() {
	arrs := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	s := make([][]int, len(arrs))
	for k := range s {
		s[k] = make([]int, len(arrs[0]))
	}
	for i := 1; i < len(arrs); i++ {
		for j := 1; j < len(arrs[0]); j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + arrs[i][j]
		}
	}
	res := s[len(arrs)-1][len(arrs[0])-1] - s[len(arrs)-1][0] - s[0][len(arrs[0])-1] + s[0][0]
	fmt.Println(s)
	fmt.Println(res)
	fmt.Println(7 + 8 + 9 + 10 + 12 + 13 + 14 + 15 + 17 + 18 + 19 + 20)
}
```

## 一维差分

`定义和基本思路`

 一维差分就是一维前缀和的逆运算

 ```text
 假设 我们有 元素 a1 a2 a3.... an
 然后我们需要构造一个数组 b : b1 b2 b3 ... bn
 使得an = b1 + b2 + b3 + ... + bn
 构造完成之后 使得 a数组就是b数组的前缀和
 一维构造很简单:
 也就是 
 b1 = a1，
 b2 = a2 - a1，
 b3 = a3 - a2, 
 bn = an - an-1
 然后有:
 b2 + b1 = a1 + a2 - a1 = a2
 b3 + b2 + b1 = a1 + a2  a1 + a3 - a2 = a3
 bn + bn - 1 + ... b1 = an
 ```

 如果我们想求原数组的话，我们对b数组求出一遍前缀和 就可以求出a数组了。 o(n)。

 打个比方 比如我们现在给定一段a数组的区间aL 到 aR

 ![3.png](https://s2.loli.net/2023/07/25/Mt8cFvsYWwAeqbG.png)

 不考虑算法问题，我们直接做的话，就是循环一遍数组，然后每个数加个nums，我们时间复杂度就是O(n)。

 那我们如果用到差分的情况呢，那我们的时间复杂度就可以降到O(1).

 那如果我们在想求原来的a数组的话，我们只需要扫描一遍b数组，就可以得到原来的a数组了。

 回到我们的a数组，假设说我们现在要让我们a数组 L~ R这段区间 都加上nums的话

 **我们只需要让bL + nums。**

 那么aL ~aR 这段区间的数都可以加上一个C 因为我们可知 aL = b1+ b2 + bL。那么aL+ nums 就是 b1+ b2 + b3 + ... + bL + nums.

 那么 aL + 1 +nums 也有 b1 + b2 + b3 + ... + bL + nums + bL + 1。

 以此类推 那么aR + nums 也会是 b1 + b2 + b3 + ... bL + nums + bL + 1 + ... + bR

 但是我们有一个问题，因为 我们 bL + nums 会导致 aR +1的数也会加一个nums，所以为了不让 aR之后的数字 也加上一个nums，我们需要让bR + 1 - nums。

 ![4](https://pic.imgdb.cn/item/64bf438b1ddac507cc6b91ec.png)

`模板和代码实现`

```go
func main() {
	arr := append([]int{0}, utilcom.BuildTestArrInt(5, 10, 0)...) // [0 6 0 6 2 6]
	nums, l, r := 1, 1, len(arr)-2
	// 构建差分b数组
	b := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		b[i] = arr[i] - arr[i-1]
	}

	b[l] += nums
	b[r+1] -= nums

	fmt.Println(arr)
	for i := 1; i < len(b); i++ {
		b[i] += b[i-1]
		fmt.Println(b[i])
	}
}
```

## 二维差分

`定义和基本思路`

 ![5](https://pic.imgdb.cn/item/64bf635f1ddac507cca844af.png)

`模板和代码实现`
 
 ![6](https://pic.imgdb.cn/item/64bf64fe1ddac507ccab9e98.png)

 回想一下我们之前讲过的二维前缀和和一维差分，我们如果想让矩阵a的一个子矩阵都加上一个数字nums，那应该怎么做呢?

 核心差分操作:

 给矩阵a以左上角a (x1,y1)，右下角为a (x2,y2) 的子矩阵每个数都加上一个数字nums的话，对差分矩阵的影响是:
 
 ```text
 b[x1, y1] += nums
 b[x1, y2 + 1] -= nums
 b[x2 + 1, y1] -= nums
 b[x2 + 1, y2 + 1] += nums
 ```

 如何构造b[i,j]的数组呢?

 在讲一维差分时，我们是直接利用公式进行构造差分数组，在这里给大家进一个更加巧妙地方式帮助大家去更好的理解如何去构造差分数组:

 **刚开始时候数组a，b全0，此时默认b是a数组的差分数组，而数组a插入a\[i][j]，要想维持数组b仍是数组a的差分数组，那么就要使得b\[i][j]处也加上a\[i][j]，这里的a\[i][j]可以理解为后面的nums。**

 如果不理解的话，我们返回来看一下一维差分:

 **如果将a数组中[l,r]部分的数据全部加上nums看作一次插入操作，那么构造的过程可以看作是将a进行了n次插入操作。第一次在[1,1]的范围插入了a[1],第二次在[2,2]范围插入a[2]，第三次在[3,3]范围插入a[3].....，以此类推，进行n次插入后，那么数组a就正好是数组b的前缀和了。**

 如果还不理解，我们也可以用公式:

 ```text
 b[i][j] = a[i][j] - a[i-1][j] - a[i][j-1] + a[i-1][j-1]

 最后求前缀和就是我们的二位前缀和公式

a[i][j] = b[i-1][j] + b[i][j-1] - b[i-1][j-1] + b[i][j]
 ```
 
`相关题目`
输入一个n行m列的整数矩阵，再输入q个操作，每个操作包含五个整数x1,y1,x2,y2,nums，其中(x1,y1)和(x2,y2)表示一个子矩阵的左上角坐标和右下角坐标。 最后返回结果

```go
func main() {
	var max, min = 100, 0
	var a, b, n, m, q, nums = [8][8]int{}, [8][8]int{}, 5, 5, 1, 3
	var x1, y1, x2, y2 = 1, 1, n - 1, m - 1
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] = rand.Intn(max-min) + min
		}
	}

	// 构造差分数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// insert(i, j, i, j, a[i][j], &b)
			b[i][j] = a[i][j] - a[i-1][j] - a[i][j-1] + a[i-1][j-1]
		}
	}

	for ; q > 0; q-- {
		insert(x1, y1, x2, y2, nums, &b)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[i][j] = b[i-1][j] + b[i][j-1] - b[i-1][j-1] + b[i][j]
		}
	}
	fmt.Println(a)
	fmt.Println(b)

}

func insert(x1, y1, x2, y2, nums int, b *[8][8]int) {
	// b[i][j] = a[i][j] - a[i-1][j] - a[i][j-1] + a[i-1][j-1]
	b[x1][y1] += nums
	b[x1][y2+1] -= nums
	b[x2+1][y1] -= nums
	b[x2+1][y2+1] += nums
}
```