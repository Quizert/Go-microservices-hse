package library

import (
	"task1/mylibrary/internal/idgenerator"
	"task1/mylibrary/model"
)

type Library interface {
	Search(title string) (*model.Book, bool)
	AddBook(book *model.Book)
	ReplaceStorage() //Меняет хранилище, но сохраняет книги
	ReplaceIdGen(gen idgenerator.Generator)
	PrintBooks()
	generateID(title string) uint32
}
