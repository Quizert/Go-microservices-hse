package storage

import (
	"task1/internal/idgenerator"
	"task1/internal/library"
	"task1/internal/model"
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

func (ms *MapStorage) RegenerateId(generator idgenerator.Generator) {
	for k, v := range ms.Storage {
		newID := generator.GenerateID(v.Title)
		delete(ms.Storage, k)
		v.Id = newID
		ms.Storage[newID] = v
	}
}

func CreateMapStorage() library.Storage {
	return &MapStorage{Storage: make(map[uint32]*model.Book)}
}
