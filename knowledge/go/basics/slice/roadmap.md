---
title: "slice 专题"
category: go/basics
domain: go
---

# slice 专题

Go 中最核心的数据结构，从基本使用到底层内存管理。

## 初级 · 基本使用

- [[go/basics/slice/slice-basics]] — 声明、切片操作、append、copy

  需要先掌握 [[go/basics/pointer]] 和数组的基本概念。

## 中级 · 扩容机制

- [[go/basics/slice/slice-expansion]] — 容量增长策略、内存分配行为

  前置：[[go/basics/slice/slice-basics]]

## 高级 · 底层结构

- [[go/basics/slice/slice-header]] — reflect.SliceHeader、与底层数组共享内存、GC 影响

  前置：[[go/basics/slice/slice-expansion]]
