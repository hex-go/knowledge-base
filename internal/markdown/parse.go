package markdown

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/hex-go/hex-go-interview/internal/types"
)

// ParseFile 读取并解析一个 .md 文件
func ParseFile(baseDir, relPath string) (*types.MarkdownFile, error) {
	fullPath := filepath.Join(baseDir, relPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	mf := &types.MarkdownFile{
		RelPath: filepath.ToSlash(relPath),
	}
	content := string(data)

	// 解析 YAML frontmatter
	if strings.HasPrefix(content, "---\n") || strings.HasPrefix(content, "---\r\n") {
		end := strings.Index(content[4:], "\n---")
		if end == -1 {
			end = strings.Index(content[4:], "\r\n---")
		}
		if end > 0 {
			fm := content[4 : 4+end]
			mf.Meta = parseFrontmatter(fm)
			content = content[4+end+4:]
			// 跳过可能的空行
			content = strings.TrimLeft(content, "\r\n")
		}
	}

	mf.Body = content
	mf.Sections = parseSections(content)
	return mf, nil
}

// parseFrontmatter 解析 YAML frontmatter 为 Meta
func parseFrontmatter(yml string) types.Meta {
	m := types.Meta{}
	lines := strings.Split(yml, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		colon := strings.Index(line, ":")
		if colon < 0 {
			continue
		}
		key := strings.TrimSpace(line[:colon])
		val := strings.TrimSpace(line[colon+1:])

		switch key {
		case "title":
			m.Title = unquote(val)
		case "category":
			m.Category = val
		case "tags":
			m.Tags = parseYAMLList(val)
		case "difficulty":
			m.Difficulty = val
		case "status":
			m.Status = val
		case "wrong_count":
			m.WrongCount, _ = strconv.Atoi(val)
		case "last_review":
			m.LastReview = val
		case "review_interval_days":
			m.ReviewInterval, _ = strconv.Atoi(val)
		case "related":
			m.Related = parseYAMLList(val)
		case "type":
			m.Type = val
		case "exercise_ref":
			m.ExerciseRef = val
		case "last_wrong_date":
			m.LastWrongDate = val
		}
	}
	return m
}

// parseSections 按 ## 标题将正文拆分为 Section 列表
func parseSections(body string) []types.Section {
	var sections []types.Section
	// 匹配 ## 标题（不匹配 ###）
	re := regexp.MustCompile(`(?m)^## (.+)$`)
	matches := re.FindAllStringIndex(body, -1)
	if matches == nil {
		return sections
	}

	for i, m := range matches {
		heading := body[m[0]+3 : m[1]]
		start := m[1] + 1 // 跳过标题行后的换行
		end := len(body)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}
		sections = append(sections, types.Section{
			Heading: strings.TrimSpace(heading),
			Content: strings.TrimSpace(body[start:end]),
		})
	}
	return sections
}

// FindSection 找到指定标题的 Section，返回 index（-1 表示未找到）
func FindSection(sections []types.Section, heading string) int {
	for i, s := range sections {
		if strings.TrimSpace(s.Heading) == heading {
			return i
		}
	}
	return -1
}

// ParseAttempts 解析 attempts.md 文件中的尝试记录
func ParseAttempts(baseDir, relPath string) ([]types.AttemptBlock, error) {
	fullPath := filepath.Join(baseDir, relPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	content := string(data)

	// 匹配 ## 尝试 N · YYYY-MM-DD · ✓/✗
	re := regexp.MustCompile(`(?m)^## 尝试 (\d+) · (\d{4}-\d{2}-\d{2}) · (.+)$`)
	matches := re.FindAllStringSubmatchIndex(content, -1)

	var blocks []types.AttemptBlock
	for i, m := range matches {
		idx, _ := strconv.Atoi(content[m[2]:m[3]])
		date := content[m[4]:m[5]]
		result := content[m[6]:m[7]]
		passed := strings.Contains(result, "✓") || strings.Contains(result, "通过")

		start := m[0]
		end := len(content)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}

		blocks = append(blocks, types.AttemptBlock{
			Index:   idx,
			Date:    date,
			Passed:  passed,
			RawText: strings.TrimSpace(content[start:end]),
		})
	}
	return blocks, nil
}

// DeleteAttemptBlock 从 attempts.md 中删除指定的尝试块
func DeleteAttemptBlock(baseDir, relPath string, attemptIndex int) error {
	fullPath := filepath.Join(baseDir, relPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}
	content := string(data)

	re := regexp.MustCompile(`(?m)^## 尝试 \d+ · \d{4}-\d{2}-\d{2} · .+$`)
	matches := re.FindAllStringIndex(content, -1)

	var blocks []struct{ start, end int }
	for i, m := range matches {
		end := len(content)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}
		blocks = append(blocks, struct{ start, end int }{m[0], end})
	}

	if attemptIndex < 0 || attemptIndex >= len(blocks) {
		return fmt.Errorf("attempt index %d out of range (0-%d)", attemptIndex, len(blocks)-1)
	}

	// 拼接删除后的内容
	before := strings.TrimRight(content[:blocks[attemptIndex].start], "\r\n")
	after := ""
	if attemptIndex+1 < len(blocks) {
		after = content[blocks[attemptIndex+1].start:]
	}
	newContent := strings.TrimRight(before, "\r\n") + "\n\n" + strings.TrimLeft(after, "\r\n")

	return os.WriteFile(fullPath, []byte(newContent), 0644)
}

func unquote(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

func parseYAMLList(s string) []string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		s = s[1 : len(s)-1]
	}
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	var result []string
	for _, p := range parts {
		p = unquote(strings.TrimSpace(p))
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}
