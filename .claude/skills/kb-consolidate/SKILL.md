---
name: kb-consolidate
description: Guides the knowledge graph consolidation stage by scheduling reviews, checking graph health, splitting or merging knowledge cards, and marking mastery. Use when user says "帮我安排今天的练习", "检测知识图谱", "修复断链", "修复孤儿节点", "拆分 [[X]]", "把 A 合并回 B", "X 我懂了", or "这题我彻底会了".
---

# kb-consolidate

## 触发词

- "帮我安排今天的练习"
- "检测知识图谱"
- "修复断链"
- "修复孤儿节点"
- "拆分 [[X]]"
- "把 A 合并回 B"
- "X 我懂了"
- "这题我彻底会了"

## 工作流

1. 扫描 `knowledge/`，找出超出复习间隔的知识点和 `wrong-book/` 中正在复习的条目。
2. 输出今日复习清单、建议新学知识点和练习量。
3. 执行健康检查：断链、孤儿、过期复习、空理解区、孤立知识点。
4. 对断链和孤儿节点给出修复建议，等待用户确认。
5. 需要拆分时，按概念子边界预览、拆文件、建目录、更新 roadmap。
6. 需要合并时，反向合并子卡片回父卡片。
7. 用户确认掌握后，将知识点和对应 wrong-book 条目标记为 `mastered`。

## 参考

- `docs/MANUAL.md` §5.2.3
