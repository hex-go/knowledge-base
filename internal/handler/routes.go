package handler

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sort"
	"time"

	"github.com/hex-go/knowledge-base/internal/markdown"
	"github.com/hex-go/knowledge-base/internal/types"
	"github.com/hex-go/knowledge-base/web"
)

// Server 管理所有 HTTP 路由
type Server struct {
	baseDir string
	tmpl    *Templates
}

// NewServer 创建服务实例
func NewServer(baseDir string) (*Server, error) {
	s := &Server{baseDir: baseDir}
	var err error
	s.tmpl, err = loadTemplates()
	if err != nil {
		return nil, fmt.Errorf("加载模板: %w", err)
	}
	return s, nil
}

// RegisterRoutes 注册所有路由
func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/static/", s.handleStatic)
	mux.HandleFunc("/", s.handleHome)
	mux.HandleFunc("/knowledge/", s.handleKnowledge)
	mux.HandleFunc("/concepts/", s.handleConcepts)
	mux.HandleFunc("/exercises/", s.handleExercises)
	mux.HandleFunc("/wrong-book/", s.handleWrongBook)
	mux.HandleFunc("/search", s.handleSearch)
}

// --- Static files ---

func (s *Server) handleStatic(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/static/")
	data, err := fs.ReadFile(web.Static, "static/"+path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	contentType := "text/plain"
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	}
	w.Header().Set("Content-Type", contentType)
	w.Write(data)
}

// --- Home / Dashboard ---

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	knFiles := s.scanFiles("knowledge")
	exFiles := s.scanFiles("exercises")
	conFiles := s.scanFiles("concepts")
	wbFiles := s.scanFiles("wrong-book")

	// 待复习错题
	reviewCount := 0
	for _, f := range wbFiles {
		if f.Meta.Status == "reviewing" {
			reviewCount++
		}
	}

	// 需要复习的知识点（超过 review_interval_days）
	var reviewItems []map[string]string
	today := time.Now().Format("2006-01-02")
	for _, f := range knFiles {
		if f.Meta.LastReview == "" || f.Meta.Status == "mastered" {
			continue
		}
		lastReview, _ := time.Parse("2006-01-02", f.Meta.LastReview)
		if time.Since(lastReview).Hours() > float64(f.Meta.ReviewInterval*24) {
			reviewItems = append(reviewItems, map[string]string{
				"Title": f.Meta.Title,
				"Link":  "/knowledge/" + f.RelPath,
				"Meta":  "上次复习: " + f.Meta.LastReview,
				"Type":  "overdue",
			})
		}
	}
	for _, f := range wbFiles {
		if f.Meta.Status == "reviewing" {
			reviewItems = append(reviewItems, map[string]string{
				"Title": f.Meta.Title,
				"Link":  "/wrong-book/" + f.RelPath,
				"Meta":  "错误 " + fmt.Sprintf("%d 次", f.Meta.WrongCount),
				"Type":  "wrong",
			})
		}
	}

	s.render(w, "home.html", map[string]interface{}{
		"ReviewCount":   reviewCount,
		"KnowledgeCount": len(knFiles),
		"ExerciseCount":  len(exFiles),
		"ConceptCount":   len(conFiles),
		"ReviewItems":    reviewItems,
		"Today":          today,
	})
}

// --- Knowledge ---

