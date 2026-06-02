# knowledge-base 使用手册

> 个人技术知识图谱 — 基于 Skill Graph + Obsidian + AI 自动维护

---

## 1. 仓库概述

### 定位

将零散的技术知识组织为一张 **可遍历的双向链接图**。人通过 Obsidian 浏览，AI 通过 [[wiki-links]] 沿图检索上下文、自动维护关系。一切围绕「写卡片 → 自动链接 → 反复复习」的核心循环。

### 核心理念

| 层 | 描述 |
|---|------|
| **碎片化卡片** | 每个文件是一个完整的、独立的思考单元（Zettelkasten 风格） |
| **语义链接** | [[wiki-links]] 嵌入自然段落，链接本身携带「为什么」 |
| **三级链接策略** | Level 1/2 自动，Level 3 手动；领域内聚，跨域隔离 |
| **AI 维护** | 人不手动维护链接关系，AI 在写入时自动分析并补链 |

### 当前领域

- Go
- K8s (Kubernetes)
- Linux / OS
- Bash
- Python
- AI / ML
- Databases (MySQL, Redis, PostgreSQL, Elasticsearch)
- Security (OAuth2, OIDC, RBAC, ABAC, PKI)
- Distributed Systems (Raft, Consensus)
- Design Patterns

---

## 2. 目录结构

```
knowledge-base/
├── knowledge/                 # 知识点卡片（按领域分子目录）
│   ├── go/
│   │   ├── roadmap.md
│   │   ├── basics/
│   │   ├── concurrency/
│   │   ├── runtime/
│   │   └── compiler/
│   ├── k8s/
│   ├── linux/
│   ├── python/
│   ├── ai/
│   ├── databases/
│   ├── security/
│   ├── distributed/
│   └── patterns/
│
├── exercises/                 # 编码题（按领域分子目录）
│   ├── go/
│   └── ...
│
├── concepts/                  # 概念问答（按领域分子目录）
│   ├── go/
│   └── ...
│
├── wrong-book/                # 错题本（统一，frontmatter 标记 domain）
│
├── cross-domain/              # 跨域 bridge 文件（/cross-link 触发，AI 生成）
│
├── cmd/server/                # Web 浏览服务
├── internal/                  # 服务端逻辑
├── web/                       # 前端模板 + 静态文件
│
├── docs/
│   └── MANUAL.md              # ← 本文件：完整使用手册
│
├── CLAUDE.md                  # AI 行为规则（opencode / Claude 读取）
├── .claude/skills/            # 项目内知识库 skills
├── AGENTS.md                  # Agent 配置（setup-matt-pocock-skills 生成）
├── README.md                  # 仓库简介（指向本手册）
│
├── go.mod / go.sum
└── .vscode/
```

### 目录职责速查

| 目录 | 内容 | 链接方向 |
|------|------|----------|
| `knowledge/` | 知识点卡片 | 域内双向 Level 2，不引用其他层 |
| `exercises/` | 编码题（problem + solution + skeleton + attempts） | 单向 Level 1 → `[[knowledge/...]]` |
| `concepts/` | 概念问答（标准答案 + attempts） | 单向 Level 1 → `[[knowledge/...]]` |
| `wrong-book/` | 错题记录（统一，按 domain 标签隔离） | 仅 frontmatter `exercise_ref`，无 [[link]] |
| `cross-domain/` | 跨域 bridge 文件 | Level 3 手动触发 |

---

## 3. 文件模板速查

### 3.1 知识点卡片 `knowledge/<domain>/<topic>.md`

```markdown
---
title: "切片扩容机制"
category: go/basics/slice
tags: [slice, 扩容, 底层原理]
difficulty: medium
status: learning           # learning | mastered
last_review: 2026-06-01
review_interval_days: 7
level: intermediate
mastery:
  wrong_count: 3
  last_wrong_date: 2026-06-01
  review_count: 5
domain: go
---
# 切片扩容机制

## 核心要点

Slice 是对 [[go/basics/array]] 的封装……

## 我的理解

用自己的语言重述。此处[[链接]]只链域内知识，不链练习/错题/概念。
```

