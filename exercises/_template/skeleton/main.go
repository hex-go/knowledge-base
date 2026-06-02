package main

import (
	"fmt"
	"reflect"
)

// TODO: 在此实现你的解法

func main() {
	tests := []struct {
		name     string
		input    any
		expected any
		// 按需调整字段
	}{
		// TODO: 添加测试用例
	}

	var total, passed int
	for _, tt := range tests {
		total++
		// TODO: 调用你的函数并比较结果
		ok := reflect.DeepEqual(nil, tt.expected)
		if ok {
			passed++
			fmt.Printf("  PASS  %s\n", tt.name)
		} else {
			fmt.Printf("  FAIL  %s\n", tt.name)
		}
	}
	fmt.Printf("\n%d/%d passed\n", passed, total)
}
