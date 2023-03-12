package db

import "errors"

type Book struct {
	Id          uint
	Name        string
	AuthorName  string
	ReleaseYear string
	Price       uint32
}

var BookList []Book

func Insert(newBook Book) {
	BookList = append(BookList, newBook)
}

func Delete(id uint) error {
	for index, item := range BookList {
		if item.Id == id {
			BookList = append(BookList[:index], BookList[index+1:]...)
			return nil
		}
	}

	return errors.New("book with given id does not exists")
}

