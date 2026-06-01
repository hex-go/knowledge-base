---
title: "GMP 调度器专题"
category: go/runtime
domain: go
---

# GMP 调度器专题

Go 的核心——M:N 协程调度模型。

## 中级 · 模型概览

- [[go/runtime/gmp/gmp-overview]] — G（goroutine）、M（machine / OS 线程）、P（processor / 逻辑处理器）三要素、M:N 调度关系

  前置：[[go/concurrency/goroutine]]

## 高级 · 调度细节

- [[go/runtime/gmp/gmp-detail]] — work-stealing 机制、sysmon 监控线程、Go 1.14+ 基于信号的异步抢占

  前置：[[go/runtime/gmp/gmp-overview]]
