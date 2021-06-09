package service

import (
	"context"

	pb "github.com/404th/contactlistgrpc/genproto"
	"github.com/404th/contactlistgrpc/repo"
	"github.com/404th/contactlistgrpc/service/grpcclient"
)

type contactService struct {
	strg repo.ContactStorageI
}

func NewContactService(client *grpcclient.GrpcClient) *contactService {
	return &contactService{
		strg: strg.
	}
}

func (cs *contactService) CreateContact(ctx context.Context, create pb.ContactCreateReq) (pb.WithID, error) {
	// cs.strg.Contact().Create(create.Name, create.Pa)
	return pb.WithID{}, nil
}

func (cs *contactService) GetContact(ctx context.Context, ctc pb.WithID) (pb.Contact, error) {
	return pb.Contact{}, nil
}

func (cs *contactService) GetAllContacts(ctx context.Context, create pb.ContactCreateReq) (pb.Contacts, error) {
	return pb.Contacts{}, nil
}

func (cs *contactService) UpdateContact(ctx context.Context, create pb.ContactCreateReq) (pb.WithID, error) {
	return pb.WithID{}, nil
}

func (cs *contactService) DeleteContact(ctx context.Context, create pb.ContactCreateReq) (pb.EmptyResponse, error) {
	return pb.EmptyResponse{}, nil
}
