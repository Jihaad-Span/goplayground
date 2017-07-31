package migrate

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/mattes/migrate/database/mysql"
	//"github.com/mattes/migrate"
)

func NewSQLDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/play?multistatements=true")

	if err != nil {
		panic(err)
	}

	//driver, _ := mysql.WithInstance(db, &mysql.Config{})
	//m, _ := migrate.NewWithDatabaseInstance(
	//	"file:///migrations",
	//	"mysql",
	//	driver,
	//)

	//m.Steps(2)
	return db
}

func ShowBooksSQL(db *sql.DB) []Book {
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

type Book struct{
	Title string
	Author string
}