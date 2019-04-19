package method

type Book struct {
	Page  int
	index int
}

func (b Book) Pages() int {
	return b.Page
}

func (b *Book) SetPages(page int) {
	b.Page = page
}

func (b *Book) Pages2() int {
	return Book.Pages(*b)
}

func (b *Book) SetIndex(index int) {
	b.index = index
}
