---
title: "自定义切片追加 — 参考答案"
category: go/exercises/slice
tags: [slice, append, 扩容]
difficulty: medium
domain: go
---
# 自定义切片追加 — 参考答案

## 核心思路

1. 计算 `newLen = len(old) + len(elems)`
2. 如果 `newLen <= cap(old)`，直接在原底层数组上操作：扩展长度 + copy 新元素
3. 否则计算新容量 → 分配新数组 → copy 老元素 + 新元素 → 更新指针

## 扩容规则

- `oldCap == 0`：新容量 = `newLen`
- `oldCap < 256`：反复翻倍直到 >= newLen
- `oldCap >= 256`：反复 +25% 直到 >= newLen

## 参考答案

```go
func AppendInt(s *[]int, elems ...int) {
	if len(elems) == 0 {
		return
	}

	old := *s
	newLen := len(old) + len(elems)

	if newLen <= cap(old) {
		// 容量足够，直接在原底层数组上追加
		old = old[:newLen]
		copy(old[len(*s):], elems)
		*s = old
		return
	}

	// 容量不够，需要扩容
	newCap := nextCap(cap(old), newLen)
	// 分配新数组，长度 = newLen，容量 = newCap
	newSlice := make([]int, newLen, newCap)
	copy(newSlice, old)
	copy(newSlice[len(old):], elems)
	*s = newSlice
}

// nextCap 按 Go 规则计算扩容后的新容量
func nextCap(oldCap, newLen int) int {
	if oldCap == 0 {
		return newLen
	}

	newCap := oldCap
	for newCap < newLen {
		if newCap < 256 {
			newCap *= 2
		} else {
			newCap += newCap / 4
		}
	}
	return newCap
}
```

## 关键点

1. **`*s` 的更新**：扩容后底层数组变了，必须通过指针更新切片头，否则调用方看到的还是旧切片。
2. **先 copy 老数据再 copy 新数据**：顺序错了会导致数据被覆盖。
3. **`make([]int, newLen, newCap)`**：第三个参数是容量，别写错。
4. **边界处理**：`elems` 空切片直接返回，`oldCap == 0` 单独处理避免 ×2 死循环。
