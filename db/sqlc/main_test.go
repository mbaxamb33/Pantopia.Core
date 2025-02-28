package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/pantopia?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	log.Println("Database connection established successfully")

	testQueries = New(conn)

	code := m.Run() // Run tests and capture exit code
	conn.Close()    // Close the DB connection before exiting
	os.Exit(code)
}
