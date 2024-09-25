package library

import (
	"task1/mylibrary/internal/idgenerator"
	"task1/mylibrary/internal/model"
	"task1/mylibrary/internal/storage"
)

type library struct {
	storage storage.Storage
	idGen   func(title string) uint32
}

func NewLibrary(s storage.Storage, idGen func(title string) uint32) Library {
	return &library{storage: s, idGen: idGen}
}

func (l *library) Search(title string) (*model.Book, bool) {
	id := l.idGen(title)
	return l.storage.Search(id)
}

func (l *library) addBook(book *model.Book) {
	book.Id = l.generateID(book.Title)
	l.storage.AddBook(book)
}

func (l *library) generateID(title string) uint32 {
	return l.idGen(title)
}

func CreateMapStorage() storage.Storage {
	return &storage.MapStorage{Storage: make(map[uint32]model.Book)}
}

func CreateSliceStorage() storage.Storage {
	return &storage.SliceStorage{Storage: make([]model.Book, 0)}
}

func CreateIdGen() func(title string) uint32 {
	return idgenerator.FnvID
}
