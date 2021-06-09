package postgres_test

import (
	"log"
	"os"
	"testing"

	"github.com/404th/Tasks/psql/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	strg storage.StorageI
)

func TestMain(m *testing.M) {

	connStr := "user=spenc dbname=test_db password=mysecretpassword host=localhost sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatalf("error connect psql: %s", err)
		return
	}
	strg = storage.NewStoragePg(db)
	os.Exit(m.Run())
}
