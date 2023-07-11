package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/767829413/algorithm_basis/utilcom"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	arr := utilcom.BuildTestArrInt(5, 200, 0)
	// arr := utilcom.BuildTestArrFloat64(5)
	//排序前
	fmt.Println(arr)
	fmt.Println(len(arr))
	//排序
	QuickSort(arr, 0, len(arr)-1)
	//排序后
	fmt.Println(arr)
	fmt.Println(len(arr))
}

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