**必填字段**：`title`, `category`, `tags`, `difficulty`, `status`, `last_review`, `review_interval_days`, `level`, `mastery`, `domain`

### 3.2 编码题 `exercises/<domain>/<topic>/`

```
exercises/go/slices/
├── problem.md         ← frontmatter + 题目描述 + 考察知识点
├── solution.md        ← 参考答案
├── skeleton/
│   └── main.go        ← 代码骨架（VS Code 手写，go run 自测）
└── attempts.md        ← 答题记录
```

**problem.md 格式**：

```markdown
---
title: "Slice 扩容陷阱"
category: go/concurrency
tags: [slice, append, 扩容]
difficulty: medium
status: learning
domain: go
---
# Slice 扩容陷阱

## 题目描述

...

## 考察

[[knowledge/go/basics/slice]]
```

**attempts.md 格式**：

```markdown
## 尝试 1 · 2026-06-01 · ✗

**AI 批改：**
- ✅ 正确：[做对了什么]
- ❌ 错误：[哪里错了 + 正确写法]
- + 改进：[可优化点]

**你的代码：**
```go
// 完整代码
```

## 尝试 2 · 2026-06-01 · ✓

**AI 批改：**
- ✅ 正确：全部通过
```
```
- **错误原因：** 分析

## 尝试 2 · 2026-06-01 · ✓ 通过
- **修改意图：**
- **变更：** ...
```

### 3.3 概念问答 `concepts/<domain>/<topic>.md`

```markdown
---
title: "谈谈 GMP 调度模型"
category: go/concurrency
tags: [gmp, goroutine, scheduler]
difficulty: hard
status: learning
domain: go
---
# GMP 调度模型

## 题目

> 简述 GMP 调度模型的核心设计。

## 标准答案

完整回答……

## 关键要点

- 要点 1
- 要点 2

## 考察

[[knowledge/go/concurrency/gmp]]
```

答题记录文件 `concepts/<domain>/<topic>-attempts.md` 追加用户回答与 AI 反馈。

### 3.4 错题本 `wrong-book/<topic>.md`

```markdown
---
exercise_ref: exercises/go/slices/append-traps
type: coding               # coding | concept
wrong_count: 2
last_wrong_date: 2026-06-01
status: reviewing
domain: go
---
# 题目名称 — 错题记录

## 错误记录

### 第 1 次 · 2026-06-01
**错误思维链：** 我当时的思路是……

**正确思维链关键链路：** 正确思路应该是……
```

注意：wrong-book prose 中不可嵌入 `[[wiki-links]]`，跨域关联用纯文本提及。

### 3.5 cross-domain bridge 文件 `cross-domain/<relation>.md`

```markdown
---
title: "Goroutine 泄漏 ↔ Pod 生命周期管理"
domains: [go, k8s]
pattern: resource-lifecycle
created: 2026-06-01
---

→ [[knowledge/go/concurrency/goroutine-leak]]
→ [[knowledge/k8s/pod-lifecycle]]

## 底层模式：资源生命周期对称性

无论是 goroutine 还是 Pod，核心问题是**创建和销毁是否对称**。
两者都遵循「谁创建谁负责销毁」的哲学。

Go 侧：defer cancel() + context.WithTimeout
K8s 侧：terminationGracePeriodSeconds + preStop hook

## 为什么这个链接有价值

理解 Go 的 goroutine 泄漏后，K8s 的 Pod 优雅终止是
同一模式在分布式层面的复现。
```

---

## 4. 知识图谱链接规范

### 价值层级

```
层 1（最高）: knowledge/          纯知识图谱，双向内链，golden source
层 2（高）：   wrong-book/         个性化错误模式，价值高于练习
层 3（中）：   exercises/ + concepts/  练习记录，单向引用 knowledge
```

