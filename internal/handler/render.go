package handler

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"regexp"
	"strings"

	"github.com/hex-go/knowledge-base/web"
)

// markdownToHTML 将 markdown 文本转为简单 HTML
func markdownToHTML(md string) template.HTML {
	// 代码块
	md = replaceCodeBlocks(md)
	// 行内代码
	md = replaceInlineCode(md)
	// 加粗
	md = replaceBold(md)
	// 段落（连续的文本行）
	md = replaceParagraphs(md)

	return template.HTML(md)
}

func replaceCodeBlocks(md string) string {
	re := regexp.MustCompile("```(?s)(.*?)```")
	return re.ReplaceAllStringFunc(md, func(match string) string {
		inner := match[3 : len(match)-3]
		// 去掉语言标识
		if idx := strings.Index(inner, "\n"); idx >= 0 {
			inner = inner[idx+1:]
		}
		return "<pre><code>" + template.HTMLEscapeString(strings.TrimRight(inner, "\r\n")) + "</code></pre>"
	})
}

func replaceInlineCode(md string) string {
	re := regexp.MustCompile("`([^`]+)`")
	return re.ReplaceAllString(md, "<code>$1</code>")
}

func replaceBold(md string) string {
	re := regexp.MustCompile(`\*\*(.+?)\*\*`)
	return re.ReplaceAllString(md, "<strong>$1</strong>")
}

func replaceParagraphs(md string) string {
	lines := strings.Split(md, "\n")
	var out []string
	var buf []string

	flush := func() {
		if len(buf) > 0 {
			out = append(out, "<p>"+strings.Join(buf, " ") +"</p>")
			buf = buf[:0]
		}
	}

	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		switch {
		case strings.HasPrefix(line, "<pre>"), strings.HasPrefix(line, "<h"),
			strings.HasPrefix(line, "<ul>"), strings.HasPrefix(line, "<ol>"),
			strings.HasPrefix(line, "<li>"), strings.HasPrefix(line, "</"):
			flush()
			out = append(out, line)
		case line == "":
			flush()
		default:
			buf = append(buf, line)
		}
	}
	flush()

	return strings.Join(out, "\n")
}

type Templates struct {
	pages map[string]*template.Template // page name -> parsed set (layout + page)
}

func loadTemplates() (*Templates, error) {
	funcs := template.FuncMap{
		"markdown": markdownToHTML,
	}

	entries, err := fs.ReadDir(web.Templates, "templates")
	if err != nil {
		return nil, err
	}

	pages := make(map[string]*template.Template)
	for _, e := range entries {
		if e.IsDir() || e.Name() == "layout.html" {
			continue
		}
		pageName := strings.TrimSuffix(e.Name(), ".html")

		// 每个页面用 ParseFS 组合 layout + page
		tmpl := template.New("layout").Funcs(funcs)
		if _, err := tmpl.ParseFS(web.Templates, "templates/layout.html", "templates/"+e.Name()); err != nil {
			return nil, fmt.Errorf("解析 %s: %w", e.Name(), err)
		}
		pages[pageName] = tmpl
	}

	return &Templates{pages: pages}, nil
}

func (t *Templates) Exec(w http.ResponseWriter, pageName string, data interface{}) error {
	tmpl, ok := t.pages[pageName]
	if !ok {
		return fmt.Errorf("未找到页面模板: %s", pageName)
	}
	return tmpl.ExecuteTemplate(w, "layout.html", data)
}
