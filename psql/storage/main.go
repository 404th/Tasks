package storage

import (
	"github.com/404th/Tasks/psql/storage/postgres"
	"github.com/404th/Tasks/psql/storage/repo"
	"github.com/jmoiron/sqlx"
)

func (s storagePg) Contact() repo.ContactStorageI {
	return s.contactRepo
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:          db,
		contactRepo: postgres.NewContactRepo(db),
	}
}

type storagePg struct {
	db          *sqlx.DB
	contactRepo repo.ContactStorageI
}

type StorageI interface {
	Contact() repo.ContactStorageI
}
