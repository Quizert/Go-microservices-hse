package library

import "task1/mylibrary/internal/model"

type Library interface {
	Search(title string) (*model.Book, bool)
	addBook(book *model.Book)
	generateID(title string) uint32
}
