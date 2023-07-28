# 高精度算法

## 定义和基本思路

 解决超过了精度范围的数做加减乘除是高精度算法的目标

 基本思路: 

 ```text
 使用类似list的数据结构来装载数据,然后就是加减乘除的进位和借位
 ```

## 加法实现
 
 让类似list这种的先来装我们的数字，结果也用类似list这种的来装我们的数字，然后诸位相加。

 思考一下小学是如何列竖式的:

  ![1.png](https://s2.loli.net/2023/07/28/J2MCXhKSoOEv1lu.png)

 若 t1 表示第一个数当前位数的大小，t2 表示第二个数当前位数的大小，next 表示进位数,从个位数开始相加，使用 t 记录 (t1 + t2 + next) 得出的结果，t % 10 为该位数确定好的元素，进行下一个位数操作时，需要 t /= 10。

 最后的进位可能会大于0，需要进行额外判断。

 **注意: 可能两数的长度不一样所以我们需要判断一下**

  ![2.png](https://s2.loli.net/2023/07/28/prClw25jYMtShfv.png)

 答案 t = (t1 + t2 + next)，next为加完之和的进位。 next = t / 10.而该位上的答案则是 t % 10

 **注意: 加法是从低位到高位相加，数据输入时是从高位到低位相加，那么我们在遍历类似list的时候要从后往前进行遍历。**

```go
func main() {
	fmt.Println("请输入a和b的值,空格隔开")
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	inputArr := strings.Split(data, " ")
	a, b := inputArr[0], inputArr[1]
	aArr, bArr := []int{}, []int{}
	for i := len(a) - 1; i >= 0; i-- {
		aArr = append(aArr, int(a[i]-'0'))
	}

	for i := len(b) - 1; i >= 0; i-- {
		bArr = append(bArr, int(b[i]-'0'))
	}

	res := add(aArr, bArr)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
}

func add(a, b []int) []int {
	if len(a) < len(b) {
		return add(b, a)
	}
	res, t := []int{}, 0
	for i := 0; i < len(b); i++ {
		t = a[i] + b[i]
		res = append(res, t%10)
		t = t / 10
	}
	if t != 0 {
		res = append(res, t)
	}
	return res
}
```

## 减法实现

 这里实现的是让大数减去小数的操作。**实际上小数减去大数也就是大数减去小数加一个减号就可以了。**
  
 还是考虑一下列竖式减法是如何做的:

  ![3.jpg](https://s2.loli.net/2023/07/28/wxYdnLRKcCBJZhz.png)

 那么我们现在知道了两个大的数字是通过容器来储存的(java的List，C++的vector,Python的list)，那么加法操作就是诸位相加。其实减法的话也很简单，就是诸位相减。

  1. 首先我们要确定两个相减的数字谁大谁小，所以我们需要自己写一个帮助函数 (方法)来确定哪个比较大。 我们默认的是第一个数字大，第二个数字小，如果第一个比第二个小则需要通过我们写的帮助函数 (方法)来交换两个数字

  2. 运用借位当作答案，也当作借位时用的借位符。公式为 Ai - Bi - t。 当然-Bi的这个操作也是可能不进行的,因为有可能会出现Bi数位不够的情况。数位不够时答案则为 Ai - t。如果 t 为负数，说明我们的A数这一位上的数要比B数这一位上的数要小，那么答案就是(t + 10) % 10 。否则就是 t。

   ![4.png](https://s2.loli.net/2023/07/28/qMZdaupLHRI3ywS.png)

   算完每一位后，判断一下进位t，如果进位为负数说明我们向高位进行了借位，那将t置为1，否则就置为0。

 **最后对于减法，还有一个特殊情况:**
    
   ![5.png](https://s2.loli.net/2023/07/28/1TauSUkZCK5hyWw.png)

   我们的答案不可能为 0003，所以需要去掉后面没有用的0，直到后面的一位不为0的时候结束。

```go
func main() {
	fmt.Println("请输入a")
	a, _ := reader.ReadString('\n')
	fmt.Println("请输入b")
	b, _ := reader.ReadString('\n')
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)
	aArr, bArr := []int{}, []int{}
	for i := len(a) - 1; i >= 0; i-- {
		aArr = append(aArr, int(a[i]-'0'))
	}

	for i := len(b) - 1; i >= 0; i-- {
		bArr = append(bArr, int(b[i]-'0'))
	}
	flag := isMoreThan(aArr, bArr)
	var res []int
	switch flag {
	case 2:
		res := sub(bArr, aArr)
		res[0] *= -1
	case 3:
		res = []int{0}
	default:
		res = sub(aArr, bArr)
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
}

func sub(a, b []int) []int {
	res, t := []int{}, 0
	for i := 0; i < len(a); i++ {
		if i < len(b) {
			t = a[i] - b[i] - t
		} else {
			t = a[i] - t
		}
		res = append(res, ((t + 10) % 10))
		if t < 0 {
			t = 1
		}else {
			t = 0
		}
	}
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	return res
}

func isMoreThan(a, b []int) int {
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] == b[i] {
				continue
			}
			if a[i] > b[i] {
				return 1
			} else {
				return 2
			}
		}
		return 3
	}
	if len(a) > len(b) {
		return 1
	} else {
		return 2
	}
}
```

## 乘法实现

 由于数据过大，暂时只实现高精度乘法针对**大数乘较小的整数**的范围

  一个大数 A 乘以一个小数 b，答案是什么?

   ![6.png](https://s2.loli.net/2023/07/28/UhzMQlf7IP6AVwF.png)
 
 结论: 

  1. A0 * b 这个位置上，这一位的结果为 A0 * b % 10
  2. 进位 t 为 A0 * b / 10
  3. 下一位上 A1 这一位的结果为(A1 * b + t) % 10

 示例: 

  ![7.png](https://s2.loli.net/2023/07/28/6xeUjoZ7muiWfdB.png)

  **那么我们的数字t的公式: t += nums[i] * b，那么每一位上的数字就是 t % 10**

  *注意: 最后假如我们乘的是0的话，答案为0，但是如果一个大数乘以0，最后可能是00000....所以需要去掉多余的0。*

```go
func main() {
	fmt.Println("请输入a")
	aStr, _ := reader.ReadString('\n')
	fmt.Println("请输入b")
	bStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)
	b, _ := strconv.Atoi(strings.TrimSpace(bStr))
	a := []int{}
	for i := len(aStr) - 1; i >= 0; i-- {
		a = append(a, int(aStr[i]-'0'))
	}
	res := mul(a, b)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
}

func mul(a []int, b int) []int {
	res, t := []int{}, 0
	for i := 0; i < len(a) || t > 0; i++ {
		if i < len(a) {
			t += a[i] * b
		}
		res = append(res, t%10)
		t = t / 10
	}
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	return res
}
```

## 除法实现
 
 和乘法类似,这里也是只实现**大数除以较小的整数**

 首先利用竖列式看一下除法情况:
  
  ![8.png](https://s2.loli.net/2023/07/28/Voaz7YILRGZc824.png)

  通过列竖式,可以总结一下:

  那么假如我们有一个数A: A5A4A3A2A1，那么每一步的求法就是:

   ![9.png](https://s2.loli.net/2023/07/28/ZjnLrpUf17NlPac.png)

   L5 = (A5 % b) * 10 + A4
  
  先求出L1,找规律,那么R4的答案不难看出是: R4 = L5 / b

  最关键的两个公式就都得到了，循环就可以得出答案了。
