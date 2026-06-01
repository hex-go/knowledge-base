---
title: "切片扩容容量计算 — 参考答案"
category: go/exercises/slice
---

## 参考实现

```go
func CalcNewCap(oldCap, numNew int) int {
    needed := oldCap + numNew

    if needed > oldCap*2 {
        return needed
    }

    const threshold = 256

    if oldCap < threshold {
        return oldCap * 2
    }

    newCap := oldCap
    for newCap < needed {
        newCap += (newCap + 3*threshold) / 4
    }
    return newCap
}
```

## 关键点

1. 先算"需求是否超翻倍"—超了直接按需分配。
2. 小切片（<256）简单翻倍。
3. 大切片（≥256）用渐进公式 `(cap + 3*256)/4` 循环逼近，而非固定乘 1.25。
