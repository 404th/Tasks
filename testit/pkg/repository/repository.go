package repository

import (
	"github.com/404th/testit/genproto"
	"github.com/jmoiron/sqlx"
)

type ContactRepo interface {
	CreateContact(ctc *genproto.ContactCreateReq) (genproto.WithID, error)
	GetContact(ctc *genproto.WithID) (genproto.Contact, error)
	GetAllContacts(ctc *genproto.EmptyResponse) (genproto.Contacts, error)
	UpdataContact(ctc *genproto.Contact) (genproto.Contact, error)
	DeleteContact(ctc *genproto.WithID) (genproto.EmptyResponse, error)
}

type Repo struct {
	ContactRepo
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		ContactRepo: NewContactRepository(db),
	}
}
