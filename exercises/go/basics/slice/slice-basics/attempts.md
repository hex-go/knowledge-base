## 尝试 1 · 2026-06-01 · ✓

**AI 批改：**
- ✅ 正确：`CloneSlice` — nil 判断 + make(len, cap) + copy，覆盖 nil/empty/normal 三种情况
- ✅ 正确：`InsertAt` — 创建新切片、分段拷贝、插入值、移位后段，不污染原切片
- ✅ 正确：`RemoveAt` — 原地移位后 `s = s[:l-1]` 截断，删去残留元素
- + 改进：`InsertAt` 用 append + 切片拼接可简化为一行（`append(s[:index], append([]int{value}, s[index:]...)...)`）

**你的代码：**
```go
func CloneSlice(s []int) []int {
	var out []int
	if s != nil {
		out = make([]int, len(s), cap(s))
		copy(out, s)
	}
	return out
}

func InsertAt(s []int, index, value int) []int {
	newL := len(s) + 1
	out := make([]int, newL)
	copy(out, s[:index])
	out[index] = value
	for i := index + 1; i < newL; i++ {
		out[i] = s[i-1]
	}
	return out
}

func RemoveAt(s []int, index int) []int {
	l := len(s)
	for i := index; i < l-1; i++ {
		s[i] = s[i+1]
	}
	s = s[:l-1]
	return s
}
```
