package controllers

/**
 * @file book-controller.go
 * @author Original author: Free Code Camp
 *		   Changes made by 0xChristopher for learning purposes
 *
 * This file handles incoming API calls, utilizes corresponding 'Book' functions for database interaction,
 * and then returns information back to the front end.
 */

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ChristopherH-eth/GoBookstore/pkg/models"
	"github.com/ChristopherH-eth/GoBookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var newBook models.Book

// The GetBook() function writes the list of all books in the database on the frontend.
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// The GetBookById() function writes a book to the frontend based on the Id.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(Id)
	response, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// The CreateBook() function creates a new book, adds it to the database, and writes it to the frontend.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	response, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// The DeleteBook() function deletes a book from the database.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(Id)
	response, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// The UpdateBook() function updates a book in the database and writes the new details to the frontend.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, db := models.GetBookById(Id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	response, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
