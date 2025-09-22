package library

import (
	"fmt"

	"github.com/YourCurseSheyme/go_homework_2025/hw1/book"
	"github.com/YourCurseSheyme/go_homework_2025/hw1/generators"
	"github.com/YourCurseSheyme/go_homework_2025/hw1/storages"
)

type Library struct {
	Data   storages.Storage
	Titles map[string][]int
	IDGen  generators.IDGenerator
}

func NewLibrary(data storages.Storage, idGen generators.IDGenerator) *Library {
	return &Library{Data: data, Titles: make(map[string][]int), IDGen: idGen}
}

func (l *Library) ReplaceIDGen(newGen generators.IDGenerator) {
	l.IDGen = newGen
}

func (l *Library) ReplaceStorage(newData storages.Storage) {
	l.Data = newData
	l.Titles = make(map[string][]int)
}

func (l *Library) Add(title, author string, year int) (book.Book, bool) {
	if l.Data == nil {
		return book.Book{}, false
	}
	if l.IDGen == nil {
		return book.Book{}, false
	}
	id := l.IDGen()
	book := book.Book{ID: id, Title: title, Author: author, Year: year}
	l.Data.AddBook(book)
	l.Titles[title] = append(l.Titles[title], id)
	return book, true
}

func (l *Library) FindByTitle(title string) []book.Book {
	if l.Data == nil {
		return nil
	}
	ids := l.Titles[title]
	if len(ids) == 0 {
		return nil
	}
	books := make([]book.Book, 0, len(ids))
	for _, id := range ids {
		if book, ok := l.Data.GetByID(id); ok {
			books = append(books, book)
		}
	}
	return books
}

func Demo() {
	fmt.Println("> Homework №1")

	test := []book.Book{
		{Title: "Мечтают ли андроиды об электроовцах?", Author: "Филип Дик", Year: 1968},
		{Title: "Дюна", Author: "Фрэнк Герберт", Year: 1963},
		{Title: "Мечтают ли андроиды об электроовцах?", Author: "Филип Дик", Year: 1968},
	}
	lib := NewLibrary(storages.NewSliceStorage(), generators.IncrementalIDGen())
	for _, item := range test {
		lib.Add(item.Title, item.Author, item.Year)
	}
	for idx := 0; idx < len(test)-1; idx++ {
		result := lib.FindByTitle(test[idx].Title)
		fmt.Printf("result len: %d\n", len(result))
		for _, item := range result {
			fmt.Println(item)
		}
		fmt.Println("+-----------")
	}

	fmt.Println("Swap ID generator")
	fmt.Println("+-----------")
	lib.ReplaceIDGen(generators.RandomIDGen(1000000))
	yaTest := []book.Book{
		{Title: "Марсианин", Author: "Энди Вейер", Year: 2011},
		{Title: "Мечтают ли андроиды об электроовцах?", Author: "Филип Дик", Year: 1968},
	}
	for _, item := range yaTest {
		lib.Add(item.Title, item.Author, item.Year)
	}
	for idx := 0; idx < len(yaTest); idx++ {
		result := lib.FindByTitle(yaTest[idx].Title)
		fmt.Printf("result len: %d\n", len(result))
		for _, item := range result {
			fmt.Println(item)
		}
		fmt.Println("+-----------")
	}

	fmt.Println("Swap storage type")
	fmt.Println("+-----------")
	lib.ReplaceStorage(storages.NewMapStorage())
	wholeTest := append(test, yaTest...)
	for _, item := range wholeTest {
		lib.Add(item.Title, item.Author, item.Year)
	}
	for idx := 1; idx < len(wholeTest)-1; idx++ {
		result := lib.FindByTitle(wholeTest[idx].Title)
		fmt.Printf("result len: %d\n", len(result))
		for _, item := range result {
			fmt.Println(item)
		}
		fmt.Println("+-----------")
	}
}
