package db

import (
	"fmt"

	_ "github.com/lib/pq"
)

const (
	createQuery          = `INSERT into books(bookID,name,author,price) values($1,$2,$3,$4)`
	selectAllQuery       = `SELECT * from books`
	selectParticularBook = `SELECT * from books where bookID = $1`
	deleteBook           = `DELETE from books where bookID = $1`
	updateBook           = `UPDATE books set name = $2, author = $3, price = $4 where bookID = $1`
)

func (database *PostgresDBMSImpl) AddBook(book *Book) error {
	res, err := database.DB.Query(createQuery, book.BookID, book.Name, book.Author, book.Price)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (database *PostgresDBMSImpl) GetAllBooks() ([]*Book, error) {
	res, err := database.DB.Query(selectAllQuery)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	fmt.Println(res)
	var books []*Book
	for res.Next() {
		var book Book
		err = res.Scan(&book.BookID, &book.Name, &book.Author, &book.Price)
		if err != nil {
			fmt.Println("Problem reading book from DB")
			continue
		}
		books = append(books, &book)
	}
	return books, nil
}

func (database *PostgresDBMSImpl) GetBook(bookID string) (*Book, error) {
	fmt.Println(bookID)
	res, err := database.DB.Query(selectParticularBook, bookID)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	defer res.Close()
	var book Book
	err = res.Scan(&book.BookID, &book.Name, &book.Author, &book.Price)
	if err != nil || book.BookID == "" {
		return nil, nil
	}
	return &book, nil
}

func (database *PostgresDBMSImpl) DeleteBook(bookID string) (int, error) {
	res, err := database.DB.Exec(deleteBook, bookID)
	if err != nil {
		return 1, err
	}
	n, err := res.RowsAffected()
	if err != nil || n == 0 {
		return 0, nil
	}
	return 1, nil
}

func (database *PostgresDBMSImpl) UpdateBook(book *Book) (int, error) {
	res, err := database.DB.Exec(updateBook, book.BookID, book.Name, book.Author, book.Price)
	if err != nil {
		return 1, err
	}
	n, err := res.RowsAffected()
	if err != nil || n == 0 {
		return 0, nil
	}
	return 1, nil
}
