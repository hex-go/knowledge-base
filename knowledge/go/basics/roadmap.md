---
title: "Go 语言基础"
category: go/basics
domain: go
---

# Go 语言基础

Go 的类型系统、数据结构、语法机制，从基本使用到底层原理。

## 初级 · 语法与日常使用

- [[go/basics/naming-conventions]] — 命名规范
- [[go/basics/variables-constants]] — 常量与变量
- [[go/basics/operators]] — 运算符
- [[go/basics/control-flow]] — 控制语句 (if、for、switch)
- [[go/basics/pointer]] — 指针
- [[go/basics/func-method]] — 函数与方法
- [[go/basics/error-handling]] — 异常处理 (error、panic、recover)
- [[go/basics/init-order]] — init 执行顺序
- [[go/basics/go-modules]] — 依赖管理
- [[go/basics/memory-allocation]] — new 与 make 的区别

## 中级 · 类型系统深入

- [[go/basics/generics]] — 泛型
- [[go/basics/string/string-internals]] — string 不可变性、底层 struct、与 []byte 转换
- [[go/basics/slice/slice-expansion]] — slice 扩容策略与内存分配
- [[go/basics/struct/struct-alignment]] — struct 内存对齐与 padding
- [[go/basics/defer/defer-internals]] — defer 执行时机、闭包陷阱、性能

## 高级 · 底层数据结构

- [[go/basics/slice/slice-header]] — slice 底层 header、与数组共享内存
- [[go/basics/map/map-internals]] — map 哈希桶、溢出桶、扩容
- [[go/basics/interface/interface-dispatch]] — iface / eface、动态派发

## 子模块

- [[go/basics/string/roadmap]] — string 专题
- [[go/basics/slice/roadmap]] — slice 专题
- [[go/basics/map/roadmap]] — map 专题
- [[go/basics/struct/roadmap]] — struct 专题
- [[go/basics/interface/roadmap]] — interface 专题
- [[go/basics/defer/roadmap]] — defer 专题
