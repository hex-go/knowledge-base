---
name: kb-roadmap
description: Guides the roadmap-based global review workflow by rendering the knowledge tree, finding weak spots, traversing dependency chains, and running interview-style reviews. Use when user says "看 X 知识脉络", "X 有哪些薄弱点", "串联 X", "面试前复习 X", or "用阿里 P7 的角度面我".
---

# kb-roadmap

## 触发词

- "看 X 知识脉络"
- "X 有哪些薄弱点"
- "串联 X"
- "面试前复习 X"
- "用阿里 P7 的角度面我"

## 工作流

1. 从对应领域或维度的 `roadmap.md` 读取全部计划知识点。
2. 对照实际文件状态，渲染知识树和进度。
3. 扫描 `wrong_count`、过期复习、空理解区和长期停滞，输出薄弱点排序。
4. 沿 `requires` 链做串联讲解，解释“为什么先学这个再学那个”。
5. 面试模式下逐主题出题：已掌握跳过，学习中追问，未创建则询问是否要学。
6. 每个维度结束时给出总结和下一步建议。

## 参考

- `docs/MANUAL.md` §5.3.1
- `docs/MANUAL.md` §5.3.2
- `docs/MANUAL.md` §5.3.3
- `docs/MANUAL.md` §5.3.4
