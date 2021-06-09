package main

import (
	"net"

	"github.com/404th/grpcd/protos"
	"github.com/404th/grpcd/server/psql"
	"github.com/404th/grpcd/server/storage"
	"github.com/404th/grpcd/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	db *sqlx.DB
}

func main() {
	db := psql.Run()
	storage := storage.NewStoragePg(db)
	taskServ := service.NewTaskService(storage)

	lis, err := net.Listen("tcp", ":3435")
	if err != nil {
		panic(err)
	}

	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	protos.RegisterServiceNameServer(gsrv, taskServ)

	if err := gsrv.Serve(lis); err != nil {
		panic(err)
	}
}
