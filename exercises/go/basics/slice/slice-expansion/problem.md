---
title: "实现切片扩容容量计算"
category: go/exercises/slice
tags: [slice, append, 扩容, capacity]
difficulty: medium
status: learning
domain: go
---
# 实现切片扩容容量计算

## 题目描述

Go 1.18+ 的切片扩容策略如下：

- 如果所需容量超过当前容量的 2 倍，直接扩容到所需容量。
- 否则，当前容量 < 256 时翻倍；当前容量 ≥ 256 时从当前容量开始，反复增长约 25% 直到满足需求。

请实现 `CalcNewCap(oldCap, numNew int) int`，根据当前容量 `oldCap` 和欲追加的元素个数 `numNew`，返回新容量。

## 输入输出要求

- `oldCap ≥ 0`
- `numNew ≥ 0`
- 返回值为扩容后的容量（不考虑内存对齐向上取整）

## 考察

[[knowledge/go/basics/slice/slice-expansion]]
