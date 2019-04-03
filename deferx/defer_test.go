package deferx

import "testing"

func TestParseDefer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("recover got it\n")
			t.Error(err)
		}
	}()

	ParseDefer()
}
