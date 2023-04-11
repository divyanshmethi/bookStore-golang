package db

//import (
//	"database/sql"
//	"fmt"
//
//	"bookStore/config"
//	"bookStore/models"
//	_ "github.com/lib/pq"
//)
//
//const (
//	createQuery          = `INSERT into books(bookID,name,author,price) values($1,$2,$3,$4)`
//	selectAllQuery       = `SELECT * from books`
//	selectParticularBook = `SELECT * from books where bookID = $1`
//	deleteBook           = `DELETE from books where bookID = $1`
//	updateBook           = `UPDATE books set name = $2, author = $3, price = $4 where bookID = $1`
//)
//
//// Todo: Explore -- How we can create an instance at the runtime and inject into the server
//// In that case we would have methods define on DB
//// We would separate the functions as an interface and let it implement them -- Currently it looks very ugly
//var DB *sql.DB
//
//func ConnectToDB() {
//	//Todo: Think about null pointer exception here
//	dbConfig := config.GetGlobalConfig().Data.Postgres
//	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
//	var err error
//	DB, err = sql.Open("postgres", psqlInfo)
//	if err != nil {
//		panic("Unable to connect to DB")
//	}
//}
//
//func (database *PostgresDBMSImpl) AddBook(book *models.Book) error {
//	res, err := DB.Query(createQuery, book.BookID, book.Name, book.Author, book.Price)
//	if err != nil {
//		return err
//	}
//	fmt.Println(res)
//	return nil
//}
//
//// Todo: DB layer would have a separate struct for Book along with a mapper
//func (database *PostgresDBMSImpl) GetAllBooks() ([]*models.Book, error) {
//	res, err := DB.Query(selectAllQuery)
//	if err != nil {
//		return nil, err
//	}
//	defer res.Close()
//	fmt.Println(res)
//	var books []*models.Book
//	for res.Next() {
//		var book models.Book
//		err = res.Scan(&book.BookID, &book.Name, &book.Author, &book.Price)
//		if err != nil {
//			fmt.Println("Problem reading book from DB")
//			continue
//		}
//		books = append(books, &book)
//	}
//	return books, nil
//}
//
//// Todo: DB layer would have a separate struct for Book along with a mapper
//func (database *PostgresDBMSImpl) GetBook(bookID string) (*models.Book, error) {
//	fmt.Println(bookID)
//	res, err := DB.Query(selectParticularBook, bookID)
//	if err != nil {
//		return nil, err
//	}
//	if res == nil {
//		return nil, nil
//	}
//	defer res.Close()
//	var book models.Book
//	err = res.Scan(&book.BookID, &book.Name, &book.Author, &book.Price)
//	if err != nil || book.BookID == "" {
//		return nil, nil
//	}
//	return &book, nil
//}
//
//// Todo: DB layer would have a separate struct for Book along with a mapper
//func (database *PostgresDBMSImpl) DeleteBook(bookID string) (int, error) {
//	res, err := DB.Exec(deleteBook, bookID)
//	if err != nil {
//		return 1, err
//	}
//	n, err := res.RowsAffected()
//	if err != nil || n == 0 {
//		return 0, nil
//	}
//	return 1, nil
//}
//
//func (database *PostgresDBMSImpl) UpdateBook(book *models.Book) (int, error) {
//	res, err := DB.Exec(updateBook, book.BookID, book.Name, book.Author, book.Price)
//	if err != nil {
//		return 1, err
//	}
//	n, err := res.RowsAffected()
//	if err != nil || n == 0 {
//		return 0, nil
//	}
//	return 1, nil
//}
