package main

import (
	"github.com/404th/contactlistgrpc/postgres"
	"github.com/404th/contactlistgrpc/repo"
	"github.com/404th/contactlistgrpc/service"
	"github.com/404th/contactlistgrpc/service/grpcclient"
)

func main() {

	db, err := postgres.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	client, err := grpcclient.New()
	if err != nil {
		panic(err)
	}

	contactService := service.NewContactService()

}
