# CLAUDE.md

本文件为 Claude Code 在此仓库中工作时提供指导。

## 仓库定位

这是用户的 Go 语言自学助手。核心思路：**Markdown 知识库（人类可读）+ 本地 Web 浏览 + Claude Code AI 陪练**。

用户的日常节奏：
- **每日**：回顾高频错题 → 刷编码题 → 学新知识点
- **每周**：全量错题复习 + 重点知识点回顾 → 模拟概念问答（面试）

## 核心行为

### 出编码题

当用户要求出编码题时：
1. 在 `exercises/` 对应领域子目录下创建题目目录
2. 生成 `problem.md`（题目正文）、`solution.md`（参考答案/解析）、`skeleton/main.go`（代码骨架：`package main` + 基础 import + 测试数据 + 空函数签名）
3. 提示用户在 VS Code 中打开 `skeleton/main.go` 手写代码
4. 用户可以用 `go run ./exercises/<path>/skeleton/` 自测

### 批改编码题

当用户提交答案后：
1. 对照 `solution.md` 审查用户的 `attempts.md` 最新记录
2. 指出遗漏、错误、可以改进的地方
3. 在 `attempts.md` 中追加本次尝试记录（diff 风格 + 修改意图 + 正确/错误标记）

### 出概念问答

当用户要求概念问答时：
1. 指定一个 `concepts/` 下的题目（如 "谈谈 GMP 调度模型"）
2. 用户在 VS Code 写好回答后，CC 对照标准答案点评
3. 补充遗漏点，记录到对应的 `concepts/<topic>-attempts.md`

### 根据错题本出同类题

1. 分析 `wrong-book/` 中 `status: reviewing` 且 `wrong_count` 高的题目
2. 找到薄弱领域，生成变体题目（同知识点、不同场景）

### 安排练习计划

1. 扫描全部 knowledge 文件，找 `last_review` 超过 `review_interval_days` 的知识点
2. 扫描 `wrong-book/` 中 `status: reviewing` 的条目
3. 输出：今日复习清单 + 建议新学知识点 + 推荐练习题量

### 标记掌握

当用户说某个题/知识点彻底会了：
1. 将对应文件的 `status` 改为 `mastered`
2. 错题本对应条目 `status` 改为 `mastered`

## 文件格式规范

### 知识点 knowledge/*.md

YAML frontmatter 必须包含：`title`、`category`、`tags`、`difficulty`、`status`、`last_review`、`review_interval_days`、`related`

正文必须包含 `## 我的理解` 区块。

### 编码题 exercises/**/problem.md

YAML frontmatter：`title`、`category`、`tags`、`difficulty`、`status`（learning/attempted/mastered）

### 编码题 exercises/**/attempts.md

每次尝试以 `## 尝试 N · 日期 · ✓/✗` 为分隔符。包含：修改意图、diff 代码块、错误原因。

`difficulty` 写错时，diff 以 `+` 表示添加行、`-` 表示删除行，并保留足够上下文行。

### 概念问答 concepts/*.md

标准答案文件：frontmatter + 自由正文。
答题记录文件（`<topic>-attempts.md`）：记录用户回答 + CC 反馈 + 补充后完整回答。

### 错题本 wrong-book/*.md

YAML frontmatter 必须包含：`exercise_ref`、`type`（coding|concept）、`wrong_count`、`last_wrong_date`、`status`

记录：每次错误思维链 + 正确思维链关键链路。

## 目录结构

```
knowledge/           # 知识点（basics / internals / concurrency / runtime / algorithms）
concepts/            # 概念问答（原"八股文"）
exercises/           # 编码题
wrong-book/          # 错题本
cmd/server/          # Web 服务入口
```

## Web 服务

```bash
go run ./cmd/server           # 默认 :8080
go run ./cmd/server -port 9090
```

Web 端用于浏览 + 小修改（编辑「我的理解」、删除无用答题记录），不做出题/批改。

## 早期测评残留

根目录 `.go` 文件 和 `go-interview-assessment/` 是加密的 TSD 格式测评数据，不要尝试解码或修改。`.vscode/settings.json` 禁用了编辑器补全（保证测评公平性），如果影响正常编码体验可恢复。
