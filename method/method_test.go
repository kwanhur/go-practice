package method

import "testing"

func TestFuncType(t *testing.T) {
	var b Book
	t.Logf("%T", b.Page)
	t.Logf("%T", b.Pages)
	t.Logf("%T", (&b).SetIndex)
	t.Logf("%T", (&b).Pages)
}

func TestBook_SetPages(t *testing.T) {
	book := Book{Page: 10}
	t.Logf("before page:%d", book.Pages())
	t.Log("set page to 100")
	book.SetPages(100)
	t.Logf("after page:%d", book.Pages())

	t.Log("call implicit declared func set and get")
	(*Book).SetPages(&book, 101)
	t.Logf("implicit call page:%d", Book.Pages(book))
}

func TestBook_Pages2(t *testing.T) {
	book := Book{Page: 100}
	t.Log(book.Pages2())
}

func TestBook_SetIndex(t *testing.T) {
	var book Book
	t.Logf("before index:%d", book.index)
	t.Logf("set index to 112")
	book.SetIndex(112)
	t.Logf("after set index:%d", book.index)
}
