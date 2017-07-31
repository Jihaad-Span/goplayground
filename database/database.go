package database

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//var title, author string
		var myBook Book
		err := db.QueryRow("select title, author from books").Scan(&myBook.Title, &myBook.Author)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(rw, "The first book is '%s' by '%s'", myBook.Title, myBook.Author)
	})
}

func ShowBooks2(db *sql.DB) []Book {
	var myBooks []Book
	var myBook Book
	rows, err := db.Query("select title, author from books")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&myBook.Title, &myBook.Author)
		myBooks = append(myBooks, myBook)
	}

	return myBooks
}

func NewDB() *sql.DB {
	fmt.Println("skdnfodsnfidsnfds")
	db, err := sql.Open("sqlite3", "testDB.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists books(title text, author text)")

	if err != nil {
		panic(err)
	}

	return db
}

type Book struct{
	Title string
	Author string
}



