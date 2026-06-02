---
title: "Slice 基本操作练习"
category: go/exercises/slice
tags: [slice, append, copy, make]
difficulty: easy
status: learning
domain: go
---
# Slice 基本操作练习

## 题目描述

实现三个 slice 操作函数，不允许依赖第三方库：

### CloneSlice(s []int) []int
返回 s 的深拷贝。修改返回值不影响原切片，反之亦然。入参为 nil 时返回 nil。

### InsertAt(s []int, index, value int) []int
在指定位置插入元素，保序。index 取值范围 [0, len(s)]。

### RemoveAt(s []int, index int) []int
删除指定位置元素，保序。index 取值范围 [0, len(s)-1]。

## 考察

[[knowledge/go/basics/slice/slice-basics]]
