package grpc_client

import (
	"log"

	"github.com/404th/grpce/protos/protos"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	connections map[string]interface{}
}

func New () (*GrpcClient, error) {
	connAuth, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error while initializing auth service: %v", err)
	}

	return &GrpcClient{
		connections: map[string]interface{}{
			"auth_service": protos.NewAuthServiceClient(connAuth),
		},
	}, nil
}

func (g *GrpcClient) AuthService() protos.AuthServiceClient {
	return g.connections["auth_service"].(protos.AuthServiceClient)
}
