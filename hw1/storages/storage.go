package storages

import "github.com/YourCurseSheyme/go_homework_2025/hw1/book"

type Storage interface {
	AddBook(book book.Book)
	GetByID(id int) (book.Book, bool)
}
