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
	// MergeSort(arr, 0, len(arr)-1)
	ShellSort(arr)
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

func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + ((r - l) >> 1)
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	temp, i, j := []int{}, l, mid+1
	for i <= mid || j <= r {
		if i > mid {
			for j <= r {
				temp = append(temp, arr[j])
				j++
			}
			break
		}
		if j > r {
			for i <= mid {
				temp = append(temp, arr[i])
				i++
			}
			break
		}
		if arr[i] < arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}
	for i := l; i <= r; i++ {
		arr[i] = temp[i-l]
	}
}

func ShellSort(arr []int) {
	incream := len(arr) >> 1
	for incream > 0 {
		for i := incream; i < len(arr); i++ {
			idx, cur := i, arr[i]
			for idx >= incream && arr[idx-incream] > cur {
				arr[idx] = arr[idx-incream]
				idx -= incream
			}
			arr[idx] = cur
		}
		incream >>= 1
	}
}

func HeapSort(arr []int) {
	h := NewHeap(arr)
	idx := 0
	for h.Size() != 0 {
		arr[idx] = h.RemoveTop()
		idx++
	}
}

type heap []int

func NewHeap(arr []int) heap {
	heap := heap(append([]int{0}, arr...))
	for i := len(heap) / 2; i > 0; i-- {
		heap.Down(i)
	}
	return heap
}

func (h heap) Top() int {
	if h.Size() == 0 {
		panic("heap is empty")
	}
	return h[1]
}

func (h heap) RemoveTop() int {
	if h.Size() == 0 {
		panic("heap is empty")
	}
	top := h[1]
	h[1] = h[len(h)-1]
	h = h[:len(h)-1]
	h.Down(1)
	return top
}

func (h heap) Size() int {
	l := len(h) - 1
	if l < 0 {
		return 0
	}
	return l
}

func (h heap) Down(idx int) {
	l, r := 2*idx, 2*idx+1
	if l >= h.Size() || r >= h.Size() {
		return
	}
	if h[l] > h[r] {
		h[idx], h[r] = h[r], h[idx]
		idx = r
	} else {
		h[idx], h[l] = h[l], h[idx]
		idx = l
	}
	h.Down(idx)
}

func (h heap) Up(idx int) {
	var f int
	if idx%2 == 0 {
		f = idx / 2

	} else {
		f = (idx - 1) / 2
	}
	if f <= 1 {
		return
	}
	if h[f] > h[idx] {
		h[f], h[idx] = h[idx], h[f]
		idx = f
		h.Up(idx)
	}
}
