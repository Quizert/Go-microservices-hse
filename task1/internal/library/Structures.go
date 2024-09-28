package library

import (
	"task1/internal/idgenerator"
	"task1/internal/storage"
)

type Library struct {
	storage storage.Storage
	idGen   idgenerator.Generator
}
