package main

import (
	"fmt"
	"task1/internal/library"
	"task1/internal/model"
)

func main() {
	//Создаем слайс книг
	book := library.CreateBook("Harry Potter", "J. K. Rowling")
	book2 := library.CreateBook("Witcher", "Andrzej Sapkowski")
	book3 := library.CreateBook("1984", "George Orwell")
	books := []*model.Book{book, book2, book3}
	//Создаем библиотеку и загружаем в нее книги
	FnvGen := library.CreateFnvGen()
	lib := library.NewMapLibrary(FnvGen)
	for _, b := range books {
		lib.AddBook(b)
	}

	findHarryPotter, ok := lib.Search(book.Title) //Книга существует
	fmt.Println("I found the", findHarryPotter.Title, ok)

	searchDoesNotExist, ok := lib.Search("bibabo")
	fmt.Println(searchDoesNotExist, ok, "Книги не существует, поэтому ok = false, чтобы мы случайно не использовали nil")
	//Заменяем функцию генератора и находим еще книгу
	CrcGen := library.CreateCrcGen()
	lib.SetGenerator(CrcGen)
	findWitcher, ok := lib.Search(book2.Title)
	fmt.Println(findWitcher, ok, "Заменили функцию генератора, поэтому библиотека не может найти книгу по новому id")
	lib.ReplaceStorage() //Заменяет хранилище с сохранением книг

	lib2 := library.NewSliceLibrary(CrcGen)

	exampleBook := library.CreateBook("bibabo", "Alen")
	lib2.AddBook(exampleBook)

	b, ok := lib2.Search(exampleBook.Title)
	fmt.Println("I found the", b.Title, ok)
}