func (s *Server) handleKnowledge(w http.ResponseWriter, r *http.Request) {
	relPath := strings.TrimPrefix(r.URL.Path, "/knowledge/")
	relPath = strings.TrimPrefix(relPath, "/")

	if r.Method == http.MethodPost {
		s.handleKnowledgePost(w, r, relPath)
		return
	}

	// 知识库根目录 -> 列出所有
	if relPath == "" || relPath == "/" {
		tree := s.buildKnowledgeTree()
		s.render(w, "knowledge.html", map[string]interface{}{
			"Tree": tree,
		})
		return
	}

	// 具体知识点
	f, err := markdown.ParseFile(s.baseDir, "knowledge/"+relPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s.render(w, "knowledge-detail.html", map[string]interface{}{
		"File":     f,
		"Sections": f.Sections,
	})
}

func (s *Server) handleKnowledgePost(w http.ResponseWriter, r *http.Request, relPath string) {
	if strings.HasSuffix(relPath, "/review") {
		relPath = strings.TrimSuffix(relPath, "/review")
		s.markReviewed("knowledge/" + relPath)
		http.Redirect(w, r, "/knowledge/"+relPath, http.StatusSeeOther)
		return
	}
	if strings.HasSuffix(relPath, "/master") {
		relPath = strings.TrimSuffix(relPath, "/master")
		s.markMastered("knowledge/" + relPath)
		http.Redirect(w, r, "/knowledge/"+relPath, http.StatusSeeOther)
		return
	}
	if strings.HasSuffix(relPath, "/edit") {
		relPath = strings.TrimSuffix(relPath, "/edit")
		r.ParseForm()
		section := r.FormValue("section")
		content := r.FormValue("content")
		s.updateSection("knowledge/"+relPath, section, content)
		http.Redirect(w, r, "/knowledge/"+relPath, http.StatusSeeOther)
		return
	}
	http.NotFound(w, r)
}

// --- Concepts ---

func (s *Server) handleConcepts(w http.ResponseWriter, r *http.Request) {
	relPath := strings.TrimPrefix(r.URL.Path, "/concepts/")
	relPath = strings.TrimPrefix(relPath, "/")

	if r.Method == http.MethodPost {
		s.handleConceptPost(w, r, relPath)
		return
	}

	// 列表
	if relPath == "" || relPath == "/" {
		files := s.scanFiles("concepts")
		s.render(w, "concepts.html", map[string]interface{}{"Files": files})
		return
	}

	// 详情
	f, err := markdown.ParseFile(s.baseDir, "concepts/"+relPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// 查找对应 attempts 文件
	base := strings.TrimSuffix(relPath, filepath.Ext(relPath))
	attemptsPath := "concepts/" + base + "-attempts.md"
	attempts, _ := markdown.ParseAttempts(s.baseDir, attemptsPath)

	s.render(w, "concept-detail.html", map[string]interface{}{
		"File":     f,
		"Sections": f.Sections,
		"Attempts": attempts,
	})
}

func (s *Server) handleConceptPost(w http.ResponseWriter, r *http.Request, relPath string) {
	if strings.HasSuffix(relPath, "/delete-attempt") {
		relPath = strings.TrimSuffix(relPath, "/delete-attempt")
		r.ParseForm()
		idx := r.FormValue("index")
		index := 0
		fmt.Sscanf(idx, "%d", &index)
		base := strings.TrimSuffix(relPath, filepath.Ext(relPath))
		markdown.DeleteAttemptBlock(s.baseDir, "concepts/"+base+"-attempts.md", index)
		http.Redirect(w, r, "/concepts/"+relPath, http.StatusSeeOther)
		return
	}
	http.NotFound(w, r)
}

// --- Exercises ---

func (s *Server) handleExercises(w http.ResponseWriter, r *http.Request) {
	relPath := strings.TrimPrefix(r.URL.Path, "/exercises/")
	relPath = strings.TrimPrefix(relPath, "/")

	if r.Method == http.MethodPost {
		s.handleExercisePost(w, r, relPath)
		return
	}

	// 列表
	if relPath == "" || relPath == "/" {
		files := s.scanExercises()
		s.render(w, "exercises.html", map[string]interface{}{"Files": files})
		return
	}

	// 详情
	relPath = strings.TrimSuffix(relPath, "/")
	dir := filepath.Dir(relPath)

	var f *types.MarkdownFile
	// 尝试作为目录处理
	fullDir := filepath.Join(s.baseDir, "exercises", relPath)
	info, err := os.Stat(fullDir)
	if err != nil || !info.IsDir() {
		// 尝试作为 .md 文件
		f, err = markdown.ParseFile(s.baseDir, "exercises/"+relPath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		s.render(w, "exercise-detail.html", map[string]interface{}{
			"File":     f,
			"Sections": f.Sections,
		})
		return
	}

	_ = dir
	f, err = markdown.ParseFile(s.baseDir, "exercises/"+relPath+"/problem.md")
	if err != nil {
		http.NotFound(w, r)
		return
	}

	attempts, _ := markdown.ParseAttempts(s.baseDir, "exercises/"+relPath+"/attempts.md")

	s.render(w, "exercise-detail.html", map[string]interface{}{
		"File":     f,
		"Sections": f.Sections,
		"Attempts": attempts,
	})
}

func (s *Server) handleExercisePost(w http.ResponseWriter, r *http.Request, relPath string) {
	if strings.HasSuffix(relPath, "/delete-attempt") {
		relPath = strings.TrimSuffix(relPath, "/delete-attempt")
		r.ParseForm()
		idx := r.FormValue("index")
		index := 0
		fmt.Sscanf(idx, "%d", &index)
		markdown.DeleteAttemptBlock(s.baseDir, "exercises/"+relPath+"/attempts.md", index)
		http.Redirect(w, r, "/exercises/"+relPath, http.StatusSeeOther)
		return
	}
	http.NotFound(w, r)
}

// --- Wrong Book ---

func (s *Server) handleWrongBook(w http.ResponseWriter, r *http.Request) {
	relPath := strings.TrimPrefix(r.URL.Path, "/wrong-book/")
	relPath = strings.TrimPrefix(relPath, "/")

	if r.Method == http.MethodPost {
		s.handleWrongBookPost(w, r, relPath)
		return
	}

	// 列表
	if relPath == "" || relPath == "/" {
		files := s.scanFiles("wrong-book")
		// 按 wrong_count 降序
		sort.Slice(files, func(i, j int) bool {
			return files[i].Meta.WrongCount > files[j].Meta.WrongCount
		})
		s.render(w, "wrongbook.html", map[string]interface{}{"Files": files})
		return
	}

	// 详情
	f, err := markdown.ParseFile(s.baseDir, "wrong-book/"+relPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s.render(w, "wrongbook-detail.html", map[string]interface{}{
		"File":     f,
		"Sections": f.Sections,
	})
}

func (s *Server) handleWrongBookPost(w http.ResponseWriter, r *http.Request, relPath string) {
	if strings.HasSuffix(relPath, "/master") {
		relPath = strings.TrimSuffix(relPath, "/master")
		s.markMastered("wrong-book/" + relPath)
		// 也标记原题为已掌握
		f, err := markdown.ParseFile(s.baseDir, "wrong-book/"+relPath)
		if err == nil && f.Meta.ExerciseRef != "" {
			s.markMastered(f.Meta.ExerciseRef + "/problem.md")
		}
		http.Redirect(w, r, "/wrong-book/"+relPath, http.StatusSeeOther)
		return
	}
	http.NotFound(w, r)
}

// --- Search ---

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		s.render(w, "search.html", map[string]interface{}{})
		return
	}

	type result struct {
		Title string
		Link  string
		Kind  string
		Meta  string
	}
	var results []result

	searchDirs := map[string]string{
		"knowledge":  "知识点",
		"concepts":   "概念",
		"exercises":  "编码题",
		"wrong-book": "错题",
	}

	for dir, kind := range searchDirs {
		files := s.scanFiles(dir)
		for _, f := range files {
			if strings.Contains(strings.ToLower(f.Meta.Title), strings.ToLower(q)) ||
				strings.Contains(strings.ToLower(f.Body), strings.ToLower(q)) {
				link := "/" + dir + "/" + f.RelPath
				if kind == "编码题" {
					link = "/" + dir + "/" + filepath.Dir(f.RelPath)
				}
				results = append(results, result{
					Title: f.Meta.Title,
					Link:  link,
					Kind:  kind,
					Meta:  f.Meta.Category,
				})
			}
		}
	}

	s.render(w, "search.html", map[string]interface{}{
		"Query":   q,
		"Results": results,
	})
}

// --- helpers ---

func (s *Server) render(w http.ResponseWriter, name string, data map[string]interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if data == nil {
		data = map[string]interface{}{}
	}
	if _, ok := data["Query"]; !ok {
		data["Query"] = ""
	}
	pageName := strings.TrimSuffix(name, ".html")
	if err := s.tmpl.Exec(w, pageName, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) scanFiles(dir string) []*types.MarkdownFile {
	var files []*types.MarkdownFile

	filepath.WalkDir(filepath.Join(s.baseDir, dir), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}
		// 跳过模板文件
		if strings.HasPrefix(d.Name(), "_template") || strings.HasPrefix(d.Name(), ".") {
			return nil
		}

		rel, _ := filepath.Rel(s.baseDir+"/"+dir, path)
		mf, err := markdown.ParseFile(s.baseDir+"/"+dir, rel)
		if err != nil {
			return nil
		}

		// 跳过 attempts 文件（由详情页单独加载）
		if strings.Contains(d.Name(), "-attempts") {
			return nil
		}

		files = append(files, mf)
		return nil
	})

	return files
}

// scanExercises 扫描编码题（每个子目录的第一个 problem.md）
func (s *Server) scanExercises() []*types.MarkdownFile {
	var files []*types.MarkdownFile

	exercisesDir := filepath.Join(s.baseDir, "exercises")
	filepath.WalkDir(exercisesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}
		if strings.Contains(d.Name(), "attempts") || strings.Contains(d.Name(), "solution") {
			return nil
		}
		if strings.Contains(path, "_template") {
			return nil
		}

		rel, _ := filepath.Rel(s.baseDir, path)
		mf, err := markdown.ParseFile(s.baseDir, rel)
		if err != nil {
			return nil
		}
		mf.RelPath = strings.TrimPrefix(filepath.ToSlash(rel), "exercises/")
		files = append(files, mf)
		return nil
	})

	return files
}

