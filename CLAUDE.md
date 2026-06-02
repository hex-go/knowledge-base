# CLAUDE.md

本文件为 Claude Code 在此仓库中工作时提供指导。

## 仓库定位

这是用户的多领域技术知识图谱。核心思路：**Markdown 知识库（人类可读）+ 本地 Web 浏览 + Claude Code AI 陪练**。

用户的日常节奏：
- **每日**：巩固（安排练习 + 回顾错题）→ 内化/学习（刷题或学新知识点 → 写「我的理解」→ AI 审核）
- **每周**：Roadmap 鸟瞰（看脉络 / 找差距 / 串联）→ 全量错题复习 → 模拟面试

## 核心行为

仓库行为由 5 个专用 skill 驱动，映射到 `docs/MANUAL.md` 第 5 章的生命周期：

| Skill | 阶段 | 触发 |
|-------|------|------|
| `kb-learn` | 学习（§5.2.1） | "学习 X" / "审一下" |
| `kb-practice` | 内化（§5.2.2） | "练习 X" / "闯关 X" |
| `kb-consolidate` | 巩固（§5.2.3） | "帮我安排今天的练习" / "这题我彻底会了" |
| `kb-roadmap` | 支线（§5.3） | "看 X 知识脉络" / "面试前复习 X" |
| `kb-crosslink` | 跨域（§5.4） | "/cross-link A B" |

所有详细工作流见 `docs/MANUAL.md` 第 5 章，命令速查见第 6 章。

## 文件格式规范

### 知识点 knowledge/*.md

YAML frontmatter 必须包含：`title`、`category`、`tags`、`difficulty`、`status`、`last_review`、`review_interval_days`、`mastery`、`level`、`domain`

正文必须包含 `## 我的理解` 区块。

### 编码题 exercises/**/problem.md

YAML frontmatter：`title`、`category`、`tags`、`difficulty`、`status`（learning/attempted/mastered）、`domain`

正文必须包含 `## 考察` 区块，其中以 `[[knowledge/...]]` 引用考察的知识点。

### 编码题 exercises/**/attempts.md

每次尝试以 `## 尝试 N · 日期 · ✓/✗` 为分隔符。内容为 AI 批改摘要（✅ 正确 / ❌ 错误 / + 改进）+ 用户完整代码。格式详见 `docs/MANUAL.md` §3.2。

### 概念问答 concepts/*.md

标准答案文件：frontmatter + 自由正文。
答题记录文件（`<topic>-attempts.md`）：记录用户回答 + CC 反馈 + 补充后完整回答。

### 错题本 wrong-book/*.md

YAML frontmatter 必须包含：`exercise_ref`、`type`（coding|concept）、`wrong_count`、`last_wrong_date`、`status`、`domain`

记录：每次错误思维链 + 正确思维链关键链路。

## 目录结构

```
knowledge/           # 知识点（按领域分组）
├── go/              # Go 领域
├── k8s/             # Kubernetes 领域
├── linux/           # Linux 领域
├── python/          # Python 领域
├── ai/              # AI 领域
├── databases/       # 数据库领域
├── security/        # 安全领域
├── distributed/     # 分布式系统领域
├── patterns/        # 设计模式领域
├── _template.md     # 领域知识文件模板
concepts/            # 概念问答（原"八股文"）
exercises/           # 编码题（按领域分组）
├── go/              # Go 编程题
wrong-book/          # 错题本（按 domain 字段过滤）
cmd/server/          # Web 服务入口
cross-domain/        # 跨域关联 bridge 文件
.claude/skills/      # 项目内 skills（跟随仓库提交）
```

## Web 服务

```bash
go run ./cmd/server           # 默认 :8080
go run ./cmd/server -port 9090
```

Web 端用于浏览 + 小修改（编辑「我的理解」、删除无用答题记录），不做出题/批改。

## 价值层级

整个仓库按价值分层，上层不引用下层：

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

## 知识图谱链接规范

### 三级链接策略

- **Level 1（自动，单向）**：练习/概念 → 知识。创建 exercises/ 或 concepts/ 文件时，AI 自动在 `## 考察` 段写入 `[[knowledge/...]]`。纯机械匹配，不询问。
- **Level 2（自动 + 可沟通）**：域内知识 ↔ 域内知识。AI 分析内容后在 prose 中嵌入语义链接。列出变更，用户可回复"保留"或"撤销第 N 条"。
- **Level 3（手动触发）**：跨域关联。仅 `/cross-link A B` 触发，生成 `cross-domain/` 下 bridge 文件。

### 领域隔离

- 用户指定"练习 Go"时，AI 仅加载 `knowledge/go/` + `exercises/go/` + `wrong-book/`（按 `domain: go` 过滤）
- 领域文件 [[wiki-links]] 仅供域内跳转，不指向其他领域
- `cross-domain/` 仅在用户显式触发时加载

### [[wiki-links]] 书写规范

- 路径相对于项目根：`[[knowledge/go/basics/slice]]`（可省略 `knowledge/` 前缀，如 `[[go/basics/slice]]`）
- 必须嵌入自然段落中，携带语义上下文
- 不可使用列表形式或 frontmatter 引用列表

### bridge 文件规范

`cross-domain/` 下的 bridge 文件是独立的知识卡片：
- YAML frontmatter 包含：title / domains / pattern / created
- 正文包含核心洞察文字 + 链接到两端知识文件的 [[wiki-links]]
- 由 /cross-link 命令手动生成

## 命令触发

全部命令速查见 `docs/MANUAL.md` 第 6 章。

关键命令（按生命周期排列）：
- "学习 X" — 生成知识卡片骨架
- "审一下" — AI 审核当前卡片
- "练习 X" — 单知识点出题
- "练习今天的知识点" — 今日批量练习
- "闯关 X" — 综合出卷
- "帮我安排今天的练习" — 复习调度
- "检测知识图谱" — 健康报告
- "这题我彻底会了" — 标记掌握
- "看 X 知识脉络" — Roadmap 概览
- "串联 X" — 依赖链讲解
- "面试前复习 X" — 模拟面试
- `/sync [domain]` — 触发 Level 1+2 链接维护
- `/cross-link A B` — 跨域联想
- `/undo` — 撤销 AI 刚做的链接或文件变更

## 模型分配策略

- **深度推理、跨域联想、复杂编码**：`ds-hex/deepseek-v4-pro`（默认模型）
- **编码题出题、方案设计、知识写作**：`xh/claude-sonnet-4-6`
- **日常问答、格式检查、简单补全**：`ds-hex/deepseek-v4-flash`
- **快速批改、标签补全、轻量任务**：`xh/claude-haiku-4-5`

## 早期测评残留

根目录 `.go` 文件 和 `go-interview-assessment/` 是加密的 TSD 格式测评数据，不要尝试解码或修改。`.vscode/settings.json` 禁用了编辑器补全（保证测评公平性），如果影响正常编码体验可恢复。
