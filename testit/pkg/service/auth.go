package service

import (
	"github.com/404th/testit/genproto"
	"github.com/404th/testit/pkg/repository"
)

type ContactSrv struct {
	repo repository.ContactRepo
}

func NewContactSrv(repo repository.ContactRepo) *ContactSrv {
	return &ContactSrv{
		repo: repo,
	}
}

func (u *ContactSrv) CreateContact(contact *genproto.ContactCreateReq) (genproto.WithID, error) {
	return u.repo.CreateContact(contact)
}

func (u *ContactSrv) GetContact(contact *genproto.WithID) (genproto.Contact, error) {
	return u.repo.GetContact(contact)
}

func (u *ContactSrv) GetAllContacts(contact *genproto.EmptyResponse) (genproto.Contacts, error) {
	return u.repo.GetAllContacts(contact)
}

func (u *ContactSrv) UpdateContact(contact *genproto.Contact) (genproto.Contact, error) {
	return u.repo.UpdataContact(contact)
}

func (u *ContactSrv) DeleteContact(contact *genproto.WithID) (genproto.EmptyResponse, error) {
	return u.repo.DeleteContact(contact)
}
