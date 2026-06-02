package main

import "fmt"

// AppendInt adds elements to the end of a slice, growing it if needed.
// It mimics Go's built-in append behavior including expansion rules.
func AppendInt(s *[]int, elems ...int) {
	// TODO: implement
}

func main() {
	// Test 1: no expansion needed
	s1 := make([]int, 0, 4)
	AppendInt(&s1, 1, 2)
	fmt.Printf("Test 1: len=%d cap=%d data=%v (expected len=2 cap=4 [1 2])\n", len(s1), cap(s1), s1)

	// Test 2: expansion triggered (oldCap < 256, double)
	s2 := make([]int, 0, 2)
	AppendInt(&s2, 1)
	AppendInt(&s2, 2)
	AppendInt(&s2, 3) // should trigger growth: 2*2=4
	fmt.Printf("Test 2: len=%d cap=%d data=%v (expected len=3 cap=4 [1 2 3])\n", len(s2), cap(s2), s2)

	// Test 3: multiple elements trigger growth
	s3 := make([]int, 0, 2)
	AppendInt(&s3, 1, 2, 3) // 3 elems into cap=2, need growth: 2*2=4 fits 3
	fmt.Printf("Test 3: len=%d cap=%d data=%v (expected len=3 cap=4 [1 2 3])\n", len(s3), cap(s3), s3)

	// Test 4: large slice growth (oldCap >= 256, +25%)
	s4 := make([]int, 256, 256)
	oldCap := cap(s4)
	AppendInt(&s4, 1)
	fmt.Printf("Test 4: oldCap=%d newCap=%d (expected ~%d, +~25%%)\n", oldCap, cap(s4), oldCap+oldCap/4)

	// Test 5: empty append (no-op)
	s5 := make([]int, 0, 2)
	AppendInt(&s5) // nothing to append
	fmt.Printf("Test 5: len=%d cap=%d data=%v (expected len=0 cap=2 [])\n", len(s5), cap(s5), s5)

	// Test 6: oldCap == 0
	s6 := make([]int, 0, 0)
	AppendInt(&s6, 1)
	fmt.Printf("Test 6: len=%d cap=%d data=%v (expected len=1 cap=1 [1])\n", len(s6), cap(s6), s6)
}
