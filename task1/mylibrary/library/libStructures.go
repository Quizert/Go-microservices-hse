package library

import (
	"task1/mylibrary/internal/storage"
)

type library struct {
	storage storage.Storage
	idGen   func(title string) uint32
}
