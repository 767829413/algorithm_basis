# 计算机基础知识

## 基本数学知识

`计算机上的一维坐标系`

 **选某一坐标为坐标原点，以某个方向为正方向，选择适当的标度建立一个坐标轴，就构成了一维坐标系，适于描述物体在一维。**

 ![1.jpg](https://s2.loli.net/2023/06/16/QmKTNjBbM1JZCpn.png)

`计算机上的二维坐标系`

 ![2.jpg](https://s2.loli.net/2023/06/16/yiMjLHAcVqZxK2u.png)

 计算机中表示二维数组的坐标系x轴是向下的，y轴是向上的。

 已知一条线经过原点(0,0)，并且和x轴所形成的夹角为45°，一次方程为: y=x。斜率就是 y/x = 1/1 = 1

 ![3.png](https://s2.loli.net/2023/06/16/tErX8lNmg2HF9nc.png)

 那么如果我们把此线段向右移动1，那么此时方程为y=x+1, 此时截距等于1

 ![4](https://pic.imgdb.cn/item/648bcec61ddac507cc82b105.png)

`四方向`

 ![5](https://pic.imgdb.cn/item/648bd03f1ddac507cc865b84.png)

 其中，中间的坐标为(x,y)，那么它上下左右四个方向的坐标为多少?

 ![6](https://pic.imgdb.cn/item/648bd1031ddac507cc889f21.png)

 那么，我们在以后做题的过程中，可能会遇到搜索四个方向的情况，此时我们就可以通过我们得到的数学关系，按照你喜欢或者题目要求的顺序，定义方向向量。
 
 **总结: x轴从上至下依次增大 (从北向南依次增大)，y轴从左至右依次增大 (从西向东依次增大);每走一格，根据方向增加或减少一个单位长度。**

`八方向`

 类似上面的四方向,也可以得出八个方向的向量

 ![7](https://pic.imgdb.cn/item/648bd2291ddac507cc8b5b82.png)

## 算法复杂度

`定义`

 **算法复杂度分为时间复杂度和空间复杂度。其作用: 时间复杂度是指执行算法所需要的计算工作量,而空间复杂度是指执行这个算法所需要的内存空间。 (算法的复杂性体现在运行该算法时的计算机所需资源的多少上，计算机资源最重要的是时间和空间(即寄存器)资源，因此复杂度分为时间和空间复杂度。)**

 ![8](https://pic.imgdb.cn/item/648bd3541ddac507cc8e1714.png)

`时间复杂度`

 **一个语句的频度是指该语句在算法中被重复执行的次数。算法中所有语句的频度之和记为T(n)，它是算法问题规模n的函数，时间复杂度主要分析T(n)的数量级。算法中基本运算(最深层循环内的语句)与T(n)同数量级，因此通常采用算法中基本运算的频度f(n)来分析我们算法的时间复杂度，记为BigO,T(n) = O(f(n))，简写为O(n)。**

 在计算机科学中，时间复杂性，又称时间复杂度，算法的时间复杂度是一个函数，它定性描述该算法的运行时间。这是一个代表算法输入值的字符串的长度的函数。时间复杂度常用大0符号表述，不包括这个函数的低阶项和首项系数。使用这种方式时，时间复杂度可被称为是渐近的，亦即考察输入值大小趋近无穷时的情况.

 在T(n) = O(f(n))这个式子中，O的含义是T(n)的数量级，其严格的数学定义是: 若T(n)和f(n)是定义在正整数集合上的两个函数，则存在正常数C和n0，使得当n >= n0时，都满足0 <= T(n) <= Cf(n)。

 **算法的时间复杂度不仅依赖于问题的规模n，也取决于待输入数据数据的性质**

 例如给定一个int类型的数组A [0... n-1] 当中，查找给定值 k 的算法大致如下:

 ```go
 var i = n - 1
 for i >= 0 && A[i] != k {
    i--
 }
 return i
 ```

 该算法中第三行语句的频度不难发现实际上和n的规模有关，并且和A的各个元素的取值和k的取值都有关系。

* 如果A的所有元素中没有和k相等的元素，则第三行的代码的频度为f(n) = n
* 如果A的最后一个元素等于k，那么第三行代码的频度f(n)= 0

 这就是所谓的最坏和最好的时间复杂度，通过最坏和最好时间复杂度的概念我们可以得到平均时间复杂度。

* **最坏时间复杂度**: 最坏时间复杂度指的是最坏情况下，算法的时间复杂度，也就是我们上述代码第三行A的元素没有和K相等情况下的复杂度
* **最好时间复杂度**: 最好时间复杂度指的是在最好的情况下，算法的时间复杂度，例如A的最后一个元素恰好等于k。
* **平均时间复杂度**: 平均时间复杂度指的是所有可能输入实例在等概率出现的情况下，算法的期望运行时间。

我们再去分析一个算法或者一段代码的时间复杂度的时候，有以下两条规则:

* 加法规则 T(n) = T1(n) + T2(n) = O(f(n)) + O(g(n)) = O(max(f(n), g(n)))
* 乘法规则: T(n)= T1(n) \* T2(n) = O(f(n)) \* O(g(n)) = O(f(n)* g(n))

常见的时间复杂度有:

O(1) < O(logn) < O(n) < O(nlogn) < O(n^2) < O(n^3) < O(2^n) < O(n!) < O(n^n)

举例1: 

```go
for i := 0; i < 2 * n; i ++ { println(i) };
```

 O(n),我们for循环循环了2n次，因此时间复杂度为O(n)，不用考虑常数空间。

举例2:

```go
for i:=0; i<n; i++ {
    for j:=0; j<n; j++ {
        println(i,j)
    }
}
```

 O(n^2), 两重循环，所以为O(n^2)

举例3:

```go
var x = 2
for x < n / 2 {
    x = 2 * x
}
```

 O(logn), x乘以了若干次2之后依旧小于 n/2，所以x大概乘以了logn次

举例4:

```go
func fact(n int) int {
    if n <= 1 {
        return 1;
    }
    return n * fact(n-1)
}
```
 
 O(n), 其本质就是从n乘到1，乘了n次

举例5:

```go
count := 0;
for i:=0; i<n; i = i * 2 {
    for j:=0; j<n; j++ {
        count++
    }
} 
```
 
 O(nlogn), 第一层logn,第二层 n

`空间复杂度`

 **空间复杂度(Space Complexity)是对一个算法在运行过程中临时占用存储空间大小的量度，记做S(n)=O(f(n))。比如直接插入排序的时间复杂度是O(n^2), 空间复杂度是O(1)。而一般的递归算法就要有o(n)的空间复杂度了，因为每次递归都要存储返回信息。一个算法的优劣主要从算法的执行时间和所需要占用的存储空间两个方面衡量。**

 算法的空间复杂度S(n)定义为该算法所耗费的存储空间，他是问题规模n的函数。记为S(n)=O(g(n))。

 一个算法或者程序在执行的时候所需要一些额外的存储空间来存放本身所用的命令、常数、变量和输入数据外，还需要要一些对数据机型操作的工作单元和存储一些为实现计算所需要信息的辅助空间。如果输入的数据所占空间只取决于问题本身，和算法无关，则只需分析除了输入和程序之外的额外空间。

 算法原地工作被称为原地修改算法，即所需的辅助空间为常量。即O(1)。

## 进制基本概念和转换

`定义`

 **程序中的所有数在计算机内存、都是以二进制的形式储存的。位运算就是直接对整数在内存中的二进制位进行操作。比如，and运算本来是一个逻辑运算符，但整数与整数之间也可以进行and运算。举个例子，6的二进制是110，11的二进制是1011，那么6 and11的结果就是2，它是二进制对应位进行逻辑运算的结果 (0表示False，1表示True,空位都当0处理)。**

`进制运算`

 **进制也就是进位计数制，是人为定义的带进位的计数方法 (有不带进位的计数方法，比如原始的结绳计数法，唱票时常用的“正”字计数法，以及类似的tally mark计数)。 对于任何一种进制---X进制，就表示每一位置上的数运算时都是逢X进一位。 十进制是逢十进一，十六进制是逢十六进一，二进制就是逢二进一，以此类推，x进制就是逢x进位**

计算机使用数制和相互转换

![1.jpg](https://s2.loli.net/2023/06/19/silE5wJF7X9W6kn.jpg)

**十进制**

*600，3/5，-7.99.....看着这些耳熟能详的数字，你有没有想太多呢? 其实这都是全世界通用的十进制，即*

 1. 满十进一，满二十进二，以此类推.....
 2. 按权展开，第一位权为10^0，第二位10^1.....以此类推，第N位10^ (N-1)，该数的数值等于每位位的数值*该位对应的权值之和。

**二进制**
*二进制 (binary)是在数学和数字电路中指以2为基数的记数系统，是以2为基数代表系统的二进位制。这一系统中，通常用两个不同的符号0 (代表零)和1 (代表一)来表示。数字电子电路中，逻辑门的实现直接应用了二进制，因此现代的计算机和依赖计算机的设备里都用到二进制。每个数字称为一个比特Bit，Binary digit的缩写。运算:*

* 加法
  * 二进制加法有四种情况: 0+0=0，0+1=1，1+0=1，1+1=10(0 进位为1)
* 乘法
  * 二进制乘法有四种情况: 0x0=0，1x0=0，0x1=0，1x1=1。
* 减法
  * 二进制减法有四种情况: 0-0=0，1-0=1，1-1=0，0-1=1。
* 除法
  * 二进制除法有两种情况(除数只能为1): 0/1=0，1/1=1

`二进制转换为十进制`

 *方法:“按权展开求和”，该方法的具体步骤是先将二进制的数写成加权系数展开式，而后根据十进制的加法规则进行求和。*

 eg: 
 
 $ {(1011)}_{2}=1\times {2}^{3}+0\times {2}^{2}+1\times {2}^{1}+1\times {2}^{0}={(11)}_{10} $

 **规律:个位上的数字的次数是0，十位上的数字的次数是1，.....依次递增,而十分位的数字次数是-1，百分位上数字的次数是-2，......，依次递减。**

`十进制转换为二进制`

 *一个十进制数转换为二进制数要分整数部分和小数部分分别转换，最后再组合到一起。*

 **整数部分采用“除2取余，逆序排列"法。具体做法是: 用2整除十进制整数，可以得到一个商和余数;再用2去除商，又会得到一个商和余数，如此进行，直到商为小于1时为止，然后把先得到的余数作为二进制数的低位有效位，后得到的余数作为二进制数的高位有效位依次排列起来。**

 eg: 173:
  
  ![3.png](https://s2.loli.net/2023/06/19/OItheSbXN89YcfD.png)

  **小数部分要使用"乘 2 取整法”。即用十进制的小数乘以 2 并取走结果的整数(必是 0 或 1)然后再用剩下的小数重复刚才的步骤，直到剩余的小数为 0 时停止，最后将每次得到的整数部分按先后顺序从左到右排列即得到所对应二进制小数。例如，将十进制小数 0.8125 转换成二进制小数过程如下:**

  ![4.png](https://www.runoob.com/wp-content/uploads/2018/11/210-3.png)

`通用的进制转换`

 *不同进制之间的转换本质就是确定各个不同权值位置上的数码。*转换正整数的进制的有一个简单算法，就是通过用目标基数作长除法;余数给出从最低位开始的“数字”。
 
 eg: 
 
* 1020304从10进制转到7进制:
 
 ```text
 1020304 / 7 = 145757 r 5 ↑ => 11446435
 145757  / 7 = 20822  r 3 |
 20822   / 7 = 2974   r 4 |
 2974    / 7 = 424    r 6 |
 424     / 7 = 60     r 4 |
 60      / 7 = 8      r 4 |
 8       / 7 = 1      r 1 |
 1       / 7 = 0      r 1 |
 ```
 
* 10110111 从2进制转到5进制:
 
 ```text
 10110111 / 101 = 100100 r 11(3)  ↑ => 1213
 100100   / 101 = 111    r 1(1)   |
 111      / 101 = 1      r 10(2)  |
 1        / 101 = 0      r 1(1)   |
 ```

`八进制`

 **Octal，缩写OCT或O，一种以8为基数的计数法，采用0，1，2，3，4，5，6，7八个数字，逢八进一。一些编程语言中常常以数字0开始表明该数字是八进制。八进制的数和二进制数可以按位对应(八进制一位对应二进制三位)，因此常应用在计算机语言中,八进制(基数为8)表示法在计算机系统中很常见，因此，我们有时能看到人们使用八进制表示法。由于十六进制一位可以对应4位二进制数字，用十六进制来表示二进制较为方便。因此，八进制的应用不如十六进制。有一些程序设计语言提供了使用八进制符号来表示数字的能力，而且还是有一些比较古老的Unix应用在使用八进制。**

 *八进制转换为十进制和二进制*

 **二进制与八进制的互相转换和二进制与十六进制的转换类似，区别在于需要操作的是三位一组而不是四位。下表列出了二进制与八进制的等效表示。**
 
 **为了把八进制数换算为二进制，将每一个八进制数字替换成表中对应的三位。**

 例如，八进制 123 换算成二进制的结果就是 001 010 011

 **为了将一个二进制数换算为八进制，只需将二进制串划分成每三个位一组(如果需要的话,在前面补零),将三位一组的位串替换为相应的八进制数字即可。如果需要将八进制数换算为十六进制，只需将八进制数换算为二进制，然后再换算为十六进制即可。**
 
 例:
  将八进制数12转换成十进制数 $ {(12)}_{8}=1\times {8}^{1}+2\times {8}^{0}={(10)}_{10} $

  将八进制数17.36转换成2进制数 $ {(17.36)}_{8}=001\ 111\ .\ 011\ 110={(1111.01111)}_{2} $

`十六进制`

 **六进制(简写为hex或下标16) 在数学中是一种逢16进1的进位制。一般用数字0到9和字母A到F (或a~f) 表示，其中:A~F表示10~15，这些称作十六进制数字。**

**例如十进制数57，在二进制写作111001，在16进制写作39。如今的16进制则普遍应用在计算机领域，这是因为将4个位元 (Bit) 化成单独的16进制数字不太困难。1字节可以表示成2个连续的16进制数字。可是，这种混合表示法容易令人混滑，因此需要一些字首、字尾或下标来显示。**

*十六进制转换为二进制*
 
 $ {(1CA)}_{16}=0001\ 1100\ 1010={(1\ 1100\ 1010)}_{2} $

 每一个16进制数字转换为2进制并起来即可。

*十六进制转换为八进制*

 先用1化4方法，将十六进制化为二进制;再用3并1方法，将二进制化为8制。
 
 例: $ {(1CA)}_{16}={(1\ 1100\ 1010)}_{2}={(712)}_{8} $
 
 说明:小数点前的高位零和小数点后的低位零可以去除!

## 位运算

`位运算符`

这是一个包含常用位运算符的表格，展示了这些运算符在不同编程语言中的使用方式，并提供了示例代码。

| 位运算符 | 描述 | Golang 示例 | C++ 示例 | PHP 示例 | Java 示例 | Python 示例 |
|----------|------|-------------|----------|----------|-----------|-------------|
| `&`      | 与运算符，将两个操作数的对应位进行逻辑与操作 | `a & b` | `a & b` | `$a & $b` | `a & b` | `a & b` |
| `\|`      | 或运算符，将两个操作数的对应位进行逻辑或操作 | `a \| b` | `a \| b` | `$a \| $b` | `a \| b` | `a \| b` |
| `^`      | 异或运算符，将两个操作数的对应位进行逻辑异或操作 | `a ^ b` | `a ^ b` | `$a ^ $b` | `a ^ b` | `a ^ b` |
| `~`      | 非运算符，对操作数的每个位进行逻辑非操作 | `~a` | `~a` | `~$a` | `~a` | `~a` |
| `<<`     | 左移运算符，将操作数的所有位向左移动指定的位数 | `a << n` | `a << n` | `$a << $n` | `a << n` | `a << n` |
| `>>`     | 右移运算符，将操作数的所有位向右移动指定的位数 | `a >> n` | `a >> n` | `$a >> $n` | `a >> n` | `a >> n` |
| `>>>`    | 无符号右移运算符，将操作数的所有位向右移动指定的位数 | N/A | N/A | N/A | N/A | N/A |

*无符号右移运算符（`>>>`）在Golang语言中没有对应的运算符。在示例中，`a`和`b`代表操作数，`n`代表移动的位数。*

`与运算 & and`

 and运算通常用于二进制的取位操作，例如一个数 and 1的结果就是取二进制的最末位。
 
 这可以用来判断一个整数的奇偶，二进制的最末位为0表示该数为偶数，最末位为1表示该数为奇数。

 **相同位的两个数字都为1，则为1; 若有一个不为1，则为0。**

 ```text
 00101
 11100 &
 --------
 00100
 ```

`或运算 | or`

 or运算通常用于二进制特定位上的无条件赋值，例如一个数or 1的结果就是二进制最未位强行变成1。如果需要把二进制最末位变成0，对这个数or 1之后再减一就可以了，其实际意义就是把这个数强行变成最接近的偶数。

 一个数或上0，还是该数，也就是说，一个数字或上0并不会改变此数字的大小。

**相同位只要一个为1即为1。**

 ```text
 00101
 11100 |
 --------
 11101
 ```

`异或运算 ^ xor (无进位加法)`

 异或的符号是^。按位异或运算,对等长二进制模式按位或二进制数的每一位执行逻辑按位异或操作.操作的结果是如果某位不同则该位为1，否则该位为0.

 xor运算的逆运算是它本身，也就是说两次异或同一个数最后结果不变，即(a xor b) xor b= a。

 xor运算可以用于简单的加密，比如你想对你MM说1314520，但怕别人知道，于是双方约定拿你的生日19880516作为密钥。1314520 xor 19880516 = 20665500，你就把20665500告诉MM。MM再次计算20665500 xor 19880516的值，得到1314520。

**相同位不同则为1，相同则为0。**

 ```text
 00101     11001
 11100 ^   11100 ^
 ------------------
 11001     00101
 ```

`非运算 ~ not`

 not运算的定义是把内存中的0和1全部取反。使用not运算时要格外小心，你需要注意整数类型有没有符号。
 
 如果not的对象是无符号整数 (不能表示负数)，那么得到的值就是它与该类型上界的差，因为无符号类型的数是用00到$FFFF依次表示的。

`原码`

 **原码:是最简单的机器数表示法。用最高位表示符号位，"1”表示负号，"0"表示正号。其他位存放该数的二进制的绝对值.**

 若以带符号位的四位二进值数为例: 

 ```text
 1010 : 最高位为'1’,表示这是一个负数，其他三位为'010'，即(0*2^2) +(1*2^1) + (0*2^0) = 2 ('^'表示幂运算符)
        所以1010表示十进制数 (-2).
 ```

 下面表格给出部份正负数的二进制原码的表示法: 

 | 十进制数 | 二进制原码 |
 |:--------:|:----------:|
 | 0        | `00000000` |
 | -0       | `10000000` |
 | 10       | `00001010` |
 | -10      | `10001010` |
 | 25       | `00011001` |
 | -25      | `10011001` |
 | 50       | `00110010` |
 | -50      | `10110010` |
 | 100      | `01100100` |
 | -100     | `11100100` |
 | 127      | `01111111` |
 | -127     | `11111111` |

 OK，原码表示法很简单有没有，虽然出现了+0和-0，但是直观易懂.但是会有问题

 ```text
 0001 + 0010 = 0011 (1 + 2 = 3) 没问题
 0000 + 1000 = 1000 (0 + -0 = 0) 也没啥问题
 0001 + 1001 = 1010 (1 + (-1) = -2) ?????????(WTF)
 ```

 1+ (-1) =-2，这属于是离谱的妈妈给离谱开门 -> 离谱到家了

 于是我们可以看到其实正数之间的加法通常是不会出错的，因为它就是一个很简单的二进制加法。

 而正数与负数相加，或负数与负数相加，就要引起莫名其妙的结果，这都是该死的符号位引起的。
 0分为+0和-0也是因他而起。

 所以原码，虽然直观易懂，易于正值转换。但用来实现加减法的话，运算规则总归是太复杂。于是反码来了。

`反码`

 我们知道，原码最大的问题就在于一个数加上他的相反数不等于零
 
 例如: 0001 + 1001=1010(1 +(-1)= -2)  0010 + 1010 = 1100(2 +(-2) = -4)

 于是反码的设计思想就是冲着解决这一点，既然一个负数是一个正数的相反数，那我们干脆用一个正数按位取反来表示负数试试。
 
 **正数的反码还是等于原码**
 **负数的反码就是他的原码除符号位外，按位取反。**

 若以带符号位的四位二进制数为例:

 ```text
 3  是正数，反码与原码相同，则可以表示为 0011 
 -3 原码是 1011，符号位保持不变，低三位(011) 按位取反得(100)
 所以 -3 的反码为 1100
 ```

 下图给出部分正负数的二进制数反码表示法:
 
 | 十进制数 | 原码 | 反码 | 十进制数 | 原码 | 反码 |
 |:------:|:------:|:------:|:------:|:------:|:------:|
 |  0     |  0000    |  0000    |  -0    |  1000    |  1111    |
 |  1     |  0001    |  0001    |  -1    |  1001    |  1110    |
 |  2     |  0010    |  0010    |  -2    |  1010    |  1101    |
 |  3     |  0011    |  0011    |  -3    |  1011    |  1100    |
 |  4     |  0100    |  0100    |  -4    |  1100    |  1011    |
 |  5     |  0101    |  0101    |  -5    |  1101    |  1010    |
 |  6     |  0110    |  0110    |  -6    |  1110    |  1001    |
 |  7     |  0111    |  0111    |  -7    |  1111    |  1000    |

 对着上图，我们再试着用反码的方式解决一下原码的问题

 ```text
 # 以下皆表示反码
 0001 + 1110 = 1111 (1+ (-1) = -0)
 互为相反数相加等于0，解决。虽然是得到的结果是 1111 也就是 -0
 ```

 再试着做一下两个负数相加

 ```text
 1110 (-1) + 1101 (-2) = 1011 (-4)
 (-1) + (-2) = (-4) ?(WTF)
 ```

 不过好像问题不大，因为1011 (是-4的反码，但是从原码来看，他其实是-3。巧合吗?)

 再看个例子

 *1110 (-1) + 1100 (-3) = 1010 (-5)*
 
 确实是巧合，看来相反数问题是解决了，但是却让两个负数相加的出错了。

 **两个正数相加和两个负数相加，其实都是一个加法问题，只是有无符号位罢了。而正数+负数才是真正的减法问题。**

 也就是说只要正数+负数不会出错，那么就没问题了。负数加负数出错没关系的，负数的本质就是正数加上一个符号位而已。

 在原码表示法中两个负数相加，其实在不溢出的情况下结果就只有符号位出错而已
 
 *1001(原码) + 1010(原码) = 0011(原码)*

 反码的负数相加出错，其实问题不大。我们只需要在实现两个负数加法时，将两个负数反码包括符号位全部按位取反相加，然后再给他的**符号位强行置“1”**就可以了。

 所以反码表示法其实已经解决了减法的问题，他不仅不会像原码那样出现两个相反数相加不为零的情况，而且对于任的一个正数加负数，如: 0001 (1) + 1101 (-2) = 1110 (-1) 计算结果是正确的。所以反码与原码比较，最大的优点，就在于解决了减法的问题。

 但是我们还是不满足为什么 0001 + 1110 = 1111 (1+ (-1) =-0) 为什么是-0

 而且虽然说两个负数相加问题不大，但是问题不大，也是问题呀。好吧，处女座。接下来就介绍我们的大boss补码。

`补码`

 ```text
 补码:  正数的补码等于他的原码
        负数的补码等于反码+1。 (这只是一种算补码的方式，很多相关书籍对于补码就是这句话)
 
 在《计算机组成原理中》，补码的另外一种算法: 负数的补码等于他的原码自低位向高位开始的第一个“1’及其右边的保持不变，左边的各位按位取反，符号位不变。还是莫名其妙有没有，为什么补码等于反码加一，为什么自低位向高位取反......?
 
 其实上面那两段话，都只是补码的求法，而不是补码的定义。很多人以为求补码就要先求反码，其实并不是。那些天才的计算机学家，并不会心血来潮的把反码+1就定义为补码。只不过是补码正好就等于反码加1罢了。所以，忘记那些书上那句负数的补码等于它的反码+1。就这句话把我们带入了理解的误区.这就是为什么《计算机组成原理》这本书中，要特意先讲补码，再讲反码!

 然后说负数的补码等于他的原码自低位向高位开始数的第一个”1”及其右边的 保持不变左边的各位按位取反，符号位不变。但是这句话，同样不是补码的定义，它只是补码的另外一种求法。它的存在，告诉我们忘记那句该死的"反码+1"它并不是必须的!
 
 如果你有兴趣了解，补码的严格说法，建议可以看一下《计算机组成原理》。它会用 "模" 和 "同余" 的概念，严谨地解释补码。
 ```

 **补码的思想:**

 ```text
 补码的思想，第一次见可能会觉得很绕，但是如果你肯停下来仔细想想，绝对会觉得非常美妙,补码的思想其实就来自于生活，只是我们没注意到而已。时钟，经纬度，，《易经》里的八卦。那我们来看下生活中的时钟!

 如果说现在时针现在停在10点钟，那么什么时候时针会停在八点钟呢? 简单，过去隔两个小时的时候，是八点钟。未来过十个小时的时候也是八点钟也就是说时间正拨10小时，或是倒拨2小时都是八点钟。也就是10 - 2 = 8，而且 10 + 10 = 8 (10+10 = 10+2+8 = 12+8 = 8) 这个时候满12说明时针在走第二圈了，又走了8小时，所以时针正好又停在八点钟,所以12在时钟运算中，称之为模，超过了12就会重新从1开始算了。也就是说，10-2和10+10从另一个角度来看是等效的，它都使时针指向了八点钟。

 既然是等效的，那在时钟运算中，减去一个数，其实就相当于加上另外一个数 (这个数与减数相加正好等于12)，也称为同余数。即: n 的同余数 = 模 n! 这就是补码所谓模运算思想的生活例子。
 ```

 **补充**

 ```text
 “模”是指一个计量系统的计数范围。如时钟等。计算机也可以看成一个计量机器，它也有一个计量范围，即都存在一个“模”。例如: 时钟的计量范围是0~ 11，模=12。
 
 表示n位的计算机计量范围是0~2^(n)-1，模=2^(n)。“模”实质上是计量器产生“溢出”的量，它的值在计量器上表示不出来，计量器上只能表示出模的余数。任何有模的计量器，均可化减法为加法运算。(把减去一个数等价转化为加上一个正数)

 在这里，我们再次强调原码，反码，补码的引入是为了解决做减法的问题。在原码，反码表示法中，我们把减法化为加法的思维是减去一个数，等于加上一个数的相反数，结果发现引入了符号位，却因为符号位造成了各种意向不到的问题.

 但是从上面的例子中，我们可以看到: 对于数值有限制、有溢出的运算(模运算)来说，减去一个数，其实也相当于加上这个数的同余数。
 也就是说，我们不引入负数的概念，就可以把减法当成加法来算。所以接下来我们聊4位二进制数的运算，也不必急于引入符号位。因为补码的思想，把减法当成加法时并不是必须要引入符号位的。
 ```

 **补码的例子**
 
 接下来我们就做一做四位二进制数的减法吧(先不引入符号位)

 *0110 (6) - 0010 (2) [6 -2 = 4，但是由于计算机中没有减法器，我们没法算]*

 这个时候，我们想想时钟运算中，减去一个数，是可以等同于加上另外一个正数 (同余数),那么这个数是什么呢? 从时钟运算中我们可以看出这个数与减数相加正好等于模。

 那么四位二进制数的模是多少呢? 也就是说四位二进制数最大容量是多少? 其实就是2^4=16=10000B, 那么2的同余数，就等于10000-0010=1110 (14)

 既然如此,  0110 (6) - 0010 (2) = 0110 (6) + 1110 (14) = 10100 (20 = 16 + 4)

 OK，我们看到按照这种算法得出的结果是10100，但是对于四位二进制数，最大只能存放4位(硬件决定了) ，如果我们低四位，正好是0100 (4) ，正好是我们想要的结果，至于最高位的'1，计算机会把他放入 psw 寄存器进位位中。8位机则会放在 cy 中，x86 会放在 cf中 (这个我们不作讨论)

 这个时候，我们再想想在四位二进制数中，减去2，就相当于加上它的同余数14 (至于它们为什么同余，还是建议看《计算机组成原理》)

 但是减去2，从另外一个角度来说，也是加上(-2) 。即加上(-2) 和加上14其实得到的二进制结果除了进位位，结果是一样的。如果我们把1110 (14)的最高位看作符号位后就是 (-2) 的补码，这可能也是为什么负数的符号位是"1"而不是"0"，而且在有符号位的四位二进制数中，能表示的只有-8~7，而无符号位数 (14)的作用和有符号数 (-2)的作用效果其实是一样的。

 那正数的补码呢?加上一个正数，加法器就直接可以实现。所以它的补码就还是它本身。

 ![1.png](https://s2.loli.net/2023/06/20/QVGJKpbSouHcxWO.png)

 到这里，我们发现原码，反码的问题，补码基本都解决了。在补码中也不存在负零了，因为1000表示-8

 这是因为根据上面的补码图，做减法时，0001 (1) + 1111 (-1) = 0000,再也不需要一个1000 来表示负 0 了，就把它规定为 -8, 负数与负数相加的问题也解决了 1111 (-1) + 1110 (-2) = 1101 (-3)
 
 可能说得有点绕，但是实在是没办法。补码还可以这样画。

 **补码的溢出**

 四位二进制表示的补码的区间为-8 ~ +7,结果超出此范围就会产生补码溢出的现象。八位二进制表示的补码的区间为-128 ~ + 127，结果超出此范围就会产生补码溢出的现象
 
 例子:

 ```text
 十进制数相加: (-15) +(-20) = (-35)
 用补码计算，就会有自然丢弃的现象。
 (-15)补= 1111 0001
 (-20)补= 1110 1100
 相加: - - - - - - - -
 可得: (1)1101 1101
 括号中的1，是进位，自然丢弃。剩下的 1101 1101，就是(-35)补。
 ```

 八位的补码，可以表示十进制数 -128~+127,运算结果超出这个范围，就溢出了。
 十进制数计算: (+80) +(+90) = (+170)用补码计算，就会有溢出。
 
 ```text
 (+80)补= 0101 0000
 (+90)补= 0101 1010
 相加: - - - - - - -
 可得: (0) 1010 1010
 ```

 进位是0。剩下的 1010 1010，却是(-86)补。为什么不是(+170)补? 是因为超出范围溢出了

 **处理溢出问题可以用变形补码，其实就是在末位加一个符号位的数即可**
 **取反运算和反码不同!!!!**

`位运算的简单应用`

 | 功能 | 示例 | 位运算(x为int列子变量名) |
 | :---------------------------: | :---------------------------------: | :------------------------------:|
 | 去掉最后一位                   | 110011 -> 11001                     | x >> 1                          |
 | 在最后加一个0                  | 110011 -> 1100110                   | x << 1                          |
 | 在最后加一个1                  | 110011 -> 1100111                   | (x << 1) + 1                    |
 | 把最后一位变成1                | 101100 -> 101101                    | x \| 1                          |
 | 把最后一位变成0                | 101101 -> 101100                    | (x \| 1) - 1                    |
 | 最后一位取反                   | 101101 -> 101100                    | x ^ 1                           |
 | 把右数第k位变成1               | 101001 -> 101101, k = 3             | x ^ (1 << (k - 1))              |
 | 把右数第k位变成0               | 101101 -> 101001, k = 3             | x & (~ (1 <<(k - 1)))           |
 | 取末三位                       | 101101 -> 101                       | x & 7                           |
 | 取右数第k位                    | 1101101 -> 1, k = 4                 | x >> (k-1) & 1                  |
 | 把末k位变成1                   | 1101101 -> 1, k = 4                 | x \| ((1 << k) - 1)             |
 | 末k位取反                      | 101001 -> 100110, k =4              | x ^ ((1 << k) - 1)              |
 | 把右边连续的1变成0             | 100101111 -> 100100000              | x & (x + 1)                     |
 | 把右起第一个0变成1             | 100101111 -> 100111111              | x \| (x + 1)                    |
 | 取右边连续的1                  | 100101111 -> 1111                   | (x ^ (x + 1)) > 1               |
 | 去掉右起第一个1的左边          | 100101000 -> 1000                   | x & (x)                          |
 | 从右边开始，把最后一个1改写成0  | 100101000->100100000                | x & (x - 1)                      |