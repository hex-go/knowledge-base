---
name: kb-practice
description: Guides the knowledge graph internalization stage by generating coding or concept exercises, grading answers, and writing wrong-book records. Use when user says "练习 X", "练习今天的知识点", "闯关 X", "给我出一道编码题", "给我出一道概念题", "根据错题本出一题", "帮我看看我的答案", "批改", or "做完了".
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
- "批改"
- "做完了"

## 工作流

1. 先判断是单知识点、今日批量、综合闯关，还是错题变体。
2. 编码题走 `exercises/`：创建 `problem.md`、`solution.md`、`skeleton/main.go`、`attempts.md`。
   `.go` 文件用 Write 工具创建（避免 PowerShell 下 heredoc 吃掉反引号）。
   生成后在 WSL 中执行 `code <path>` 打开 `skeleton/main.go`。
   skeleton 模板必须在 `main()` 末尾包含 `fmt.Printf("\n%d/%d passed\n", passed, len(tests))` 测试统计行。
3. 概念题走 `concepts/` 标准答案文件与 `-attempts.md` 答题记录，自动写入 `## 考察` 的 `[[knowledge/...]]`。
4. 用户说"做完了"/"批改"时，AI 直接读取 `skeleton/main.go` 内容，对照 `solution.md` 批改。输出格式：✅ 正确 / ❌ 错误 / + 改进，附用户完整代码。
5. 批改完成后自动将本次尝试写入 `attempts.md`。分隔符 `## 尝试 N · 日期 · ✓/✗`，内容为 AI 批改摘要 + 用户完整代码，按 `docs/MANUAL.md` §3.2 格式。
6. 通过则更新知识点的掌握计数。
7. 失败则写入 `wrong-book/`，记录错误思维链和正确链路。
8. 如果同一知识点反复出错，提示回到学习阶段补强理解。

## 参考

- `docs/MANUAL.md` §5.2.2
- `docs/MANUAL.md` §3.2
- `docs/MANUAL.md` §3.3
- `docs/MANUAL.md` §3.4
- `docs/MANUAL.md` §4
