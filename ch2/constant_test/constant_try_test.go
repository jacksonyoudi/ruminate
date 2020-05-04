package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstanTry1(t *testing.T) {
	t.Log(Readable)
	t.Log(Writeable)
	t.Log(Wednesday)
}
