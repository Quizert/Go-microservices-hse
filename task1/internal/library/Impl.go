package library

import (
	"task1/internal/idgenerator"
	"task1/internal/model"
	"task1/internal/storage"
)

func (l *Library) Search(title string) (*model.Book, bool) {
	id := l.idGen.GenerateID(title)
	return l.storage.Search(id)
}

func (l *Library) AddBook(book *model.Book) {
	book.Id = l.idGen.GenerateID(book.Title)
	l.storage.AddBook(book)
}

func (l *Library) ReplaceStorage() {
	l.storage = l.storage.ReplaceStorage()
}

func (l *Library) SetGenerator(gen idgenerator.Generator) {
	l.idGen = gen
	l.storage.RegenerateId(l.idGen)
}

func CreateBook(title, author string) *model.Book {
	return &model.Book{Title: title, Author: author}
}

func CreateLibrary(storage storage.Storage, idGen idgenerator.Generator) *Library {
	return &Library{storage: storage, idGen: idGen}
}
