package grpcclient

import (
	pb "github.com/404th/contactlistgrpc/genproto"
	"google.golang.org/grpc"
)

type GrpcClientI interface {
	ContactService() pb.ContactServiceClient
}

type GrpcClient struct {
	connections map[string]interface{}
}

func New() (*GrpcClient, error) {
	connContact, err := grpc.Dial("contact service dial host: localhost port: 8045", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		connections: map[string]interface{}{
			"contact_service": pb.NewContactServiceClient(connContact),
		},
	}, nil
}

func (g *GrpcClient) ContactService() pb.ContactServiceClient {
	return g.connections["contact_service"].(pb.ContactServiceClient)
}
