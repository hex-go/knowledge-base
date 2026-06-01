---
title: "Go 编译器与工具链"
category: go/compiler
domain: go
---

# Go 编译器与工具链

从源码到可执行文件的完整过程，以及编译器的深度优化。

## 初级 · 工具链

- [[go/compiler/go-build]] — 编译过程、交叉编译、go build / run / install

## 高级 · 编译器优化

- [[go/compiler/escape-analysis]] — 逃逸分析：决定变量分配到堆还是栈
- [[go/compiler/inlining]] — 内联优化：减少函数调用开销

## 专家 · 底层

- [[go/compiler/ssa]] — 静态单赋值 (SSA)、中间表示、编译阶段
- [[go/compiler/assembly]] — Go 汇编 (Plan 9)、go tool compile -S
