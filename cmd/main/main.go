package main

/**
 * @file main.go
 * @author Original author: Free Code Camp
 *		   Changes made by 0xChristopher for learning purposes
 *
 * This application demonstrates simple web server functionality by handling multiple routes via the Gorilla
 * Mux package. It also uses the Gorm package to allow API calls to interact with a MySQL database. A user
 * can view the current book entries, view a book by its id, add a book to the list, update a current book
 * entry, as well as delete books from the list.
 *
 * The main.go file handles router initialization and webserver creation.
 */

import (
	"log"
	"net/http"

	"github.com/ChristopherH-eth/GoBookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Initialize new router and attach Bookstore routes; server startup
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
