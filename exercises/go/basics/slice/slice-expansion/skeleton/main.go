package main

import "fmt"

// CalcNewCap 计算 Go 扇片扩容后的新宺量.
// oldCap：当前容量，numNew：欲追加的元素个数。
func CalcNewCap(oldCap, numNew int) int {
	// TODO: 实现扩容适辑
	return 0
}

func main() {
	tests := []struct {
		oldCap, numNew, want int
		desc                 string
	}{
		{4, 10, 14, "need-big: direct"},
		{4, 1, 8, "small: double"},
		{128, 1, 256, "small-edge: double"},
		{256, 1, 512, "big-edge: grow"},
		{512, 1, 832, "big: grow"},
		{1024, 1, 1472, "big: grow"},
		{10, 5, 20, "multi: double"},
		{512, 100, 832, "multi-big: grow"},
	}
	passed := 0
	for _, tt := range tests {
		got := CalcNewCap(tt.oldCap, tt.numNew)
		status := "PASS"
		if got != tt.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] %s | oldCap=%d numNew=%d got=%d want=%d\n",
				status, tt.desc, tt.oldCap, tt.numNew, got, tt.want)
		if status == "PASS" {
			passed++
		}
	}
	fmt.Printf("\n%d/%d passed\n", passed, len(tests))
}
