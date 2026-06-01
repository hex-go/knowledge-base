---
title: "GC 专题"
category: go/runtime
domain: go
---

# GC 专题

Go 的垃圾回收——三色标记 + 并发回收。

## 中级 · GC 原理

- [[go/runtime/gc/gc-basics]] — 三色标记法、写屏障、STW 阶段、并发标记与清扫

## 高级 · 调优

- [[go/runtime/gc/gc-tuning]] — GOGC 参数、GOMEMLIMIT、GC trace 分析、pprof heap profile

  前置：[[go/runtime/gc/gc-basics]]
