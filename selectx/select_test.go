package selectx

import "testing"

func TestSelect(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic with err:%s", err)
		}
	}()

	Select()
}

func TestDefault(t *testing.T) {
	Default()
}

func TestReturnFirst(t *testing.T) {
	ReturnFirst()
}
