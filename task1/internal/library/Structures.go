package library

import (
	"task1/internal/idgenerator"
	"task1/internal/model"
)

type Storage interface {
	Search(id uint32) (*model.Book, bool)
	AddBook(book *model.Book) //Storage works with an internal book
	RegenerateId(generator idgenerator.Generator)
}

type Library struct {
	storage Storage
	idGen   idgenerator.Generator
}
