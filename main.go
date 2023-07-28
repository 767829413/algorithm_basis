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
	res := div(a, b)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
}

func div(a []int, b int) []int {
	res, t := []int{}, 0
	for i := len(a) - 1; i >= 0; i-- {
		t = t*10 + a[i]
		res = append([]int{t / b}, res...)
		t = t % b
	}
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	return res
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
		} else {
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

func add(a, b []int) []int {
	if len(a) < len(b) {
		return add(b, a)
	}
	res, t := []int{}, 0
	for i := 0; i < len(a); i++ {
		if i >= len(b) {
			t = a[i]
		} else {
			t = a[i] + b[i]
		}

		res = append(res, t%10)
		t = t / 10
	}
	if t != 0 {
		res = append(res, t)
	}
	return res
}