**铁律**：
- `knowledge/` 不挂任何通往 exercises/ concepts/ wrong-book/ 的 [[link]]
- `wrong-book/` 用 frontmatter 的 `exercise_ref` 关联原题（非 [[link]]），prose 中纯文本提及概念
- `exercises/` 和 `concepts/` 可单向引用 `[[knowledge/...]]`，不可反向
- 跨域提及在 prose 中用纯文本，不加 [[ ]]

### 三级链接策略

| 级别 | 方向 | 范围 | 触发方式 | 产物 | 是否需要知识卡片 |
|------|------|------|----------|------|----------------|
| Level 1 | 单向 | 练习/概念 → 知识 | 自动（创建练习/概念时） | `## 考察` 段中的 `[[knowledge/...]]` | 否 |
| Level 2 | 双向 | 域内知识 ↔ 域内知识 | 自动 + 可沟通 | prose 中嵌入语义 [[link]] | 否 |
| Level 3 | 双向 | 跨域关联 | 手动 `/cross-link A B` | `cross-domain/` 下 bridge 文件 | **是** |

### Level 1 — 全自动单向（不询问）

**规则**：创建 `exercises/` 或 `concepts/` 文件时，AI 自动匹配知识点，在 `## 考察` 段写入 `[[knowledge/...]]`。纯机械匹配，不询问。

```markdown
## 考察

[[knowledge/go/basics/slice]]
```

### Level 2 — 自动 + 可沟通（域内双向）

**规则**：AI 分析知识文件内容，在 prose 中嵌入域内语义链接。可沟通修改。

```
Go slice 底层依赖 [[go/basics/array]] 的连续内存，
这也是为什么对 slice 切片后修改会影响到原数组。
```

### Level 3 — 手动触发（跨域）

**规则**：`/cross-link <domain-a> <domain-b>` 后，AI 扫描双方知识库讨论候选关联，确认后写入 `cross-domain/` 下的 bridge 文件。

### 领域隔离规则

- 用户指定"练习 Go"时，AI **仅加载** `knowledge/go/` + `exercises/go/` + `wrong-book/`（按 `domain: go` 过滤）
- 领域文件 [[wiki-links]] 仅供域内跳转，不指向其他领域
- `cross-domain/` 仅在用户显式触发时加载

### [[wiki-links]] 书写规范

- 路径相对于项目根：`[[knowledge/go/basics/slice]]`
- **必须嵌入自然段落中**，不可使用列表形式或 frontmatter 引用列表
- 链接所在的句子应能独立推断关系类型（依赖 / 应用场景 / 延伸 / 易错点等）

✅ 正确写法：
```
Slice 底层依赖 [[go/basics/array]] 的连续内存布局。
```

❌ 错误写法：
```
相关知识点：[[array]], [[pointer]]
```

---

## 5. 工作流

### 5.1 每日学习节奏

```
1. "帮我安排今天的练习"  →  巩固：扫描过期 + 错题本，输出复习清单
2. 回顾高频错题        →  内化：错题反哺，发现薄弱知识点
3. 刷一道题/学一个知识点  →  内化/学习：按当天状态选择
4. 写「我的理解」       →  学习：填充后将触发审核→补链
5. "审一下" / "更新知识图谱"  →  学习：审核通过后自动补链
```

### 5.2 知识点生命周期（主线）

知识卡片经历 3 个阶段：学习 → 内化 → 巩固。

#### 5.2.1 学习 — 从骨架到补链

