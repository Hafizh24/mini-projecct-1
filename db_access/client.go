package db_access

import "bookstore_system/model"

type DBClient interface {
	AddBook(b model.Book)
	GetBook(id int) (model.Book, error)

}

type BookStore struct {
	books []model.Book
}



func (bs *BookStore) AddBook(b model.Book) {
	bs.books = append(bs.books, b)
}

func (bs *BookStore) GetBook(id int) (model.Book, error) {
	for _, book := range bs.books {

		if book.ID == id {
			return book, nil
		}
	}
	return model.Book{}, nil
}