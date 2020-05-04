package operation_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 3, 4}
	c := [...]int{1, 3, 3, 4, 5}
	d := [...]int{1, 3, 3, 4, 6}

	// 数组和类型相同可以比较
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(c == d)
}
