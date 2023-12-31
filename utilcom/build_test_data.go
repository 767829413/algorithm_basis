package utilcom

import (
	"fmt"
	"math/rand"
	"time"
)

func BuildTestArrInt(num, max, min int) []int {
	if max < min {
		panic(" max should be less than min")
	}
	rand.Seed(time.Now().UnixNano())
	//构建随机[]int 最大值 63
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		randNum := rand.Intn(max-min) + min
		arr[i] = randNum
	}
	return arr
}

func BuildTestArrFloat64(num int) []float64 {
	rand.Seed(time.Now().UnixNano())
	arr := make([]float64, num)
	for i := 0; i < num; i++ {
		randNum := rand.Float64()
		arr[i] = randNum
	}
	return arr
}

func PrintSortComparisonInt(arr []int, f func(arr []int)) {
	//排序前
	fmt.Println(arr)
	fmt.Println(len(arr))
	//排序
	f(arr)
	//排序后
	fmt.Println(arr)
	fmt.Println(len(arr))
}

func PrintSortComparisonFloat64(arr []float64, f func(arr []float64)) {
	//排序前
	fmt.Println(arr)
	fmt.Println(len(arr))
	//排序
	f(arr)
	//排序后
	fmt.Println(arr)
	fmt.Println(len(arr))
}
