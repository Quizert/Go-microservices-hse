package storage

import (
	"task1/mylibrary/internal/model"
)

type Storage interface {
	Search(id uint32) (*model.Book, bool)
	AddBook(book *model.Book)
}
