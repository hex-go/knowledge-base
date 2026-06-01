---
title: "Go 运行时"
category: go/runtime
domain: go
---

# Go 运行时

调度器、垃圾回收、内存分配——支撑 Go 程序运行的三大底层系统。

## 中级 · 运行时概览

- [[go/runtime/memory-allocation]] — 内存分配器（mspan / mcache / mcentral / mheap）
- [[go/runtime/gmp/gmp-overview]] — GMP 调度模型、M:N 调度

## 高级 · 深入

- [[go/runtime/gmp/gmp-detail]] — work-stealing、sysmon、抢占式调度（1.14+ 异步抢占）
- [[go/runtime/gc/gc-basics]] — 三色标记法、STW、写屏障

## 专家 · 调优

- [[go/runtime/gc/gc-tuning]] — GOGC、GOMEMLIMIT、trace 分析、pprof
- [[go/runtime/netpoller]] — 网络轮询器、epoll、异步 I/O

## 子模块

- [[go/runtime/gmp/roadmap]] — GMP 调度器专题
- [[go/runtime/gc/roadmap]] — 垃圾回收专题
