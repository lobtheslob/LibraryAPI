package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	// ===================================
	// connect to database
	// ===================================

	// connect to the dbs
	mySQLDB := connectToDB()
	if migErr := handleMigrations(mySQLDB); migErr != nil {
		log.Error("migration error: ", migErr)
	}

	repo := NewRepository(mySQLDB)

	// initialize a panic watcher
	defer recover()

	// Init router
	r := mux.NewRouter()

	//TODO - implement repo in remaining HandleFuncs with closure like getBooks
	// Route handles & endpoints
	r.HandleFunc("/books", getBooks(repo)).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))

}