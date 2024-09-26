package library

import (
	"fmt"
	"task1/mylibrary/internal/idgenerator"
	"task1/mylibrary/internal/storage"
	"task1/mylibrary/model"
)

func NewMapLibrary(gen idgenerator.Generator) Library {
	return &library{storage: storage.CreateMapStorage(), idGen: gen}
}

func NewSliceLibrary(gen idgenerator.Generator) Library {
	return &library{storage: storage.CreateSliceStorage(), idGen: gen}
}

func (l *library) Search(title string) (*model.Book, bool) {
	id := l.idGen(title)
	return l.storage.Search(id)
}

func (l *library) AddBook(book *model.Book) {
	book.Id = l.generateID(book.Title)
	l.storage.AddBook(book)
}

func (l *library) generateID(title string) uint32 {
	return l.idGen(title)
}

func (l *library) ReplaceStorage() {
	switch store := l.storage.(type) {
	case *storage.MapStorage:
		newSliceStore := storage.CreateSliceStorage()
		for _, book := range store.Storage {
			newSliceStore.AddBook(book)
		}
		l.storage = newSliceStore
	case *storage.SliceStorage:
		newMapStore := storage.CreateMapStorage()
		for _, book := range store.Storage {
			newMapStore.AddBook(book)
		}
		l.storage = newMapStore
	default:
		_ = fmt.Errorf("error, strange storage type")
	}
}

func (l *library) PrintBooks() {
	l.storage.PrintBooks()
}

func CreateBook(title, author string) *model.Book {
	return &model.Book{Title: title, Author: author}
}

func CreateFnvGen() func(title string) uint32 {
	return idgenerator.FnvID
}

func CreateCrcGen() func(title string) uint32 {
	return idgenerator.CrcID
}

func (l *library) ReplaceIdGen(gen idgenerator.Generator) {
	l.idGen = gen
}
