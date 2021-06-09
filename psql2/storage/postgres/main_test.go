package postgres_test

import (
	"log"
	"testing"

	"github.com/404th/psql2"
	"github.com/404th/psql2/storage"
)

var (
	strg storage.StorageI
)

func TestMain(t *testing.M) {
	cfg := psql2.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "spenc",
		Password: "mysecretpassword",
		DBName:   "test_db",
		SSLMode:  "disable",
	}

	db, err := psql2.Run(cfg)
	if err != nil {
		log.Fatalf("error occured while connecting db: %s", err.Error())
	}

	defer db.Close()
}
