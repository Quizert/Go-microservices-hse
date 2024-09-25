package storage

import (
	"task1/mylibrary/internal/model"
)

type MapStorage struct {
	Storage map[uint32]model.Book
}

func (ms *MapStorage) Search(id uint32) (*model.Book, bool) {
	book, ok := ms.Storage[id]
	return &book, ok
}

func (ms *MapStorage) AddBook(book *model.Book) {
	ms.Storage[book.Id] = *book
}

func FillMap(ms MapStorage) {

}
