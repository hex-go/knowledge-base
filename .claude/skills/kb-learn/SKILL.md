---
name: kb-learn
description: Guides the knowledge graph learning stage by generating card skeletons, auditing understanding, and applying approved semantic links. Use when user says "学习 X", "生成知识卡片 X", "审一下", "审核 X", "保留", "撤销第 N 条", "更新知识图谱", or "/sync".
---

# kb-learn

## 触发词

- "学习 X"
- "生成知识卡片 X"
- "审一下"
- "审核 X"
- "保留"
- "撤销第 N 条"
- "更新知识图谱"
- "/sync [domain]"

## 工作流

1. 定位用户要学习的知识点路径。
2. 读取对应 `roadmap.md`，确认 `level`、`requires`、父级主题和文件名。文件名即图谱节点标签，命名时遵循 CLAUDE.md 规范（精简、短横线连接、无领域前缀）。
3. 创建知识卡片，写入 frontmatter、`## 标准知识`、`## 我的理解` 空区块。从 roadmap 提取中文名填入 `aliases`。生成后在 WSL 中执行 `code <path>` 打开文件。
4. 允许用户边写边问，但不提前补链。
5. 用户触发 "审一下" 后，逐段审核准确性、遗漏和表达清晰度。
6. 审核通过后，按 `docs/MANUAL.md` 的 Level 2 规则做语义补链。
7. 列出候选链接，等待用户 "保留" 或 "撤销第 N 条"。

## 参考

- `docs/MANUAL.md` §5.2.1
- `docs/MANUAL.md` §3.1
- `docs/MANUAL.md` §4
