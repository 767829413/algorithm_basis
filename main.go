package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/767829413/algorithm_basis/utilcom"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	arr := utilcom.BuildTestArrInt(9, 100, -50)
	// arr := utilcom.BuildTestArrFloat64(5)
	//排序前
	fmt.Println(arr)
	fmt.Println(len(arr))
	//排序
	ShellSort(arr)
	//排序后
	fmt.Println(arr)
	fmt.Println(len(arr))
}

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
