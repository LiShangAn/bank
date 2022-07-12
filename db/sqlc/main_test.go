package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // A pure Go postgres driver for Go's database/sql package
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:admin@127.0.0.1:5432/postgres?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	log.Println("run test ...")
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
