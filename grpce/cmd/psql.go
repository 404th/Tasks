package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	host string
	port string
	user string
	dbname string
	password string
	sslmode string
}

func NewPsqlDB (cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.host, cfg.port, cfg.user, cfg.password, cfg.password, cfg.sslmode))

	if err != nil {
		log.Fatalf("error while connecting to psql: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error while connecting to db: %v", err)
		return nil, err
	}

	return db, nil
}