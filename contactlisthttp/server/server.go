package server

import (
	"fmt"
	"net"

	contact "github.com/404th/contactlistgrpc/genproto"
	"github.com/404th/contactlistgrpc/service"
	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) Run(port int) {
	fmt.Println("gRPC server is running!")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	contactService := service.ContactService{}

	grpcServer := grpc.NewServer()
	contact.RegisterContactServiceServer(grpcServer, &contactService)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("gRPC server is ready!")
}
