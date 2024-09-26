package storage

import (
	"fmt"
	"task1/mylibrary/model"
)

type SliceStorage struct {
	Storage []*model.Book
}

func (ss *SliceStorage) Search(id uint32) (*model.Book, bool) {
	for _, book := range ss.Storage {
		if book.Id == id {
			return book, true
		}
	}
	return nil, false
}

func (ss *SliceStorage) AddBook(book *model.Book) {
	ss.Storage = append(ss.Storage, book)
}

func (ss *SliceStorage) PrintBooks() {
	for i, book := range ss.Storage {
		fmt.Printf("%d) Title: %s, Author: %s, Id: %d\n", i+1, book.Title, book.Author, book.Id)
	}
}
func CreateSliceStorage() Storage {
	return &SliceStorage{Storage: make([]*model.Book, 0)}
}
