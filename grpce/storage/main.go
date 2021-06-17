package storage

import (
	"github.com/404th/grpce/storage/postgres"
	"github.com/404th/grpce/storage/repo"
	"github.com/jmoiron/sqlx"
)

type storagePg struct {
	db *sqlx.DB
	repo repo.UserStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db: db,
		repo: postgres.NewRepo(db),
	}
}

func (strg storagePg) Auth() repo.UserStorageI {
	return strg.repo
}

type StorageI interface {
	Auth() repo.UserStorageI
}
