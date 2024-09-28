package storage

import (
	"task1/internal/idgenerator"
	"task1/internal/model"
)

type Storage interface {
	Search(id uint32) (*model.Book, bool)
	AddBook(book *model.Book) //Storage works with an internal book
	ReplaceStorage() Storage
	RegenerateId(generator idgenerator.Generator)
}
