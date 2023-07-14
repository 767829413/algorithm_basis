# 排序

```text
 排序是计算机内经常进行的一种操作，其目的是将一组“无序”的记录序列调整为“有序”的记录序列。分内部排序和外部排序，若整个排序过程不需要访问外存便能完成则称此类排序问题为内部排序。反之，若参加排序的记录数量很大，整个序列的排序过程不可能在内存中完成，则称此类排序问题为外部排序。内部排序的过程是一个逐步扩大记录的有序序列长度的过程。

 总而言之: 将杂乱无章的数据元素，通过一定的方法按照关键字顺序排列的过程叫做排序。
```

十大经典排序分别是:
冒泡排序，插入排序，选择排序，希尔排序，计数排序，基数排序，桶排序，快速排序，归并排序和堆排序。

快速排序，堆排序，选择排序和希尔排序是不稳定的排序算法，

基数排序，冒泡排序，插入排序，计数排序，归并排序和桶排序为稳定的排序算法。

## 算法的分类
 
 十种常见排序算法可以分为两大类

 1. 比较类排序: 通过比较来决定元素间的相对次序，由于其时间复杂度不能突破O(nlogn),因此也称为非线性时间比较类排序。
 2. 非比较类排序: 不通过比较来决定元素间的相对次序，它可以突破基于比较排序的时间下界，以线性时间运行，因此也称为线性时间非比较类排序。

   ![排序算法.png](https://s2.loli.net/2023/07/07/4UfYA7L6KaPxshu.png)

## 排序的稳定性

 排序按照稳定性分为稳定的排序算法和不稳定的排序算法。

 1. 稳定排序: 假设在待排序的文件中，存在两个或两个以上的记录具有相同的关键字，在用某种排序法排序后，若这些相同关键字的元素的相对次序仍然不变，则这种排序方法是稳定。
 2. 不稳定排序: 假设在待排序的文件中，存在两个或两个以上的记录具有相同的关键字在用某种排序法排序后，若这些相同关键字的元素的相对次序发生了变化，则这种排序方法是不稳定的

 | 排序方法 | 时间复杂度(平均) | 时间复杂度(最坏) | 时间复杂度(最好) | 空间复杂度 | 稳定性 |
 | :------: | :-------------: | :-------------: | :-------------: |:-------------: | :-------------: |
 | 插入排序 | O(n²) | O(n²) | O(n) | O(1) | 稳定 |
 | 希尔排序 | O(n¹·³) | O(n²) | O(n) | O(1) | 不稳定 |
 | 选择排序 | O(n²) | O(n²) | O(n²) | O(1) | 不稳定 |
 | 堆排序 | O(nlog₂n) | O(nlog₂n) | O(nlog₂n) | O(1) | 不稳定 |
 | 冒泡排序 | O(n²) | O(n²) | O(n) | O(1) | 稳定 |
 | 快速排序 | O(nlog₂n) | O(n²) | O(nlog₂n) | O(1) | 不稳定 |
 | 归并排序 | O(nlog₂n) | O(nlog₂n) | O(nlog₂n) | O(n) | 稳定 |
 |  |  |  |  |  |  |
 |  |  |  |  |  |  |
 | 计数排序 | O(n+k) | O(n+k) | O(n+k) | O(n+k) | 稳定 |
 | 桶排序 | O(n+k) | O(n²) | O(n) | O(n+k) | 稳定 |
 | 插入排序 | O(n\*k) | O(n\*k) | O(n\*k) | O(n+k) | 稳定 |

## 冒泡排序

`定义`

 ```text
 冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素如果它们的顺序错误就把它们交换过来。走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢”浮”到数列的顶端。
 ```

`算法思路`
 
 1. 比较相邻的元素。如果第一个比第二个大，就交换它们两个;
 2. 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数;
 3. 针对所有的元素重复以上的步聚，除了最后一个
 4. 重复1~3步骤,直到排序完成

`代码实现`

```go
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ { // n - 1 循环,最后一个不管
		for j := 0; j < len(arr)-1-i; j++ { // 第二轮指针,每一轮都会排好一个数
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
```

## 选择排序

`定义`

 ```text
 选择排序(Selection-sort)是一种简单直观的排序算法。它的工作原理: 首先在未排序序列中找到最小(大)元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小(大)元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
 ```

`算法思路`
 
 n个记录的直接选择排序可经过n-1趟直接选择排序得到有序结果。具体算法描述如下:

 1. 初始状态: 无序区为R[1..n]，有序区为空;
 2. 第i趟排序(i=1,2,3...n-1)开始时，当1个记录R交换，使 R\[1...i\]和 R[i+1...n)分别变为记录个数增加1个的新有序区和记录个数减少1个的新无序区;
 3. n-1 趟结束,数组有序化

`代码实现`

```go
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}
```

## 插入排序

`定义`

 ```text
 插入排序 (Insertion-Sort) 的算法描述是一种简单直观的排序算法。它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
 ```

`算法思路`
 
 一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下:

 1. 从第一个元素开始，该元素可以认为已经被排序;
 2. 取出下一个元素，在已经排序的元素序列中从后向前扫描;
 3. 如果该元素(已排序) 大于新元素，将该元素移到下一位置;
 4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
 5. 将新元素插入到该位置后;
 6. 重复步骤2~5。

`代码实现`

```go
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		preIdx, cur := i-1, arr[i]
		for ; preIdx >= 0 && arr[preIdx] > cur; preIdx-- {
			arr[preIdx+1] = arr[preIdx]
		}
		arr[preIdx+1] = cur
	}
}
```

## 希尔排序

`定义`

 ```text
 1959年Sshell发明，第一个突破O(n^2)的排序算法，是简单插入排序的改进版。它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫缩小增量排序
 ```

`算法思路`
 
 先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述:

 1. 选择一个增量序列t1，t2，...，tk，其中ti > tj，tk = 1;
 2. 按增量序列个数k，对序列进行 k 趟排序;
 3. 每趟排序，根据对应的增量ti，将待排序列分割成若千长度为 m 的子序列，分别对各子表进行直接插入排序。仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度;

`代码实现`

```go
func ShellSort(arr []int) {
	increase := len(arr) >> 1
	for increase > 0 {
		// 选择排序
		for i := increase; i < len(arr); i++ {
			idx, cur := i, arr[i]
			pre := idx - increase
			for idx >= increase && cur < arr[pre] {
				arr[idx] = arr[pre]
				idx -= increase
			}
			arr[idx] = cur
		}
		increase >>= 1
	}
}
```

## 计数排序

`定义`

 ```text
 计数排序不是基于比较的排序算法，其核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数
 ```

`算法思路`
 
 1. 找出待排序数组中最大和最小的元素;
 2. 统计数组中每个值为i的元素出现的次数，存入数组C的第i项;
 3. 对所有的计数累加 (从数组C的第一个元素开始，每一项和前一项相加);
 4. 反向填充目标数组: 将每个元素i放在新数组的第C(i)项，每一个元素就将C(i)减去1;

`代码实现`

```go
func CountingSort(arr []int, maxNumer int) {
	bucket := make([]int, maxNumer+1)
	for _, v := range arr {
		bucket[v] += 1
	}
	idx := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			arr[idx] = i
			idx++
			bucket[i]--
		}
	}
}
```

## 桶排序

`定义`

 ```text
 桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。桶排序(Bucket sort)的工作的原理: 假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序(有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排)。
 ```

`算法思路`

 1. 设置一个定量的数组当作空桶;
 2. 遍历输入数据，并且把数据一个一个放到对应的桶里去;
 3. 对每个不是空的桶进行排序;
 4. 从不是空的桶里把排好序的数据拼接起来;

`代码实现`

```go
func BucketSort(arr []float64) {
	bucketSize := len(arr)
	buckets := make([][]float64, bucketSize)
	for i := 0; i < bucketSize; i++ {
		bucketIdx := int(float64(bucketSize) * arr[i])
		buckets[bucketIdx] = append(buckets[bucketIdx], arr[i])
	}

	for i := 0; i < bucketSize; i++ {
		sort.Float64s(buckets[i])
	}

	idx := 0
	for i := 0; i < bucketSize; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[idx] = buckets[i][j]
			idx++
		}
	}
}
```

## 基数排序

`定义`

 ```text
 基数排序是按照低位先排序，然后收集;再按照高位排序，然后再收集;依次类推，直到最高位。有时候有些属性是有优先级顺序的，先按低优先级排序，再按高优先级排序。最后的次序就是高优先级高的在前，高优先级相同的低优先级高的在前。
 ```

`算法思路`

 1. 取得数组中的最大数，并取得位数;
 2. arr为原始数组，从最低位开始取每个位组成radix数组;
 3. 对radix进行计数排序(利用计数排序适用于小范围数的特点);

`代码实现`

```go
var MaxDigit = 0
func RadixSort(arr []int) {
	buckets := [10]*Queue{}
	for i := range buckets { // 定义十个桶表示 0~9
		buckets[i] = NewQueue()
	}
	// 取位
	mod, div, l := 10, 1, len(arr)
	for i := 0; i < MaxDigit; i++ {
		for j := 0; j < l; j++ {
			digit := (arr[j] / div) % mod
			buckets[digit].Push(arr[j])
		}
		idx :=0
		for i:=0;i<10;i++ {
			for !buckets[i].IsEmpty(){
				arr[idx] = buckets[i].Pop()
				idx++
			}
		}
		div *= 10
	}
}

type Queue struct {
	data []int
}

// 初始化队列
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// 入队
func (q *Queue) Push(x int) {
	q.data = append(q.data, x)
}

// 从队列中取出队头元素
func (q *Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[0]
}

// 从队列中弹出队头元素
func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool { return len(q.data) == 0 }

// 统计队列的大小
func (q *Queue) Size() int { return len(q.data) }
```

## 快速排序

`定义`

 ```text
 快速排序的基本思想: 通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。
 ```

`算法思路`

 快速排序使用分治法来把一个串 (list) 分为两个子串 (sub-lists) 。具体算法描述如下:

 1. 从数列中挑出一个元素，称为“基准”(pivot);
 2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面(相同的数可以到任一边) 。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区 (partition) 操作;
 3. 递归地 (recursive) 把小于基准值元素的子数列和大于基准值元素的子数列排序;
 
`代码实现`

```go
func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	i, j, x := l, r, arr[(l+r)>>1]
	for i < j {
		for arr[i] < x {
			i++
		}
		for arr[j] > x {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	QuickSort(arr, l, j)
	QuickSort(arr, j+1, r)
}
```

## 归并排序

`定义`

 ```text
 归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法 (Divideand Conquer)的一个非常典型的应用。将已有序的子序列合并，得到完全有序的序列;即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并
 ```

`算法思路`

 归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法 (Divide andConquer) 的一个非常典型的应用。将已有序的子序列合并，得到完全有序的序列;即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。

 1. 把长度为n的输入序列分成两个长度为n/2的子序列;
 2. 对这两个子序列分别采用归并排序;
 3. 将两个排序好的子序列合并成一个最终的排序序列;
 
`代码实现`

```go
func NormalizedSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := ((r - l) >> 1) + l
	NormalizedSort(arr, l, mid)
	NormalizedSort(arr, mid+1, r)
	mergeArr := []int{}
	for i, j := l, mid+1; i <= mid || j <= r; {
		if i > mid {
			for j <= r {
				mergeArr = append(mergeArr, arr[j])
				j++
			}
			break
		}
		if j > r {
			for i <= mid {
				mergeArr = append(mergeArr, arr[i])
				i++
			}
			break
		}
		if arr[i] < arr[j] {
			mergeArr = append(mergeArr, arr[i])
			i++
		} else {
			mergeArr = append(mergeArr, arr[j])
			j++
		}
	}
	for idx, i := 0, l; i <= r; i, idx = i+1, idx+1 {
		arr[i] = mergeArr[idx]
	}
}
```

## 堆排序

`定义`

 ```text
 堆排序(Heapso)是指利用堆这种数结构所设计的一种排序算法，堆积是一个近似完全二又树的结构，并同时满足堆积的性质: 即子结点的键值或索引总是小于等于 (或者大于等于) 它的父节点。
 ```

`堆的基本知识`
 
 用到了数据结构堆。学会堆排序之前首先要知道什么是大顶堆和小顶堆。

 1. 堆是一颗完全二叉树;
 2. 小根堆: 每一个点的元素的大小都是小于等于左右子节点的,所以根节点就是整个堆里最小的值;
 3. 小根堆是数组来存储元素的。数组的第一个位置装的是根节点。如果根节点的是第i个点的话，那么左儿子在数组中的下标位置为2 \* i，右儿子的坐标位置为 2 \* i + 1，*注意我们的下标从1开始，这样就不会有 2 \* 0 = 0 的情况了，比较方便*;
 4. 堆的操作: 
	* 插入一个数
	* 求堆中的最小值
	* 删除最小值
	* 删除任意一个堆中的元素
	* 修改任意一个堆中的元素
 5. ![1.jpg](https://s2.loli.net/2023/07/12/OGtQZMYuiDcs3ho.png)

`堆的操作`

 堆常用的操作有两个，一个是上移操作(up),下移操作 (down)

 1. **下移操作 (down)**
	* ![2.jpg](https://s2.loli.net/2023/07/12/kEVDPJaGMjgzmh2.png) 
 2. **上移操作(up)**
	* ![3.jpg](https://s2.loli.net/2023/07/12/hQlX52G6VboAK4O.png) 
 3. 堆的其他操作
	* 插入一个数 num : 插入一个数就是在我们的数组最后插入一个数，然后在up一下就可以。

	```text
	heap[++ size] = num, up(num);
	```

	* 求堆中的最小值: 数组中第一个元素就是我们的最小值也就是根节点

	```text
	heap[1];
	```

	* 删除最小值: 因为小根堆最麻烦的就是删除最小值，所以我们需要用一些小trick去解决，我们目前这里只有这样解决:
		1. 把堆最后面的一个元素赋值给堆顶元素;
		2. 删除堆最后面的那个元素;
		3. 将新的根节点进行down操作;
	
	```text
	heap[1] = heap[size]，size -- ,down(1);
	```

	* 删除任意一个堆中的元素: 和删除最小值类似，也是让堆的最后的一个元素赋值给当前要删除的元素的节点，然后分情况判断是up还是down。
		1. 如果heap\[x\]的值变大了，说明需要进行down操作;
		2. 如果heap\[x\]的值变小了，说明需要进行up操作;
		3. 任然保持小顶堆规则,不进行操作

	```text
	heap[x] = heap[size],size --，down(x) up(x);
	```

	* 修改任意一个堆中元素为 num: 这个和删除一样,只是不需要进行最后一个元素的赋值

	```text
	heap[x] = num, down(x) up(x);
	```

`算法思路`

 堆排序也是一种排序，跟我们的快速排序，归并排序是一样的，我们每次都输出我们的堆顶元素，然后删除堆顶元素，在输出新的堆顶元素，那么这样我们就可以从小到大输出我们的序列了。

 1. 输出堆顶元素，那也就是输出heap\[1\]就可以了，时间复杂度就是O(1);
 2. 需要删除堆顶元素，需要实现down操作。down操作时间复杂度跟堆层数有关系，时间复杂度为O(logn);
 3. 需要建堆，如何将一个数组变成堆的操作,目前有两种方式:
	* 直接建堆，把每个数组中的元素都down一遍，每一次down的操作是logn的时间复杂度，一共n个数，那么时间复杂度为O(nlogn)。
	* 可以从中间开始建堆: 时间复杂度为O(n)。为什么呢? 来看一个图:
		* ![4.png](https://s2.loli.net/2023/07/12/iWazhJAE9cZKqYd.png)
 
`代码实现`

```go
func HeapSort(arr []int) {
	h := NewHeap(arr)
	idx:=0
	for !h.IsEmpty() {
		arr[idx] = h.Pop()
		idx++
	}
}

type Heap struct {
	data []int
	cap  int
}

func NewHeap(arr []int) *Heap {
	heap := &Heap{
		data: append([]int{0}, arr...),
		cap:  len(arr),
	}
	for i := heap.cap / 2; i > 0; i-- {
		heap.Down(i)
	}
	return heap
}

func (h *Heap) Down(root int) {
	min, childL, childR := root, root*2, root*2+1
	if childL <= h.cap && h.data[childL] < h.data[min] {
		min = childL
	}
	if childR <= h.cap && h.data[childR] < h.data[min] {
		min = childR
	}

	if root != min {
		h.data[root], h.data[min] = h.data[min], h.data[root]
		// 递归下去维护小根堆
		h.Down(min)
	}
}

func (h *Heap) Pop() int {
	if h.IsEmpty() {
		panic("Heap is empty")
	}
	res := h.data[1]
	defer func() {
		h.data[1] = h.data[h.cap]
		h.cap--
		h.Down(1)
	}()
	return res
}

func (h *Heap) IsEmpty() bool {
	return h.cap == 0
}
```