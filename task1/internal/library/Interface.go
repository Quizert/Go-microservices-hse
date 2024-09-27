package library

import (
	"task1/internal/model"
)

type Storage interface {
	Search(id uint32) (*model.Book, bool)
	AddBook(book *model.Book) //Storage works with an internal book
}

type Generator interface {
	GenerateID(title string) uint32
	GetFunc() func(title string) uint32
}

type Library interface {
	Search(title string) (*model.Book, bool)
	AddBook(book *model.Book)
	SetGenerator(gen Generator)
	ReplaceStorage()
}
