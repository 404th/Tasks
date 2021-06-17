package main

import (
	"fmt"
	"log"
	"net"

	"github.com/404th/grpce/protos/protos"
	"github.com/404th/grpce/service"
	"github.com/404th/grpce/service/grpc_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// pb "github.com/404th/grpce/protos/protos"
)

func main () {
	var (
		port string = ":5051"
	)

	cfg := &Config {
		host: "localhost",
		port: "5432",
		user: "spenc",
		dbname: "test_db",
		password: "mysecretpassword",
		sslmode: "disable",
	}

	connDB, err := NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("error while connection todb %v", err)
		return
	}

	client, err := grpc_client.New()

	authService := service.NewAuthService(connDB, client)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}

	s := grpc.NewServer()

	protos.RegisterAuthServiceServer(s, authService)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error while serving; %v", err)
	}
	fmt.Printf("Server is running on port: %v", port)
}