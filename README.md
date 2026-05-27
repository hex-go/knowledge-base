# Go 学习助手

本仓库是 Go 语言自学体系的环境。核心思路：**Markdown 知识库（人类可读）+ 本地 Web 服务（浏览/编辑）+ Claude Code（AI 陪练）**。

## 快速开始

```bash
# 启动本地 Web 服务（默认 http://localhost:8080）
go run ./cmd/server

# 指定端口
go run ./cmd/server -port 9090
```

## 目录结构

```
knowledge/          # 知识点库
  basics/           #   基础语法
  internals/        #   底层原理（内存对齐、字符串、切片、map、struct、channel）
  concurrency/      #   并发编程（原语、易错点、使用场景）
  runtime/          #   运行时（GMP、GC、内存分配、netpoller）
  algorithms/       #   常见算法（排序、二分、DP、递归）

concepts/           # 概念问答（原"八股文"）
exercises/          # 编码题（含题目、参考答案、代码骨架）
wrong-book/         # 错题本（关联回原题，记录错误思维链 + 正确思维链）

cmd/server/         # Web 服务入口
```

每一个知识点/题目/错题都是一个或多个 `.md` 文件，人类可直接打开阅读。

## Markdown 文件格式

每个文件以 YAML frontmatter 开头，正文自由书写。

### 知识点文件

```markdown
---
title: "切片扩容机制"
category: internals/slice-array
tags: [slice, 扩容, 底层原理]
difficulty: medium
status: learning          # learning | mastered
last_review: 2026-05-27
review_interval_days: 7
related: [internals/array-vs-slice]
---
# 切片扩容机制

## 标准知识

> Go 1.18 之前，容量 < 1024 时翻倍，>= 1024 时增长 25%……

## 我的理解

这里写自己的认知感悟。可以和标准知识不同，重要的是自己的语言重述。
```

### 编码题文件

```text
exercises/
└── concurrency/
    └── concurrent-counter/
        ├── problem.md      # 题目正文
        ├── solution.md     # 参考答案/解析
        ├── skeleton/
        │   └── main.go     # 代码骨架（VS Code 打开编辑）
        └── attempts.md     # 答题记录（diff 风格）
```

`attempts.md` 格式：

```markdown
## 尝试 1 · 2026-05-27 · ✗ 错误
- **修改意图：** 给 Counter.Increment 加锁
- **变更：**
  ```diff
  func (c *Counter) Increment() {
  +   c.mu.Lock()
      c.count++
  +   c.mu.Unlock()
  }
  ```
- **错误原因：** 漏了读也要加锁

## 尝试 2 · 2026-05-27 · ✓ 通过
- **修改意图：** 读写都加锁
- **变更：**
  ```diff
  func (c *Counter) Value() int {
  +   c.mu.Lock()
  +   defer c.mu.Unlock()
      return c.count
  }
  ```
```

### 概念问答文件

```text
concepts/gmp.md          # 标准答案
concepts/gmp-attempts.md # 答题记录
```

### 错题本文件

```markdown
---
exercise_ref: exercises/concurrency/concurrent-counter
type: coding               # coding | concept
wrong_count: 3
last_wrong_date: 2026-05-27
status: reviewing          # reviewing | mastered
---
# 并发安全计数器 — 错题记录

## 错误记录

### 第 1 次 · 2026-05-20
**错误思维链：**
我把 Lock 放错位置了，放在了 goroutine 外面。

**正确思维链关键链路：**
Lock 必须保护临界区，粒度刚好覆盖共享状态的访问。读写都要上锁。

### 第 2 次 · 2026-05-25
错误思维链：……
正确思维链关键链路：……
```

## 日常使用节奏

### 每日

1. 打开 Web 服务，看一眼今日待复习错题
2. 回顾上次/高频错题（Web 端按错误次数排序）
3. 刷一道编码题：CC 出题 → 生成骨架 → VS Code 手写 → `go run .` 自测 → CC 批改
4. 学一个新知识点，写「我的理解」

### 每周（周考模式）

1. 全量错题复习 + 高频知识点回顾
2. CC 抽 3-5 道概念问答，模拟面试
3. CC 对自己回答做对照点评
4. 清理无价值的旧答题记录

## 与 Claude Code 协作

在仓库根目录给 CC 发消息即可触发以下行为：

| 指令示例 | CC 做什么 |
|---|---|
| "给我出一道并发编码题" | 在 `exercises/concurrency/` 下生成题目目录、骨架 `main.go`、`problem.md`、`solution.md` |
| "帮我看看我的答案" | 对照 `solution.md` 批改 `attempts.md` 中的最新记录，指出遗漏和错误 |
| "这个概念题问的什么？" | 出一道概念问答，等你在 VS Code 写完回答后，帮你对比标准答案 |
| "根据我的错题本，出一道同类题" | 分析 `wrong-book/` 找到你的薄弱点，生成变体题目 |
| "帮我安排今天的练习" | 根据 `last_review` 和 `wrong_count`，输出今日练习计划 |
| "这题我彻底会了" | 把 `status` 改为 `mastered`，从高频复习中移除 |

编排原则：**CC 负责生成和批改，Web 服务负责浏览和小修改，VS Code 是你手写代码的地方。**

## 元数据字段速查

| 字段 | 含义 | 可选值 |
|---|---|---|
| `title` | 标题 | 自由文本 |
| `category` | 所属分类路径 | `basics` / `internals/slice-array` / `concurrency` / `runtime/gmp` / `algorithms/sorting` 等 |
| `tags` | 标签列表 | `[goroutine, channel, select]` 等 |
| `difficulty` | 难度 | `easy` / `medium` / `hard` |
| `status` | 学习/掌握状态 | `learning` / `reviewing` / `mastered` |
| `wrong_count` | 错误次数 | 整数 |
| `last_review` | 最近复习日期 | `YYYY-MM-DD` |
| `review_interval_days` | 复习间隔 | 整数（天） |
| `related` | 关联知识点 | 路径数组 |
| `type` | 题目类型（仅错题本） | `coding` / `concept` |
| `exercise_ref` | 关联原题（仅错题本） | 原题目录路径 |

## Web 服务路由

| 路由 | 功能 |
|---|---|
| `/` | 仪表盘 |
| `/knowledge/` | 知识点目录 |
| `/knowledge/{path}` | 知识点详情 + 编辑「我的理解」 |
| `/concepts/` | 概念问答列表 |
| `/exercises/` | 编码题列表 |
| `/exercises/{path}` | 编码题详情（答案默认折叠） |
| `/wrong-book/` | 错题本（按时间/频率排序） |
| `/search?q=` | 全站搜索 |
