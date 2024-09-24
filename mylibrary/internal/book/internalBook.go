package book

type Book struct {
	UserBook
	Id uint32
}

func CreateInternalBook(t, author string, id uint32) *Book {
	return &Book{UserBook{t, author}, id}
}
