# 递归和二分

## 递归

`定义`
 
 ```text
 程序调用自身的编程技巧称为递归 (recursion)。递归作为一种算法在程序设计语言中广泛应用。一个过程或函数在其定义或说明中有直接或间接调用自身的一种方法它通常把一个大型复杂的问题层层转化为一个与原问题相似的规模较小的问题来求解，递归策略只需少量的程序就可描述出解题过程所需要的多次重复计算，大大地减少了程序的代码量。递归的能力在于用有限的语句来定义对象的无限集合。一般来说，递归需要有边界条件、递归前进段和递归返回段。当边界条件不满足时，递归前进，当边界条件满足时，递归返回。
 ```

 **递归并不是一个算法，而是一种程序语言中的一个特性，是对函数或者方法自身的调用.**

`递归的应用条件`

 递归实际上就是在某个方法或者函数运行的过程中调用自己。构成递归需要具备如下条件:

* 子问题必须与原始问题为同样的事，且更为简单。
* 问题解法按递归算法实现
* 数据结构的形式是按照递归定义的。如二叉树等，由于结构本身就有递归的特性，因此树也经常被用于递归形象化的模型

 递归的缺点: 
 
 ```text
 递归解决问题运行效率较低，因此应该尽量避免使用递归，除非没有更好的算法或者某种特定情况，递归更为合适的时候选择递归解决问题。
 在递归调用的过程中系统为每一层的返回点，局部量等开辟了栈来储存，递归次数过多可能会造成栈溢任。
 所以递归解决问题也叫暴力搜索
 ```

