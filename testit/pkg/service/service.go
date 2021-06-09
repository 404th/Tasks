package service

import (
	"github.com/404th/testit/genproto"
	"github.com/404th/testit/pkg/repository"
)

type ContactServiceI interface {
	CreateContact(contact *genproto.ContactCreateReq) (genproto.WithID, error)
	GetContact(contact *genproto.WithID) (genproto.Contact, error)
	GetAllContacts(contact *genproto.EmptyResponse) (genproto.Contacts, error)
	UpdateContact(contact *genproto.Contact) (genproto.Contact, error)
	DeleteContact(contact *genproto.WithID) (genproto.EmptyResponse, error)
}

type Service struct {
	ContactServiceI
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		ContactServiceI: NewContactSrv(repo),
	}
}
