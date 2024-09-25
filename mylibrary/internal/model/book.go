package model

type Book struct {
	Title  string
	Author string
	Id     uint32
}

func CreateBook(title, author string, id uint32) *Book {
	return &Book{title, author, id}
}
