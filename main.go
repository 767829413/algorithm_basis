package main

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

}

const lenHash = 99997

type bucket struct {
	key int
	val int
}

type MyHashMap struct {
	l    int
	data [99997][]*bucket
}

func Constructor() MyHashMap {
	return MyHashMap{
		l:    lenHash,
		data: [99997][]*bucket{},
	}
}

func (this *MyHashMap) Put(key int, value int) {
	idx := key % this.l
	if this.data[idx] == nil {
		this.data[idx] = []*bucket{}
	}
	for _, v := range this.data[idx] {
		if v.key == key {
			if v.val == value {
				return
			} else {
				v.val = value
				return
			}
		}
	}
	this.data[idx] = append(this.data[idx], &bucket{key, value})
}

func (this *MyHashMap) Get(key int) int {
	idx := key % this.l
	if this.data[idx] == nil {
		return -1
	}
	for _, v := range this.data[idx] {
		if v.key == key {
			return v.val
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := key % this.l
	if this.data[idx] == nil {
		return
	}
	indexs := []int{}
	for k, v := range this.data[idx] {
		if v.key == key {
			indexs = append(indexs, k)
		}
	}
	if len(indexs) == 0 {
		return
	}
	for _, index := range indexs {
		switch index {
		case 0:
			this.data[idx] = this.data[idx][1:]
		case len(this.data[idx]) - 1:
			this.data[idx] = this.data[idx][:len(this.data[idx])-1]
		default:
			this.data[idx] = append(this.data[idx][:index], this.data[idx][index:]...)
		}
	}
}
