# knowledge-base

个人技术知识图谱。基于 **Skill Graph + Obsidian + AI 自动维护**。

## 快速开始

```bash
# Web 浏览
go run ./cmd/server              # http://localhost:8080

# Obsidian 浏览
Obsidian → Open folder as vault → 选本目录

# Opencode 陪练
opencode
```

## 目录全景

```
knowledge/       知识点卡片（go / k8s / linux / python / ai / databases / security / distributed / patterns）
exercises/       编码题（按领域子目录）
concepts/        概念问答
wrong-book/      错题本（统一，domain 标签隔离）
cross-domain/    跨域 bridge 文件
```

## 完整手册

参见 [`docs/MANUAL.md`](docs/MANUAL.md)，包含：

- 文件模板速查
- 知识图谱链接规范（三级链接策略 + 领域隔离）
- 所有工作流与命令速查
- 模型分配策略
- 工具链指南（opencode / Obsidian / Web / VS Code）
- Skill 清单

## AI 行为规则

参见 [`CLAUDE.md`](CLAUDE.md)，定义 AI 在仓库中的行为规范。
