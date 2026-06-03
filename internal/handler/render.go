package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"regexp"
	"strings"

	"github.com/hex-go/knowledge-base/web"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func markdownToHTML(md string) template.HTML {
	var buf bytes.Buffer
	mdEngine := goldmark.New(
		goldmark.WithExtensions(
			extension.Table,
			extension.Strikethrough,
			extension.TaskList,
			extension.Linkify,
		),
	)
	if err := mdEngine.Convert([]byte(md), &buf); err != nil {
		return template.HTML(template.HTMLEscapeString(md))
	}
	html := replaceMermaid(replaceWikiLinks(buf.String()))
	return template.HTML(html)
}

func replaceMermaid(html string) string {
	re := regexp.MustCompile(`<pre><code class="language-mermaid">([\s\S]*?)</code></pre>`)
	return re.ReplaceAllString(html, `<pre class="mermaid">$1</pre>`)
}

func replaceWikiLinks(md string) string {
	re := regexp.MustCompile(`\[\[([^\]|]+)(?:\|([^\]]+))?\]\]`)
	return re.ReplaceAllStringFunc(md, func(match string) string {
		sub := re.FindStringSubmatch(match)
		path := sub[1]
		text := sub[2]
		if text == "" {
			text = path
		}
		return fmt.Sprintf(`<a href="/%s.md">%s</a>`, path, text)
	})
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
