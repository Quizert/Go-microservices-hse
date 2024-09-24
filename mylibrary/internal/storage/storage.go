package storage

import (
	"task1/mylibrary/internal/book"
)

type Storage interface {
	Search(id uint32) (*book.UserBook, bool)
	AddBook(book *book.Book)
}
