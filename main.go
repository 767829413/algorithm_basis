package main

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	println((4 + 2)>>1)
}
