package pointer_routine

import (
	"testing"
	"unsafe"
	"time"
	"syscall"
)

func TestForRoutine(t *testing.T) {
	var ta []*testA

	ta = append(ta, &testA{A: "kwanhur"})
	ta = append(ta, &testA{A: "hello"})

	for _, a := range ta {
		go func(tea *testA) {
			t.Logf("address:%v, val:%s", unsafe.Pointer(tea), tea.A)
		}(a)
	}

	time.Sleep(1 * time.Second)
}

func TestGid(t *testing.T) {
	t.Log(syscall.Getgid())
}