```
触发: "学习 Go slice 扩容机制"
      "生成知识卡片 go/basics/slice/slice-basics"

AI 行为:
1. 确认知识点路径后，在 roadmap.md 中读取 level、requires 等信息
2. 创建 .md 文件，填入 frontmatter (title, category, tags, difficulty, status: learning, level, domain, last_review: 今日, review_interval_days, mastery)
3. 写入 ## 标准知识 和 ## 我的理解 空区块
4. 如果 roadmap 声明了 requires，在文件中嵌入引用

人 (或在 AI 辅助下) 填充:
  - ## 标准知识：从源码/文档/博客学习
  - ## 我的理解：用自己的话重述
  - 填充过程中可随时求助 AI：解释、对比、举例

下一步 → AI 审核（触发 "审一下" / "审核 go/basics/slice/slice-expansion"）:
1. AI 逐段检查：理解是否准确、有无遗漏、措辞是否清晰
2. 输出审核报告（准确 ✓ / 偏差 ✗ / 补充建议）
3. 人结合反馈修改，可多轮迭代
4. 审核通过后，AI 自动做补链分析:
   例: "slice 扩容需要分配新的底层数组"
       → 识别 "底层数组" 与 [[go/basics/array]] 的关联
5. 列出所有建议链接:
   [1] slice-expansion → basics/array    (新增)
   [2] slice-expansion → basics/pointer  (新增)
6. 人确认: "保留" / "撤销第 N 条"
7. 确认后写入文件
```

审核通过前的修改都是迭代，不做补链以免浪费；审核通过后一次性补链，人只需要确认一次。

#### 5.2.2 内化 — 从练习到反馈

```
三种练习模式:

1. 单知识点出题
   触发: "练习 go/basics/slice/slice-expansion"

2. 今日所学批量出题
   触发: "练习今天的知识点"
   → AI 扫描今日生成的卡片或 last_review 为今天的卡片
   → 综合出一组题（编码 + 概念问答混合）

3. 综合闯关
   触发: "闯关 Go basics 中级"
   → AI 按难度 (level: intermediate) 混合出卷
   → 覆盖该维度下多个知识点

做题 → 提交答案 → AI 批改:
  ✓ 通过 → 对应知识 mastery.review_count +1
  ✗ 错误 → 写入 wrong-book（记录错误思维链 + 正确链路）

错题反哺知识理解:
  如果某个知识点在 wrong-book 反复出错（wrong_count ≥ 3），
  AI 提示："你在 [知识点] 上反复出错，建议回到 5.2.1 补充 ## 我的理解"
  形成"学习 → 练习 → 发现不足 → 再学习"的闭环

出题类型:
  ● 编码题（带 skeleton + solution）
  ● 概念问答题
  ● 代码纠错题
  ● 场景设计题
```

**概念问答**和**编码练习**不做严格区分，统一归入"内化"阶段——同一个知识点可以编码题也行、概念题也行，依 user 想练什么而定。

#### 5.2.3 巩固 — 从复习到掌握

```
复习调度（每日自动 + 手动触发）:
  触发: "帮我安排今天的练习"
  → 扫描 knowledge/ 中 last_review 超过 review_interval_days 的节点
  → 扫描 wrong-book/ 中 status: reviewing 的条目
  → 输出：今日复习清单 + 建议新学知识点 + 推荐练习题量

健康检测（按需触发）:
  触发: "检测知识图谱" / "检测 Go 域链接健康"

  AI 输出报告格式:

  ┌─────────────────────────────────────────┐
  │ Go 域 · 知识图谱健康报告                  │
  │ 2026-06-01                              │
  ├─────────────────────────────────────────┤
  │ 质量指标                                │
  │  文件总数:    12        (+2 自上周)       │
  │  断链数量:    2        ⚠ 需要修复         │
  │  孤儿节点:    1        ⚠ 无入链           │
  │  过期复习:    3        ⚠ 超过间隔          │
  │  空理解区:    1        ⚠ 待补充            │
  │  孤立知识点:   0        ✅                  │
  ├─────────────────────────────────────────┤
  │ 断链详情                                │
  │  go/basics/memory-allocation →          │
  │    [[go/basics/pointer]]  ⚠ 文件不存在     │
  │  go/runtime/gc/gc-basics →              │
  │    [[go/runtime/memory]]  ⚠ 文件不存在     │
  ├─────────────────────────────────────────┤
  │ 孤儿节点                                │
  │  go/basics/defer/defer-internals         │
  │    ← 没有任何其他知识点引用它               │
  ├─────────────────────────────────────────┤
  │ 建议操作                                │
  │  1. 修复 2 处断链（创建目标文件或修正链接）    │
  │  2. 给 defer-internals 添加入链            │
  │  3. 复习 3 个过期节点                      │
  │  4. 补充 slice-header 的「我的理解」         │
  └─────────────────────────────────────────┘

  检测项目: 断链 / 孤儿 / 过期复习 / 空理解区 / 孤立知识点
  修复: "修复断链" / "修复孤儿节点" → AI 逐条确认方案

拆分合并（子方向成熟时）:
  触发: "拆分 [[go/basics/slice]]" / "slice 这部分我想拆开细看"
  AI: 解析文件 → 识别概念子边界 → 输出预览 → 人确认 → 建子目录 + 拆文件 + 写双向 link + 更新 roadmap
  每个子卡片继承父级 mastery 数据，保持「标准知识 + 我的理解」结构

  反向: "把 [A] 合并回 [B]"
  当子概念之间交叉引用过多、单独复习总是在翻上下文时

标记掌握（归档）:
  触发: "[知识点] 我懂了" / "这题我彻底会了"
  → 知识文件 status → mastered
  → review_interval_days → 30
  → 关联 wrong-book 条目 status → mastered
  → 不再推送复习，但可手动查阅

  什么不算掌握:
  - ## 我的理解 还是空的
  - 学习时间不足 review_interval_days 的 3 倍
  - 对应的 wrong-book 条目仍存在
```

