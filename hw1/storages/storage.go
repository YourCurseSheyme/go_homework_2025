package storages

import (
	"fmt"

	"github.com/YourCurseSheyme/go_homework_2025/hw1/book"
)

var ErrorBookNotFound = fmt.Errorf("book doesn't exist")

type Storage interface {
	AddBook(book book.Book)
	GetByID(id int) (book.Book, error)
	RemoveByID(id int) error
}
