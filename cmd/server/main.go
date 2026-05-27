package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hex-go/hex-go-interview/internal/handler"
)

func main() {
	port := flag.Int("port", 8080, "监听端口")
	flag.Parse()

	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取工作目录失败: %v", err)
	}

	// 确保从仓库根目录执行
	if _, err := os.Stat(filepath.Join(baseDir, "go.mod")); err != nil {
		log.Fatalf("请从仓库根目录运行: go run ./cmd/server")
	}

	srv, err := handler.NewServer(baseDir)
	if err != nil {
		log.Fatalf("初始化服务: %v", err)
	}

	mux := http.NewServeMux()
	srv.RegisterRoutes(mux)

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Go 学习助手已启动 → http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
