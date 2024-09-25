package storage

import "task1/mylibrary/internal/model"

type SliceStorage struct {
	Storage []model.Book
}

func (ss *SliceStorage) Search(id uint32) (*model.Book, bool) {
	for _, book := range ss.Storage {
		if book.Id == id {
			return &book, true
		}
	}
	return nil, false
}

func (ss *SliceStorage) AddBook(book *model.Book) {
	ss.Storage = append(ss.Storage, *book)

}

func CreateSliceStorage() Storage {
	return &SliceStorage{Storage: make([]model.Book, 0)}
}
