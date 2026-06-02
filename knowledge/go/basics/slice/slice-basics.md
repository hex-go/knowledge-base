---
title: 'Slice 基本使用'
category: go/basics/slice
tags: [slice, append, copy, 切片, make]
difficulty: easy
status: learning
last_review: 2026-06-01
review_interval_days: 7
level: beginner
domain: go
mastery:
  wrong_count: 0
  last_wrong_date: ''
  review_count: 1
---
# Slice 基本使用

## 标准知识

### 声明与初始化

```go
var s []int                  // nil slice, len=0, cap=0
s := []int{1, 2, 3}          // 字面量初始化, len=3, cap=3
s := make([]int, 5)          // len=5, cap=5, 零值填充
s := make([]int, 3, 5)       // len=3, cap=5, 预分配容量
```

### 切片操作

```go
s := []int{0,1,2,3,4}
s[1:3]   // [1,2]  下标 1 到 3-1
s[:2]    // [0,1]  从开头到 2-1
s[2:]    // [2,3,4] 从 2 到末尾
s[:]     // [0,1,2,3,4] 全切片
```

切片不复制数据，与原数组/切片共享底层内存。修改切片元素会影响原切片。

### append

```go
s := []int{1, 2}
s = append(s, 3)          // [1,2,3]
s = append(s, 4, 5, 6)    // [1,2,3,4,5,6]
s = append(s, other...)   // 合并另一个切片
```

append 返回新切片头。如果容量不足会自动扩容（分配新的底层数组）。必须接收返回值，否则可能丢数据。

### copy

```go
src := []int{1, 2, 3}
dst := make([]int, 2)
n := copy(dst, src)  // n=2, dst=[1,2]
```

copy 按 min(len(dst), len(src)) 复制元素，不会扩容目标切片。返回实际复制的元素个数。

### nil slice vs empty slice

```go
var s1 []int        // nil, len=0, s1 == nil → true
s2 := []int{}       // empty, len=0, s2 != nil
s3 := make([]int,0) // empty, len=0, s3 != nil
```

nil slice 和 empty slice 的 len/cap 都是 0，对 append/len/cap/range 行为一致。区别在于 JSON 序列化：nil → null，empty → []。

### 值类型 vs 引用语义

slice 本身是值类型（一个包含 ptr/len/cap 的 struct），赋值和传参会复制这个 header，但底层数组共享。所以通过一个切片修改元素会影响另一个，但 append 触发的扩容只改变当前切片的 header。

## 我的理解

### 声明
2种方式：
- var 关键字声明（声明后的变量是nil，需要赋值后，才分配内存空间）
- := 语法糖声明（用的范围广、方便，但是只能在函数内使用，不能在包级别使用）
  - 搭配 `字面量`，声明和赋值一步到位
  - 搭配 `make关键字`, 类型、长度必填，长度0值填充，容量为预分配空间，容量省略后跟长度相同）
### 特性
- 切片是对array的引用，改变数组的值，切片也会相应发生改变
- append操作，有可能触发[[go/basics/slice/slice-expansion|切片扩容]]，切片一旦扩容，则分配新的底层数组，旧数据拷贝过去，因此切片、原始array的引用关系发生改变，值改变不再传导；
- append操作可以追加多个元素或另一个切片，但一次append操作只会统一触发一次扩容。
- copy操作，目标在前源在后，如果目标切片容量不够，并不会触发扩容，源多余的元素会被丢掉。
- var声明的是 `nil切片`, 字面量空、make是长度容量=0的是`empty切片`
  - 相同点：len/cap都是0，append/len/cap/range的行为都一致；
  - 不同点：json序列化，`nil切片` 是 null，`empty切片` 是 []
- 切片本身是值类型（包含len、cap、ptr的struct），赋值、传参都会复制这个结构体。ptr指向的底层数组是同一个，所以对切片的修改会影响另一个。