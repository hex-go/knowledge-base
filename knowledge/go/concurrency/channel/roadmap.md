---
title: "channel 专题"
category: go/concurrency
domain: go
---

# channel 专题

Go 的 CSP 并发模型核心——"不要通过共享内存来通信，而要通过通信来共享内存。"

## 初级 · 基本操作

- [[go/concurrency/channel/channel-basics]] — 创建（有缓冲 / 无缓冲）、发送、接收、关闭、for range

  需要先理解 [[go/concurrency/goroutine]]。

## 中级 · 模式与陷阱

- [[go/concurrency/channel/channel-close]] — 关闭规则：谁创建谁关闭、向已关闭 channel 发送 = panic、从已关闭 channel 接收 = 零值 + false、nil channel 的阻塞行为

  前置：[[go/concurrency/channel/channel-basics]]

- [[go/concurrency/channel/channel-patterns]] — 生产者消费者、扇入 (fan-in)、扇出 (fan-out)、or-done channel、done channel

  前置：[[go/concurrency/channel/channel-basics]]

## 高级 · 底层结构

- [[go/concurrency/channel/channel-internals]] — hchan 结构体、环形队列 (ring buffer)、sendq / recvq 等待队列、阻塞与唤醒机制

  前置：[[go/concurrency/channel/channel-basics]]、[[go/concurrency/goroutine]]
