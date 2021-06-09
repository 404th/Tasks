package main

import (
	"github.com/404th/testit/pkg/handler"
	"github.com/404th/testit/pkg/repository"
	"github.com/404th/testit/pkg/service"
	"github.com/404th/testit/server"
	_ "github.com/lib/pq"
)

func main() {
	cfg := repository.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "spenc",
		Password: "mysecretpassword",
		DBName:   "test_db",
		SSLMode:  "disable",
	}

	db := repository.NewPostgresDB(cfg)

	repo := repository.NewRepo(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	err := srv.Run("3454", handlers.InitRoute())
	if err != nil {
		panic(err)
	}
}
