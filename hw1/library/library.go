package library

import (
	"fmt"

	"github.com/YourCurseSheyme/go_homework_2025/hw1/book"
	"github.com/YourCurseSheyme/go_homework_2025/hw1/generators"
	"github.com/YourCurseSheyme/go_homework_2025/hw1/storages"
)

var (
	ErrorEmptyStorage     = fmt.Errorf("storage is empty")
	ErrorEmptyIDGenerator = fmt.Errorf("id generator is empty")
)

type Library struct {
	Data   storages.Storage
	Titles map[string][]int
	IDGen  generators.IDGenerator
}

func NewLibrary(data storages.Storage, idGen generators.IDGenerator) *Library {
	return &Library{
		Data:   data,
		Titles: make(map[string][]int),
		IDGen:  idGen,
	}
}

func (l *Library) ReplaceIDGen(newGen generators.IDGenerator) {
	l.IDGen = newGen
}

func (l *Library) ReplaceStorage(newData storages.Storage) {
	l.Data = newData
	l.Titles = make(map[string][]int)
}

func (l *Library) Add(title, author string, year int) (book.Book, error) {
	if l.Data == nil {
		return book.Book{}, ErrorEmptyStorage
	}
	if l.IDGen == nil {
		return book.Book{}, ErrorEmptyIDGenerator
	}
	id := l.IDGen()
	item := book.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Year:   year,
	}
	l.Data.AddBook(item)
	l.Titles[title] = append(l.Titles[title], id)
	return item, nil
}

func (l *Library) Remove(id int) error {
	if l.Data == nil {
		return ErrorEmptyStorage
	}
	item, err := l.Data.GetByID(id)
	if err != nil {
		return err
	}
	if err := l.Data.RemoveByID(id); err != nil {
		return err
	}
	ids := l.Titles[item.Title]
	newIDs := make([]int, 0, len(ids))
	for _, stored := range ids {
		if stored != id {
			newIDs = append(newIDs, stored)
		}
	}
	if len(newIDs) == 0 {
		delete(l.Titles, item.Title)
	} else {
		l.Titles[item.Title] = newIDs
	}
	return nil
}

func (l *Library) FindByTitle(title string) ([]book.Book, error) {
	if l.Data == nil {
		return nil, ErrorEmptyStorage
	}
	ids := l.Titles[title]
	if len(ids) == 0 {
		return nil, storages.ErrorBookNotFound
	}
	books := make([]book.Book, 0, len(ids))
	for _, id := range ids {
		if item, err := l.Data.GetByID(id); err == nil {
			books = append(books, item)
		} else {
			fmt.Println("> error: ", err)
		}
	}
	return books, nil
}

func Demo() {
	fmt.Println("> Homework №1")

	fmt.Println("Fill 'n' print the library")
	fmt.Println("+-----------")
	test := []book.Book{
		{Title: "Мечтают ли андроиды об электроовцах?", Author: "Филип Дик", Year: 1968},
		{Title: "Дюна", Author: "Фрэнк Герберт", Year: 1963},
		{Title: "Мечтают ли андроиды об электроовцах?", Author: "Филип Дик", Year: 1968},
	}
	lib := NewLibrary(storages.NewSliceStorage(), generators.IncrementalIDGen())
	for _, item := range test {
		_, err := lib.Add(item.Title, item.Author, item.Year)
		if err != nil {
			fmt.Println("> error:", err)
		}
	}
	for idx := 0; idx < len(test)-1; idx++ {
		result, err := lib.FindByTitle(test[idx].Title)
		if err != nil {
			fmt.Println("> error:", err)
			continue
		}
		fmt.Printf("result len: %d\n", len(result))
		for _, item := range result {
			fmt.Println(item)
		}
		fmt.Println("+-----------")
	}

	fmt.Println("Delete book")
	fmt.Println("+-----------")
	err := lib.Remove(2)
	if err != nil {
		fmt.Println("> error:", err)
	} else {
		fmt.Println("book id=1 removed")
	}
	fmt.Println("+-----------")

	fmt.Println("Find deleted book")
	fmt.Println("+-----------")
	_, err = lib.FindByTitle("Дюна")
	if err != nil {
		fmt.Println("> error:", err)
	} else {
		fmt.Println("book id=1 found, unexpected behavior")
	}
	fmt.Println("+-----------")

	fmt.Println("Swap ID generator")
	fmt.Println("+-----------")
	lib.ReplaceIDGen(generators.RandomIDGen(1000000))
	yaTest := []book.Book{
		{Title: "Марсианин", Author: "Энди Вейер", Year: 2011},
		{Title: "Мечтают ли андроиды об электроовцах?", Author: "Филип Дик", Year: 1968},
	}
	for _, item := range yaTest {
		_, err := lib.Add(item.Title, item.Author, item.Year)
		if err != nil {
			fmt.Println("> error:", err)
		}
	}
	for idx := 0; idx < len(yaTest); idx++ {
		result, err := lib.FindByTitle(yaTest[idx].Title)
		if err != nil {
			fmt.Println("> error:", err)
			continue
		}
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
		_, err := lib.Add(item.Title, item.Author, item.Year)
		if err != nil {
			fmt.Println("> error:", err)
		}
	}
	for idx := 1; idx < len(wholeTest)-1; idx++ {
		result, err := lib.FindByTitle(wholeTest[idx].Title)
		if err != nil {
			fmt.Println("> error:", err)
			continue
		}
		fmt.Printf("result len: %d\n", len(result))
		for _, item := range result {
			fmt.Println(item)
		}
		fmt.Println("+-----------")
	}
}
