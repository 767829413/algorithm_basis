# 并查集算法

## 并查集的定义和作用

 ```text
 并查集，在一些有N个元素的集合应用问题中，我们通常是在开始时让每个元素构成一个单元素的集合，然后按一定顺序将属于同一组的元素所在的集合合并，其间要反复查找一个元素在哪个集合中。这一类问题近几年来反复出现在信息学的国际国内赛题中。其特点是看似并不复杂，但数据量极大，若用正常的数据结构来描述的话，往往在空间上过大，计算机无法承受，即使在空间上勉强通过，运行的时间复杂度也极高，根本就不可能在比赛规定的运行时间 (1~3秒)内计算出试题需要的结果，只能用并查集来描述。
 并查集是一种树型的数据结构，用于处理一些不相交集合 (disjoint sets) 的合并及查询问题。常常在使用中以森林来表示。
 ```

 并查集本质上可以做如下两个操作:

  1. 将两个集合合并
  2. 查看某两个元素是否在同一个集合当中

  如果暴力去做的话，查看某两个元素是否在同一个集合的操作并不难，我们拿一个数组 arr[] 存储每一个元素属于哪一个集合。比如 arr[x] == a 表示数字x属于集合a，那么查看连续个元素是否属于一个集合的代码就是 arr[x] == arr[y] 。这个操作的时间复杂度不难看出来是O(1)。

  但是将两个集合合并的话，这个操作就非常的麻烦了，需要扩容并且一个一个把其中一个集合的元素放入另一个集合中，非常麻烦，时间复杂度极高，并不适合。这就得看并查集了!

  **并查集本质是一棵多叉树!!!**

## 并查集的相关操作

并查集的思路如下: 

每个集合用一个多叉树表示，根节点的编号就是整个集合的编号.

1. 初始化: 把每个元素的所在集合都设置为自己，也就是每个元素占一个集合

2. int f[n]; for (int i = 0; i < n; i++) f[i] == i;
	![0.png](https://s2.loli.net/2023/07/26/o6iEvtXVlLTcahf.png)

3. 查找: 查找元素所在的集合,即根节点。

```C
int find(int i) { //查询祖先
	if (f[i] == i) return i;
	else return find(f[i]);
}
```

4. 合并: 将两个元素所在的集合合并为一个集合。

***通常来说，合并之前，应先判断两个元素是否属于同一集合，这可用上面的“查找”操作实现。***

```C
 void union(int i, int j) {
	int i = find(i);
	int j = find(j);
	f[i] = j;
}
```

假如说们合并的是union(5,4)，那么结果就是这样的:
 
 ![1.png](https://s2.loli.net/2023/07/26/aRynt5PDVGNzckS.png)

继续合并操作: union(4,3),union(3,2)

会得到如下图所示: 

 ![2.png](https://s2.loli.net/2023/07/26/34crtHbedul765z.png)

## 路径压缩

 我们发现其实我们的find函数 (方法)并没有那么的高效，因为我们在找父节点的时候，是一层一层朝上找的，最多会有n层的节点我们需要往上走，那么有没有什么好的办法我们直接一次性的找到节点呢? 这里我们就有一个对并查集的优化操作，叫做**路径压缩**。

 ![3.png](https://s2.loli.net/2023/07/26/U6cSiVmZl9jQaM5.png)
  
 假设我们开始的时候树是图左所示, 我们通过路径压缩，每次在找根节点的时候让路径上的节点直接都指向跟根节点，变成图右的形式。

 ```C
 int find(int i) {
	if (f[i] == i) return i;
	else {
		f[i] = find(f[i]);
		return f[i];
	}
 }

 int find(int i) {
	if (f[i] != i) f[i] = find(f[i]);
	return f[i];
 }
 ```

## 并查集代码实践

 **一共有 n 个数，编号是 1~n，最开始每个数各自在一个集合中。第一行输入数字n和m，表示n个数字和m次询问，询问操作共有两种: 1.M a b，将编号为数字a和b的两个集合合并。如果已经合并，则忽略此操作。2.Q a b询问编号a和b的两个数字是否在一个集合里。是的话输出yes，不是输出No。**

```go
var reader = bufio.NewReader(os.Stdin)
var f []int

func main() {
	fmt.Println("请输入n和m的值")
	data, _ := reader.ReadString('\n')
	dataArr := strings.Split(strings.TrimSpace(data), " ")
	n, _ := strconv.Atoi(strings.TrimSpace(dataArr[0]))
	m, _ := strconv.Atoi(strings.TrimSpace(dataArr[1]))
	// 初始化
	f = make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = i
	}

	for ; m > 0; m-- {
		fmt.Println("请输入操作 M a b 或者 Q a b")
		input, _ := reader.ReadString('\n')
		inputArr := strings.Split(strings.TrimSpace(input), " ")
		op := strings.TrimSpace(inputArr[0])
		a, _ := strconv.Atoi(strings.TrimSpace(inputArr[1]))
		b, _ := strconv.Atoi(strings.TrimSpace(inputArr[2]))
		if op == "M" {
			f[find(a)] = find(b)
		} else {
			if find(a) == find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

func find(x int) int {
	if f[x] != x {
		f[x] = find(f[x])
	}
	return f[x]
}
```
