package type_test

import (
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	var c MyInt
	b = int64(a)
	c = MyInt(a)
	t.Log(b)
	t.Log(c)

}

func TestPoint(t *testing.T) {
	// 指针不支持运算的
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(aPtr)
}

func TestString(t *testing.T) {
	var s string // 默认初始化 ""
	t.Log("*" + s + "*")
	t.Log(len(s))
}
