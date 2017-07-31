package database

import (
	"testing"
	"os"
	"flag"
	"database/sql"
)

var databaseFlag = flag.Bool("database", false, "run database integration tests")
var messageQueue = flag.Bool("messageQueue", false, "run message queue integration tests")
var integration = flag.Bool("integration", false, "run all integration tests")
var db *sql.DB

func TestMain(m *testing.M) {
	flag.Parse()

	if *databaseFlag || *integration {
		setupDatabase()
	}

	if *messageQueue || *integration {
		setupDatabase()
	}

	//if !testing.Short() {
	//}
	result := m.Run()

	//if *databaseFlag || *integration {
	//	teardownDatabase()
	//}
	//
	//if *messageQueue || *integration {
	//	teardownDatabase()
	//}
	//if !testing.Short() {
	//	teardownDatabase()
	//}
	os.Exit(result)
}


func TestDatabaseGet(t *testing.T) {
	if !*databaseFlag && !*integration {
		t.Skip()
	}

	res := ShowBooks2(db)

	if res[0].Title != "Meep" {
		t.Error("Expected Span")
	}

	if res[0].Author != "Jihaad" {
		t.Error("Expected Jihaad")
	}
}

func setupDatabase() {
	db = NewDB()
}
