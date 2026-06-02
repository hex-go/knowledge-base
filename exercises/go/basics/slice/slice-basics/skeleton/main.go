package main

import (
	"fmt"
	"reflect"
)

func CloneSlice(s []int) []int {
	// TODO
	var out []int
	if s != nil {
		out = make([]int, len(s), cap(s))
		copy(out, s)
	}

	return out
}

func InsertAt(s []int, index, value int) []int {
	// TODO
	newL := len(s) + 1
	out := make([]int, newL)
	copy(out, s[:index])
	out[index] = value
	for i := index + 1; i < newL; i++ {
		out[i] = s[i-1]
	}
	return out
}

func RemoveAt(s []int, index int) []int {
	l := len(s)
	for i := index; i < l-1; i++ {
		s[i] = s[i+1]
	}
	s = s[:l-1]
	return s
}

func main() {
	// CloneSlice
	s1 := []int{1, 2, 3}
	s2 := CloneSlice(s1)
	s1[0] = 99
	pass("CloneSlice 深拷贝", !reflect.DeepEqual(s1, s2))
	pass("CloneSlice nil", CloneSlice(nil) == nil)
	pass("CloneSlice empty", len(CloneSlice([]int{})) == 0)

	// InsertAt
	ins := InsertAt([]int{1, 3, 4}, 1, 2)
	pass("InsertAt 中间插入", reflect.DeepEqual(ins, []int{1, 2, 3, 4}))
	pass("InsertAt 头部插入", reflect.DeepEqual(InsertAt([]int{2, 3}, 0, 1), []int{1, 2, 3}))
	pass("InsertAt 尾部插入", reflect.DeepEqual(InsertAt([]int{1, 2}, 2, 3), []int{1, 2, 3}))

	// RemoveAt
	rem := RemoveAt([]int{1, 2, 3}, 1)
	pass("RemoveAt 中间删除", reflect.DeepEqual(rem, []int{1, 3}))
	pass("RemoveAt 头部删除", reflect.DeepEqual(RemoveAt([]int{1, 2, 3}, 0), []int{2, 3}))
	pass("RemoveAt 尾部删除", reflect.DeepEqual(RemoveAt([]int{1, 2, 3}, 2), []int{1, 2}))
}

var total, passed int

func pass(name string, ok bool) {
	total++
	if ok {
		passed++
		fmt.Printf("  PASS  %s\n", name)
	} else {
		fmt.Printf("  FAIL  %s\n", name)
	}
}

func init() {
	fmt.Println("Slice 基本操作练习")
	fmt.Println("==================")
}
