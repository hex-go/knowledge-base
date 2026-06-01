---
name: kb-practice
description: Guides the knowledge graph internalization stage by generating coding or concept exercises, grading answers, and writing wrong-book records. Use when user says "练习 X", "练习今天的知识点", "闯关 X", "给我出一道编码题", "给我出一道概念题", "根据错题本出一题", or "帮我看看我的答案".
---

# kb-practice

## 触发词

- "练习 X"
- "练习今天的知识点"
- "闯关 X"
- "给我出一道编码题"
- "给我出一道概念题"
- "根据错题本出一题"
- "帮我看看我的答案"

## 工作流

1. 先判断是单知识点、今日批量、综合闯关，还是错题变体。
2. 编码题走 `exercises/`：创建 `problem.md`、`solution.md`、`skeleton/main.go`、`attempts.md`。
3. 概念题走 `concepts/` 标准答案文件与 `-attempts.md` 答题记录，自动写入 `## 考察` 的 `[[knowledge/...]]`。
4. 批改时对照标准答案或 `solution.md`，指出遗漏、错误和可改进处。
5. 通过则更新知识点的掌握计数。
6. 失败则写入 `wrong-book/`，记录错误思维链和正确链路。
7. 如果同一知识点反复出错，提示回到学习阶段补强理解。

## 参考

- `docs/MANUAL.md` §5.2.2
- `docs/MANUAL.md` §3.2
- `docs/MANUAL.md` §3.3
- `docs/MANUAL.md` §3.4
- `docs/MANUAL.md` §4
