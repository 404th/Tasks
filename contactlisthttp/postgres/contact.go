package postgres

import (
	"context"

	pb "github.com/404th/contactlistgrpc/genproto"
	"github.com/404th/contactlistgrpc/repo"
	"github.com/404th/contactlistgrpc/service/grpcclient"
)

type contactRepo struct {
	client *grpcclient.GrpcClient
}

// NewOrderRepo is ...
func NewContactRepo(client *grpcclient.GrpcClient) repo.ContactStorageI {
	return &contactRepo{
		client: client,
	}
}

func (r *contactRepo) CreateContact(ctx context.Context, ctc *pb.ContactCreateReq) (*pb.WithID, error) {
	return &pb.WithID{}, nil
}

func (r *contactRepo) GetContact(ctx context.Context, id *pb.WithID) (*pb.Contact, error) {
	return &pb.Contact{}, nil
}

func (r *contactRepo) GetAllContacts(ctx context.Context, ctc *pb.EmptyResponse) (*pb.Contacts, error) {
	return &pb.Contacts{}, nil
}

func (r *contactRepo) UpdateContact(ctx context.Context, ctc *pb.Contact) (*pb.WithID, error) {
	return &pb.WithID{}, nil
}

func (r *contactRepo) DeleteContact(ctx context.Context, ctc *pb.WithID) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, nil
}
