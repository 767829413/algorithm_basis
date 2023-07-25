package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// var arr := utilcom.BuildTestArrInt(9, 100, 0)
// var limit := utilcom.BuildTestArrInt(2, 9, 0)
var reader = bufio.NewReader(os.Stdin)

func main() {
	var max, min = 100, 0
	var a, b, n, m, q, nums = [8][8]int{}, [8][8]int{}, 5, 5, 1, 3
	var x1, y1, x2, y2 = 1, 1, n - 1, m - 1
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] = rand.Intn(max-min) + min
		}
	}

	// 构造差分数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// insert(i, j, i, j, a[i][j], &b)
			b[i][j] = a[i][j] - a[i-1][j] - a[i][j-1] + a[i-1][j-1]
		}
	}

	for ; q > 0; q-- {
		insert(x1, y1, x2, y2, nums, &b)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[i][j] = b[i-1][j] + b[i][j-1] - b[i-1][j-1] + b[i][j]
		}
	}
	fmt.Println(a)
	fmt.Println(b)

}

func insert(x1, y1, x2, y2, nums int, b *[8][8]int) {
	// b[i][j] = a[i][j] - a[i-1][j] - a[i][j-1] + a[i-1][j-1]
	b[x1][y1] += nums
	b[x1][y2+1] -= nums
	b[x2+1][y1] -= nums
	b[x2+1][y2+1] += nums
}
