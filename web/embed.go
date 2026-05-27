package web

import "embed"

// Templates 嵌入 HTML 模板文件
//
//go:embed templates/*.html
var Templates embed.FS

// Static 嵌入静态文件（CSS 等）
//
//go:embed static/*
var Static embed.FS
