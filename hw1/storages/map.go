package storages

import "github.com/YourCurseSheyme/go_homework_2025/hw1/book"

type MapStorage struct {
	Data map[int]book.Book
}

func NewMapStorage() *MapStorage {
	return &MapStorage{Data: make(map[int]book.Book)}
}

func (s *MapStorage) AddBook(book book.Book) {
	s.Data[book.ID] = book
}

func (s *MapStorage) GetByID(id int) (book.Book, bool) {
	item, ok := s.Data[id]
	return item, ok
}
