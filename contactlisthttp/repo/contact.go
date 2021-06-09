package repo

import (
	"context"

	pb "github.com/404th/contactlistgrpc/genproto"
)

type ContactStorageI interface {
	CreateContact(context.Context, *pb.ContactCreateReq) (*pb.WithID, error)
	GetContact(context.Context, *pb.WithID) (*pb.Contact, error)
	GetAllContacts(context.Context, *pb.EmptyResponse) (*pb.Contacts, error)
	UpdateContact(context.Context, *pb.Contact) (*pb.WithID, error)
	DeleteContact(context.Context, *pb.WithID) (*pb.EmptyResponse, error)
}
