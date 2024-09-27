package library

import (
	"fmt"
	"task1/internal/idgenerator"
	"task1/internal/model"
	"task1/internal/storage"
)

func NewMapLibrary(gen Generator) Library {
	return &library{storage: CreateMapStorage(), idGen: gen.GetFunc()}
}

func NewSliceLibrary(gen Generator) Library {
	return &library{storage: CreateSliceStorage(), idGen: gen.GetFunc()}
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
		newSliceStore := CreateSliceStorage()
		for _, book := range store.Storage {
			newSliceStore.AddBook(book)
		}
		l.storage = newSliceStore
	case *storage.SliceStorage:
		newMapStore := CreateMapStorage()
		for _, book := range store.Storage {
			newMapStore.AddBook(book)
		}
		l.storage = newMapStore
	default:
		_ = fmt.Errorf("error, strange storage type")
	}
}

func CreateBook(title, author string) *model.Book {
	return &model.Book{Title: title, Author: author}
}

func CreateFnvGen() Generator {
	return &idgenerator.GeneratorStruct{GenFunc: idgenerator.FnvID}
}
func CreateCrcGen() Generator {
	return &idgenerator.GeneratorStruct{GenFunc: idgenerator.CrcID}
}

func (l *library) SetGenerator(gen Generator) {
	l.idGen = gen.GetFunc()
}

func CreateMapStorage() Storage {
	return &storage.MapStorage{Storage: make(map[uint32]*model.Book)}
}
func CreateSliceStorage() Storage {
	return &storage.SliceStorage{Storage: make([]*model.Book, 0)}
}
