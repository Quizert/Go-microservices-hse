package storage

import (
	"task1/mylibrary/model"
)

type Storage interface {
	Search(id uint32) (*model.Book, bool)
	AddBook(book *model.Book) //Storage works with an internal book
	PrintBooks()
}
