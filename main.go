package main

import (
	"fmt"
	g "./graphql"
	//db "github.com/goplayground/database"
	mysqlDB "github.com/goplayground/migrate"
)

func main() {
	//database := db.NewDB()
	//results := db.ShowBooks2(database)

	database := mysqlDB.NewSQLDatabase()
	results := mysqlDB.ShowBooksSQL(database)
	fmt.Println("results: ", results)
	fmt.Println(g.Play())
}
