package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/767829413/algorithm_basis/utilcom"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	arr := utilcom.BuildTestArrInt(9, 100, 0)
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

func BubblingSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		preIdx, cur := i-1, arr[i]
		for preIdx >= 0 && arr[preIdx] > cur {
			arr[preIdx+1] = arr[preIdx]
			preIdx--
		}
		arr[preIdx+1] = cur
	}
}

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid, i, j := arr[l+((r-l)>>1)], l, r
	for i < j {
		for arr[i] < mid {
			i++
		}
		for arr[j] > mid {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	QuickSort(arr, l, j)
	QuickSort(arr, j+1, r)
}
