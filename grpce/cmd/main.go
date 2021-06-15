package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type Server struct {
	db *sqlx.DB
}

func main () {
	var (
		port string = ":5051"
	)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error while serving; %v", err)
	}
	fmt.Printf("Server is running on port: %v", port)
}