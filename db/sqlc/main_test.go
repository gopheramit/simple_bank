package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gopheramit/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testdb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// var err error
	testdb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db", err)

	}
	testQueries = New(testdb)
	os.Exit(m.Run())
}
