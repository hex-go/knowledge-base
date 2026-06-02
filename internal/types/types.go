package types

// Meta 表示 Markdown 文件的 YAML frontmatter 元数据
type Meta struct {
	Title           string   `json:"title"`
	Category        string   `json:"category"`
	Tags            []string `json:"tags"`
	Difficulty      string   `json:"difficulty"`
	Status          string   `json:"status"`
	Domain          string   `json:"domain,omitempty"`
	Level           string   `json:"level,omitempty"`
	WrongCount      int      `json:"wrong_count"`
	LastReview      string   `json:"last_review"`
	ReviewInterval  int      `json:"review_interval_days"`
	Related         []string `json:"related"`
	Type            string   `json:"type"`            // 仅错题本: coding | concept
	ExerciseRef     string   `json:"exercise_ref"`    // 仅错题本: 原题路径
	LastWrongDate   string   `json:"last_wrong_date"` // 仅错题本
	ReviewCount     int      `json:"review_count,omitempty"`
}

// MarkdownFile 表示一个完整的 .md 文件的解析结果
type MarkdownFile struct {
	Meta     Meta              `json:"meta"`
	Body     string            `json:"body"`    // frontmatter 之后的全部正文
	Sections []Section         `json:"sections"` // 按 ## 分隔的区块
	RelPath  string            `json:"rel_path"` // 相对于仓库根的路径
}

// Section 表示正文中一个 ## 标题引出的区块
type Section struct {
	Heading string `json:"heading"` // 不含 ## 前缀的标题文本
	Content string `json:"content"` // 区块内容
}

// AttemptBlock 表示答题记录中的一次尝试（解析自 attempts.md）
type AttemptBlock struct {
	Index  int    // 尝试序号
	Date   string // YYYY-MM-DD
	Passed bool   // 是否通过
	RawText string // 整个 ## 尝试 N 区块的原始文本
}

// CategoryNode 目录树节点
type CategoryNode struct {
	Name     string         `json:"name"`
	Path     string         `json:"path"` // 相对路径
	IsLeaf   bool           `json:"is_leaf"`
	File     *MarkdownFile  `json:"file,omitempty"`
	Children []CategoryNode `json:"children,omitempty"`
}
