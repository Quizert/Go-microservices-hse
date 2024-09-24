package storage

import (
	"task1/mylibrary/internal/book"
)

type MapStorage struct {
	MapStorage map[uint32]book.Book
}

func (ms *MapStorage) Search(id uint32) (*book.UserBook, bool) {
	v, ok := ms.MapStorage[id]
	newUserBook := book.CreateUserBook(v.GetTitle(), v.GetAuthor())
	return newUserBook, ok
}

func (ms *MapStorage) AddBook(book *book.Book) {

	ms.MapStorage[book.Id] = *book
}
