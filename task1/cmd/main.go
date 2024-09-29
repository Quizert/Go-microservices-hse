package main

import (
	"fmt"
	"task1/internal/idgenerator"
	"task1/internal/library"
	"task1/internal/model"
	"task1/internal/storage"
)

func main() {
	//Создаем слайс книг
	book := library.CreateBook("Harry Potter", "J. K. Rowling")
	book2 := library.CreateBook("Witcher", "Andrzej Sapkowski")
	book3 := library.CreateBook("1984", "George Orwell")
	books := []*model.Book{book, book2, book3}
	//Создаем библиотеку и загружаем в нее книги
	FnvGen := idgenerator.CreateFnvGen()
	mapStore := storage.CreateMapStorage()
	lib := library.CreateLibrary(mapStore, FnvGen)
	for _, b := range books {
		lib.AddBook(b)
	}

	findHarryPotter, ok := lib.Search(book.Title) //Книга существует
	fmt.Println("I found the", findHarryPotter.Title, ok)

	searchDoesNotExist, ok := lib.Search("bibabo")
	fmt.Println(searchDoesNotExist, ok, "Книги не существует, поэтому ok = false, чтобы мы случайно не использовали nil")
	//Заменяем функцию генератора и находим еще книгу
	CrcGen := idgenerator.CreateCrcGen()
	lib.SetGenerator(CrcGen)
	findWitcher, ok := lib.Search(book2.Title)
	fmt.Println(findWitcher, ok, "Библиотека перегенерировала ID и нашла книгу")

	newSliceStorage := storage.CreateSliceStorage()
	lib.ReplaceStorage(newSliceStorage) //Заменяет хранилище с сохранением книг
	lib.AddBook(library.CreateBook("newbook", "alice"))
	lib.AddBook(library.CreateBook("tokyoghoul", "somebody"))
	lib.SetGenerator(FnvGen)
	fmt.Println(lib.Search("newbook")) //Все работает
}