### 5.3 全局复习（支线）— 通过 Roadmap 鸟瞰

不循环每个知识点，而是通过 Roadmap 树从更高层次审视知识结构。

#### 5.3.1 概览 — 看知识脉络

```
触发: "看 Go 知识脉络"
      "看 Go 并发知识脉络"

AI 行为:
1. 从对应维度 roadmap.md 读取全部计划知识点（含 level、requires）
2. 检查实际 .md 文件的存在状态和 frontmatter status
3. 渲染知识树:

   Go 基础 (basics)                ● 3/28  ═══════════════════════════
   ├── slice                       ● 1/3   ═══════════
   │   ├── slice-basics            ●       ⚪ 已生成，待补充
   │   ├── slice-expansion         ✅      ● 已通过审核，可练习
   │   └── slice-header            ⬜      ○ 未创建
   ├── map                         ⬜ 0/3  ═══
   ├── string                      ⬜ 0/3  ═══
   ...

   Go 并发 (concurrency)           ⬜ 0/13 ═══════════════════════════

  状态标记:
    ⬜ 未创建  ● 待补充  ⚪ 待审核  ✅ 已掌握
    🔄 学习中  ✗ 需修正  ◉ 待复习
```

#### 5.3.2 差距 — 找出薄弱点

```
触发: "Go 有哪些薄弱点"
      "并发方面我哪块比较弱"

AI 行为:
1. 扫描该领域全部文件，综合以下指标:
   - wrong_count > 0（来自 wrong-book）
   - last_review 过期
   - ## 我的理解 为空或极短
   - status: learning > 7 天未推进
2. 按严重程度排序输出:
   ⚠ 高优先级: wrong_count ≥ 3 且过期
   ⚡ 中优先级: 空理解区或长期未推进
   ℹ 低优先级: 轻微过期
3. 输出建议动作:
   "建议优先补充 slice-expansion 的 ## 我的理解，
    然后练习 go/basics/slice 的相关题目。"
```

#### 5.3.3 串联 — 梳理依赖链

```
触发: "串联 Go 并发"
      "slice 的依赖关系是怎样的"

AI 行为:
1. 沿 roadmap.md 的 requires 链正向/反向遍历
2. 以叙事方式讲解逻辑依赖:
   "GMP 调度是 goroutine 的实现基础，而 channel 和 sync 包
    都是在 GMP 框架上构建的并发原语。先学 GMP 再学 channel
    更容易理解阻塞和唤醒的底层机制。"
3. 不堆知识点，讲逻辑关系和"为什么要先学这个"
4. 可沿某条链深入做复习:
   "你要不要沿着这条链逐个过一遍？"
```

#### 5.3.4 面试复习 — 快速过一遍

