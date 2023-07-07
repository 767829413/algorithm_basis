package main

import (
	"bufio"
	"os"

	"github.com/767829413/algorithm_basis/utilcom"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	arr := utilcom.BuildTestArr(5, 100, 1)
	utilcom.PrintSortComparison(arr, BubbleSort)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ { // n - 1 循环,最后一个不管
		for j := 0; j < len(arr)-1-i; j++ { // 第二轮指针,每一轮都会排好一个数
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