func (s *Server) buildKnowledgeTree() []types.CategoryNode {
	var roots []types.CategoryNode

	filepath.WalkDir(filepath.Join(s.baseDir, "knowledge"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(s.baseDir+"/knowledge", path)
		rel = filepath.ToSlash(rel)
		log.Printf("[DEBUG] WalkDir knowledge: rel=%s name=%s isDir=%v", rel, d.Name(), d.IsDir())

		if strings.HasPrefix(filepath.Base(rel), "_") || strings.HasPrefix(filepath.Base(rel), ".") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}
		if strings.Contains(d.Name(), "roadmap") || strings.Contains(d.Name(), "-attempts") {
			return nil
		}

		mf, err := markdown.ParseFile(s.baseDir+"/knowledge", rel)
		if err != nil {
			log.Printf("[DEBUG] ParseFile failed for %s: %v", rel, err)
			return nil
		}

		parts := strings.Split(rel, "/")
		current := &roots
		for i, part := range parts {
			if i == len(parts)-1 {
				// 文件节点
				*current = append(*current, types.CategoryNode{
					Name:   part,
					Path:   rel,
					IsLeaf: true,
					File:   mf,
				})
			} else {
				// 目录节点
				found := false
				for j := range *current {
					if (*current)[j].Name == part {
						current = &(*current)[j].Children
						found = true
						break
					}
				}
				if !found {
					*current = append(*current, types.CategoryNode{
						Name:     part,
						IsLeaf:   false,
						Children: nil,
					})
					current = &(*current)[len(*current)-1].Children
				}
			}
		}
		return nil
	})

	return roots
}