```
触发: "面试前复习 Go"
      "面试前复习 Go 并发"

AI 行为:
1. 从域 roadmap 顶层开始，逐主题出概念问答
2. 三种状态:
   ✅ 已掌握 → 快速跳过（或出两道确认）
   🔄 学习中 → 出典型面试题验证理解
   ⬜ 未创建 → 询问要不要学（"这是一个常见面试题，要听吗？"）
3. 每个维度结束时总结:
   "你的 Go 基础部分掌握较好，并发方面 channel 的 close 原则
    不够清晰，建议面试前再练两道题。"
4. 可切换面试官风格:
   "用阿里 P7 的角度面我"
   "用字节跳动的方式面我"
   "用大厂面试官的语气追问"
```

### 5.4 跨域联想

```
/cross-link go k8s

AI 输出候选关联列表（如 goroutine 泄漏 ↔ Pod 生命周期），
逐条讨论，确认后生成 cross-domain/ bridge 文件。
```
---

## 6. 命令速查表

### 6.1 自然语言命令

| 你说 | AI 做什么 |
|------|----------|
| "学习 Go slice 扩容机制" | 生命周期·学习：创建骨架，等待填充和审核 |
| "生成知识卡片 go/basics/slice/slice-basics" | 同上（别名） |
| "审一下" | 生命周期·学习：审核当前卡片，通过后自动补链 |
| "审核 go/basics/slice/slice-basics" | 同上，指定卡片 |
| "保留" | 确认 AI 自动生成的链接列表 |
| "撤销第 N 条" | 删除 AI 刚生成的第 N 条链接 |
| "更新知识图谱" / /sync | 全量同步 Level 1 + Level 2 链接 |
| "练习 go/basics/slice/slice-expansion" | 生命周期·内化：单知识点出题 |
| "练习今天的知识点" | 生命周期·内化：今日卡片综合出题 |
| "闯关 Go basics 中级" | 生命周期·内化：按难度综合出卷 |
| "给我出一道 Go 编码题" | 生命周期·内化：在 `exercises/go/` 下随机出题 |
| "给我出一道 [领域] 编码题" | 同上，指定领域目录 |
| "根据错题本出一题" | 生命周期·内化：分析 wrong-book/ 变体出题 |
| "帮我看看我的答案" | 生命周期·内化：对照 solution.md 批改 |
| "批改" | 生命周期·内化：读取 skeleton/main.go，对照 solution 批改 |
| "做完了" | 同上（别名） |
| "这题我彻底会了" | 生命周期·巩固：status → mastered |
| "[知识点] 我懂了" | 生命周期·巩固：知识点 status → mastered |
| "帮我安排今天的练习" | 生命周期·巩固：扫描过期节点，输出复习清单 |
| "检测知识图谱" | 生命周期·巩固：输出健康报告 |
| "修复断链" | 生命周期·巩固：AI 逐条修补方案 |
| "修复孤儿节点" | 生命周期·巩固：AI 建议入链位置 |
| "拆分 [[go/basics/slice]]" | 生命周期·巩固：按概念边界拆分 |
| "把 [A] 合并回 [B]" | 生命周期·巩固：反向合并 |
| "看 Go 知识脉络" | 支线·概览：渲染知识树 + 进度 |
| "Go 有哪些薄弱点" | 支线·差距：扫描 wrong-book + 空理解区 + 过期 |
| "串联 Go 并发" | 支线·串联：沿 requires 链讲解逻辑依赖 |
| "面试前复习 Go" | 支线·面试：逐主题模拟问答 |
| "练习 Go" | 进入 Go 域隔离练习模式 |
| "联想一下 [A] 和 [B]" | 跨域：触发跨域关联讨论（等效 /cross-link） |

### 6.2 斜杠命令

| 命令 | 作用 |
|------|------|
| `/init` | 初始化/更新 AGENTS.md |
| `/models` | 查看可用模型 |
| `/new` | 新会话 |
| `/undo` | 撤销上一步（含文件变更） |
| `/redo` | 重做 |
| `/sessions` | 切换会话 |
| `/compact` | 压缩上下文 |
| `/share` | 分享会话 |
| `/themes` | 切换主题 |
| `/thinking` | 显示/隐藏推理过程 |
| `/help` | 帮助 |
| `/editor` | 打开外部编辑器撰写消息 |

