---
title: "Slice 基本操作 — 参考答案"
category: go/exercises/slice
---

## 参考实现

```go
func CloneSlice(s []int) []int {
	if s == nil {
		return nil
	}
	dst := make([]int, len(s))
	copy(dst, s)
	return dst
}

func InsertAt(s []int, index, value int) []int {
	return append(s[:index], append([]int{value}, s[index:]...)...)
}

func RemoveAt(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
```

## 关键点

- **CloneSlice**：`make` + `copy` 是标准深拷贝。nil 入参是显式判断，`make([]int, 0)` 返回的是 empty slice 而不是 nil。
- **InsertAt**：利用切片和 append 在中间插入一个小数组。
- **RemoveAt**：`append(s[:index], s[index+1:]...)` 拼接跳过被删元素。
- 三个函数都可能改变底层数组，但通过返回值体现结果。
