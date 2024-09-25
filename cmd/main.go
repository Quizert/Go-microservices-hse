package main

import (
	"fmt"
	"task1/mylibrary/library"
)

func main() {
	mapStorage := library.CreateMapStorage()
	idGen := library.CreateIdGen()
	lib := library.NewLibrary(mapStorage, idGen)
	fmt.Println(a, ok)
}