### 6.3 自定义命令

| 命令 | 作用 |
|------|------|
| `/sync [domain]` | 触发 Level 1 + Level 2 链接维护（可选领域限定） |
| `/cross-link <domain-a> [domain-b]` | 手动跨域联想，生成 bridge 文件 |

### 6.4 opencode CLI

```bash
opencode                                 # 启动 TUI
opencode /path/to/project                # 指定项目目录
opencode run "消息"                      # 非交互模式
opencode run -m ds-hex/deepseek-v4-pro "消息"  # 指定模型
opencode serve                           # headless API
opencode web                             # Web 界面
opencode agent list                      # 列出 agent
opencode stats --days 7                  # 近 7 天用量
opencode export [sessionID]             # 导出会话
```

---

## 7. 模型分配策略

| 场景 | 推荐模型 | 理由 |
|------|----------|------|
| 深度推理、跨域联想 | `ds-hex/deepseek-v4-pro` | 默认模型，强推理能力 |
| 编码题出题、方案设计 | `xh/claude-sonnet-4-6` | 写作质量适中，编码能力强 |
| 每日问答、格式检查 | `ds-hex/deepseek-v4-flash` | 轻量快速，低成本 |
| 快速批改、标签补全 | `xh/claude-haiku-4-5` | 最轻量，反复调用不心疼 |

指定模型示例：

```
在 Opencode CLI：opencode run -m xh/claude-haiku-4-5 "快速批改这道题"
在对话中：用 /models 查看，切换前 @ 指定
```

---

## 8. 工具链

### 8.1 opencode

当前配置（`~/.config/opencode/opencode.json`）：

- **2 个供应商**：xh（muyuan.do，Claude 模型）、ds-hex（DeepSeek）
- **4 个模型**：deepseek-v4-pro（默认）、deepseek-v4-flash、claude-sonnet-4-6、claude-haiku-4-5

TUI 常用快捷键：

| 快捷键 | 作用 |
|--------|------|
| `Tab` | 切换 Plan / Build 模式 |
| `@` | 模糊搜索文件引用 |
| `!` | 行首写 shell 命令 |
| `ctrl+x u` | 撤销 |
| `ctrl+x r` | 重做 |
| `ctrl+x n` | 新会话 |
| `ctrl+x l` | 切换会话 |
| `ctrl+x m` | 查看模型 |
| `ctrl+x c` | 压缩上下文 |
| `ctrl+x t` | 切换主题 |
| `ctrl+p` | 命令面板 |

### 8.2 Obsidian

将项目根目录作为 Obsidian Vault 打开：

```
Obsidian → Open folder as vault → 选 knowledge-base/
```

收益：

| 功能 | 用途 |
|------|------|
| Graph View | 可视化整张知识图谱，发现孤立节点和密集区域 |
| Backlinks | 查看哪个文件引用了当前节点 |
| Local Graph | 只看当前节点 + 直接邻居 |
| 搜索 | 按标签、标题、内容快速检索 |

**建议**：`.gitignore` 中忽略 `.obsidian/workspace.json`（每次打开都会变化），保留 `.obsidian/` 下的配置和插件设置。

### 8.3 Web 服务

```bash
cd knowledge-base
go run ./cmd/server            # 默认 http://localhost:8080
go run ./cmd/server -port 9090 # 指定端口
```

路由表：

| 路由 | 功能 |
|------|------|
| `/` | 仪表盘（复习概览） |
| `/knowledge/` | 知识点浏览 |
| `/knowledge/{path}` | 知识点详情 + 编辑「我的理解」 |
| `/concepts/` | 概念问答列表 |
| `/exercises/` | 编码题列表 |
| `/exercises/{path}` | 编码题详情 |
| `/wrong-book/` | 错题本（按时间/频率排序） |
| `/search?q=` | 全站搜索 |

