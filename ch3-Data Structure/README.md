# 数据结构

 ```text
 数据结构(data structure)是带有结构特性的数据元素集合，它研究的是数据的逻辑结构和数据的物理结构以及它们之间的相互关系，并对这种结构定义相适应的运算，设计出相应的算法，并确保经过这些运算以后所得到的新结构仍保持原来的结构类型。简而言之，数据结构是相互之间存在一种或多种特定关系的数据元素的集合，即带“结构”的数据元素的集合。“结构”就是指数据元素之间存在的关系，分为逻辑结构和存储结构。
 ```

## 数组

`定义`

```text
数组(Array)是有序的元素序列。若将有限个类型相同的变量的集合命名,那么这个名称为数组名。组成数组的各个变最称为数组的分量,也称为数组的元素,有时也称为下标变量,用于区分数组的各个元素的数字编号称为下标。数组是在程序设计中,为了处理方便,把具有相同类型的若干元素按有序的形式组织起来的一种形式。 这些有序排列的同类数据元素的集合称为数组。

数组是用于储存多个相同类型数据的集合。
```

`数组的数据类型`

 **数组元素并非只能是基本数据类型,还可以是枚举,结构体和类。**

`数组的结构形式`

 1. **栈内存**: *在方法中定义的一些基本类型的变量和对象的引用变量都在方法的栈内存中分配,当在一段代码中定义一个变量时,在栈内存中为这个变量分配内存空间,当超出变量的作用域后,会自动释放掉为该变量所分配的内存空间。*

 2. **堆内存**: *在堆中分配的内存,在堆中创建了一个数组或对象后,同时还在栈内存中定义一个特殊的变量。让栈内存中的这个变量的取值等于数组或者对象在堆内存中的首地址,栈中的这个变量就成了数组或对象的引用变量,引用变量实际上保存的是数组或对象在堆内存中的地址(也称为对象的句柄),以后就可以在程序中使用栈的引用变量来访问堆中的数组或对象。*

`一维数组`

 *一维数组是最简单的数组,其逻辑结构是线性表。要使用一维数组,需经过定义、初始化和应用等过程*

`二维数组`

 *前面介绍的数组只有一个下标,称为一维数组,其数组元素也称为单下标变量,在实际问题中有很多量是二维的或多维的,因此C语言允许构造多维数组,多维数组元素有多个下标,以标识它在数组中的位置,所以也称为多下标变量。多维数组可由二维数组类推而得到,*

## 链表

