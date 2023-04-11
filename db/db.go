package db

import (
	"database/sql"
	"fmt"

	"bookStore/config"
)

type PostgresDBMSImpl struct {
	DB *sql.DB
}

// Book holds all info related to a book
type Book struct {
	BookID string  `json:"bookId"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

//go:generate mockery --name=DBMS --case=underscore
type DBMS interface {
	AddBook(book *Book) error
	GetAllBooks() ([]*Book, error)
	GetBook(bookID string) (*Book, error)
	DeleteBook(bookID string) (int, error)
	UpdateBook(book *Book) (int, error)
}

func NewPostgresDBMSImpl() DBMS {
	dbConfig := config.GetGlobalConfig().Data.Postgres
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("Unable to connect to DB")
	}
	return &PostgresDBMSImpl{DB: database}
}
