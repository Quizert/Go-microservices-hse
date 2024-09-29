package storage

import (
	"task1/internal/idgenerator"
	"task1/internal/library"
	"task1/internal/model"
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

func (ss *SliceStorage) RegenerateId(generator idgenerator.Generator) {
	for _, book := range ss.Storage {
		newID := generator.GenerateID(book.Title)
		book.Id = newID
	}
}
func CreateSliceStorage() library.Storage {
	return &SliceStorage{Storage: make([]*model.Book, 0)}
}
