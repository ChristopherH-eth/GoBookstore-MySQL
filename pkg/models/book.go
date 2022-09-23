package models

/**
 * @file book.go
 * @author Original author: Free Code Camp
 *		   Changes made by 0xChristopher for learning purposes
 *
 * This file defines the Book struct as well as migrating the current database connection to the 'books'
 * table. It also defines the ways in which operations can be used with a Book and the database.
 */

import (
	"github.com/ChristopherH-eth/GoBookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// The init() function initializes a connection to the database.
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// The CreateBook() function creates and returns a new Book object
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// The GetAllBooks() function returns a list of all books in the database
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// The GetBookById() function returns a book from the database based on its Id
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

// The DeleteBook() function deletes a book from the database based on its Id
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(book)
	return book
}