// --- file mutation helpers ---

func (s *Server) markReviewed(filePath string) {
	f, err := markdown.ParseFile(s.baseDir, filePath)
	if err != nil {
		return
	}
	updateFrontmatterString(s.baseDir, filePath, "last_review", time.Now().Format("2006-01-02"))
	_ = f
}

func (s *Server) markMastered(filePath string) {
	updateFrontmatterString(s.baseDir, filePath, "status", "mastered")
}

func (s *Server) updateSection(filePath, sectionName, content string) {
	fullPath := filepath.Join(s.baseDir, filePath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return
	}

	text := string(data)

	// 找到 ## sectionName 区块并替换
	re := regexp.MustCompile(`(?m)^## ` + regexp.QuoteMeta(sectionName) + `\s*\n([\s\S]*?)(\n## |\z)`)
	loc := re.FindStringSubmatchIndex(text)
	if loc != nil {
		newText := text[:loc[2]] + content + text[loc[3]:]
		os.WriteFile(fullPath, []byte(newText), 0644)
	} else {
		// 区块不存在，追加到文件末尾
		newSection := "\n\n## " + sectionName + "\n\n" + content
		os.WriteFile(fullPath, []byte(text+newSection), 0644)
	}
}

func updateFrontmatterString(baseDir, filePath, key, value string) {
	_ = baseDir
	fullPath := filepath.Join(baseDir, filePath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return
	}
	text := string(data)

	// 在 frontmatter 中找到 key 并替换
	fmEnd := strings.Index(text[4:], "\n---")
	if fmEnd < 0 {
		return
	}
	fm := text[4 : 4+fmEnd]

	// 已有该 key
	re := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(key) + `:\s*.*$`)
	loc := re.FindStringIndex(fm)
	if loc != nil {
		newFm := fm[:loc[0]] + key + ": " + value + fm[loc[1]:]
		newText := "---\n" + newFm + text[4+fmEnd:]
		os.WriteFile(fullPath, []byte(newText), 0644)
	} else {
		// 新增 key
		newFm := strings.TrimRight(fm, "\r\n") + "\n" + key + ": " + value
		newText := "---\n" + newFm + text[4+fmEnd:]
		os.WriteFile(fullPath, []byte(newText), 0644)
	}
}

// --- JSON API (for potential future use) ---

func (s *Server) handleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
