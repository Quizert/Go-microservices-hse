package main

import (
	"fmt"
	"task1/mylibrary/library"
)

func main() {
	mapStorage := library.CreateMapStorage()
	idGen := library.CreateIdGen()

	lib := library.NewLibrary(mapStorage, idGen)

	book := library.CreateBook("aaaa", "bbbb")
	book2 := library.CreateBook("ccc", "ddd")

	lib.AddBook(book)
	lib.AddBook(book2)
	alice, ok := lib.Search("aaaa")
	biba, ok := lib.Search("ccc")

	fmt.Println(alice.GetTitle(), ok)
	fmt.Println(biba.GetTitle(), ok)
	fmt.Println(lib.Search("ASDAD"))
}
