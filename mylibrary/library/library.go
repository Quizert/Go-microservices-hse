package library

import "task1/mylibrary/internal/book"

type Library interface {
	Search(title string) (*book.UserBook, bool)
	AddBook(book *book.UserBook)
	generateID(title string) uint32
}
