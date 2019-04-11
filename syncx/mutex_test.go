package syncx

import "testing"

func TestUserAges_Add(t *testing.T) {
	u := UserAges{
		ages: make(map[string]int),
	}
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	u.Add("kwanhur", 20)
	t.Log(u.Get("kwanhur"))
}