Web 端用于浏览 + 轻量修改（编辑「我的理解」、删除无用答题记录），**不做**出题/批改。

### 8.4 VS Code

用于：
- 手写编码题的 `skeleton/main.go`
- 编写概念问答的回答
- Git 操作

Obsidian 和 VS Code 共存编辑同一套 `.md` 文件，不会冲突。

---

## 9. Skill 清单

### 9.1 项目内知识库 skills

项目内 skill 存放在 `.claude/skills/`，跟随仓库提交。它们是本知识库的主工作流入口，详细流程仍以本手册第 5 章为准。

| Skill | 对应工作流 | 触发 |
|-------|------------|------|
| `kb-learn` | 学习：骨架生成、AI 审核、审核后补链 | "学习 X" / "审一下" / "更新知识图谱" |
| `kb-practice` | 内化：出题、批改、错题反哺 | "练习 X" / "闯关 X" / "帮我看看我的答案" |
| `kb-consolidate` | 巩固：复习调度、健康检测、拆分合并、标记掌握 | "帮我安排今天的练习" / "检测知识图谱" / "这题我彻底会了" |
| `kb-roadmap` | 全局复习支线：概览、差距、串联、模拟面试 | "看 X 知识脉络" / "串联 X" / "面试前复习 X" |
| `kb-crosslink` | 跨域联想：候选关联讨论、bridge 文件生成 | `/cross-link A B` / "联想一下 A 和 B" |

### 9.2 已安装通用 skills

当前已安装的通用 skill，在知识库中的角色如下：

| Skill | 在知识库中的角色 |
|-------|----------------|
| `neat-freak` | **核心** — 会话结束后同步知识文件与记忆，检查断链、过期 review |
| `handoff` | 每个学习阶段后将讨论沉淀为知识点卡片 |
| `grill-me` | 复习时对设计方案或概念进行压力测试 |
| `write-a-skill` | 创建自定义知识库技能（如 go-linker） |
| `setup-matt-pocock-skills` | 初始化 AGENTS.md，建立 issue tracker 的文档结构 |
| `skill-auditor` | 审计所有 skill 在当前环境的可用性 |
| `diagnose` | 编码题中遇到难以定位的 Bug 时系统性诊断 |
| `prototype` | 快速验证某个 Go 设计方案 |
| `tdd` | 红-绿-重构 TDD 练习编码题 |
| `caveman` | 极简输出模式（复习时省 token） |
| `zoom-out` | 理解不熟悉的代码段时获取全局上下文 |
| `aihot` | 跟进 AI 领域动态 |
| `khazix-writer` | 将知识点整理为公众号文章 |
| `hv-analysis` | 深度学习某个技术（如 K8s 调度器）出研究报告 |
| `improve-codebase-architecture` | 代码架构分析（主要用于 Go 项目） |
| `find-skills` | 搜索可安装的新 skill |
| `to-issues` / `to-prd` / `triage` | 计划管理（暂未启用） |
| `microsoft-foundry` / `capacity` / `deploy-model` / `preset` / `customize` | Azure OpenAI 部署（暂未启用） |
| `jenkins-pipeline` | CI/CD 流水线（暂未启用） |

---

## 10. 注意事项

### Git

- 提交前检查 `git status` + `git diff`
- 多领域改动可分批提交：`git add knowledge/go/` 一次，`git add knowledge/k8s/` 另一次

### .gitignore 建议

```gitignore
.obsidian/workspace.json
# 或整个 .obsidian/（看是否要共享 Obsidian 配置）
```

### 加密测评残留

根目录 `.go` 文件和旧 `go-interview-assessment/` 目录是加密 TSD 格式测评数据，**不要尝试解码或修改**。

### VS Code 补全

`.vscode/settings.json` 禁用了编辑器补全（原本用于测评公平性）。如果影响编码体验可恢复。

### API Key 安全

API Key 已在 `~/.config/opencode/opencode.json` 中管理，不要在其他文件中硬编码。
