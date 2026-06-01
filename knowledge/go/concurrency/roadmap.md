---
title: "Go 并发编程"
category: go/concurrency
domain: go
---

# Go 并发编程

goroutine、channel、同步原语、context——Go 并发的四大支柱。

## 初级 · 并发原语

- [[go/concurrency/goroutine]] — 创建、调度、泄漏排查
- [[go/concurrency/select]] — 多路复用、default、超时控制
- [[go/concurrency/channel/channel-basics]] — 创建、发送、接收、关闭
- [[go/concurrency/sync-mutex]] — 互斥锁 Lock / Unlock
- [[go/concurrency/sync-waitgroup]] — Add / Done / Wait
- [[go/concurrency/sync-map]] — 并发安全的 map
- [[go/concurrency/context]] — WithCancel / WithTimeout / WithValue
- [[go/concurrency/timer-ticker]] — time.Timer / time.Ticker

## 中级 · 模式与陷阱

- [[go/concurrency/sync-rwmutex]] — 读写锁，适用场景
- [[go/concurrency/sync-once]] — 单例、线程安全初始化
- [[go/concurrency/goroutine-pool]] — 协程池模式
- [[go/concurrency/channel/channel-close]] — 关闭规则：谁关、何时关、panic 场景
- [[go/concurrency/channel/channel-patterns]] — 生产者消费者、扇入扇出、or-done
- [[go/concurrency/context-internals]] — 树结构、Done channel、传播机制

## 高级 · 底层与无锁

- [[go/concurrency/atomic]] — atomic 包、CAS、lock-free 编程
- [[go/concurrency/sync-map-internals]] — 读写分离、dirty / read、原子操作
- [[go/concurrency/channel/channel-internals]] — hchan 结构、ring buffer、阻塞唤醒

## 子模块

- [[go/concurrency/channel/roadmap]] — channel 专题
