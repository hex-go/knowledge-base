---
title: "defer 专题"
category: go/basics
domain: go
---

# defer 专题

Go 的延迟执行机制，从简单用法到编译器优化。

## 初级 · 基本使用

- [[go/basics/defer/defer-basics]] — 基本用法、LIFO 执行顺序、参数求值时机

## 中级 · 内部实现

- [[go/basics/defer/defer-internals]] — 堆分配 vs 栈分配（Go 1.14+ 优化）、闭包陷阱、性能开销

  前置：[[go/basics/defer/defer-basics]]
