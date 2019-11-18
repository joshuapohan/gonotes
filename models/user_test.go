package models

import(
	"log"
	"testing"
	"database/sql"
	"github.com/lib/pq"
)

const DB_URL = "postgres://postgres:admin@localhost:5432/go-notes?sslmode=disable"
const clearDBStmt = "delete from users where 1=1"

var db *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func TestSaveNewUser(t *testing.T) {
	db = ConnectDB()

	db.QueryRow(clearDBStmt)

	u := &User{name: "jp", email: "jp@test.com", pwd:"12345"}
	err := u.SaveNewUser(db)
	if err != nil {
		t.Error(err)
		t.Error("Failed to insert new user to database")
	}
}