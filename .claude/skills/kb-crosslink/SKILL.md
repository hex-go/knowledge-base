---
name: kb-crosslink
description: Handles manual cross-domain association by comparing two domains, proposing bridge candidates, and generating cross-domain bridge files. Use when user says "/cross-link A B" or "联想一下 A 和 B".
---

# kb-crosslink

## 触发词

- "/cross-link A B"
- "联想一下 A 和 B"

## 工作流

1. 加载两个领域的 `knowledge/` 文件。
2. 识别候选关联：共通模式、概念映射、问题复现。
3. 逐条给出候选关联并说明理由。
4. 等待用户确认后，生成 `cross-domain/` bridge 文件。
5. bridge 文件写清 `domains`、`pattern`、核心洞察，并链接两端知识点。

## 参考

- `docs/MANUAL.md` §5.4
- `docs/MANUAL.md` §3.5
