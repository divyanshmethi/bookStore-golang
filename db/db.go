package db

import (
	"database/sql"
	"fmt"

	"bookStore/config"
	"bookStore/models"
)

type PostgresDBMSImpl struct {
	DB *sql.DB
}

//go:generate mockery --name=DBMS --case=underscore
type DBMS interface {
	AddBook(book *models.Book) error
	GetAllBooks() ([]*models.Book, error)
	GetBook(bookID string) (*models.Book, error)
	DeleteBook(bookID string) (int, error)
	UpdateBook(book *models.Book) (int, error)
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
