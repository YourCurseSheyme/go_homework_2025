package storages

import "github.com/YourCurseSheyme/go_homework_2025/hw1/book"

type SliceStorage struct {
	Data []book.Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{Data: make([]book.Book, 0)}
}

func (s *SliceStorage) AddBook(book book.Book) {
	s.Data = append(s.Data, book)
}

func (s *SliceStorage) GetByID(id int) (book.Book, error) {
	for idx := range s.Data {
		if s.Data[idx].ID == id {
			return s.Data[idx], nil
		}
	}
	return book.Book{}, ErrorBookNotFound
}

func (s *SliceStorage) RemoveByID(id int) error {
	for idx, item := range s.Data {
		if item.ID == id {
			s.Data = append(s.Data[:idx], s.Data[idx+1:]...)
			return nil
		}
	}
	return ErrorBookNotFound
}
