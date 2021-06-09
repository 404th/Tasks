package postgres

import "github.com/jmoiron/sqlx"

func NewPostgresDB() (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=spenc dbname=test_db password=mysecretpassword sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db, nil
}
