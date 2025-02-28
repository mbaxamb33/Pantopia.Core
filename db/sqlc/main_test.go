package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/mbaxamb33/pantopia/util"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	log.Println("Database connection established successfully")

	testQueries = New(conn)

	code := m.Run() // Run tests and capture exit code
	conn.Close()    // Close the DB connection before exiting
	os.Exit(code)
}
