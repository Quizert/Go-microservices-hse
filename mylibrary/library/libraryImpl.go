package library

import (
	"task1/mylibrary/internal/book"
	"task1/mylibrary/internal/idgenerator"
	"task1/mylibrary/internal/storage"
)

type library struct {
	storage storage.Storage
	idGen   func(title string) uint32
}

func NewLibrary(s storage.Storage, idGen func(title string) uint32) Library {
	return &library{storage: s, idGen: idGen}
}

func (l *library) Search(title string) (*book.UserBook, bool) {
	id := l.idGen(title)
	return l.storage.Search(id)
}

func (l *library) AddBook(b *book.UserBook) {
	newInternalBook := book.CreateInternalBook(b.GetTitle(), b.GetAuthor(), l.idGen(b.GetTitle()))
	l.storage.AddBook(newInternalBook)
}

func (l *library) generateID(title string) uint32 {
	return l.idGen(title)
}

func CreateMapStorage() storage.Storage {
	return &storage.MapStorage{MapStorage: make(map[uint32]book.Book)}
}
func CreateIdGen() func(title string) uint32 {
	return idgenerator.FnvID
}

func CreateBook(title string, author string) *book.UserBook {
	return book.CreateUserBook(title, author)
}
