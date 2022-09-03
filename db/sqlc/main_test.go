package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/LiShangAn/bank/util"
	_ "github.com/lib/pq" // A pure Go postgres driver for Go's database/sql package
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://postgres:admin@127.0.0.1:5432/postgres?sslmode=disable"
// )

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	fmt.Println("start test")

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	log.Println("run test ...")
	testDB, err = sql.Open(config.DBdriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
