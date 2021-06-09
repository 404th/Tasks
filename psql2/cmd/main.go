package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/404th/psql2"
	"github.com/404th/psql2/storage"
	"github.com/404th/psql2/storage/models"
	"github.com/google/uuid"
)

func main() {

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

	strg := storage.NewStoragePg(db)

	// CREATE
	id, err := strg.Task().Create(models.Task{
		ID:    "",
		Title: uuid.NewString(),
		Body:  "Ertaga tog'ga chiqamiz.Hamma tayyor bo'lib tursin.",
		Done:  rand.Intn(2) == 1,
	})
	if err != nil {
		log.Fatalf("Error: can't create => %v", err)
	}
	fmt.Printf("After creating => %s\n", id)

	// GET
	task, err := strg.Task().Get(id)
	if err != nil {
		log.Fatalf("Error: can't get => %v", err)
	}
	fmt.Printf("After getting a task => %v\n", task)

}