`递归的思想内涵`

 正如上面所讲，递归实际上就是有去(递)有回(归)，正如英雄联盟的两句台词:“我随疾风前行。“和“疾风亦有归途。”。如下图所示。

 ![1.png](https://s2.loli.net/2023/06/21/HM5PxnsRrX6OCY1.png)
 
 **“有去”**是指: 递归问题必须可以分解为若干个规模较小，与原问题形式相同的子问题，这些子问题可以用相同的解题思路来解决;
 
 **“有回”**是指:这些问题的演化过程是一个从大到小，由近及远的过程，并且会有一个明确的终点(临界点)，一旦到达了这个临界点，就不用再往更小、更远的地方走下去。最后从这个临界点开始，原路返回到原点，原问题解决。

 更直接地说，递归的基本思想就是把规模大的问题转化为规模小的相似的子问题来解决。特别地，在函数实现时，因为解决大问题的方法和解决小问题的方法往往是同一个方法，所以就产生了函数调用它自身的情况，这也正是递归的定义所在。格外重要的是，这个解决问题的函数必须有明确的结束条件，否则就会导致无限递归的情况。

`用归纳法来理解递归`

 递归在数学上的模型是什么? 观察递归，我们会发现，递归的数学模型其实就是 **数学归纳法** ，这个在高中的数列里面是最常用的了，下面回忆一下数学归纳法。

 数学归纳法适用于将解决的原问题转化为解决它的子问题，而它的子问题又变成子问题的子问题，而且我们发现这些问题其实都是一个模型，也就是说存在相同的逻辑归纳处理项。当然有一个是例外的，也就是归纳结束的那-一个处理方法不适用于我们的归纳处理项，当然也不能适用，否则我们就无穷归纳了。总的来说，归纳法主要包含以下三个关键要素:
 
* 步进表达式:      问题蜕变成子问题的表达式
* 结束条件:        什么时候可以不再使用步进表达式
* 直接求解表达式:  在结束条件下能够直接计算返回值的表达式

 事实上，这也正是某些数学中的数列问题在利用编程的方式去解决时可以使用递归的原因比如著名的斐波那契数列问题。

`递归的三要素`

* 明确递归终止条件;
  * 递归就是有去有回，既然这样，那么必然应该有一个明确的临界点，程序一旦到达了这个临界点，就不用继续往下递去而是开始实实在在的归来。换句话说，该临界点就是一种简单情境，可以防止无限递归。
* 给出递归终止时的处理办法;
  * 在递归的临界点存在一种简单情境，在这种简单情境下，我们应该直接给出问题的解决方案。一般地，在这种情境下，问题的解决方案是直观的、容易的。
* 提取重复的逻辑，缩小问题规模;
  * 递归问题必须可以分解为若干个规模较小、与原问题形式相同的子问题，这些子问题可以用相同的解题思路来解决。从程序实现的角度而言，我们需要抽象出一个干净利落的重复的逻辑，以便使用相同的方式解决子问题。

`递归模板`

* 递的过程中解决问题

```C
function recursion(大问题) {
  if (end_condition) { // 递归终止条件
    end; // 递归终止时的处理方法
  } else {
    solove; // 递
    recursion(小问题); // 递归到最深处,归来
  }
}
```

* 归的过程中解决问题

```C
function recursion(大问题) {
  if (end_condition) { // 递归终止条件
    end;
  } else {
    recursion(小问题); // 递
    solove; // 归
  }
}
```

`递归应用场景`
 
 实际学习工作中，递归算法一般用于解决三类问题:

* 问题的定义是按递归定义的 (Fibonacci函数，阶乘，...) ;
* 问题的解法是递归的(有些问题只能使用递归方法来解决，例如，汉诺塔问题，...);
* 数据结构是递归的(链表、树等的操作，包括树的遍历，树的深度，...)

`递归与循环的区别`

 递归与循环是两种不同的解决问题的典型思路。递归通常很直白地描述了一个问题的求解过程，因此也是最容易被想到解决方式。循环其实和递归具有相同的特性，即做重复任务，但有时使用循环的算法并不会那么清晰地描述解决问题步骤。
 
 单从算法设计上看，递归和循环并无优劣之别。然而，在实际开发中，因为函数调用的开销，递归常常会带来性能问题，特别是在求解规模不确定的情况下;
 而循环因为没有函数调用开销，所以效率会比递归高。递归求解方式和循环求解方式往往可以互换，也就是说，如果用到递归的地方可以很方便使用循环替换，而不影响程序的阅读，那么替换成循环往往是好的。
 
 问题的递归实现转换成非递归实现一般需要两步工作: 

* 自己建立“堆栈(一些局部变量)”来保存这些内容以便代替系统栈，比如树的三种非递归遍历方式;
* 把对递归的调用转变为对循环处理

```go
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
```

## 整数二分

`定义`

 **对于区间[a，b]上连续不断且f(a) • f (b) <0 的函数y=f (x) ，通过不断地把函数f(x)的零点所在的区间一分为二，使区间的两个端点逐步逼近零点，进而得到零点近似值的方法叫二分法。**

`二分模板`

```go
// 第一个模板在区间[L,R],被切割成为[L，MID] 和 [MID + 1，R]
func binary_search_1(left, right int) int {
  for left < right {
    mid := (left + right) >> 1
    if check(mid) {
      right = mid
    }else {
      left = mid + 1
    }
  }
  return left
}

// 第二个模板在区间[L,R]，被切割成[L，MID - 1] 和 [MID，R]
func binary_search_1(left, right int) int {
  for left < right {
    mid := (left + right + 1) >> 1
    if check(mid) {
      left = mid
    } else {
      right = mid - 1
    }
  }
  return left
}
```

`二分本质`
 
 二分的本质其实并不是单调性，如果某个数组存在单调性的话，是一定可以用二分的，但是用二分的题目，不一定非要具有单调性。 

 我们来看这么一个区间:
  
  ![3.png](https://s2.loli.net/2023/06/25/CrJvRbOlHfBaS1D.png)

 **我们有这么个区间，并且我们在区间定义了某种性质。使得这种性质在右半边区间是满足的，在左半边区间是不满足的。**
 
 **假如说我们可以找到一个这样的性质，能把整个区间一分为二，然后一块区间满足某种性质，一块不满足，那么二分就可以寻找这个性质的边界。我们既可以找到红色的这个边界也可以找到蓝色的这个边界。**
 
* 我们先来看这个红色的边界点(如何找红色边界点):

   1. 先找一个中间值: mid = (left + right + 1) / 2
   2. 然后检查一下这个 mid 是否满足红色区间的这个性质
   3. 如果成立: 也就是说满足红色区间的性质，那么我们想找红色边界点的话，我们就可以把 left 更新成 mid, 也就是把 [left, right] 更新成 [mid, right]
      ![4.png](https://s2.loli.net/2023/06/25/jdKygTJ2RkS5cGl.png)
   4. 如果不成立: 那么就是 mid 并不满足红色区间，那么mid就在蓝色区间，那么我们要找红色的边界，right 就可以更新成 mid - 1。mid 一定不是边界点。那么 right 更新成 mid - 1 也就是区间变成了 [left, mid - 1]
      ![5.png](https://s2.loli.net/2023/06/25/5Jr2lQPSkdGYWBm.png)

   ```text
   为什么要加1呢? 是为了防止发生死循环!
   打个比方，我们知道 java c++ python 里除法是向下取整，那么假如说我们当前 left = right - 1 的话，那么 mid = (left + right + 1) / 2 = left ,等于 left，然后我们在检查 mid 恰好满足红色区间的时候，left = mid,此时 mid = left,那么就等于没有更改区间，那么就会一直在 left 到 right 这个区间里，发生死循环。所以要加上1，跳出次循环。
   ```

* 我们再来看这个蓝色的边界点(如何找蓝色边界点):

   1. 先找一个中间值: mid = (left + right) / 2
   2. 然后检查一下这个 mid 是否满足蓝色区间的这个性质
   3. 如果成立: 也就是说满足蓝色区间的性质，那么我们想找蓝色边界点的话，我们就可以把 right 更新成 mid, 也就是把 [left, right] 更新成 [left, mid]
      ![5.png](https://s2.loli.net/2023/06/25/5Jr2lQPSkdGYWBm.png)
   4. 如果不成立: 那么就是 mid 并不满足蓝色区间，那么 mid 就在红色区间，那么我们要找蓝色的边界，left 就可以更新成 mid + 1。mid 一定不是边界点。那么区间变成了 [mid + 1, right]
      ![4.png](https://s2.loli.net/2023/06/25/jdKygTJ2RkS5cGl.png)
