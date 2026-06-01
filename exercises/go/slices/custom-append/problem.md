---
title: "自定义切片追加（扩容模拟）"
category: internals/slice-array
tags: [slice, append, 扩容]
difficulty: medium
status: learning
---
# 自定义切片追加（扩容模拟）

## 题目描述

Go 内置的 `append` 函数会自动处理切片扩容，但你真的理解它是怎么做的吗？

请实现函数 `AppendInt(s *[]int, elems ...int)`，它模拟 Go 内置 `append` 的行为：

1. 如果原切片容量足够（`len + len(elems) <= cap`），直接在原底层数组上追加
2. 如果容量不够，分配新的底层数组，容量按以下规则计算：
   - 旧容量 **< 256**：翻倍（×2），如果还不够就继续翻倍直到装下
   - 旧容量 **>= 256**：每次增加 25%（`+ oldCap/4`），如果还不够就继续增加直到装下
   - 如果旧容量为 0，新容量直接取 `newLen`
3. 将旧元素拷贝到新数组，再追加新元素
4. 通过指针 `*s` 更新切片

## 输入/输出示例

```go
s := make([]int, 0, 2)
AppendInt(&s, 1)
// s = [1], len=1, cap=2

AppendInt(&s, 2, 3)  // 容量不够，触发扩容：oldCap=2 < 256，翻倍得 4
// s = [1,2,3], len=3, cap=4

AppendInt(&s, 4, 5)  // 容量又不够：oldCap=4 < 256，翻倍得 8
// s = [1,2,3,4,5], len=5, cap=8

big := make([]int, 256, 256)
AppendInt(&big, 1)  // oldCap=256 >= 256，扩容 25% → 320
// cap=320
```

## 约束

- 使用标准库，不引入第三方依赖
- 不要调用内置 `append`，自己实现扩容逻辑
- 处理好边界：`elems` 为空、`oldCap == 0`、一次追加多个元素
