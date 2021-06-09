package psql2

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Run(cfg Config) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatalf("error while connecting to psql: %v", err.Error())
		return nil, err
	}
	if err != nil {
		log.Fatalf("error while closing psql: %v", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error while connecting to psql in PING: %v", err.Error())
	}

	return db, nil
}
