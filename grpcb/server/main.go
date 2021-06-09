package main

import (
	"context"
	"net"

	"github.com/404th/grpcb/proto"
	"google.golang.org/grpc"
)

type Server struct {
}

func main() {
	lis, err := net.Listen("tcp", ":3537")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterContactServiceServer(srv, &Server{})

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *Server) Plus(ctx context.Context, req *proto.Req) (*proto.Resp, error) {
	res := req.GetA() + req.GetB()

	return &proto.Resp{Res: res}, nil
}

func (s *Server) Mult(ctx context.Context, req *proto.Req) (*proto.Resp, error) {
	res := req.GetA() * req.GetB()

	return &proto.Resp{Res: res}, nil
}

func (s *Server) Div(ctx context.Context, req *proto.Req) (*proto.Resp, error) {
	res := req.GetA() / req.GetB()

	return &proto.Resp{Res: res}, nil
}

func (s *Server) Min(ctx context.Context, req *proto.Req) (*proto.Resp, error) {
	res := req.GetA() - req.GetB()

	return &proto.Resp{Res: res}, nil
}
