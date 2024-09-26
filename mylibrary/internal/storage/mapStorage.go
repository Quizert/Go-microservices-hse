package storage

import (
	"fmt"
	"task1/mylibrary/model"
)

type MapStorage struct {
	Storage map[uint32]*model.Book
}

func (ms *MapStorage) Search(id uint32) (*model.Book, bool) {
	book, ok := ms.Storage[id]
	return book, ok
}

func (ms *MapStorage) AddBook(book *model.Book) {
	ms.Storage[book.Id] = book
}

func (ms *MapStorage) PrintBooks() {
	i := 1
	for _, book := range ms.Storage {
		fmt.Printf("%d) Title: %s, Author: %s, Id: %d\n", i, book.Title, book.Author, book.Id)
		i++
	}
}
func CreateMapStorage() Storage {
	return &MapStorage{Storage: make(map[uint32]*model.Book)}
}
