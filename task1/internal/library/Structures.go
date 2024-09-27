package library

type library struct {
	storage Storage
	idGen   func(title string) uint32
}
