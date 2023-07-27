package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var arr := utilcom.BuildTestArrInt(9, 100, 0)
// var limit := utilcom.BuildTestArrInt(2, 9, 0)
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