`链表定义`

 **链表是一种物理存储单元上非连续、非顺序的存储结构,数据元素的逻顺序是通过链表中的指针链接次序实现的。链表由一系列结点(链表中每一个元素称为结点)组成,结点可以在运行时动态生成。每个结点包括两个部分: 一个是存储数据元索的数据域,另一个是存储下一个结点地址的指针域。相比于线性表(比如数组),接作复杂。由于不用按顺序存储,链表在插入的时候可以达到 O(1) 的复杂度,比另一种线性表顺序表快得多,但是查找一个节点或者访问特定编号的节点则需要O(n)的时间,而线性表和顺序表相应的时间复杂度分别是 O(logn) 和 O(1)。**

  ![1.png](https://s2.loli.net/2023/06/27/qGfUgvnw3AhyJ5o.png)

`动态链表`
 
 **动态链表定义**

  *这种链表在初始时不一定分配足够的空间,但是在后续插入的时候需要动态申请存储空间,并且存储空间不一定连续,在进行插入和删除时则不需要移动元素。修改指针城即可,所以仍然具有链表的主要优点,链表结构可以是动态地分配存储的,即在需要时才开辟结点的存储空间,实现动态链接。说的直白一点:也就是我们在工程上使用的链表结构 => 指针+结构体的链表,我们需要创建节点的时候,才会创建,也就是我们经常使用的链表。*

`单链表`

 **单链表定义**

  *单链表是一种链式存取的教据结构,用一组地址任意的存储单元存放线性表中的教据元素。链表中的教据是以结点来表示的,每个结点的构成:元素(教据元素的映象)+指针(指示后继元素存储位置),元素就是存储数据的存储单元,指针就是连接每个结点的地址数据。*

 **单链表相关操作**

```go
type Node struct {
	Value int
	Next  *Node
}

type MyLinkedList struct {
	Head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	s, i := this.Head, 0
	
	for i < index && s != nil {
		s, i = s.Next, i+1
	}
	if s == nil {
		return -1
	}
	return s.Value
}

func (this *MyLinkedList) AddAtHead(val int) {
	new := &Node{
		Value: val,
		Next:  this.Head,
	}
	this.Head = new
}

func (this *MyLinkedList) AddAtTail(val int) {
	new := &Node{
		Value: val,
		Next:  nil,
	}
	s := this.Head
	if s == nil {
		this.Head = new
		return
	}
	for s.Next != nil {
		s = s.Next
	}
	s.Next = new
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	new := &Node{
		Value: val,
		Next:  nil,
	}

	s := this.Head
	for i := 0; i < index-1; i++ {
		if s == nil {
			return
		}
		s = s.Next
	}
	if s == nil {
		return
	}
	new.Next = s.Next
	s.Next = new
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.Head == nil {
			return
		}
		this.Head = this.Head.Next
		return
	}
	s := this.Head
	for i := 0; i < index-1; i++ {
		if s == nil || s.Next == nil {
			return
		}
		s = s.Next
	}
	if s.Next == nil {
		return
	}
	s.Next = s.Next.Next
}
```

`双向链表`
 
 **双向链表定义**

 ```text
 双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向直接后继和直接前驱。所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前驱结点和后继结点。一般我们都构造双向循环链表。
 ```
 
 **双向链表相关操作**
 
 *用双链表来实现一下单链表的那些方法,双链表比单链表快得多，测试用例花费的时间比单链表快了两倍。但是它更加复杂，它包含了 size，记录链表元素个数，和伪头伪尾。这里我们就不设置头结点和尾结点为null了，可以初始化成dummy虚拟头结点和尾巴节点这样可以先把二者关联起来。*

```go
type node struct {
	Value int
	Next  *node
	Prev  *node
}

type MyLinkedList struct {
	Size      int
	DummyHead *node
	DummyTail *node
}

func Constructor() MyLinkedList {
	m := MyLinkedList{
		Size:      0,
		DummyHead: &node{Value: -1},
		DummyTail: &node{Value: -1},
	}
	m.DummyHead.Next = m.DummyTail
	m.DummyTail.Prev = m.DummyHead
	return m
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	var cur *node
	if index+1 < this.Size-index {
		cur = this.DummyHead
		for i := 0; i < index+1; i++ {
			cur = cur.Next
		}
	} else {
		cur = this.DummyTail
		for i := 0; i < this.Size-index; i++ {
			cur = cur.Prev
		}
	}
	return cur.Value
}

func (this *MyLinkedList) AddAtHead(val int) {
	new := &node{Value: val}
	frist, second := this.DummyHead, this.DummyHead.Next
	new.Prev, new.Next = frist, second
	frist.Next, second.Prev = new, new
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	new := &node{Value: val}
	frist, second := this.DummyTail.Prev, this.DummyTail
	new.Prev, new.Next = frist, second
	frist.Next, second.Prev = new, new
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	}
	if index > this.Size {
		return
	}
	new := &node{Value: val}
	var frist, second *node
	if index+1 < this.Size-index {
		frist = this.DummyHead
		for i := 0; i < index; i++ {
			frist = frist.Next
		}
		second = frist.Next
	} else {
		second = this.DummyTail
		for i := 0; i < this.Size-index; i++ {
			second = second.Prev
		}
		frist = second.Prev
	}
	new.Next = second
	new.Prev = frist
	frist.Next = new
	second.Prev = new
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	var frist, second *node
	if index+1 < this.Size-index {
		frist = this.DummyHead
		for i := 0; i < index; i++ {
			frist = frist.Next
		}
		second = frist.Next.Next
	} else {
		second = this.DummyTail
		for i := 0; i < this.Size-index-1; i++ {
			second = second.Prev
		}
		frist = second.Prev.Prev
	}
	frist.Next = second
	second.Prev = frist
	this.Size--
}
```

## 栈

`栈的定义`

 ```text
 栈 (stack) 又名堆栈，它是一种运算受限的线性表。限定仅在表尾进行插入和删除操作的线性表。这一端被称为栈顶，相对地，把另一端称为栈底。向一个栈插入新元素又称作进栈、入栈或压栈，它是把新元素放到栈顶元素的上面，使之成为新的栈顶元素;从一个栈删除元素又称作出栈或退栈，它是把栈顶元素删除掉，使其相邻的元素成为新的栈顶元素。

 栈作为一种数据结构，是一种只能在一端进行插入和删除操作的特殊线性表。它按照后进先出的原则存储数据，先进入的数据被压入栈底，最后的数据在栈顶，需要读数据的时候从栈顶开始弹出数据(最后一个数据被第一个读出来)。栈具有记忆作用，对栈的插入与删除操作中，不需要改变栈底指针。

 栈是允许在同一端进行插入和删除操作的特殊线性表。允许进行插入和删除操作的一端称为栈顶(top)，另一端为栈底(bottom); 底固定，而栈顶浮动;栈中元素个数为零时称为空栈。插入一般称为进栈 (PUSH) ，删除则称为退栈 (POP) 。栈也称为先进后出表
 
 栈可以用来在函数调用的时候存储断点，做递归时要用到栈!
 ```

 栈的结构是先进后出，类似于我们玩的汉诺塔，计算机结构如图:

  ![1.png](https://s2.loli.net/2023/06/29/wbVu8GTIEJ1NQr7.png)

`栈的基本操作和实现`

```go
type Stack struct {
	data []int
}

// 初始化栈
func NewStack() Stack {
	return Stack{
		data: []int{},
	}
}

// 将元素压进栈
func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

// 将元素弹栈
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 拿到栈顶元素
func (s *Stack) Top() int {
	if s.IsEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}

// 统计栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}
```

`单调栈`
 
 ```text
 单调栈是一个概念，只不过存的元素的顺序是单调递增或者是递减的(或者说是具有单调性的)。
 ```

 **单调栈问题: 接雨水**

  ![2.png](https://s2.loli.net/2023/06/29/tlEcuh3veasgZPM.png)

 分析:

  ![3.png](https://s2.loli.net/2023/06/29/2ZPuA5wd46DfXj7.png)

 ```text
 假设这是我们要算的雨水块1，2，3
 我们设定: 单调栈的单调性是单调递减的，栈中储存的是每个高度块的下标.
 依次从左往右存储高度块，看看是否形成雨水块。

 加到下标为4的高度块的时候单调递减栈开始不满足条件，开始弹伐，把弹出去的上一个栈顶元素记为last。
 雨水块1的高度 height = 当前栈顶元素的高度 - 上一个栈顶元素的高度
 雨水块1的宽度 width = 当前要加入的元素到栈顶元素的距离-1。
 雨水块1的面积 = 高 x 宽

 不难发现雨水块2和1的求法是很像的,再来看一下雨水块3。这个用1和2的办法就不行了，就不是我们想要的高度了。
 雨水块3的高度 height = 当前要加入的元素的高度 - 上一个栈顶元素的高度。
 雨水块3的宽度 width = 当前加入的元素到栈顶元素的距离 - 1。
 雨水块3的面积= 高 x 宽

 然后再把 雨水块123 的面积加起来就好了
 ```

```go
// 采用单调栈
func trap(height []int) int {
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
```

## 队列

`队列定义`

```text
队列是一种特殊的线性表，特殊之处在于它只允许在表的前端 (front) 进行删除操作，而在表的后端(rear) 进行插入操作，和栈一样，队列是一种操作受限制的线性表。进行插入操作的端称为队尾，进行删除操作的端称为队头。

队列中没有元素时，称为空队列。队列的数据元素又称为队列元素。在队列中插入一个队列元素称为入队，从队列中删除一个队列元素称为出队。因为队列只允许在一端插入，在另一端删除，所以只有最早进入队列的元素才能最先从队列中删除，故队列又称为先进先出 (FIFO-first in first out) 线性表。
```

 ![4.png](https://s2.loli.net/2023/06/29/8Au9kwmgP3nHNMl.png)

`队列的基本操作和实现`

```go
type Queue struct {
	data []int
}

// 初始化队列
func NewQueue() Queue {
	return Queue{
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

`双端队列的基本操作和实现`

 *deque (double-ended queue，双端队列) 是一种具有队列和栈的性质的数据结构。双端队列中的元素可以从两端弹出，其限定插入和删除操作在表的两端进行。*

```go
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

`单调队列`

 *单调队列也和单调栈的原理一模一样，只不过这个队列里存的元素是具有单调性的。单调队列最经典的应用就是著名的《滑动窗口》问题*

 **滑动窗口:单调队列解决**

  ![1.png](https://s2.loli.net/2023/06/29/IA6vzrYRchXBGs7.png)

 分析:

 ```text
 假设窗口这样设置: 刚开始无任何数放入窗口里:
 |* * *| 1 3 -1 -3 5 3 6 7
 然后窗口向右移动,开始装数
 |* * 1| 3 -1 -3 5 3 6 7
 数字没有装满前,是不会计算最大值的
 |* 1 3| -1 -3 5 3 6 7
 装到了k个,发现最大值为3,装入数组
 |1 3 -1| -3 5 3 6 7
 后面一样,最终可以得到 3 3 5 5 6 7
 ```

 如果不采用单调队列的话,时间复杂度根据循环最外层和每次循环窗口得出 O(n*k)

 考虑一下如何用单调队列的思想来优化一下:
 
 **假设我们每次从队尾加进来的数字都要比右边也就是队尾元素大的话我们就可以把队尾元素删掉，一直删直到我们要加的元素小于队尾元素。**
 
 那么我们就可以确保我们的对头元素一定是最大的。
 
 每次窗口从对头出，从队尾插或者除，那么我们前后都需要操作，这时候我们可以用到双端队列，双端队列和普通队列的区别就是我们对队头和队尾的元素
 
 都可以进行操作。
  
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

`用栈实现队列`

 ![1.png](https://s2.loli.net/2023/06/30/S8kde9BUOgvzAEx.png)

 分析: 
 
 两个栈,pop的时候一个出一个进,然后返回,peek同理

```go
type MyQueue struct {
	data *Stack
	tmp  *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		data: NewStack(),
		tmp:  NewStack(),
	}

}

func (this *MyQueue) Push(x int) {
	this.data.Push(x)
}

func (this *MyQueue) Pop() int {
	for !this.data.IsEmpty() {
		this.tmp.Push(this.data.Pop())
	}
	v := this.tmp.Pop()
	for !this.tmp.IsEmpty() {
		this.data.Push(this.tmp.Pop())
	}
	return v
}

func (this *MyQueue) Peek() int {
	for !this.data.IsEmpty() {
		this.tmp.Push(this.data.Pop())
	}
	v := this.tmp.Top()
	for !this.tmp.IsEmpty() {
		this.data.Push(this.tmp.Pop())
	}
	return v
}

func (this *MyQueue) Empty() bool {
	return this.data.IsEmpty()
}

type Stack struct {
	data []int
}

// 初始化栈
func NewStack() *Stack {
	return &Stack{
		data: []int{},
	}
}

// 将元素压进栈
func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

// 将元素弹栈
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

// 拿到栈顶元素
func (s *Stack) Top() int {
	if s.IsEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 统计栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}
```

`用队列实现栈`

 ![3.png](https://s2.loli.net/2023/06/30/XUcg81lEs6pAxuL.png)

 分析:

 两个队列,pop的时候一个出一个进,然后返回,peek同理,记住pop的时候要比对拿出的

```go
type MyStack struct {
	data *Queue
	tmp  *Queue
}

func Constructor() MyStack {
	return MyStack{
		data: NewQueue(),
		tmp:  NewQueue(),
	}
}

func (this *MyStack) Push(x int) {
	this.data.Push(x)
}

func (this *MyStack) Pop() int {
	var v = -1
	for !this.data.IsEmpty() {
		v = this.data.Pop()
		this.tmp.Push(v)
	}
	for !this.tmp.IsEmpty() {
		t := this.tmp.Pop()
		if v != t {
			this.data.Push(t)
		}
	}
	return v
}

func (this *MyStack) Top() int {
	var v = -1
	for !this.data.IsEmpty() {
		v = this.data.Pop()
		this.tmp.Push(v)
	}
	for !this.tmp.IsEmpty() {
		this.data.Push(this.tmp.Pop())
	}
	return v
}

func (this *MyStack) Empty() bool {
	return this.data.IsEmpty()
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
	q.data = append([]int{x}, q.data...)
}

// 从队列中取出队头元素
func (q *Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[len(q.data)-1]
}

// 从队列中弹出队头元素
func (q *Queue) Pop() int {
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

## 树

`树定义`
 
 ```text
 树是一种数据结构，它是由n(n>=0)个有限节点组成一个具有层次关系的集合。把它叫做“树”是因为它看起来像一棵倒挂的树，也就是说它是根朝上，而叶朝下的。它具有以下的特点: 
 
 每个节点有零个或多个子节点;没有父节点的节点称为根节点;每一个非根节点有且只有一个父节点;除了根节点外，每个子节点可以分为多个不相交的子树。

 树 (tree) 是包含 n(n 2 0) 个节点，当 n=0 时，称为空树，非空树中 (n - 1) 条边的有穷集，在非空树中:
 (1)每个元素称为节点 (node)。
 (2)有一个特定的节点被称为根节点或树根 (root)
 (3) 除根节点之外的其余数据元素被分为 m (m >= 0)个互不相交的集合 T1,T2,..... Tm - 1，其中每一个集合 Ti(m >= i >= 1) 本身也是一棵树，被称作原树的子树 (subtree)

 树也可以这样定义: 树是由根节点和若干颗子树构成的。树是由一个集合以及在该集合上定义的一种关系校成的。集合中的元素称为树的节点，所定义的关系称为父子关系。父子关系在树的节点之间建立了一个层次结构。在这种层次结构中有一个节点具有特殊的地位，这个节点称为该树的根节点，或称为树根。
 ```

 就是一颗树 根节点位root 其余的称为node :

  ![树结构.png](https://s2.loli.net/2023/07/04/pKzF3SYLrbR1Pft.png)

`树的递归定义`

 ```text
 单个节点是一棵树，树根就是该节点本身。

 设T1,T2,T3....Tk是树，它们的根节点分别为n1,n2,n3,....nk。用一个新的节点n 作为n1,n2,n3,...nk 的父亲，则得到一课新树，节点 n 就是新树的根。我们称n1,n2,n3...,nk 为一组兄弟节点，它们都是节点n的子节点。我们还称T1,T2,... Tk为节点n的子树。
 ```

`树的学术名词`

* 空集合也是树，称为空树。空树中没有节点;
* 孩子节点或子节点: 一个节点含有的子树的根节点称为该节点的子节点;
* 节点的度: 一个节点含有的子节点的个数称为该节点的度;
* 叶节点或终端节点: 度为0的节点称为叶节点;
* 非终端节点或分支节点: 度不为0的节点;
* 双亲节点或父节点: 若一个节点含有子节点，则这个节点称为其子节点的父节点;
* 兄弟节点: 具有相同父节点的节点互称为兄弟节点;
* 树的度: 一棵树中，最大的节点的度称为树的度;
* 节点的层次: 从根开始定义起，根为第1层，根的子节点为第2层，以此类推;
* 树的高度或深度: 树中节点的最大层次;
* 堂兄弟节点: 双亲在同一层的节点互为堂兄弟;
* 节点的祖先: 从根到该节点所经分支上的所有节点;
* 子孙: 以某节点为根的子树中任一节点都称为该节点的子孙;
* 森林: 由 n 棵互不相交的树的集合称为森林。

`树的种类`

* 无序树: 树中任意节点的子结点之间没有顺序关系，这种树称为无序树,也称为自由树;
* 有序树: 树中任意节点的子结点之间有顺序关系，这种树称为有序树;
* 二叉树: 每个节点最多含有两个子树的树称为二叉树;
* 满二叉树: 叶节点除外的所有节点均含有两个子树的树被称为满二叉树
* 完全二叉树: 除最后一层外，所有层都是满节点，且最后一层缺右边连续节点的二叉树称为完全二叉树;
* 哈夫曼树(最优二叉树): 带权路径最短的二叉树称为哈夫曼树或最优二叉树。

`二叉树`

 **定义和基本形态**

 ```text
 二叉树(Binary tree) 是树形结构的一个重要类型，许多实际问题抽象出来的数据结构往往是二叉树形式，即使是一般的树也能简单地转换为二叉树，而且二叉树的存储结构及其算法都较为简单，因此二叉树显得特别重要。二叉树特点是每个结点最多只能有两棵子树，且有左右之分。
 
 二叉树是n个有限元素的集合，该集合或者为空、或者由一个称为根(root)的元素及两个不相交的、被分别称为左子树和右子树的二叉树组成，是有序树。当集合为空时，称该二叉树为空二叉树。在二叉树中，一个元素也称作一个结点。

 二叉树(binary tree) 是指树中节点的度不大于2的有序树，它是一种最简单且最重要的树。
 二叉树的送定义为: 二叉树是一棵空树，或者是一棵由一个根节点和两棵互不相交的，分别称作根的左子树和右子树组成的非空树;左子树和右子树又同样都是二叉树。
 ```

 二叉树是递归定义的，其结点有左右子树之分，逻辑上二叉树有五种基本形态: 

  ![1.jpg](https://s2.loli.net/2023/07/04/v3hwWOjGQEBRUrq.png)

 1. 空二叉树
 2. 只有一个根结点的二叉树
 3. 只有左子树
 4. 只有右子树
 5. 完全二叉树

 **二叉树的性质**

  1. 性质1: 二叉树的第i层上至多有2^(i-1) (i >= 1) 个节点
  2. 性质2: 深座为h的二叉树中至多合有(2^h)-1个节点
  3. 性质3: 若在任意一棵二叉树中，有n0个叶子节点，有n2个度为2的节点，则必有n0=n2+1。
  4. 性质4: 具有n个节点的完全二叉树深为log2x+1 (其中x表示不大于n的最大整数)。

 **二叉树的遍历**
 
 ```text
 所谓遍历(Traversal)是指沿着某条搜索路线，依次对树中每个结点均做一次且仅做一次访问。访问结点所做的操作依赖于具体的应用问题。 遍历是二叉树上最重要的运算之一，是二叉树上进行其它运算之基础。

 遍历是对树的一种最基本的运草，所谓遍历二叉树，就是按一定的规则和顺序走遍二叉树的所有结点，使每一个结点都被访问一次，而目只被访问一次。由于二叉树是非线性结构，因些树的遍历实质上是将二叉树的各个结点转换成为一个线性序列来表示。

 ```

 树是递归定义的，最简单也最快的实现方法就是递归遍历整个树，但是除了递归之外，迭代也是可以遍历树的，那么也就是树的宽度优先搜索。
 
 **树的深度优先DFS**

 深度优先搜索属于图算法的一种，英文缩写为DFS即Depth First Searh,其过程简要来说是对每一个可的分支路径深入到不能再深入为止，而且每个节点只能访问一次.
 
 深度优先遍历图的方法是，从图中某顶点v出发:

 1. 访问顶点v;
 2. 依次从v的未被访问的邻接点出发，对图进行深度优先遍历;直至图中和v有路径相通的顶点都被访问;
 3. 若此时图中尚有顶点未被访问，则以一个未被访问的顶点出发，重新进行深度优先遍历，直到图中所有顶点均被访问过为止， 当然，当人们刚刚掌握深度优先搜索的时候常常用它来走迷宫,事实上我们还有别的方法，那就是广度优先搜索(BFS).
 
 深度优先搜索用栈 (stack) 来实现，整个过程可以想象成一个倒立的树形:

 1. 把根节点压入栈中。
 2. 每次从栈中弹出一个元素，按索所有在它下一级的元素，把这些元索压入栈中。并把这个元素记为它下一级元素的前驱。
	* 前驱节点:对一棵二叉树进行中序遍历，遍历后的序，当前节点的前一个节点为该节点的前驱节点
	* 后继节点:对一棵二叉树进行中序遍历，遍历后的顺序，当前节点的后一个节点为该节点的后继节点
 3. 找到所要找的元素时结束程序.
 4. 如果遍历整个树还没有找到，结束程序。

 **树的宽度优先BFS**

 BFS，其英文全称是Breadth First Search。BFS并不使用经验法则算法。从算法的观点，所有因为展开节点而得到的子节点都会被加进一个先进先出的队列中。一般的实验里，其邻居节点尚未被检验过节点会被放置在一个被称为 open 的容器中(例如队列或是链表)，而被检验过的节点则被放置在被你为 closed 的容器中。 (open-closed表)
 
 广度优先搜索使用队列(queue) 来实现，整个过程也可以看做一个倒立的树形:

 1. 把根节点放到队列的末尾。
 2. 每次从队列的头部取出一个元素，查看这个元素所有的下一级元素，把它们放到队列的末尾。并把这个元素记为它下一级元素的前驱。
 3. 找到所要找的元素时结束程序。
 4. 如果遍历整个树还没有找到，结束程序

 **算法实现**

 从二叉树的递归定义可知，一棵非空的二叉树由根结点及左、右子树这三个基本部分组成因此，在任一给定结点上，可以按某种次序执行三个操作:

 1. 访问节点的本身 (Node)
 2. 遍历该节点的左子树 (L)
 3. 遍历该节点的右子树 (R)

 以上三种操作拥有六种执行顺序: NLR，LNR，LRN，NRL，RNL，RLN。

 *但是注意前三种次序和后三种次序对称，所以我们只学习前三种次序*

 遍历命名,根据访问结点操作发生位置命名:

 1. NLR: 二叉树的前序遍历(Preorder Traversal 亦称 (先序遍历) )
	* 访问根结点的操作发生在遍历其左右子树之前。
 2. LNR: 二叉树的中序遍历(lnorder Traversal)
	* 访问根结点的操作发生在遍历其左右子树之中 (间)。
 3. LRN: 二叉树的后序遍历(Postorder Traversal) 
	* 访问根结点的操作发生在遍历其左右子树之后。

 **递归的思路**

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var res = []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func inorderTraversal(root *TreeNode) []int {
	var res = []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func postorderTraversal(root *TreeNode) []int {
	var res = []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}
```
 
 **迭代的思路**

 1. 前序遍历
	* 建立一个栈，用来存储节点
	* 根据左右根的顺序将节点依次压入栈中，在压入栈中的同时，按照顺序把节点里的元素依次压入栈中。输出完毕之后按照顺序弹栈。
	* 输出答案。
 2. 中序遍历
    * 建立一个栈，用来存储节点
	* 根据左根右的顺序将节点依次压入栈中，在压入栈中的同时，按照顺序把节点里的元素依次压入栈中。输出完毕之后按照顺序弹栈。
	* 输出答案。
 3. 后序遍历
    * 同前序遍历,先右再左
	* 翻转结果数组
	* 输出答案。

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var res = []int{}
	var stack = NewStack()
	for !stack.IsEmpty() || root != nil {
		for root != nil {
			// 先访问根节点
			res = append(res, root.Val)
			// 一直寻找左子树
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop().Right
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	var res = []int{}
	var stack = NewStack()
	for !stack.IsEmpty() || root != nil {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		if !stack.IsEmpty() {
			root = stack.Pop()
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	var res = []int{}
	var stack = NewStack()
	for !stack.IsEmpty() || root != nil {
		for root != nil {
			// 放入根节点
			res = append(res, root.Val)
			stack.Push(root)
			// 一直找右
			root = root.Right
		}
		// 再寻找左节点
		if !stack.IsEmpty() {
			root = stack.Pop().Left
		}
	}
	// 得到 nrl 的结果
	tmp := []int{}
	for i := len(res) - 1; i >= 0; i-- {
		tmp = append(tmp, res[i])
	}
	return tmp
}

type Stack struct {
	data []*TreeNode
}

// 初始化栈
func NewStack() Stack {
	return Stack{
		data: []*TreeNode{},
	}
}

// 将元素压进栈
func (s *Stack) Push(x *TreeNode) {
	s.data = append(s.data, x)
}

// 将元素弹栈
func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 拿到栈顶元素
func (s *Stack) Top() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// 统计栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}
```

 **二叉树的层序遍历**

 二叉树的层序遍历用到的就是: 宽度优先搜索，也就是BFS，BFS使用时要记住一点:就是宽搜只能用在边权值固定的图中，例如二叉树的层序遍历，我们在每一层每一层的遍历，边权都是1，因为每一层和每一层的间距都是一定的，我们也是每一层每一层的遍历。

 [1.png](https://s2.loli.net/2023/07/05/d2p7VOwomcSeGCf.png)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	q := NewQueue()
	if root != nil {
		q.PushTail(root)
	}
	for !q.IsEmpty() {
		tmp := []int{}
		childNums := q.Size()
		for ; childNums > 0; childNums-- {
			node := q.RemoveFrist()
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q.PushTail(node.Left)
			}
			if node.Right != nil {
				q.PushTail(node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

type Queue struct {
	data []*TreeNode
}

// 初始化队列
func NewQueue() Queue {
	return Queue{
		data: []*TreeNode{},
	}
}

// 从左边入队
func (q *Queue) PushFrist(x *TreeNode) {
	q.data = append([]*TreeNode{x}, q.data...)
}

// 从右边入队
func (q *Queue) PushTail(x *TreeNode) {
	q.data = append(q.data, x)
}

// 获取队头元素
func (q *Queue) PeekFrist() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	return q.data[0]
}

// 从队头弹出元素
func (q *Queue) RemoveFrist() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

// 获取队尾元素
func (q *Queue) PeekTail() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	return q.data[len(q.data)-1]
}

// 从队尾弹出元素
func (q *Queue) RemoveTail() *TreeNode {
	if q.IsEmpty() {
		return nil
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

  **二叉树的垂序遍历**

   ![3.png](https://s2.loli.net/2023/07/05/bmdPnItRHx13OQZ.png)
 
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
	res := [][]int{}
	m := make(map[int][]*col)
	var dfs func(root *TreeNode, x, y int)
	dfs = func(root *TreeNode, x, y int) {
		if root == nil {
			return
		}
		_, ok := m[y]
		if !ok {
			m[y] = []*col{}
		}
		m[y] = append(m[y], &col{x, root.Val})
		dfs(root.Left, x+1, y-1)
		dfs(root.Right, x+1, y+1)
	}
	dfs(root, 0, 0)
	tmpY := []int{}
	for k, _ := range m {
		tmpY = append(tmpY, k)
	}
	sort.Ints(tmpY)
	for _, key := range tmpY {
		arrs := m[key]
		if len(arrs) < 2 {
			res = append(res, []int{arrs[0].val})
		} else {
			sort.Slice(arrs, func(i, j int) bool {
				if arrs[i].x == arrs[j].x {
					return arrs[i].val < arrs[j].val
				} else {
					return arrs[i].x < arrs[j].x
				}
			})
			tmpArr := []int{}
			for _, v := range arrs {
				tmpArr = append(tmpArr, v.val)
			}
			res = append(res, tmpArr)
		}
	}
	return res
}

type col struct {
	x   int
	val int
}
```

`多叉树遍历`

 N叉树是没有中序遍历的!!!因为二叉树的中序遍历是左 根 右的节点顺序进行遍历的，由于N叉树无法分辨左右子树，所以没有中序遍历，

 1. 多叉树的前序遍历

	* ![1.png](https://s2.loli.net/2023/07/06/SXBwp5D3OdfbKgR.png)

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
 // 递归的写法
func preorder(root *Node) []int {
	res := []int{}

	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		for _, v := range root.Children {
			dfs(v)
		}
	}
	dfs(root)
	return res
}

// 迭代的写法
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := NewStack()
	stack.Push(root)
	for !stack.IsEmpty() {
		n := stack.Pop()
		res = append(res, n.Val)
		for i := len(n.Children) - 1; i >= 0; i-- {
			stack.Push(n.Children[i])
		}
	}
	return res
}

type Stack struct {
	data []*Node
}

// 初始化栈
func NewStack() Stack {
	return Stack{
		data: []*Node{},
	}
}

// 将元素压进栈
func (s *Stack) Push(x *Node) {
	s.data = append(s.data, x)
}

// 将元素弹栈
func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 拿到栈顶元素
func (s *Stack) Top() *Node {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// 统计栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}
```

 2. 多叉树的后序遍历

	* ![2.png](https://pic.imgdb.cn/item/64a636081ddac507cc463920.png)

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
 // 递归版本
func postorder(root *Node) []int {
	res := []int{}
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root != nil {
			for _, v := range root.Children {
				dfs(v)
			}
			res = append(res, root.Val)
		}
	}
	dfs(root)
	return res
}

// 迭代版本
func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := NewStack()
	stack.Push(root)
	for !stack.IsEmpty() {
		n := stack.Pop()
		res = append(res, n.Val)
		for _, v := range n.Children {
			stack.Push(v)
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

type Stack struct {
	data []*Node
}

// 初始化栈
func NewStack() Stack {
	return Stack{
		data: []*Node{},
	}
}

// 将元素压进栈
func (s *Stack) Push(x *Node) {
	s.data = append(s.data, x)
}

// 将元素弹栈
func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 拿到栈顶元素
func (s *Stack) Top() *Node {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// 统计栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}
```

 2. 多叉树的层序遍历

	* ![3.png](https://s2.loli.net/2023/07/06/9JrvIHEkB1W8Qes.png)

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := NewQueue()
	q.PushTail(root)
	for !q.IsEmpty() {
		l := q.Size()
		tmp := []int{}
		for ; l > 0; l-- {
			n := q.RemoveFrist()
			tmp = append(tmp, n.Val)
			for _, v := range n.Children {
				q.PushTail(v)
			}
		}
		res = append(res, tmp)
	}
	return res
}

type Queue struct {
	data []*Node
}

// 初始化队列
func NewQueue() Queue {
	return Queue{
		data: []*Node{},
	}
}

// 从左边入队
func (q *Queue) PushFrist(x *Node) {
	q.data = append([]*Node{x}, q.data...)
}

// 从右边入队
func (q *Queue) PushTail(x *Node) {
	q.data = append(q.data, x)
}

// 获取队头元素
func (q *Queue) PeekFrist() *Node {
	if q.IsEmpty() {
		return nil
	}
	return q.data[0]
}

// 从队头弹出元素
func (q *Queue) RemoveFrist() *Node {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

// 获取队尾元素
func (q *Queue) PeekTail() *Node {
	if q.IsEmpty() {
		return nil
	}
	return q.data[len(q.data)-1]
}

// 从队尾弹出元素
func (q *Queue) RemoveTail() *Node {
	if q.IsEmpty() {
		return nil
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

`Trie树(字典树,前缀树)`
 
 ```text
 又称单词查找树，Trie树，是一种树形结构，是一种哈希树的变种。典型应用是用于统计，排序和保存大显的字符串(但不仅限于字符串)，所以经常被搜索引擎系统用于文本词频统计。它的优点是，利用字符串的公共前缀来减少查询时间，最大限度地减少无谓的字符串比较，查询效率比哈希树高
 ```

 **Trie树的性质**
 
 它有3个基本性质:
 
 1. 根节点不包含字符，除根节点外每一个节点都只包含一个学符;
 2. 从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串;
 3. 每个节点的所有子节点包含的字符都不相同。

 总而言之 Trie树就是用来快速的存储字符申集合的数据结构。

 **Trie树的基本操作**
 
 其基本操作有: 查找、插入和删除,当然删除操作比较少见。

  ![4.png](https://s2.loli.net/2023/07/06/NT38jZSMClbHhgD.png)

 **实现Trie树**

```go
type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	next [26]*TrieNode
	end  bool
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			next: [26]*TrieNode{},
			end:  false,
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	b := []byte(word)
	for _, v := range b {
		idx := v - 'a'
		if cur.next[idx] == nil {
			cur.next[idx] = &TrieNode{
				next: [26]*TrieNode{},
				end:  false,
			}
		}
		cur = cur.next[idx]
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	b := []byte(word)
	for _, v := range b {
		idx := v - 'a'
		if cur.next[idx] == nil {
			return false
		}
		cur = cur.next[idx]
	}
	return cur.end
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	b := []byte(prefix)
	for _, v := range b {
		idx := v - 'a'
		if cur.next[idx] == nil {
			return false
		}
		cur = cur.next[idx]
	}
	return true
}
```

## 哈希表

`定义` 
 
 ```text
 散列表 (Hash table，也叫哈希表) ，是根据关键码值(Key value)而直接进行访问的数据结构。也就是说，它通过把关键码值映射到表中一个位置来访问记录，以加快查找的速度。这个映射函数叫做散列函数，存放记录的数组叫做散列表

 给定表M，存在函数f(key)，对任意给定的关键字值key，代入函数后若能得到包含该关键字的记录在表中的地址，则称表M为哈希 (Hash) 表，函数f(key)为哈希(Hash)函数

 哈希表结构一般是将庞大的数据量存储到一个较小的集合中
 ```

`基本概念`
 
 1. 若关键字为n，则其值存放在f(n)的存储位置上。由此，不需比较便可直接取得所查记录。称这个对应关系f为散列函数，按这个思想建立的表为散列表。
 2. 对不同的关键字可能得到同一散列地址，即n1 != n2，而f(n1)==f(n2)，这种现象称为冲突。具有相同函数值的关键字对该散列函数来说称做同义词。综上所述，根据散列函数f(n)和处理冲突的方法将一组关键字映射到一个有限的连续的地址集 (区间)上，并以关键字在地址集中的“像”作为记录在表中的存储位置，这种表便称为散列表，这一映射过程称为散列造表或散列，所得的存储位置称散列地址。
 3. 若对于关键字集合中的任一个关键字，经散列函数映象到地址集合中任何一个地址的概率是相等的，则称此类散列函数为均匀散列函数(Uniform Hash function)，这就是使关键字经过散列函数得到一个“随机的地址”，从而减少冲突
 4. 在设计哈希表的时候，最需要注意两个基本因素:一个是哈希函数的编写，一个是键冲突解决算法

`散列函数`

 ```text
 一般的线性表，树中，记录在结构中的相对位置是随机的，即和记录的关键字之间不存在确定的关系，因此，在结构中查找记录时需进行一系列和关键字的比较。这一类查找方法建立在“比较”的基础上，查找的效率依赖于查找过程中所进行的比较次数。 理想的情况是能直接找到需要的记录，因此必须在记录的存储位置和它的关键字之间建立一个确定的对应关系f，使每个关键字和结构中一个唯一的存储位置相对应。
 ```

 哈希表中元素的位置是由哈希函数确定的。将数据元素的关键字n作为自变量，通过一定的函数关系计算出的值，即为该元素的存储地址。

`哈希冲突和处理的办法`

 在哈希表中，不同的关键字值对应到同一个存储位置的现象。即关键字n1 != n2，但f(n1) == f (n2)。均匀的哈希函数可以减少冲突，但不能避免冲突。发生冲突后，必须解决;也即必须寻找下一个可用地址。

 1. 拉链法 (最常用的解决哈希冲突的算法)
	* 将具有同一散列地址的记录存储在一条线性链表中。例，除留余数法中，设关键字为(18,14,01,68,27,55,79) ，除数为13。散列地址为 (5,1,1,3,1,3,1) , 模数尽量为质数,远离 2^n
 
 2. 开放定址法

`哈希集合`
 
 Hash_set 被称为集合，该容器中只能存放不重复的对象。 
 
 ![5.png](https://s2.loli.net/2023/07/06/WEYqm8JeNfTM3Qv.png)

```go
type MyHashSet struct {
	l    int
	data [99997][]int
}

func Constructor() MyHashSet {
	return MyHashSet{
		data: [99997][]int{},
		l:    99997,
	}
}

func (this *MyHashSet) Add(key int) {
	idx := key % this.l
	if this.data[idx] == nil {
		this.data[idx] = []int{key}
		return
	}
	for _, v := range this.data[idx] {
		if v == key {
			return
		}
	}
	this.data[idx] = append(this.data[idx], key)
}

func (this *MyHashSet) Remove(key int) {
	idx := key % this.l
	if this.data[idx] == nil {
		return
	}
	index := -1
	for k, v := range this.data[idx] {
		if v == key {
			index = k
		}
	}
	if index == -1 {
		return
	}
	switch index {
	case 0:
		this.data[idx] = this.data[idx][1:]
	case len(this.data[idx]) - 1:
		this.data[idx] = this.data[idx][:len(this.data[idx])-1]
	default:
		this.data[idx] = append(this.data[idx][:index], this.data[idx][index:]...)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	idx := key % this.l
	if this.data[idx] == nil {
		return false
	}
	for _, v := range this.data[idx] {
		if v == key {
			return true
		}
	}
	return false
}
```

`哈希映射`

 Hash_map 存放的是键值对，我们可以通过键 (key) 来找到对应的值 (value)。

 ![6.png](https://s2.loli.net/2023/07/06/LOPwqcnMbJzXsuE.png)

```go
const lenHash = 99997

type bucket struct {
	key int
	val int
}

type MyHashMap struct {
	l    int
	data [99997][]*bucket
}

func Constructor() MyHashMap {
	return MyHashMap{
		l:    lenHash,
		data: [99997][]*bucket{},
	}
}

func (this *MyHashMap) Put(key int, value int) {
	idx := key % this.l
	if this.data[idx] == nil {
		this.data[idx] = []*bucket{}
	}
	for _, v := range this.data[idx] {
		if v.key == key {
			if v.val == value {
				return
			} else {
				v.val = value
				return
			}
		}
	}
	this.data[idx] = append(this.data[idx], &bucket{key, value})
}

func (this *MyHashMap) Get(key int) int {
	idx := key % this.l
	if this.data[idx] == nil {
		return -1
	}
	for _, v := range this.data[idx] {
		if v.key == key {
			return v.val
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := key % this.l
	if this.data[idx] == nil {
		return
	}
	indexs := []int{}
	for k, v := range this.data[idx] {
		if v.key == key {
			indexs = append(indexs, k)
		}
	}
	if len(indexs) == 0 {
		return
	}
	for _, index := range indexs {
		switch index {
		case 0:
			this.data[idx] = this.data[idx][1:]
		case len(this.data[idx]) - 1:
			this.data[idx] = this.data[idx][:len(this.data[idx])-1]
		default:
			this.data[idx] = append(this.data[idx][:index], this.data[idx][index:]...)
		}
	}
}
```
