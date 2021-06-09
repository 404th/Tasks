package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/404th/grpcc/protos"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type server struct {
	db *sqlx.DB
}

func main() {
	//
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=spenc dbname=test_db password=mysecretpassword sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	//
	lis, err := net.Listen("tcp", ":3433")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running on port: 3433")

	srv := grpc.NewServer()
	protos.RegisterContactServiceServer(srv, &server{
		db: db,
	})

	fmt.Println("Successfully connected to Psql!")
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *server) Create(ctx context.Context, req *protos.Contact) (*protos.WithName, error) {
	query := `INSERT INTO contacts (name, username) VALUES ($1, $2)`
	fmt.Print()
	_, err := s.db.Exec(query, req.GetName(), req.GetUsername())
	if err != nil {
		log.Fatalf("error while creating: %v", err)
		return nil, err
	}

	return &protos.WithName{
		Name: req.GetName(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *protos.WithName) (*protos.Contact, error) {
	var prt protos.Contact
	query := `SELECT name, username FROM contacts WHERE name=$1`
	row := s.db.QueryRow(query, req.GetName())
	if err := row.Scan(&prt.Name, &prt.Username); err != nil {
		log.Fatalf("error while getting: %v", err)
		return nil, err
	}

	return &prt, nil
}

func (s *server) Delete(ctx context.Context, req *protos.WithName) (*protos.Contact, error) {
	var prt protos.Contact

	query1 := `SELECT name, username FROM contacts WHERE name=$1`
	row := s.db.QueryRow(query1, req.GetName())
	if err := row.Scan(&prt.Name, &prt.Username); err != nil {
		log.Fatalf("error while getting: %v", err)
		return nil, err
	}

	query := `DELETE FROM contacts WHERE name=$1`
	_, err := s.db.Exec(query, req.GetName())
	if err != nil {
		log.Fatalf("error while deleting: %v", err)
		return nil, err
	}

	return &prt, nil
}
